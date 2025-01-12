package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/app/repository/database/sql/role"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

type IRoleResource interface {
	FirstRole(query *types.QuerySQL) (*sql.Role, error)
}

type RoleResource struct {
	RoleDatabaseSQLRepository role.IRoleDatabaseSQLRepository
}

func InitRoleResource(roleDatabaseSQLRepository role.IRoleDatabaseSQLRepository) IRoleResource {
	return &RoleResource{
		RoleDatabaseSQLRepository: roleDatabaseSQLRepository,
	}
}
