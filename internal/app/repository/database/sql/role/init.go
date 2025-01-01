package role

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/config"

	pkg_constant "github.com/rodericusifo/employee-management-api/pkg/constant"
	pkg_types "github.com/rodericusifo/employee-management-api/pkg/types"
)

type IRoleDatabaseSQLRepository interface {
	FirstRole(query *pkg_types.QuerySQL) (*sql.Role, error)
}

type RoleDatabaseSQLRepository struct {
	db      *gorm.DB
	model   sql.Role
	dialect pkg_constant.DialectDatabaseSQL
}

func InitMysqlRoleDatabaseSQLRepository(db config.MysqlDatabaseSQLConnection) IRoleDatabaseSQLRepository {
	return &RoleDatabaseSQLRepository{
		db:      db,
		model:   sql.Role{},
		dialect: pkg_constant.MYSQL,
	}
}
