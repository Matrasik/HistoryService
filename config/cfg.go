package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	CurrPairs  []string
	UpdateTime int
}

func New() *Config {
	return &Config{
		CurrPairs:  getEnvSlice("CURR_PAIRS", []string{}, ","),
		UpdateTime: getEnvInt("UPDATE_TIME", 60),
	}
}

func getEnvSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)
	return val
}

func getEnvInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnv(key string, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultVal
}
