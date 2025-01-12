package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

func (r *RoleResource) FirstRole(query *types.QuerySQL) (*sql.Role, error) {
	return r.RoleDatabaseSQLRepository.FirstRole(query)
}
