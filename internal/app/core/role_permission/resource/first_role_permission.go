package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/employee-management-api/pkg/types"
)

func (r *RolePermissionResource) FirstRolePermission(query *pkg_types.QuerySQL) (*sql.RolePermission, error) {
	return r.RolePermissionDatabaseSQLRepository.FirstRolePermission(query)
}
