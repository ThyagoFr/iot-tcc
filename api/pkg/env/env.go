package env

import (
  "os"
"strconv"
)

func GetOrDefaultString(key, defaultValue string) string {
  value := os.Getenv(key)
  if value == "" {
    return defaultValue
  }
  return value
}

func GetOrDefaultInt(key string, defaultValue int) int {
  value := os.Getenv(key)
  if value == "" {
    return defaultValue
  }
  intValue, _ := strconv.Atoi(value)
  return intValue
}
