package router

import (
	"backend/internal/dao"
	"backend/internal/model"
	"backend/internal/model/entity"
	"backend/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

type sMicroservice struct {
}

func init() {
	service.RegisterMicroservice(&sMicroservice{})
}

func (s *sMicroservice) CreateMicroservice(ctx context.Context, in model.CreateMicroserviceInput) (id int64, err error) {
	var (
		user = service.BizCtx().Get(ctx).User
	)
	// 添加微服务
	id, err = dao.Microservice.Ctx(ctx).OmitEmpty().InsertAndGetId(entity.Microservice{
		ProjectId:               in.ProjectId,
		Ip:                      in.Ip,
		Port:                    in.Port,
		MicroserviceName:        in.MicroserviceName,
		MicroserviceDescription: in.MicroserviceDescription,
		CreatorId:               user.UserId,
		MicroserviceToken:       "mis_" + guid.S(),
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
	// 检测权限
	count, err := dao.Project.Ctx(ctx).Where("project_id", in.ProjectId).Where("creator_id", user.UserId).Count()
	if count == 0 || err != nil {
		return out, gerror.New(fmt.Sprintf("无权限访问项目%d", in.ProjectId))
	}
	if in.MicroserviceName != "" {
		in.MicroserviceName = "%" + in.MicroserviceName + "%"
	}
	if in.MicroserviceDescription != "" {
		in.MicroserviceDescription = "%" + in.MicroserviceDescription + "%"
	}
	if in.Ip != "" {
		in.Ip = "%" + in.Ip + "%"
	}
	// 获取项目列表
	err = dao.Microservice.Ctx(ctx).OmitEmpty().Where(g.Map{
		"project_id":                    in.ProjectId,
		"microservice_id":               in.MicroserviceId,
		"ip like":                       in.Ip,
		"port":                          in.Port,
		"microservice_name like":        in.MicroserviceName,
		"microservice_description like": in.MicroserviceDescription,
		"creator_id":                    user.UserId,
	}).Page(in.Page, in.PageSize).ScanAndCount(&out.Microservices, &out.Total, false)
	return
}
