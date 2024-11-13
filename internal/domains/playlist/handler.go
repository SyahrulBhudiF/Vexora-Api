package playlist

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/playlist/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/helpers"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/services"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/gofiber/fiber/v2"
)

type PlaylistHandler struct {
	service *services.SpotifyService
}

func NewPlaylistHandler(service *services.SpotifyService) *PlaylistHandler {
	return &PlaylistHandler{service: service}
}

func (p *PlaylistHandler) GetRecommendations(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit", 10)

	recommendations, err := p.service.GetRandomRecommendations(limit)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	}

	return ctx.JSON(types.WebResponse[entity.SpotifyResponse]{Message: "success", Success: true, ShouldNotify: false, Data: *recommendations})
}

func (p *PlaylistHandler) GetAvailableGenres(ctx *fiber.Ctx) error {
	genres, err := p.service.GetGenreSeeds()
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"data":    genres,
	})
}
