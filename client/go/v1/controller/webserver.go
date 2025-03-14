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
	ServiceToken string `json:"serviceToken"`
	Type         string `json:"type"`
	Comment      string `json:"comment"`
	Sec          int    `json:"sec"`
	Gc           bool   `json:"gc"`
}

func (c *Controller) WebServerStart() {
	// 注册处理函数
	http.HandleFunc("/runtest", enableCORS(c.receiveHandler))
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
	fmt.Println(r.Body)
	// 解析请求体
	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	fmt.Println(requestData)
	// 检查token
	if requestData.ServiceToken != c.Config.ServiceToken {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid token"})
		return
	}
	// 处理接收到的数据
	err = c.CaptureProfileDataAndUpload(r.Context(), ProfileInput{
		Type:    requestData.Type,
		Comment: requestData.Comment,
		Sec:     requestData.Sec,
		Gc:      requestData.Gc,
	})
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

// 添加 CORS 中间件
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") // 允许的请求头
		w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    // 允许携带 Cookie
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             // 允许的请求方法
		w.Header().Set("content-type", "application/json;charset=UTF-8")                                              // 返回数据格式
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}
