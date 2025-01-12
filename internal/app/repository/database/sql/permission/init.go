package permission

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/config"
	"github.com/rodericusifo/employee-management-api/internal/pkg/constant"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

type IPermissionDatabaseSQLRepository interface {
	FirstPermission(query *types.QuerySQL) (*sql.Permission, error)
}

type PermissionDatabaseSQLRepository struct {
	db      *gorm.DB
	model   sql.Permission
	dialect constant.DialectDatabaseSQL
}

func InitMysqlPermissionDatabaseSQLRepository(db config.MysqlDatabaseSQLConnection) IPermissionDatabaseSQLRepository {
	return &PermissionDatabaseSQLRepository{
		db:      db,
		model:   sql.Permission{},
		dialect: constant.MYSQL,
	}
}
