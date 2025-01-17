package model

import "github.com/gogf/gf/v2/net/ghttp"

type Context struct {
	Session *ghttp.Session // Session in context.
	User    *ContextUser   // User in context.
	// Admin   *ContextAdmin  //Admin in context.
}

type ContextUser struct {
	UserId         int    // User ID.
	UserEmail      string // 用户邮箱
	UserPermission string // 用户权限
}

// type ContextAdmin struct {
// 	UserId        int    // User ID.
// 	UserName      string // User Name.
// 	UserPhone     string // 用户手机号
// 	UserRole      string // 用户角色
// 	UserSno       string // 工号或学号
// 	UserEmail     string // 用户邮箱
// 	UserMajorName string // 用户专业
// 	UserMajorCode string //用户专业编号
// }
