package employee

import (
	internal_app_core_employee_service "github.com/rodericusifo/employee-management-api/internal/app/core/employee/service"
	internal_registry_resource_employee "github.com/rodericusifo/employee-management-api/internal/registry/resource/employee"
)

func EmployeeService() internal_app_core_employee_service.IEmployeeService {
	iEmployeeResource := internal_registry_resource_employee.EmployeeResource()
	iEmployeeService := internal_app_core_employee_service.InitEmployeeService(iEmployeeResource)
	return iEmployeeService
}
