package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/employee-management-api/pkg/types"
)

func (r *PermissionResource) FirstPermission(query *pkg_types.QuerySQL) (*sql.Permission, error) {
	return r.PermissionDatabaseSQLRepository.FirstPermission(query)
}
