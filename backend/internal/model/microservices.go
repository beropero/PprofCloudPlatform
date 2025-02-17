package model

import "backend/internal/model/entity"

type CreateMicroserviceInput struct {
	ProjectId               int    `json:"ProjectId"   `
	Ip                      string `json:"Ip"   `
	Port                    int    `json:"Port"   `
	MicroserviceName        string `json:"MicroserviceName"   `
	MicroserviceDescription string `json:"MicroserviceDescription" `
}

type UpdateMicroserviceInput struct {
	ProjectId               int    `json:"ProjectId"   `
	Ip                      string `json:"Ip"   `
	Port                    int    `json:"Port"   `
	MicroserviceId          int    `json:"MicroserviceId"   `
	MicroserviceName        string `json:"MicroserviceName"   `
	MicroserviceDescription string `json:"MicroserviceDescription" `
}

type DeleteMicroserviceByIdsInput struct {
	MicroserviceIds []int `json:"MicroserviceIds"   `
}

type GetMicroserviceByPageUserInput struct {
	Page                    int    `json:"page"   `
	PageSize                int    `json:"pageSize"   `
	MicroserviceId          int    `json:"MicroserviceId"   `
	ProjectId               int    `json:"ProjectId"   `
	Ip                      string `json:"Ip"   `
	Port                    int    `json:"Port"   `
	MicroserviceName        string `json:"MicroserviceName"   `
	MicroserviceDescription string `json:"MicroserviceDescription" `
}

type GetMicroserviceByPageUserOutput struct {
	Total         int                   `json:"total"   `
	Microservices []entity.Microservice `json:"Microservices"   `
}
