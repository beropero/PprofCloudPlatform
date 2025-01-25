package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	UploadURL string `yaml:"upload_url"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
