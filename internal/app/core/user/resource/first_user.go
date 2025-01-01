package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/employee-management-api/pkg/types"
)

func (r *UserResource) FirstUser(query *pkg_types.QuerySQL) (*sql.User, error) {
	return r.UserDatabaseSQLRepository.FirstUser(query)
}
