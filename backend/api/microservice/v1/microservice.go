package v1

import (
	"backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type DeleteMicroserviceByIdsReq struct {
	g.Meta `path:"/deletemicroservicebyids" method:"post" tags:"Microservice" summary:"删除Microservice"`
	Ids    []int `json:"ids" v:"required#ID不能为空"`
}

type DeleteMicroserviceByIdsRes struct {
	Msg string `json:"msg"`
}

type CreateMicroserviceReq struct {
	g.Meta                  `path:"/createmicroservice" method:"post" tags:"Microservice" summary:"创建Microservice"`
	ProjectId               int    `json:"ProjectId"   `
	Ip                      string `json:"Ip"   `
	Port                    int    `json:"Port"   `
	MicroserviceName        string `json:"MicroserviceName"   `
	MicroserviceDescription string `json:"MicroserviceDescription" `
}

type CreateMicroserviceRes struct {
	Msg string `json:"msg"`
}

type UpdateMicroserviceReq struct {
	g.Meta                  `path:"/updatemicroservice" method:"post" tags:"Microservice" summary:"更新Microservice"`
	ProjectId               int    `json:"ProjectId"   `
	Ip                      string `json:"Ip"   `
	Port                    int    `json:"Port"   `
	MicroserviceId          int    `json:"MicroserviceId"   `
	MicroserviceName        string `json:"MicroserviceName"   `
	MicroserviceDescription string `json:"MicroserviceDescription" `
}

type UpdateMicroserviceRes struct {
	Msg string `json:"msg"`
}

type GetMicroserviceByPageUserReq struct {
	g.Meta                  `path:"/getmicroservicebypageuser" method:"post" tags:"Microservice" summary:"分页获取Microservice"`
	Page                    int    `json:"page" d:"1"`
	PageSize                int    `json:"pageSize" d:"10"`
	MicroserviceId          int    `json:"MicroserviceId"   `
	ProjectId               int    `json:"ProjectId"   `
	Ip                      string `json:"Ip"   `
	Port                    int    `json:"Port"   `
	MicroserviceName        string `json:"MicroserviceName"   `
	MicroserviceDescription string `json:"MicroserviceDescription" `
}

type GetMicroserviceByPageUserRes struct {
	Msg           string                `json:"msg"`
	Total         int                   `json:"total"   `
	Microservices []entity.Microservice `json:"microservices"   `
}
