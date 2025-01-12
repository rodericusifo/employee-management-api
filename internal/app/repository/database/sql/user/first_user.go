package user

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/builder"
)

func (r *UserDatabaseSQLRepository) FirstUser(query *types.QuerySQL) (*sql.User, error) {
	user := new(sql.User)

	q := r.db

	if query != nil {
		q = builder.BuildQuerySQL(r.model.TableName(), q, query, r.dialect)
	}

	if err := q.Table(r.model.TableName()).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
