package _interface

import "github.com/gofiber/fiber/v2"

type PlaylistHandlerInterface interface {
	// GetRecommendations handles the request to get track recommendations
	GetRecommendations(ctx *fiber.Ctx) error

	// GetAvailableGenres handles the request to get available genre seeds
	GetAvailableGenres(ctx *fiber.Ctx) error

	// GetSearch handles the search request for tracks
	GetSearch(ctx *fiber.Ctx) error
}
