package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
)

func (r *UserResource) SaveUser(payload *sql.User) error {
	return r.UserDatabaseSQLRepository.SaveUser(payload)
}
