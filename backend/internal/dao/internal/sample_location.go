// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SampleLocationDao is the data access object for table sample_location.
type SampleLocationDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SampleLocationColumns // columns contains all the column names of Table for convenient usage.
}

// SampleLocationColumns defines and stores column names for table sample_location.
type SampleLocationColumns struct {
	SampleId   string //
	LocationId string //
	OrderIndex string //
}

// sampleLocationColumns holds the columns for table sample_location.
var sampleLocationColumns = SampleLocationColumns{
	SampleId:   "sample_id",
	LocationId: "location_id",
	OrderIndex: "order_index",
}

// NewSampleLocationDao creates and returns a new DAO object for table data access.
func NewSampleLocationDao() *SampleLocationDao {
	return &SampleLocationDao{
		group:   "default",
		table:   "sample_location",
		columns: sampleLocationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SampleLocationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SampleLocationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SampleLocationDao) Columns() SampleLocationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SampleLocationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SampleLocationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SampleLocationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
