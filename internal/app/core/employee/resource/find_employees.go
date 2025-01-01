package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/employee-management-api/pkg/types"
)

func (r *EmployeeResource) FindEmployees(query *pkg_types.QuerySQL) ([]*sql.Employee, error) {
	return r.EmployeeDatabaseSQLRepository.FindEmployees(query)
}
