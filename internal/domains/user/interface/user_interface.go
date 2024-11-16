package _interface

import "github.com/gofiber/fiber/v2"

type UserHandlerInterface interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	GetProfile(ctx *fiber.Ctx) error
	UpdateProfile(ctx *fiber.Ctx) error
	UploadProfilePicture(ctx *fiber.Ctx) error
	ChangePassword(ctx *fiber.Ctx) error
	RefreshToken(ctx *fiber.Ctx) error
}
