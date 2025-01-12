package runner

import (
	"github.com/rodericusifo/employee-management-api/internal/app/repository/database-seeder/sql/role"
	"github.com/rodericusifo/employee-management-api/internal/app/repository/database-seeder/sql/user"
	"github.com/rodericusifo/employee-management-api/internal/pkg/config"
	"github.com/rodericusifo/employee-management-api/internal/pkg/constant"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/getter"
)

func RunDatabaseSeederSQL(dialect constant.DialectDatabaseSQL) {
	switch dialect {
	case constant.POSTGRES:
	case constant.MYSQL:
		role.ExecuteMysqlRoleDatabaseSeederRepository(config.Env.DatabaseSeederMysqlRoleIsRebuildData, getter.GetMysqlDatabaseSQLConnection())
		user.ExecuteMysqlUserDatabaseSeederRepository(config.Env.DatabaseSeederMysqlUserIsRebuildData, getter.GetMysqlDatabaseSQLConnection())
	}
}
