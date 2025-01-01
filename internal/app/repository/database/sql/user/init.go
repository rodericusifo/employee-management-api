package user

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/config"

	pkg_constant "github.com/rodericusifo/employee-management-api/pkg/constant"
	pkg_types "github.com/rodericusifo/employee-management-api/pkg/types"
)

type IUserDatabaseSQLRepository interface {
	SaveUser(payload *sql.User) error
	FirstUser(query *pkg_types.QuerySQL) (*sql.User, error)
}

type UserDatabaseSQLRepository struct {
	db      *gorm.DB
	model   sql.User
	dialect pkg_constant.DialectDatabaseSQL
}

func InitMysqlUserDatabaseSQLRepository(db config.MysqlDatabaseSQLConnection) IUserDatabaseSQLRepository {
	return &UserDatabaseSQLRepository{
		db:      db,
		model:   sql.User{},
		dialect: pkg_constant.MYSQL,
	}
}
