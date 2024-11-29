package _interface

import "github.com/gofiber/fiber/v2"

type HistoryHandlerInterface interface {
	// GetRecommendations handles the request to get track recommendations
	GetRecommendations(ctx *fiber.Ctx) error

	// GetTrackByID handles the request to get a track by its ID
	GetTrackByID(ctx *fiber.Ctx) error

	// GetSearch handles the search request for tracks
	GetSearch(ctx *fiber.Ctx) error

	// MoodDetect handles the request to detect mood from an image
	MoodDetect(ctx *fiber.Ctx) error

	// GetHistory handles the request to get user's history
	GetHistory(ctx *fiber.Ctx) error
}
