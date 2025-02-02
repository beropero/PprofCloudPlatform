package upload

import (
	"bytes"
	"context"
	"fmt"
	"io"

	v1 "backend/api/upload/v1"

	"github.com/google/pprof/profile"
)

func (c *ControllerV1) UploadTotal(ctx context.Context, req *v1.UploadTotalReq) (res *v1.UploadTotalRes, err error) {
	// 使用 github.com/google/pprof 包解析 Profile 数据
	memoryFile, _ := req.Memory.Open()
	// 读取文件内容
	memorydata, err := io.ReadAll(memoryFile)
	if err != nil {
		return nil, fmt.Errorf("读取文件内容失败: %w", err)
	}
	prof, err := profile.Parse(bytes.NewReader(memorydata))
	if err != nil {
		return nil, fmt.Errorf("failed to parse profile: %v", err)
	}
	// fmt.Print(req.Memory)
	fmt.Printf("prof: %v\n", prof)
	return nil, nil
}
