package config

import (
	"fmt"
	"time"
)

type Config struct {
	Token     string
	Interval  time.Duration
	Timeout   time.Duration
	UploadUrl string
	Port      int
}

// DefaultConfig returns the default config
func DefaultConfig() *Config {
	return &Config{
		Interval:  30 * time.Second,
		Timeout:   10 * time.Second,
		UploadUrl: "https://api.pprof.control.dev/v1/upload",
		Port:      8080,
	}
}

// Validate checks if the config is valid
func (c *Config) Validate() error {
	defaultConfig := DefaultConfig()
	if c.Token == "" {
		return fmt.Errorf("token must be specified")
	}
	if c.UploadUrl == "" {
		c.UploadUrl = defaultConfig.UploadUrl
	}
	if c.Interval <= 0 {
		c.Interval = defaultConfig.Interval
	}
	if c.Timeout <= 0 {
		c.Timeout = defaultConfig.Timeout
	}
	if c.Port <= 0 {
		c.Port = defaultConfig.Port
	}
	return nil
}
