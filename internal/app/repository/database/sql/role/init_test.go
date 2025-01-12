package role

import (
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/rodericusifo/employee-management-api/internal/pkg/constant"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/mocker"
)

var (
	roleDatabaseSQLRepository IRoleDatabaseSQLRepository
	mockQuery                 sqlmock.Sqlmock
)

var (
	mockDateTime             time.Time
	mockUUID, mockDateString string
)

func SetupTestMysqlRoleDatabaseSQLRepository() {
	dialect := constant.MYSQL
	db, mock := mocker.MockDatabaseSQLConnection(dialect)

	roleDatabaseSQLRepository = InitMysqlRoleDatabaseSQLRepository(db)
	mockQuery = mock

	layoutFormat := constant.DEFAULT_TIME_LAYOUT.(string)

	mockDateString = "2015-09-02 08:04:00"
	mockDateTime, _ = time.Parse(layoutFormat, mockDateString)

	mockUUID = "ac0d6ce3-ff02-4024-896b-ea0ceba32182"
}
