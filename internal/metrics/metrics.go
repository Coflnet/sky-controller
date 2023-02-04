package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start() {
  go func() {
    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":2112", nil)  
  }()

  prometheus.MustRegister(ActiveSubscriptionsForProduct)
}
