package user

import (
	"github.com/sirupsen/logrus"

	"github.com/rodericusifo/employee-management-api/internal/pkg/config"

	gorm_seeder "github.com/kachit/gorm-seeder"
)

func ExecuteMysqlUserDatabaseSeederRepository(isRebuildData config.IsRebuildDataDBSeederMysqlUser, db config.MysqlDatabaseSQLConnection) {
	mysqlUserDatabaseSeederSQLRepository := InitMysqlUserDatabaseSeederSQLRepository(gorm_seeder.SeederConfiguration{})
	seedersStack := gorm_seeder.NewSeedersStack(db)
	seedersStack.AddSeeder(mysqlUserDatabaseSeederSQLRepository)

	if isRebuildData {
		err := seedersStack.Clear()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"message": "clear user fail",
				"detail":  err,
			}).Errorln("[EXECUTE MYSQL USER DATABASE SEEDER REPOSITORY]")
			return
		}
		logrus.WithFields(logrus.Fields{
			"message": "clear user success",
		}).Infoln("[EXECUTE MYSQL USER DATABASE SEEDER REPOSITORY]")
	}

	err := seedersStack.Seed()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"message": "seed user fail",
			"detail":  err,
		}).Errorln("[EXECUTE MYSQL USER DATABASE SEEDER REPOSITORY]")
		return
	}
	logrus.WithFields(logrus.Fields{
		"message": "seed user success",
	}).Infoln("[EXECUTE MYSQL USER DATABASE SEEDER REPOSITORY]")
}
