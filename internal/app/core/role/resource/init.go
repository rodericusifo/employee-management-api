package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/app/repository/database/sql/role"

	pkg_types "github.com/rodericusifo/employee-management-api/pkg/types"
)

type IRoleResource interface {
	FirstRole(query *pkg_types.QuerySQL) (*sql.Role, error)
}

type RoleResource struct {
	RoleDatabaseSQLRepository role.IRoleDatabaseSQLRepository
}

func InitRoleResource(roleDatabaseSQLRepository role.IRoleDatabaseSQLRepository) IRoleResource {
	return &RoleResource{
		RoleDatabaseSQLRepository: roleDatabaseSQLRepository,
	}
}
