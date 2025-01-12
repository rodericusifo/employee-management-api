package role

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/builder"
)

func (r *RoleDatabaseSQLRepository) FirstRole(query *types.QuerySQL) (*sql.Role, error) {
	role := new(sql.Role)

	q := r.db

	if query != nil {
		q = builder.BuildQuerySQL(r.model.TableName(), q, query, r.dialect)
	}

	if err := q.Table(r.model.TableName()).First(role).Error; err != nil {
		return nil, err
	}

	return role, nil
}
