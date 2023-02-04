package utils

import (
	"os"

	"github.com/rs/zerolog/log"
)

func getEnv(key string) string {
  v := os.Getenv(key)
  if v == "" {
    log.Panic().Msgf("Environment variable %s is not set", key)
  }
  return v
}

func ProxyBaseURL() string {
  return getEnv("PROXY_BASE_URL")
}

func PaymentBaseURL() string {
  return getEnv("PAYMENTS_BASE_URL")
}

func TFMUserCountURL() string {
  return "https://sky.coflnet.com/tfm/backend/online_tfm_users"
}
