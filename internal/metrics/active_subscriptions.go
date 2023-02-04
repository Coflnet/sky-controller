package metrics

import (
  "github.com/prometheus/client_golang/prometheus"
)

var (
  ActiveSubscriptionsForProduct = prometheus.NewGaugeVec(prometheus.GaugeOpts{
    Name: "sky_controller_active_subscriptions_for_product",
    Help: "The current number of active subscriptions for a product",
  }, []string{"product"})
)

func UpdateActiveSubscriptions(usersPerProduct map[string]int) {

  for product, users := range usersPerProduct {
    ActiveSubscriptionsForProduct.With(prometheus.Labels{
      "product": product,
    }).Set(float64(users))
  }
}
