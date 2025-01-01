package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/employee-management-api/pkg/types"
)

func (r *RoleResource) FirstRole(query *pkg_types.QuerySQL) (*sql.Role, error) {
	return r.RoleDatabaseSQLRepository.FirstRole(query)
}
