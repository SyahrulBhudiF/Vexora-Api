package playlist

import (
	"errors"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/playlist/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/helpers"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/services"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *services.SpotifyService
}

func NewPlaylistHandler(service *services.SpotifyService) *Handler {
	return &Handler{service: service}
}

func (p *Handler) GetRecommendations(ctx *fiber.Ctx) error {
	recommendations, err := p.service.GetRecommendations(10, nil)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	}

	return ctx.JSON(types.WebResponse[entity.PlaylistResponse]{Message: "success", Success: true, ShouldNotify: false, Data: *recommendations})
}

func (p *Handler) GetSearch(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	if search == "" {
		return helpers.ErrorResponse(ctx, fiber.StatusBadRequest, true, errors.New("search query is required"))
	}

	searchQuery := "track:" + search
	result, err := p.service.SearchTracks(searchQuery)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	}

	return ctx.JSON(types.WebResponse[entity.PlaylistResponse]{Message: "success", Success: true, ShouldNotify: false, Data: *result})
}

func (p *Handler) GetTrackByID(ctx *fiber.Ctx) error {
	ctxID := ctx.Params("id")

	result, err := p.service.GetTrackByID(ctxID)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	}

	return ctx.JSON(types.WebResponse[entity.PlaylistResponse]{Message: "success", Success: true, ShouldNotify: false, Data: *result})
}
