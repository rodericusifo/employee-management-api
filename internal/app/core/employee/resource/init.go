package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/app/repository/database/sql/employee"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

type IEmployeeResource interface {
	SaveEmployee(payload *sql.Employee) error
	DeleteEmployee(payload *sql.Employee) error
	FindEmployees(query *types.QuerySQL) ([]*sql.Employee, error)
	FirstEmployee(query *types.QuerySQL) (*sql.Employee, error)
	CountEmployees(query *types.QuerySQL) (int64, error)
}

type EmployeeResource struct {
	EmployeeDatabaseSQLRepository employee.IEmployeeDatabaseSQLRepository
}

func InitEmployeeResource(employeeDatabaseSQLRepository employee.IEmployeeDatabaseSQLRepository) IEmployeeResource {
	return &EmployeeResource{
		EmployeeDatabaseSQLRepository: employeeDatabaseSQLRepository,
	}
}
