// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SampleNumlabelDao is the data access object for table sample_numlabel.
type SampleNumlabelDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SampleNumlabelColumns // columns contains all the column names of Table for convenient usage.
}

// SampleNumlabelColumns defines and stores column names for table sample_numlabel.
type SampleNumlabelColumns struct {
	Id         string //
	SampleId   string //
	KeyStr     string //
	ValueIndex string //
	NumValue   string //
	UnitStr    string //
}

// sampleNumlabelColumns holds the columns for table sample_numlabel.
var sampleNumlabelColumns = SampleNumlabelColumns{
	Id:         "id",
	SampleId:   "sample_id",
	KeyStr:     "key_str",
	ValueIndex: "value_index",
	NumValue:   "num_value",
	UnitStr:    "unit_str",
}

// NewSampleNumlabelDao creates and returns a new DAO object for table data access.
func NewSampleNumlabelDao() *SampleNumlabelDao {
	return &SampleNumlabelDao{
		group:   "default",
		table:   "sample_numlabel",
		columns: sampleNumlabelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SampleNumlabelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SampleNumlabelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SampleNumlabelDao) Columns() SampleNumlabelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SampleNumlabelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SampleNumlabelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SampleNumlabelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
