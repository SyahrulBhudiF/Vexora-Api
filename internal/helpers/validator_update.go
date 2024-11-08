package helpers

import (
	"github.com/gofiber/fiber/v2"
	"reflect"
)

func UpdateEntity(ctx *fiber.Ctx, request any, entity any) error {
	requestValue := reflect.ValueOf(request).Elem()
	entityValue := reflect.ValueOf(entity).Elem()

	for i := 0; i < requestValue.NumField(); i++ {
		field := requestValue.Field(i)
		if field.Kind() == reflect.String && field.String() != "" {
			entityField := entityValue.FieldByName(requestValue.Type().Field(i).Name)
			if entityField.IsValid() && entityField.CanSet() {
				entityField.SetString(field.String())
			}
		}
	}

	return nil
}
