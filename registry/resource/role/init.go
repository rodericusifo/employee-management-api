package role

import (
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/getter"

	internal_app_core_role_resource "github.com/rodericusifo/employee-management-api/internal/app/core/role/resource"
	internal_app_repository_database_sql_role "github.com/rodericusifo/employee-management-api/internal/app/repository/database/sql/role"
)

func RoleResource() internal_app_core_role_resource.IRoleResource {
	mysqlDatabaseSQLConnection := getter.GetMysqlDatabaseSQLConnection()
	iRoleDatabaseSQLRepository := internal_app_repository_database_sql_role.InitMysqlRoleDatabaseSQLRepository(mysqlDatabaseSQLConnection)
	iRoleResource := internal_app_core_role_resource.InitRoleResource(iRoleDatabaseSQLRepository)
	return iRoleResource
}
