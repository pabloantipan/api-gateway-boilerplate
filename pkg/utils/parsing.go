package utils

import (
	"os"
	"strconv"
)

func ParseEnvFloat64(key string) float64 {
	value, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil {
		return 0.0
	}
	return value
}
