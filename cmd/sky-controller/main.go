package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Coflnet/sky-controller/internal/usecase"
)

func main() {
	proxyScaler := usecase.ProxyScaler{
		Interval: 1 * time.Minute,
	}
  proxyScaler.Start()

	// wait for exit signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("Shutting down...")
}
