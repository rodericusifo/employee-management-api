package role_permission

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/config"

	pkg_constant "github.com/rodericusifo/employee-management-api/pkg/constant"
	pkg_types "github.com/rodericusifo/employee-management-api/pkg/types"
)

type IRolePermissionDatabaseSQLRepository interface {
	FirstRolePermission(query *pkg_types.QuerySQL) (*sql.RolePermission, error)
}

type RolePermissionDatabaseSQLRepository struct {
	db      *gorm.DB
	model   sql.RolePermission
	dialect pkg_constant.DialectDatabaseSQL
}

func InitMysqlRolePermissionDatabaseSQLRepository(db config.MysqlDatabaseSQLConnection) IRolePermissionDatabaseSQLRepository {
	return &RolePermissionDatabaseSQLRepository{
		db:      db,
		model:   sql.RolePermission{},
		dialect: pkg_constant.MYSQL,
	}
}
