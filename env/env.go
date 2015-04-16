package env

import (
	"os"
	"strconv"
)

func envS(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}

	return v
}

func envI(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}

	vi, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	return vi
}
