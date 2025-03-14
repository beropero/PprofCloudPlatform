package upload

import (
	"backend/internal/model"
	"backend/internal/service"
	"backend/utility"
	"context"
	"fmt"
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"
)

type sUpload struct {
}

func init() {
	service.RegisterUpload(&sUpload{})
}

// 保存文件
func saveTempFile(ctx context.Context, file *ghttp.UploadFile) (err error) {
	// 保存到临时文件夹
	// 获取临时文件夹
	tempDir := g.Cfg().MustGet(ctx, "server.tempDir").String()
	// 创建临时文件夹
	err = os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		return
	}
	// 保存文件
	_, err = file.Save(tempDir, false)
	return
}

// 保存文件到临时文件夹然后上传到oss
func saveFileToTempThenOss(ctx context.Context, file *ghttp.UploadFile, ftype string, token string, comment string) (err error) {
	if file == nil {
		return fmt.Errorf("block file is nil")
	}
	// 保存到临时文件夹
	file.Filename = ftype + "-" + guid.S() + ".pprof"
	err = saveTempFile(ctx, file)
	if err != nil {
		return
	}
	defer func() {
		// 删除临时文件
		tempDir := g.Cfg().MustGet(ctx, "server.tempDir").String()
		os.Remove(tempDir + "/" + file.Filename)
	}()
	// 创建profile
	err = service.Profile().CreateProfile(ctx, model.CreateProfileInput{
		Comment:      comment,
		OssPath:      file.Filename,
		Ptype:        ftype,
		ServiceToken: token,
	})
	if err != nil {
		return
	}
	// 上传到oss
	err = utility.UploadFileOss(ctx, file.Filename)
	if err != nil {
		return
	}

	return
}

// Upload 上传文件
func (s *sUpload) Upload(ctx context.Context, in model.UploadInput) (err error) {
	// 捕获指定类型
	err = saveFileToTempThenOss(ctx, in.File, in.Type, in.Token, in.Comment)
	if err != nil {
		return
	}
	return
}
