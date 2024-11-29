package _interface

import "github.com/gofiber/fiber/v2"

type IMusicHandler interface {
	GetMusic(ctx *fiber.Ctx) error
}
