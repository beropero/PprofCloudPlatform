package v1

import (
	"github.com/beropero/PprofCloudPlatform/client/go/v1/config"
	"github.com/beropero/PprofCloudPlatform/client/go/v1/controller"
)

type Client struct {
	config     *config.Config
	controller *controller.Controller
}

func NewClient(cfg *config.Config) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &Client{
		config:     cfg,
		controller: &controller.Controller{Config: cfg},
	}, nil
}

func (c *Client) Start() error {
	c.controller.WebServerStart()
	return nil
}
