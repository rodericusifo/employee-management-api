package user

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/constant"

	gorm_seeder "github.com/kachit/gorm-seeder"
)

type UserDatabaseSeederSQLRepository struct {
	gorm_seeder.SeederAbstract
	models struct {
		sql.User
		sql.Role
	}
	dialect constant.DialectDatabaseSQL
}

func InitMysqlUserDatabaseSeederSQLRepository(cfg gorm_seeder.SeederConfiguration) *UserDatabaseSeederSQLRepository {
	return &UserDatabaseSeederSQLRepository{gorm_seeder.NewSeederAbstract(cfg), struct {
		sql.User
		sql.Role
	}{}, constant.MYSQL}
}
