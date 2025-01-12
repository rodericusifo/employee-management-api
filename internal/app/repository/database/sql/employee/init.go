package employee

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/config"
	"github.com/rodericusifo/employee-management-api/internal/pkg/constant"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

type IEmployeeDatabaseSQLRepository interface {
	SaveEmployee(payload *sql.Employee) error
	DeleteEmployee(payload *sql.Employee) error
	FindEmployees(query *types.QuerySQL) ([]*sql.Employee, error)
	FirstEmployee(query *types.QuerySQL) (*sql.Employee, error)
	CountEmployees(query *types.QuerySQL) (int64, error)
}

type EmployeeDatabaseSQLRepository struct {
	db      *gorm.DB
	model   sql.Employee
	dialect constant.DialectDatabaseSQL
}

func InitMysqlEmployeeDatabaseSQLRepository(db config.MysqlDatabaseSQLConnection) IEmployeeDatabaseSQLRepository {
	return &EmployeeDatabaseSQLRepository{
		db:      db,
		model:   sql.Employee{},
		dialect: constant.MYSQL,
	}
}
