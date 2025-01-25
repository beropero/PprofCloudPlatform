package main

// import (
//     "log"
//     "time"

// )

// func main() {
//     // 加载配置
//     cfg, err := config.LoadConfig("config.yaml")
//     if err != nil {
//         log.Fatalf("Failed to load config: %v", err)
//     }

//     // 初始化 SDK
//     pprofClient := client.NewPProfClient(cfg.UploadURL)

//     // 采集并上传 pprof 数据
//     if err := pprofClient.CollectAndUpload(10 * time.Second); err != nil {
//         log.Fatalf("Failed to collect and upload pprof data: %v", err)
//     }
// }