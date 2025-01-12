package resource

import (
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/app/repository/database/sql/user"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

type IUserResource interface {
	SaveUser(payload *sql.User) error
	FirstUser(query *types.QuerySQL) (*sql.User, error)
}

type UserResource struct {
	UserDatabaseSQLRepository user.IUserDatabaseSQLRepository
}

func InitUserResource(userDatabaseSQLRepository user.IUserDatabaseSQLRepository) IUserResource {
	return &UserResource{
		UserDatabaseSQLRepository: userDatabaseSQLRepository,
	}
}
