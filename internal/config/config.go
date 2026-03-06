package config

import (
	"fmt"
	"os"
)

type Config struct {
	Stage          string
	OpenSearchURL  string
	OpenSearchUser string
	OpenSearchPass string
}

func Init() Config {
	c := Config{
		Stage:          getEnvValue("STAGE", "dev"),
		OpenSearchURL:  getEnvironmentVariables("OPENSEARCH_URL"),
		OpenSearchUser: getEnvironmentVariables("OPENSEARCH_USERNAME"),
		OpenSearchPass: getEnvironmentVariables("OPENSEARCH_PASSWORD"),
	}
	return c
}

func getEnvironmentVariables(k string) string {
	v := os.Getenv(k)
	if v == "" {
		panic(fmt.Sprintf("missing required env var: %s", k))
	}
	return v
}

func getEnvValue(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}
