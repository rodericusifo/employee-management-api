package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/app/repository/database/sql/role_permission"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

type IRolePermissionResource interface {
	FirstRolePermission(query *types.QuerySQL) (*sql.RolePermission, error)
}

type RolePermissionResource struct {
	RolePermissionDatabaseSQLRepository role_permission.IRolePermissionDatabaseSQLRepository
}

func InitRolePermissionResource(rolePermissionDatabaseSQLRepository role_permission.IRolePermissionDatabaseSQLRepository) IRolePermissionResource {
	return &RolePermissionResource{
		RolePermissionDatabaseSQLRepository: rolePermissionDatabaseSQLRepository,
	}
}
