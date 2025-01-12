package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/app/repository/database/sql/permission"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

type IPermissionResource interface {
	FirstPermission(query *types.QuerySQL) (*sql.Permission, error)
}

type PermissionResource struct {
	PermissionDatabaseSQLRepository permission.IPermissionDatabaseSQLRepository
}

func InitPermissionResource(permissionDatabaseSQLRepository permission.IPermissionDatabaseSQLRepository) IPermissionResource {
	return &PermissionResource{
		PermissionDatabaseSQLRepository: permissionDatabaseSQLRepository,
	}
}
