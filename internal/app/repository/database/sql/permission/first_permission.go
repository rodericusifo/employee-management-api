package permission

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/builder"
)

func (r *PermissionDatabaseSQLRepository) FirstPermission(query *types.QuerySQL) (*sql.Permission, error) {
	permission := new(sql.Permission)

	q := r.db

	if query != nil {
		q = builder.BuildQuerySQL(r.model.TableName(), q, query, r.dialect)
	}

	if err := q.Table(r.model.TableName()).First(permission).Error; err != nil {
		return nil, err
	}

	return permission, nil
}
