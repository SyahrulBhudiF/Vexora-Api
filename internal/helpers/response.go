package helpers

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/gofiber/fiber/v2"
)

func SuccessResponse[T any](ctx *fiber.Ctx, code int, shouldNotify bool, msg string, data T) error {
	return ctx.Status(code).JSON(types.WebResponse[T]{
		Success:      true,
		ShouldNotify: shouldNotify,
		Message:      msg,
		Data:         data,
	})
}

func ErrorResponse(ctx *fiber.Ctx, code int, shouldNotify bool, err error) error {
	return ctx.Status(code).JSON(types.WebResponse[error]{
		Success:      false,
		ShouldNotify: shouldNotify,
		Message:      err.Error(),
	})
}
