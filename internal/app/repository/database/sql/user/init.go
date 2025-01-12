package user

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/config"
	"github.com/rodericusifo/employee-management-api/internal/pkg/constant"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

type IUserDatabaseSQLRepository interface {
	SaveUser(payload *sql.User) error
	FirstUser(query *types.QuerySQL) (*sql.User, error)
}

type UserDatabaseSQLRepository struct {
	db      *gorm.DB
	model   sql.User
	dialect constant.DialectDatabaseSQL
}

func InitMysqlUserDatabaseSQLRepository(db config.MysqlDatabaseSQLConnection) IUserDatabaseSQLRepository {
	return &UserDatabaseSQLRepository{
		db:      db,
		model:   sql.User{},
		dialect: constant.MYSQL,
	}
}
