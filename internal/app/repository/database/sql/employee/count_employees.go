package employee

import (
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/builder"
)

func (r *EmployeeDatabaseSQLRepository) CountEmployees(query *types.QuerySQL) (int64, error) {
	count := int64(0)

	q := r.db

	defaultQuery := &types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
		},
	}

	if query != nil {
		query.Selects = append(query.Selects, defaultQuery.Selects...)
		q = builder.BuildQuerySQL(r.model.TableName(), q, query, r.dialect)
	} else {
		q = builder.BuildQuerySQL(r.model.TableName(), q, defaultQuery, r.dialect)
	}

	q = q.Table(r.model.TableName()).Count(&count)

	if err := q.Error; err != nil {
		return count, err
	}

	return count, nil
}
