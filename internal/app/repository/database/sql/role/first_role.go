package role

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/employee-management-api/pkg/types"
	pkg_util_builder "github.com/rodericusifo/employee-management-api/pkg/util/builder"
)

func (r *RoleDatabaseSQLRepository) FirstRole(query *pkg_types.QuerySQL) (*sql.Role, error) {
	role := new(sql.Role)

	q := r.db

	if query != nil {
		q = pkg_util_builder.BuildQuerySQL(r.model.TableName(), q, query, r.dialect)
	}

	if err := q.Table(r.model.TableName()).First(role).Error; err != nil {
		return nil, err
	}

	return role, nil
}
