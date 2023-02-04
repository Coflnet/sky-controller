package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Coflnet/sky-controller/internal/usecase"
	"github.com/Coflnet/sky-controller/internal/metrics"
)

func main() {

  go metrics.Start()

	proxyScaler := usecase.ProxyScaler{
		Interval: 1 * time.Minute,
	}
  proxyScaler.Start()

  activeSubscriptionsWatcher := usecase.ActiveSubscriptionsWatcher{
    Interval: 1 * time.Minute,
    ProductUpdateInterval: 1 * time.Hour,
  }
  activeSubscriptionsWatcher.Start()

	// wait for exit signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("Shutting down...")
}
