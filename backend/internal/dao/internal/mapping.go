// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MappingDao is the data access object for table mapping.
type MappingDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns MappingColumns // columns contains all the column names of Table for convenient usage.
}

// MappingColumns defines and stores column names for table mapping.
type MappingColumns struct {
	Id              string //
	ProfileId       string //
	Start           string //
	Limit           string //
	Offset          string //
	File            string //
	BuildId         string //
	HasFunctions    string //
	HasFilenames    string //
	HasLineNumbers  string //
	HasInlineFrames string //
}

// mappingColumns holds the columns for table mapping.
var mappingColumns = MappingColumns{
	Id:              "id",
	ProfileId:       "profile_id",
	Start:           "start",
	Limit:           "limit",
	Offset:          "offset",
	File:            "file",
	BuildId:         "build_id",
	HasFunctions:    "has_functions",
	HasFilenames:    "has_filenames",
	HasLineNumbers:  "has_line_numbers",
	HasInlineFrames: "has_inline_frames",
}

// NewMappingDao creates and returns a new DAO object for table data access.
func NewMappingDao() *MappingDao {
	return &MappingDao{
		group:   "default",
		table:   "mapping",
		columns: mappingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MappingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MappingDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MappingDao) Columns() MappingColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MappingDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MappingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MappingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
