package validator

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

type IRequestBody interface {
	CustomValidateRequestBody() error
}

func ValidateRequestBody(ctx *fiber.Ctx, req IRequestBody) error {
	validator := types.InitValidator()

	if err := ctx.BodyParser(req); err != nil {
		logrus.WithFields(logrus.Fields{
			"message": "bind request body fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST BODY]")
		return err
	}
	if err := validator.Validate(req); err != nil {
		logrus.WithFields(logrus.Fields{
			"message": "validate request body fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST BODY]")
		return err
	}
	if err := req.CustomValidateRequestBody(); err != nil {
		return err
	}
	return nil
}
