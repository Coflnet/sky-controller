package usecase

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/attribute"

	"github.com/Coflnet/sky-controller/internal/metrics"
	"github.com/Coflnet/sky-controller/internal/utils"
	api "github.com/Coflnet/sky-controller/target/payment"
)

var (
  paymentClient *api.Client
)

const subUpdateName = "subscriber_update"

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
  paymentClient, err = api.NewClient(utils.PaymentBaseURL(), api.WithTracerProvider(otel.GetTracerProvider()))

  if err != nil {
    log.Panic().Err(err).Msg("Failed to create payment client")
  }
}

func (w *ActiveSubscriptionsWatcher) updateProducts() error {

  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()
  _, span := otel.Tracer("active_subscriptions").Start(ctx, "update-products")
  defer span.End()

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
  // do up to 2 concurrent requests
  // wait for all requests to finish

  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  var span trace.Span
  ctx, span = otel.Tracer(subUpdateName).Start(ctx, "update-products")
  defer span.End()


  sem := make(chan int, 2)
  wg := sync.WaitGroup{}
  w.UsersPerSlug = sync.Map{}

  for _, slug := range w.Slugs {

    sem <- 1
    wg.Add(1)

    go func (ctx context.Context, slug string) {

      var s trace.Span
      _, s = otel.Tracer(subUpdateName).Start(ctx, "update-product")
      s.SetAttributes(attribute.String("slug", slug))
      defer s.End()

      defer func() { 
        <-sem 
        wg.Done()
      }()

      count, err := w.updateActiveUsersPerProduct(ctx, slug)
      if err != nil {
        log.Error().Err(err).Msgf("Error while updating active users for %s", slug)
      }

      w.UsersPerSlug.Store(slug, count)
    }(ctx, slug)
  }

  // update tfm user count
  tfmUserCount, err := w.TFMUserCount(ctx)
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

func (w *ActiveSubscriptionsWatcher) updateActiveUsersPerProduct(ctx context.Context, slug string) (int32, error) {
  count, err := paymentClient.ProductsServiceServiceSlugCountGet(ctx, api.ProductsServiceServiceSlugCountGetParams{
    ServiceSlug: slug,
  })

  return count, err
}

func (w *ActiveSubscriptionsWatcher) TFMUserCount(ctx context.Context) (int32, error) {

  var span trace.Span
  ctx, span = otel.Tracer(subUpdateName).Start(ctx, "tfm-user-count")
  defer span.End()

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

