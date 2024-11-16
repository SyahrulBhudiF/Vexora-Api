package playlist

import (
	"errors"
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
	recommendations, err := p.service.GetRecommendations(10, nil)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	}

	return ctx.JSON(types.WebResponse[entity.PlaylistResponse]{Message: "success", Success: true, ShouldNotify: false, Data: *recommendations})
}

func (p *PlaylistHandler) GetAvailableGenres(ctx *fiber.Ctx) error {
	genres, err := p.service.GetGenreSeeds()
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	}

	return ctx.JSON(types.WebResponse[any]{Message: "success", Success: true, ShouldNotify: false, Data: genres})
}

func (p *PlaylistHandler) GetSearch(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	if search == "" {
		return helpers.ErrorResponse(ctx, fiber.StatusBadRequest, true, errors.New("search query is required"))
	}

	searchQuery := "track:" + search
	result, err := p.service.SearchTracks(searchQuery)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	}

	return ctx.JSON(types.WebResponse[any]{Message: "success", Success: true, ShouldNotify: false, Data: *result})
}
