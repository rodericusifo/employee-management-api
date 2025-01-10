package employee

import (
	internal_pkg_util_getter "github.com/rodericusifo/employee-management-api/internal/pkg/util/getter"

	internal_app_core_employee_resource "github.com/rodericusifo/employee-management-api/internal/app/core/employee/resource"
	internal_app_core_employee_service "github.com/rodericusifo/employee-management-api/internal/app/core/employee/service"
	internal_app_repository_database_sql_employee "github.com/rodericusifo/employee-management-api/internal/app/repository/database/sql/employee"
)

func EmployeeService() internal_app_core_employee_service.IEmployeeService {
	mysqlDatabaseSQLConnection := internal_pkg_util_getter.GetMysqlDatabaseSQLConnection()
	iEmployeeDatabaseSQLRepository := internal_app_repository_database_sql_employee.InitMysqlEmployeeDatabaseSQLRepository(mysqlDatabaseSQLConnection)
	iEmployeeResource := internal_app_core_employee_resource.InitEmployeeResource(iEmployeeDatabaseSQLRepository)
	iEmployeeService := internal_app_core_employee_service.InitEmployeeService(iEmployeeResource)
	return iEmployeeService
}
