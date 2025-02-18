package router

import (
	"backend/internal/dao"
	"backend/internal/model"
	"backend/internal/model/entity"
	"backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type sProfile struct {
}

func init() {
	service.RegisterProfile(&sProfile{})
}

func (s *sProfile) CreateProfile(ctx context.Context, in model.CreateProfileInput) (err error) {
	_, err = dao.Profile.Ctx(ctx).Insert(entity.Profile{
		Comment:        in.Comment,
		MicroserviceId: in.MicroserviceId,
		ProjectId:      in.ProjectId,
		Ptype:          in.Ptype,
	})
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
	err = dao.Profile.Ctx(ctx).OmitEmpty().Where(g.Map{
		"Project_id":      in.ProjectId,
		"Comment LIKE ?":  in.Comment,
		"Microservice_id": in.MicroserviceId,
		"Ptype":           in.Ptype,
		"Profile_id":      in.ProfileId,
	}).Page(in.Page, in.PageSize).ScanAndCount(&out.Profiles, &out.Total, false)
	return
}
