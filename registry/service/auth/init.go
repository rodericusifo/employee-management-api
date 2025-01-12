package auth

import (
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/getter"

	internal_app_core_auth_service "github.com/rodericusifo/employee-management-api/internal/app/core/auth/service"
	internal_app_core_role_resource "github.com/rodericusifo/employee-management-api/internal/app/core/role/resource"
	internal_app_core_user_resource "github.com/rodericusifo/employee-management-api/internal/app/core/user/resource"
	internal_app_repository_database_sql_role "github.com/rodericusifo/employee-management-api/internal/app/repository/database/sql/role"
	internal_app_repository_database_sql_user "github.com/rodericusifo/employee-management-api/internal/app/repository/database/sql/user"
)

func AuthService() internal_app_core_auth_service.IAuthService {
	mysqlDatabaseSQLConnection := getter.GetMysqlDatabaseSQLConnection()
	iUserDatabaseSQLRepository := internal_app_repository_database_sql_user.InitMysqlUserDatabaseSQLRepository(mysqlDatabaseSQLConnection)
	iUserResource := internal_app_core_user_resource.InitUserResource(iUserDatabaseSQLRepository)
	iRoleDatabaseSQLRepository := internal_app_repository_database_sql_role.InitMysqlRoleDatabaseSQLRepository(mysqlDatabaseSQLConnection)
	iRoleResource := internal_app_core_role_resource.InitRoleResource(iRoleDatabaseSQLRepository)
	iAuthService := internal_app_core_auth_service.InitAuthService(iUserResource, iRoleResource)
	return iAuthService
}
