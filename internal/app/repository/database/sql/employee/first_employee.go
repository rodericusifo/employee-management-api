package employee

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/builder"
)

func (r *EmployeeDatabaseSQLRepository) FirstEmployee(query *types.QuerySQL) (*sql.Employee, error) {
	employee := new(sql.Employee)

	q := r.db

	if query != nil {
		q = builder.BuildQuerySQL(r.model.TableName(), q, query, r.dialect)
	}

	if err := q.Table(r.model.TableName()).First(employee).Error; err != nil {
		return nil, err
	}

	return employee, nil
}
