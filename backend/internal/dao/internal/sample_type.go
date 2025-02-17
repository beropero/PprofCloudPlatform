// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SampleTypeDao is the data access object for table sample_type.
type SampleTypeDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns SampleTypeColumns // columns contains all the column names of Table for convenient usage.
}

// SampleTypeColumns defines and stores column names for table sample_type.
type SampleTypeColumns struct {
	Id        string //
	ProfileId string //
	Type      string //
	Unit      string //
}

// sampleTypeColumns holds the columns for table sample_type.
var sampleTypeColumns = SampleTypeColumns{
	Id:        "id",
	ProfileId: "profile_id",
	Type:      "type",
	Unit:      "unit",
}

// NewSampleTypeDao creates and returns a new DAO object for table data access.
func NewSampleTypeDao() *SampleTypeDao {
	return &SampleTypeDao{
		group:   "default",
		table:   "sample_type",
		columns: sampleTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SampleTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SampleTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SampleTypeDao) Columns() SampleTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SampleTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SampleTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SampleTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
