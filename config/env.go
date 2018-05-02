package config

import (
	"os"
	"strings"
)

type EnvConfig struct {
	pairs map[string]string
}

func NewEnvConfig() *EnvConfig {
	config := &EnvConfig{
		pairs: make(map[string]string),
	}

	return config
}

func (c *EnvConfig) Load() error {
	for _, e := range os.Environ() {
		arr := strings.Split(e, "=")
		c.pairs[arr[0]] = arr[1]
	}
}

func (c *EnvConfig) GetValue(k string) string {
	if result, ok := c.pairs[k]; ok {
		return result
	}

	return ""
}
