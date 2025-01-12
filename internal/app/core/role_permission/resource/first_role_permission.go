package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

func (r *RolePermissionResource) FirstRolePermission(query *types.QuerySQL) (*sql.RolePermission, error) {
	return r.RolePermissionDatabaseSQLRepository.FirstRolePermission(query)
}
