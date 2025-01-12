package role

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/constant"

	gorm_seeder "github.com/kachit/gorm-seeder"
)

type RoleDatabaseSeederSQLRepository struct {
	gorm_seeder.SeederAbstract
	model   sql.Role
	dialect constant.DialectDatabaseSQL
}

func InitMysqlRoleDatabaseSeederSQLRepository(cfg gorm_seeder.SeederConfiguration) *RoleDatabaseSeederSQLRepository {
	return &RoleDatabaseSeederSQLRepository{gorm_seeder.NewSeederAbstract(cfg), sql.Role{}, constant.MYSQL}
}
