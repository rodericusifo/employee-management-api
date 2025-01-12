package validator

import (
	"github.com/sirupsen/logrus"

	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

type IPayload interface {
	CustomValidatePayload() error
}

func ValidatePayload(payload IPayload) error {
	validator := types.InitValidator()

	if err := validator.Validate(payload); err != nil {
		logrus.WithFields(logrus.Fields{
			"message": "validate payload fail",
			"detail":  err,
		}).Errorln("[VALIDATE PAYLOAD]")
		return err
	}
	if err := payload.CustomValidatePayload(); err != nil {
		return err
	}
	return nil
}
