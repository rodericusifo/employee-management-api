package role

import (
	"github.com/sirupsen/logrus"

	"github.com/rodericusifo/employee-management-api/internal/pkg/config"

	gorm_seeder "github.com/kachit/gorm-seeder"
)

func ExecuteMysqlRoleDatabaseSeederRepository(isRebuildData config.IsRebuildDataDBSeederMysqlRole, db config.MysqlDatabaseSQLConnection) {
	mysqlRoleDatabaseSeederSQLRepository := InitMysqlRoleDatabaseSeederSQLRepository(gorm_seeder.SeederConfiguration{})
	seedersStack := gorm_seeder.NewSeedersStack(db)
	seedersStack.AddSeeder(mysqlRoleDatabaseSeederSQLRepository)

	if isRebuildData {
		err := seedersStack.Clear()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"message": "clear role fail",
				"detail":  err,
			}).Errorln("[EXECUTE MYSQL ROLE DATABASE SEEDER REPOSITORY]")
			return
		}
		logrus.WithFields(logrus.Fields{
			"message": "clear role success",
		}).Infoln("[EXECUTE MYSQL ROLE DATABASE SEEDER REPOSITORY]")
	}

	err := seedersStack.Seed()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"message": "seed role fail",
			"detail":  err,
		}).Errorln("[EXECUTE MYSQL ROLE DATABASE SEEDER REPOSITORY]")
		return
	}
	logrus.WithFields(logrus.Fields{
		"message": "seed role success",
	}).Infoln("[EXECUTE MYSQL ROLE DATABASE SEEDER REPOSITORY]")
}
