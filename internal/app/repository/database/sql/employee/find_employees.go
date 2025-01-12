package employee

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/builder"
)

func (r *EmployeeDatabaseSQLRepository) FindEmployees(query *types.QuerySQL) ([]*sql.Employee, error) {
	employees := make([]*sql.Employee, 0)

	q := r.db

	if query != nil {
		q = builder.BuildQuerySQL(r.model.TableName(), q, query, r.dialect)
	}

	q = q.Table(r.model.TableName()).Find(&employees)

	if err := q.Error; err != nil {
		return nil, err
	}

	return employees, nil
}
