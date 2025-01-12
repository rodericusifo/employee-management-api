package permission

import (
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/getter"

	internal_app_core_permission_resource "github.com/rodericusifo/employee-management-api/internal/app/core/permission/resource"
	internal_app_repository_database_sql_permission "github.com/rodericusifo/employee-management-api/internal/app/repository/database/sql/permission"
)

func PermissionResource() internal_app_core_permission_resource.IPermissionResource {
	mysqlDatabaseSQLConnection := getter.GetMysqlDatabaseSQLConnection()
	iPermissionDatabaseSQLRepository := internal_app_repository_database_sql_permission.InitMysqlPermissionDatabaseSQLRepository(mysqlDatabaseSQLConnection)
	iPermissionResource := internal_app_core_permission_resource.InitPermissionResource(iPermissionDatabaseSQLRepository)
	return iPermissionResource
}
