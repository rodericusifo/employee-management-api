package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

func (r *EmployeeResource) CountEmployees(query *types.QuerySQL) (int64, error) {
	return r.EmployeeDatabaseSQLRepository.CountEmployees(query)
}
