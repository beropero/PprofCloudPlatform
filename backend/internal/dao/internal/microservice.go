// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MicroserviceDao is the data access object for table microservice.
type MicroserviceDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns MicroserviceColumns // columns contains all the column names of Table for convenient usage.
}

// MicroserviceColumns defines and stores column names for table microservice.
type MicroserviceColumns struct {
	MicroserviceId          string //
	ProjectId               string //
	CreatorId               string //
	MicroserviceName        string //
	Ip                      string //
	Port                    string //
	MicroserviceDescription string //
	CreatedAt               string //
	UpdatedAt               string //
	DeletedAt               string //
}

// microserviceColumns holds the columns for table microservice.
var microserviceColumns = MicroserviceColumns{
	MicroserviceId:          "microservice_id",
	ProjectId:               "project_id",
	CreatorId:               "creator_id",
	MicroserviceName:        "microservice_name",
	Ip:                      "ip",
	Port:                    "port",
	MicroserviceDescription: "microservice_description",
	CreatedAt:               "created_at",
	UpdatedAt:               "updated_at",
	DeletedAt:               "deleted_at",
}

// NewMicroserviceDao creates and returns a new DAO object for table data access.
func NewMicroserviceDao() *MicroserviceDao {
	return &MicroserviceDao{
		group:   "default",
		table:   "microservice",
		columns: microserviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MicroserviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MicroserviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MicroserviceDao) Columns() MicroserviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MicroserviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MicroserviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MicroserviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
