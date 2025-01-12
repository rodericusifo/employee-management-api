package role

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/config"
	"github.com/rodericusifo/employee-management-api/internal/pkg/constant"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

type IRoleDatabaseSQLRepository interface {
	FirstRole(query *types.QuerySQL) (*sql.Role, error)
}

type RoleDatabaseSQLRepository struct {
	db      *gorm.DB
	model   sql.Role
	dialect constant.DialectDatabaseSQL
}

func InitMysqlRoleDatabaseSQLRepository(db config.MysqlDatabaseSQLConnection) IRoleDatabaseSQLRepository {
	return &RoleDatabaseSQLRepository{
		db:      db,
		model:   sql.Role{},
		dialect: constant.MYSQL,
	}
}
