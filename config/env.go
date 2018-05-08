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
	return nil
}

func (c *EnvConfig) GetValue(k string) string {
	if result, ok := c.pairs[k]; ok {
		return result
	}

	return ""
}

func (c *EnvConfig) Review() []string {
	var result []string
	for k, v := range c.pairs {
		result = append(result, "[env] - "+k+"="+v)
	}

	return result
}
