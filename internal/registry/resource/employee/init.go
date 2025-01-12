package employee

import (
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/getter"

	internal_app_core_employee_resource "github.com/rodericusifo/employee-management-api/internal/app/core/employee/resource"
	internal_app_repository_database_sql_employee "github.com/rodericusifo/employee-management-api/internal/app/repository/database/sql/employee"
)

func EmployeeResource() internal_app_core_employee_resource.IEmployeeResource {
	mysqlDatabaseSQLConnection := getter.GetMysqlDatabaseSQLConnection()
	iEmployeeDatabaseSQLRepository := internal_app_repository_database_sql_employee.InitMysqlEmployeeDatabaseSQLRepository(mysqlDatabaseSQLConnection)
	iEmployeeResource := internal_app_core_employee_resource.InitEmployeeResource(iEmployeeDatabaseSQLRepository)
	return iEmployeeResource
}
