package uploader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/beropero/PprofCloudPlatform/client/go/v1/config"
)

// 上传json数据到后台.
func UploadJsonData(data interface{}, url string) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	// 设置请求头
	req.Header.Set("User-Agent", "pprof-control-client")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		fmt.Printf("upload failed with status: %s", resp.Status)
		return fmt.Errorf("upload failed with status: %s", resp.Status)
	}
	return nil
}

// 上传form-data数据到后台.
func UploadFormData(payload *bytes.Buffer, writer *multipart.Writer, url string, config *config.Config) error {
	method := "POST"
	// 设置超时的HTTP客户端
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// 先完成writer的所有写入
	if err := writer.Close(); err != nil {
		return fmt.Errorf("关闭writer失败: %w", err)
	}

	// 创建新请求
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	req.Header.Set("User-Agent", "pprof-control-client")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Profile-Token", config.ServiceToken)
	// 发送请求
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求失败: %w", err)
	}
	defer res.Body.Close()

	// 读取响应
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("读取响应失败: %w", err)
	}
	fmt.Println(string(body))
	return nil
}
