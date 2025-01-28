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
	}
}

// Validate checks if the config is valid
func (c *Config) Validate() error {
	if c.Token == "" {
		return fmt.Errorf("token must be specified")
	}
	if c.UploadUrl == "" {
		return fmt.Errorf("upload url must be specified")
	}
	if c.Interval <= 0 {
		return fmt.Errorf("interval must be greater than 0")
	}
	if c.Timeout <= 0 {
		return fmt.Errorf("timeout must be greater than 0")
	}
	return nil
}
