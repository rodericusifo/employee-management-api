package validator

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

type IRequestQuery interface {
	CustomValidateRequestQuery() error
}

func ValidateRequestQuery(ctx *fiber.Ctx, req IRequestQuery) error {
	validator := types.InitValidator()

	if err := ctx.QueryParser(req); err != nil {
		logrus.WithFields(logrus.Fields{
			"message": "bind request query fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST QUERY]")
		return err
	}
	if err := validator.Validate(req); err != nil {
		logrus.WithFields(logrus.Fields{
			"message": "validate request query fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST QUERY]")
		return err
	}
	if err := req.CustomValidateRequestQuery(); err != nil {
		return err
	}
	return nil
}
