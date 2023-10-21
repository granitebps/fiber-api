package core

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/fiber-api/pkg/utils"
)

type AppValidator struct {
	Validator *validator.Validate
}

type ErrorField struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func SetupValidator() *AppValidator {
	v := validator.New()

	// You can add your custom validator in here
	// v.RegisterValidation()

	return &AppValidator{
		Validator: v,
	}
}

func (v *AppValidator) Validate(c *fiber.Ctx, payload any) (fields []ErrorField, err error) {
	err = c.BodyParser(payload)
	if err != nil {
		return
	}

	err = v.Validator.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var elem ErrorField
			field := utils.CamelToSnake(err.Field())
			elem.Field = field
			elem.Message = validationRuleMessage(field, err.Tag(), err.Error())

			fields = append(fields, elem)
		}

		err = errors.New(fields[0].Message)
		return
	}

	return
}

func validationRuleMessage(field string, rule string, def string) string {
	switch rule {
	case "required":
		return fmt.Sprintf("The %s field is required.", field)
	default:
		return def
	}
}
