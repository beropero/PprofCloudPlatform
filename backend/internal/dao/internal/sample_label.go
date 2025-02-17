// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SampleLabelDao is the data access object for table sample_label.
type SampleLabelDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns SampleLabelColumns // columns contains all the column names of Table for convenient usage.
}

// SampleLabelColumns defines and stores column names for table sample_label.
type SampleLabelColumns struct {
	Id       string //
	SampleId string //
	KeyStr   string //
	ValueStr string //
}

// sampleLabelColumns holds the columns for table sample_label.
var sampleLabelColumns = SampleLabelColumns{
	Id:       "id",
	SampleId: "sample_id",
	KeyStr:   "key_str",
	ValueStr: "value_str",
}

// NewSampleLabelDao creates and returns a new DAO object for table data access.
func NewSampleLabelDao() *SampleLabelDao {
	return &SampleLabelDao{
		group:   "default",
		table:   "sample_label",
		columns: sampleLabelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SampleLabelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SampleLabelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SampleLabelDao) Columns() SampleLabelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SampleLabelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SampleLabelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SampleLabelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
