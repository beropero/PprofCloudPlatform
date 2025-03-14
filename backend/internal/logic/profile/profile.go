package router

import (
	"backend/internal/dao"
	"backend/internal/model"
	"backend/internal/model/entity"
	"backend/internal/service"
	"backend/utility"
	"bytes"
	"context"
	"os"
	"path/filepath"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	pprof "github.com/google/pprof/profile"
)

type sProfile struct {
}

func init() {
	service.RegisterProfile(&sProfile{})
}

func (s *sProfile) CreateProfile(ctx context.Context, in model.CreateProfileInput) (err error) {
	// 插入数据库
	var (
		microservice entity.Microservice
	)
	err = dao.Microservice.Ctx(ctx).OmitEmpty().Where(entity.Microservice{
		MicroserviceToken: in.ServiceToken,
	}).Scan(&microservice)
	if err != nil {

		return
	}
	g.Dump(microservice)

	_, err = dao.Profile.Ctx(ctx).Insert(entity.Profile{
		Comment:        in.Comment,
		OssPath:        in.OssPath,
		Ptype:          in.Ptype,
		MicroserviceId: microservice.MicroserviceId,
		ProjectId:      microservice.ProjectId,
	})
	if err != nil {
		return
	}
	return
}

func (s *sProfile) DeleteProfileByIds(ctx context.Context, in model.DeleteProfileByIdsInput) (err error) {
	_, err = dao.Profile.Ctx(ctx).WhereIn("microservice_id", in.ProfileIds).Delete()
	return
}

func (s *sProfile) UpdateProfile(ctx context.Context, in model.UpdateProfileInput) (err error) {
	_, err = dao.Profile.Ctx(ctx).OmitEmpty().Where("project_id", in.ProfileId).Update(entity.Profile{
		Comment: in.Comment,
	})
	return
}

func (s *sProfile) GetProfileByPageUser(ctx context.Context, in model.GetProfileByPageUserInput) (out model.GetProfileByPageUserOutput, err error) {
	if in.Comment != "" {
		in.Comment = "%" + in.Comment + "%"
	}
	if in.Ptype != "" {
		in.Ptype = "%" + in.Ptype + "%"
	}
	err = dao.Profile.Ctx(ctx).OmitEmpty().Where(g.Map{
		"Project_id":      in.ProjectId,
		"Comment LIKE ?":  in.Comment,
		"Microservice_id": in.MicroserviceId,
		"Ptype like ?":    in.Ptype,
		"id":              in.ProfileId,
	}).Page(in.Page, in.PageSize).ScanAndCount(&out.Profiles, &out.Total, false)
	return
}

// 获取文件内容
func (s *sProfile) GetProfileContent(ctx context.Context, in model.GetProfileContentInput) (out model.GetProfileContentOutput, err error) {
	var (
		profile entity.Profile
		user    = service.BizCtx().Get(ctx).User
	)
	err = dao.Profile.Ctx(ctx).Where("id", in.Id).Scan(&profile)

	if err != nil {
		return model.GetProfileContentOutput{}, err
	}
	count, err := dao.Microservice.Ctx(ctx).OmitEmpty().Where(
		entity.Microservice{
			CreatorId:      user.UserId,
			MicroserviceId: profile.MicroserviceId,
		},
	).Count()
	if err != nil {
		return model.GetProfileContentOutput{}, err
	}
	if count == 0 {
		return model.GetProfileContentOutput{}, gerror.New("没有权限")
	}
	// 拼装文件路径。
	tempDir := g.Cfg().MustGet(ctx, "server.tempDir").String()
	localFileName := filepath.Join(tempDir, profile.OssPath)
	utility.DownloadFileOss(ctx, profile.OssPath, localFileName)
	defer func() {
		// 删除临时文件
		os.Remove(localFileName)
	}()
	// 读取文件
	ost, err := os.ReadFile(localFileName)
	if err != nil {
		return model.GetProfileContentOutput{}, err
	}
	// 解析数据
	out.FileContent, err = pprof.Parse(bytes.NewReader(ost))
	if err != nil {
		return model.GetProfileContentOutput{}, err
	}
	return out, nil
}
