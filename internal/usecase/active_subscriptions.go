package usecase

import (
	"context"
	"net/http"
	"sync"
	"time"
  "encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/samber/lo"

	"github.com/Coflnet/sky-controller/internal/metrics"
	"github.com/Coflnet/sky-controller/internal/utils"
	api "github.com/Coflnet/sky-controller/target/payment"
)

var (
  paymentClient *api.Client
)

type ActiveSubscriptionsWatcher struct {
  Interval time.Duration
  ProductUpdateInterval time.Duration

  Slugs []string
  UsersPerSlug sync.Map
}

func (w *ActiveSubscriptionsWatcher) Start() {
  w.init()

  go func () {
    for {
      err := w.update()

      if err != nil {
        log.Error().Err(err).Msgf("Error while watching active subscriptions")
      }
      time.Sleep(w.Interval)
    }
  }()

  go func () {
    for {
      err := w.updateProducts()

      if err != nil {
        log.Error().Err(err).Msgf("Error while updating products")
      }
      time.Sleep(w.ProductUpdateInterval)
    }
  }()
} 

func (w *ActiveSubscriptionsWatcher) init() {
  var err error
  paymentClient, err = api.NewClient(utils.PaymentBaseURL())

  if err != nil {
    log.Panic().Err(err).Msg("Failed to create payment client")
  }
}

func (w *ActiveSubscriptionsWatcher) updateProducts() error {
  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  products, err := paymentClient.ProductsGet(ctx, api.ProductsGetParams{
    Amount: api.NewOptInt32(100),
    Offset: api.NewOptInt32(0),
  })

  if err != nil {
    return err
  }

  
  w.Slugs = lo.FilterMap[api.PurchaseableProduct, string](products, func (product api.PurchaseableProduct, _ int) (string, bool) {
    if product.Slug.Set {
      return product.Slug.Value, true
    }
    return "", false
  })

  return nil
}

func (w *ActiveSubscriptionsWatcher) update() error {

  // update all product slugs
  // do up to 10 concurrent requests
  // wait for all requests to finish

  sem := make(chan int, 10)
  wg := sync.WaitGroup{}
  w.UsersPerSlug = sync.Map{}

  for _, slug := range w.Slugs {

    sem <- 1
    wg.Add(1)

    go func (slug string) {
      defer func() { 
        <-sem 
        wg.Done()
      }()

      count, err := w.updateActiveUsersPerProduct(slug)
      if err != nil {
        log.Error().Err(err).Msgf("Error while updating active users for %s", slug)
      }

      w.UsersPerSlug.Store(slug, count)
    }(slug)
  }

  // update tfm user count
  tfmUserCount, err := w.TFMUserCount()
  if err != nil {
    log.Error().Err(err).Msgf("Error while updating tfm user count")
  } else {
    log.Debug().Msgf("TFM user count: %d", tfmUserCount)
    w.UsersPerSlug.Store("tfm", tfmUserCount)
  }


  wg.Wait()
  metrics.UpdateActiveSubscriptions(w.UsersPerSlug)

  log.Debug().Msgf("Updated active subscriptions")

  return nil
}

func (w *ActiveSubscriptionsWatcher) updateActiveUsersPerProduct(slug string) (int32, error) {
  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  count, err := paymentClient.ProductsServiceServiceSlugCountGet(ctx, api.ProductsServiceServiceSlugCountGetParams{
    ServiceSlug: slug,
  })

  return count, err
}

func (w *ActiveSubscriptionsWatcher) TFMUserCount() (int32, error) {
  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  // make http get request
  req, err := http.NewRequestWithContext(ctx, "GET", utils.TFMUserCountURL(), nil)
  if err != nil {
    return 0, err
  }

  resp, err := http.DefaultClient.Do(req)

  if err != nil {
    return 0, err
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    return 0, err
  }


  var tfmResponse TFMResponse
  err = json.NewDecoder(resp.Body).Decode(&tfmResponse)

  if err != nil {
    return 0, err
  }

  return int32(tfmResponse.Users), nil
}

	
type TFMResponse struct {
	Users    int             `json:"users"`
}

