package main

import (
	_ "net/http/pprof"
	"net/http"

	client "github.com/beropero/PprofCloudPlatform/client/go/v1"
	"github.com/beropero/PprofCloudPlatform/client/go/v1/config"
)

func main() {
	// 示例代码
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
	client, err := client.NewClient(&config.Config{
		Port:         8090,
		ServiceToken: "mis_tsjsz30o6s0d7yjenqj8dt0300l72ov0",
	})
	if err != nil {
		panic(err)
	}
	client.Start()
}
