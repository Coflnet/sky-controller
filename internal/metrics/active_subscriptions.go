package metrics

import (
  "sync"
  "github.com/prometheus/client_golang/prometheus"
)

var (
  ActiveSubscriptionsForProduct = prometheus.NewGaugeVec(prometheus.GaugeOpts{
    Name: "sky_controller_active_subscriptions_for_product",
    Help: "The current number of active subscriptions for a product",
  }, []string{"product"})
)

func UpdateActiveSubscriptions(usersPerProduct sync.Map) {

  usersPerProduct.Range(func (key, value interface{}) bool {
    ActiveSubscriptionsForProduct.With(prometheus.Labels{
      "product": key.(string),
    }).Set(float64(value.(int)))

    return true
  })
}
