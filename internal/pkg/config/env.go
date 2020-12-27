package config

import "os"

func getEnv(name string, fallback string) string {
  val := os.Getenv(name)
  if len(val) == 0 {
    return fallback
  }
  return val
}
