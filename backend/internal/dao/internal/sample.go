// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SampleDao is the data access object for table sample.
type SampleDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns SampleColumns // columns contains all the column names of Table for convenient usage.
}

// SampleColumns defines and stores column names for table sample.
type SampleColumns struct {
	Id        string //
	ProfileId string //
	Value     string //
}

// sampleColumns holds the columns for table sample.
var sampleColumns = SampleColumns{
	Id:        "id",
	ProfileId: "profile_id",
	Value:     "value",
}

// NewSampleDao creates and returns a new DAO object for table data access.
func NewSampleDao() *SampleDao {
	return &SampleDao{
		group:   "default",
		table:   "sample",
		columns: sampleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SampleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SampleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SampleDao) Columns() SampleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SampleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SampleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SampleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
