package user

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/builder"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/patcher"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/validator"
)

func (r *UserDatabaseSeederSQLRepository) Seed(db *gorm.DB) error {
	role := new(sql.Role)

	q := db

	query := &types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "slug", Operator: "=", Value: "super_admin"},
			},
		},
	}

	q = builder.BuildQuerySQL(r.models.Role.TableName(), q, query, r.dialect)

	err := q.Table(r.models.Role.TableName()).First(role).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.WithFields(logrus.Fields{
			"message": "get role fail",
			"detail":  err,
		}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [SEED]")
		return err
	}

	users := make([]*sql.User, 0)
	for _, UserSeed := range UserSeedData {
		err := validator.ValidatePayload(UserSeed)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"message": fmt.Sprintf("validation failed: user with xid %s", UserSeed.XID),
				"detail":  err,
			}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}

		user := new(sql.User)

		q := db

		query := &types.QuerySQL{
			Selects: []types.SelectQuerySQLOperation{
				{Field: "id"},
			},
			Searches: [][]types.SearchQuerySQLOperation{
				{
					{Field: "xid", Operator: "=", Value: UserSeed.XID},
					{Field: "email", Operator: "=", Value: UserSeed.Email},
				},
			},
		}

		q = builder.BuildQuerySQL(r.models.User.TableName(), q, query, r.dialect)

		err = q.Table(r.models.User.TableName()).First(user).Error

		if err != nil && err != gorm.ErrRecordNotFound {
			logrus.WithFields(logrus.Fields{
				"message": "get user fail",
				"detail":  err,
			}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}
		if user.ID != 0 {
			logrus.WithFields(logrus.Fields{
				"message": fmt.Sprintf("user with xid %s and email %s already registered", UserSeed.XID, UserSeed.Email),
			}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}

		hashedPassword, err := patcher.GenerateHashFromPassword(UserSeed.Password)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"message": fmt.Sprintf("hash password fail: user with xid %s", UserSeed.XID),
				"detail":  err,
			}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}

		users = append(users, &sql.User{
			XID:      UserSeed.XID,
			Name:     UserSeed.Name,
			Email:    UserSeed.Email,
			Password: hashedPassword,
			RoleID:   role.ID,
		})
	}
	return db.CreateInBatches(users, len(users)).Error
}
