package service

import (
	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/resource"
	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/service/dto/output"

	pkg_types "github.com/rodericusifo/employee-management-api/pkg/types"
)

type IEmployeeService interface {
	CreateEmployee(payload *input.CreateEmployeeDTO) error
	UpdateEmployee(payload *input.UpdateEmployeeDTO) error
	DeleteEmployee(payload *input.DeleteEmployeeDTO) error
	GetEmployees(payload *input.GetEmployeesDTO) (output.GetEmployeesDTO, *pkg_types.Meta, error)
	GetEmployee(payload *input.GetEmployeeDTO) (output.GetEmployeeDTO, error)
}

type EmployeeService struct {
	EmployeeResource resource.IEmployeeResource
}

func InitEmployeeService(employeeResource resource.IEmployeeResource) IEmployeeService {
	return &EmployeeService{
		EmployeeResource: employeeResource,
	}
}
