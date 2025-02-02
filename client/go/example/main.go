package main

import (
	"time"

	client "github.com/beropero/PprofCloudPlatform/client/go/v1"
	"github.com/beropero/PprofCloudPlatform/client/go/v1/config"
)

func main() {
	// 示例代码
	client, err := client.NewClient(&config.Config{
		Interval:  30 * time.Second,
		Timeout:   10 * time.Second,
		Port:      8090,
		UploadUrl: "http://127.0.0.1:8000/upload/uploadtotal",
		Token:     "123456",
	})
	if err != nil {
		panic(err)
	}
	client.Start()
}
