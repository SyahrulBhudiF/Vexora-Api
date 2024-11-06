package middleware

import (
	"fmt"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"reflect"
)

var Validator = validator.New()

func EnsureJsonValidRequest[T any](ctx *fiber.Ctx) error {
	body := new(T)
	err := ctx.BodyParser(&body)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.ErrBadRequest.Code, false, err)
	}

	validationErr := Validator.Struct(body)

	if validationErr != nil {
		var errStr string
		for i, e := range validationErr.(validator.ValidationErrors) {
			if i > 0 {
				errStr += ", "
			}

			fieldName := e.Field()
			field, _ := reflect.TypeOf(body).Elem().FieldByName(fieldName)
			fieldJSONName, _ := field.Tag.Lookup("json")
			errStr += fieldJSONName + " " + e.Tag()
		}

		return helpers.ErrorResponse(ctx, fiber.ErrBadRequest.Code, false, fmt.Errorf("invalid request: %s", errStr))
	}

	ctx.Locals("body", body)

	return ctx.Next()
}
