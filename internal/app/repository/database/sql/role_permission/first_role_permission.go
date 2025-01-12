package role_permission

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/builder"
)

func (r *RolePermissionDatabaseSQLRepository) FirstRolePermission(query *types.QuerySQL) (*sql.RolePermission, error) {
	rolePermission := new(sql.RolePermission)

	q := r.db

	if query != nil {
		q = builder.BuildQuerySQL(r.model.TableName(), q, query, r.dialect)
	}

	if err := q.Table(r.model.TableName()).First(rolePermission).Error; err != nil {
		return nil, err
	}

	return rolePermission, nil
}
