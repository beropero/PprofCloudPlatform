package router

import (
	"backend/internal/dao"
	"backend/internal/model"
	"backend/internal/model/entity"
	"backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type sMicroservice struct {
}

func init() {
	service.RegisterMicroservice(&sMicroservice{})
}

func (s *sMicroservice) CreateMicroservice(ctx context.Context, in model.CreateMicroserviceInput) (err error) {
	// 添加微服务
	_, err = dao.Microservice.Ctx(ctx).Insert(entity.Microservice{
		ProjectId:               in.ProjectId,
		Ip:                      in.Ip,
		Port:                    in.Port,
		MicroserviceName:        in.MicroserviceName,
		MicroserviceDescription: in.MicroserviceDescription,
	})
	return
}

func (s *sMicroservice) DeleteMicroserviceByIds(ctx context.Context, in model.DeleteMicroserviceByIdsInput) (err error) {
	// 删除项目
	_, err = dao.Microservice.Ctx(ctx).WhereIn("microservice_id", in.MicroserviceIds).Delete()
	return
}

func (s *sMicroservice) UpdateMicroservice(ctx context.Context, in model.UpdateMicroserviceInput) (err error) {
	// 更新项目
	_, err = dao.Microservice.Ctx(ctx).OmitEmpty().Where("project_id", in.MicroserviceId).Update(entity.Microservice{
		ProjectId:               in.ProjectId,
		Ip:                      in.Ip,
		Port:                    in.Port,
		MicroserviceName:        in.MicroserviceName,
		MicroserviceDescription: in.MicroserviceDescription,
	})
	return
}

func (s *sMicroservice) GetMicroserviceByPageUser(ctx context.Context, in model.GetMicroserviceByPageUserInput) (out model.GetMicroserviceByPageUserOutput, err error) {
	var (
		user = service.BizCtx().Get(ctx).User
	)
	if in.MicroserviceName != "" {
		in.MicroserviceName = "%" + in.MicroserviceName + "%"
	}
	if in.MicroserviceDescription != "" {
		in.MicroserviceDescription = "%" + in.MicroserviceDescription + "%"
	}
	// 获取项目列表
	err = dao.Microservice.Ctx(ctx).OmitEmpty().Where(g.Map{
		"project_id":                    in.ProjectId,
		"microservice_id":               in.MicroserviceId,
		"ip":                            in.Ip,
		"port":                          in.Port,
		"microservice_name like":        in.MicroserviceName,
		"microservice_description like": in.MicroserviceDescription,
		"creator_id":                    user.UserId,
	}).Page(in.Page, in.PageSize).ScanAndCount(&out.Microservices, &out.Total, false)
	return
}
