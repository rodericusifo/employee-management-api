package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

func (r *UserResource) FirstUser(query *types.QuerySQL) (*sql.User, error) {
	return r.UserDatabaseSQLRepository.FirstUser(query)
}
