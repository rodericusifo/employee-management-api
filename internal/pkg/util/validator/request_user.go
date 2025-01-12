package validator

import (
	"github.com/sirupsen/logrus"

	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

type IRequestUser interface {
	CustomValidateRequestUser() error
}

func ValidateRequestUser(req IRequestUser) error {
	validator := types.InitValidator()

	if err := validator.Validate(req); err != nil {
		logrus.WithFields(logrus.Fields{
			"message": "validate request user fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST USER]")
		return err
	}
	if err := req.CustomValidateRequestUser(); err != nil {
		return err
	}
	return nil
}
