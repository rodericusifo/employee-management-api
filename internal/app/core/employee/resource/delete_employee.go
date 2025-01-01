package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
)

func (r *EmployeeResource) DeleteEmployee(payload *sql.Employee) error {
	return r.EmployeeDatabaseSQLRepository.DeleteEmployee(payload)
}
