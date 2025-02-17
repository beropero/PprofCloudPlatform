package router

import (
	"backend/internal/dao"
	"backend/internal/model"
	"backend/internal/model/entity"
	"backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type sProject struct {
}

func init() {
	service.RegisterProject(&sProject{})
}

func (s *sProject) CreateProject(ctx context.Context, in model.CreateProjectInput) (err error) {
	var (
		user = service.BizCtx().Get(ctx).User
	)
	// 创建项目
	_, err = dao.Project.Ctx(ctx).Insert(entity.Project{
		CreatorId:          user.UserId,
		ProjectName:        in.ProjectName,
		ProjectDescription: in.ProjectDescription,
	})
	return
}

func (s *sProject) DeleteProjectByIds(ctx context.Context, in model.DeleteProjectByIdsInput) (err error) {
	// 删除项目
	_, err = dao.Project.Ctx(ctx).WhereIn("Project_id", in.ProjectIds).Delete()
	return
}

func (s *sProject) UpdateProject(ctx context.Context, in model.UpdateProjectInput) (err error) {
	// 更新项目
	_, err = dao.Project.Ctx(ctx).OmitEmpty().Where("Project_id", in.ProjectId).Update(entity.Project{
		ProjectName:        in.ProjectName,
		ProjectDescription: in.ProjectDescription,
	})
	return
}

func (s *sProject) GetProjectByPageUser(ctx context.Context, in model.GetProjectByPageUserInput) (out model.GetProjectByPageUserOutput, err error) {
	var (
		user = service.BizCtx().Get(ctx).User
	)
	// 获取项目列表
	if in.ProjectName != "" {
		in.ProjectName = "%" + in.ProjectName + "%"
	}
	if in.ProjectDescription != "" {
		in.ProjectDescription = "%" + in.ProjectDescription + "%"
	}
	err = dao.Project.Ctx(ctx).OmitEmpty().Where(g.Map{
		"Project_id":                 in.ProjectId,
		"Creator_id":                 user.UserId,
		"Project_name LIKE ?":        in.ProjectName,
		"Project_description LIKE ?": in.ProjectDescription,
	}).Page(in.Page, in.PageSize).ScanAndCount(&out.Projects, &out.Total, false)
	return
}
