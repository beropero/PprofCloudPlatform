package config

import (
	"fmt"
	"time"
)

type Config struct {
	ProjectToken string
	ServiceToken string
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
		UploadUrl: "http://127.0.0.1:8086/upload/uploadfile",
	}
}

// Validate checks if the config is valid
func (c *Config) Validate() error {
	defaultConfig := DefaultConfig()
	if c.ServiceToken == "" {
		return fmt.Errorf("ServiceToken must be specified")
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
