package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

// 定义接收数据的结构体
type RequestData struct {
	Token string   `json:"token"`
	Types []string `json:"types"`
}

func (c *Controller) WebServerStart() {
	// 注册处理函数
	http.HandleFunc("/receive", c.receiveHandler)
	// 检查端口是否可用
	if isPortAvailable(c.Config.Port) {
		log.Printf("Port %d is available.", c.Config.Port)
	} else {
		log.Printf("Port %d is already in use.", c.Config.Port)
	}
	// 启动服务器
	log.Printf("Server starting on :%d...", c.Config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", c.Config.Port), nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

func (c *Controller) receiveHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 解析请求体
	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	// 检查token
	if requestData.Token != c.Config.Token {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid token"})
		return
	}
	// 处理接收到的数据
	err = c.CaptureProfileDataAndUpload(requestData.Types)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to upload profile data"})
		return
	}
	// 返回成功响应
	response := map[string]interface{}{
		"status":  "success",
		"message": "Data received successfully",
	}
	json.NewEncoder(w).Encode(response)
}

// 检查端口是否可用
func isPortAvailable(port int) bool {
	timeout := 1 * time.Second
	target := fmt.Sprintf("localhost:%d", port)

	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return true // 连接失败说明端口可用
	}
	defer conn.Close()
	return false
}
