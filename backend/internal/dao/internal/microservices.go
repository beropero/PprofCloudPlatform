// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MicroservicesDao is the data access object for table microservices.
type MicroservicesDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns MicroservicesColumns // columns contains all the column names of Table for convenient usage.
}

// MicroservicesColumns defines and stores column names for table microservices.
type MicroservicesColumns struct {
	MicroserviceId          string //
	ProjectId               string //
	MicroserviceName        string //
	Ip                      string //
	Port                    string //
	MicroserviceDescription string //
	CreatedAt               string //
	UpdatedAt               string //
	DeletedAt               string //
}

// microservicesColumns holds the columns for table microservices.
var microservicesColumns = MicroservicesColumns{
	MicroserviceId:          "microservice_id",
	ProjectId:               "project_id",
	MicroserviceName:        "microservice_name",
	Ip:                      "ip",
	Port:                    "port",
	MicroserviceDescription: "microservice_description",
	CreatedAt:               "created_at",
	UpdatedAt:               "updated_at",
	DeletedAt:               "deleted_at",
}

// NewMicroservicesDao creates and returns a new DAO object for table data access.
func NewMicroservicesDao() *MicroservicesDao {
	return &MicroservicesDao{
		group:   "default",
		table:   "microservices",
		columns: microservicesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MicroservicesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MicroservicesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MicroservicesDao) Columns() MicroservicesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MicroservicesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MicroservicesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MicroservicesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
