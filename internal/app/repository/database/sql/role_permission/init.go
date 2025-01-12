package role_permission

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/config"
	"github.com/rodericusifo/employee-management-api/internal/pkg/constant"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

type IRolePermissionDatabaseSQLRepository interface {
	FirstRolePermission(query *types.QuerySQL) (*sql.RolePermission, error)
}

type RolePermissionDatabaseSQLRepository struct {
	db      *gorm.DB
	model   sql.RolePermission
	dialect constant.DialectDatabaseSQL
}

func InitMysqlRolePermissionDatabaseSQLRepository(db config.MysqlDatabaseSQLConnection) IRolePermissionDatabaseSQLRepository {
	return &RolePermissionDatabaseSQLRepository{
		db:      db,
		model:   sql.RolePermission{},
		dialect: constant.MYSQL,
	}
}
