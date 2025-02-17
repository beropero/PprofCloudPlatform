// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FunctionDao is the data access object for table function.
type FunctionDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns FunctionColumns // columns contains all the column names of Table for convenient usage.
}

// FunctionColumns defines and stores column names for table function.
type FunctionColumns struct {
	Id         string //
	ProfileId  string //
	Name       string //
	SystemName string //
	Filename   string //
	StartLine  string //
}

// functionColumns holds the columns for table function.
var functionColumns = FunctionColumns{
	Id:         "id",
	ProfileId:  "profile_id",
	Name:       "name",
	SystemName: "system_name",
	Filename:   "filename",
	StartLine:  "start_line",
}

// NewFunctionDao creates and returns a new DAO object for table data access.
func NewFunctionDao() *FunctionDao {
	return &FunctionDao{
		group:   "default",
		table:   "function",
		columns: functionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FunctionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FunctionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FunctionDao) Columns() FunctionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FunctionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FunctionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FunctionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
