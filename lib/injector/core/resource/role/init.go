package role

import (
	internal_app_core_role_resource "github.com/rodericusifo/employee-management-api/internal/app/core/role/resource"
	internal_app_repository_database_sql_role "github.com/rodericusifo/employee-management-api/internal/app/repository/database/sql/role"
	internal_pkg_util_getter "github.com/rodericusifo/employee-management-api/internal/pkg/util/getter"
)

func RoleResource() internal_app_core_role_resource.IRoleResource {
	mysqlDatabaseSQLConnection := internal_pkg_util_getter.GetMysqlDatabaseSQLConnection()
	iRoleDatabaseSQLRepository := internal_app_repository_database_sql_role.InitMysqlRoleDatabaseSQLRepository(mysqlDatabaseSQLConnection)
	iRoleResource := internal_app_core_role_resource.InitRoleResource(iRoleDatabaseSQLRepository)
	return iRoleResource
}
