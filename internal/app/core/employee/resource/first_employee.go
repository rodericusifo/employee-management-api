package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

func (r *EmployeeResource) FirstEmployee(query *types.QuerySQL) (*sql.Employee, error) {
	return r.EmployeeDatabaseSQLRepository.FirstEmployee(query)
}
