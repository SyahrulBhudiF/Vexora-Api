package _interface

import "github.com/gofiber/fiber/v2"

type IRoute interface {
	InitV1()
	initializeUserRoutes(router fiber.Router)
	initializeSpotifyRoutes(router fiber.Router)
	initializeHistoryRoutes(router fiber.Router)
	initializeMusicRoutes(router fiber.Router)
}
