package model

import "github.com/gogf/gf/v2/frame/g"

type CheckEmailCodeInput struct {
	Email string `json:"email" ` // 邮箱
	Code  string `json:"code"`
}

type SaveEmailCodeInput struct {
	Email string `json:"email" ` // 邮箱
	Code  string `json:"code"`
}

type SaveJwtTokenInput struct {
	Token  string `json:"token" `
	UserId string `json:"userId"`
	// UserRole string `json:"userRole"`
}

type SaveLikeRecordInput struct {
	LikeId string `json:"like_id"`
	UserId int    `json:"user_id"`
	Type   string `json:"type"`
}

type CheckLikeRecordInput struct {
	LikeId string `json:"like_id"`
	UserId int    `json:"user_id"`
	Type   string `json:"type"`
}

type SaveStudentInput struct {
	Code           string `json:"code"`
	StudentSnoList g.Array
}
