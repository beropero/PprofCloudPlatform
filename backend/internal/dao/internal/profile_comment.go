// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProfileCommentDao is the data access object for table profile_comment.
type ProfileCommentDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns ProfileCommentColumns // columns contains all the column names of Table for convenient usage.
}

// ProfileCommentColumns defines and stores column names for table profile_comment.
type ProfileCommentColumns struct {
	Id         string //
	ProfileId  string //
	Comment    string //
	OrderIndex string //
}

// profileCommentColumns holds the columns for table profile_comment.
var profileCommentColumns = ProfileCommentColumns{
	Id:         "id",
	ProfileId:  "profile_id",
	Comment:    "comment",
	OrderIndex: "order_index",
}

// NewProfileCommentDao creates and returns a new DAO object for table data access.
func NewProfileCommentDao() *ProfileCommentDao {
	return &ProfileCommentDao{
		group:   "default",
		table:   "profile_comment",
		columns: profileCommentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProfileCommentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProfileCommentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProfileCommentDao) Columns() ProfileCommentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProfileCommentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProfileCommentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProfileCommentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
