package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

func (r *PermissionResource) FirstPermission(query *types.QuerySQL) (*sql.Permission, error) {
	return r.PermissionDatabaseSQLRepository.FirstPermission(query)
}
