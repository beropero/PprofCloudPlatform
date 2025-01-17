// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TestRecordsDao is the data access object for table test_records.
type TestRecordsDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns TestRecordsColumns // columns contains all the column names of Table for convenient usage.
}

// TestRecordsColumns defines and stores column names for table test_records.
type TestRecordsColumns struct {
	RecordId       string //
	CpuUsage       string //
	MemoryUsage    string //
	DiskUsage      string //
	NetworkUsage   string //
	MicroserviceId string //
	CreatedAt      string //
	UpdatedAt      string //
	DeletedAt      string //
}

// testRecordsColumns holds the columns for table test_records.
var testRecordsColumns = TestRecordsColumns{
	RecordId:       "record_id",
	CpuUsage:       "cpu_usage",
	MemoryUsage:    "memory_usage",
	DiskUsage:      "disk_usage",
	NetworkUsage:   "network_usage",
	MicroserviceId: "microservice_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewTestRecordsDao creates and returns a new DAO object for table data access.
func NewTestRecordsDao() *TestRecordsDao {
	return &TestRecordsDao{
		group:   "default",
		table:   "test_records",
		columns: testRecordsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TestRecordsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TestRecordsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TestRecordsDao) Columns() TestRecordsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TestRecordsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TestRecordsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TestRecordsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
