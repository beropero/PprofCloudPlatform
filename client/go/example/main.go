package main

import (
	client "github.com/beropero/PprofCloudPlatform/client/go/v1"
	"github.com/beropero/PprofCloudPlatform/client/go/v1/config"
)

func main() {
	// 示例代码
	client, err := client.NewClient(&config.Config{
		Port:         8090,
		ServiceToken: "mis_tsjsz30o6s0d7yjenqj8dt0300l72ov0",
	})
	if err != nil {
		panic(err)
	}
	client.Start()
}
