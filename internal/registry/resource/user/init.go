package user

import (
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/getter"

	internal_app_core_user_resource "github.com/rodericusifo/employee-management-api/internal/app/core/user/resource"
	internal_app_repository_database_sql_user "github.com/rodericusifo/employee-management-api/internal/app/repository/database/sql/user"
)

func UserResource() internal_app_core_user_resource.IUserResource {
	mysqlDatabaseSQLConnection := getter.GetMysqlDatabaseSQLConnection()
	iUserDatabaseSQLRepository := internal_app_repository_database_sql_user.InitMysqlUserDatabaseSQLRepository(mysqlDatabaseSQLConnection)
	iUserResource := internal_app_core_user_resource.InitUserResource(iUserDatabaseSQLRepository)
	return iUserResource
}
