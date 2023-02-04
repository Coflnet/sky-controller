package metrics

import (
  "github.com/prometheus/client_golang/prometheus"
)

var (
  proxyReplicasGauge = prometheus.NewGauge(prometheus.GaugeOpts{
    Name: "sky_controller_proxy_replicas",
    Help: "The current number of replicas of the proxy",
  })
)

func UpdateProxyReplicas(replicas int) {
  proxyReplicasGauge.Set(float64(replicas))
}
