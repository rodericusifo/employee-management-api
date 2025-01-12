package role_permission

import (
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/getter"

	internal_app_core_role_permission_resource "github.com/rodericusifo/employee-management-api/internal/app/core/role_permission/resource"
	internal_app_repository_database_sql_role_permission "github.com/rodericusifo/employee-management-api/internal/app/repository/database/sql/role_permission"
)

func RolePermissionResource() internal_app_core_role_permission_resource.IRolePermissionResource {
	mysqlDatabaseSQLConnection := getter.GetMysqlDatabaseSQLConnection()
	iRolePermissionDatabaseSQLRepository := internal_app_repository_database_sql_role_permission.InitMysqlRolePermissionDatabaseSQLRepository(mysqlDatabaseSQLConnection)
	iRolePermissionResource := internal_app_core_role_permission_resource.InitRolePermissionResource(iRolePermissionDatabaseSQLRepository)
	return iRolePermissionResource
}
