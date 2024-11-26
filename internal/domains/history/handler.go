package history

import (
	"encoding/json"
	"errors"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/history/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/history/repository"
	entity2 "github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/helpers"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/services"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

type Handler struct {
	service     *services.SpotifyService
	clientUrl   string
	clientKey   string
	historyRepo *repository.HistoryRepository
}

func NewPlaylistHandler(service *services.SpotifyService, clientUrl string, clientKey string, repo *repository.HistoryRepository) *Handler {
	return &Handler{service: service, clientUrl: clientUrl, clientKey: clientKey, historyRepo: repo}
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

func (p *Handler) MoodDetect(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("image")
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusBadRequest, true, errors.New("file is required"))
	}

	if err := helpers.ValidateImageFile(file); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusBadRequest, true, err)
	}

	resultCh := make(chan *entity.MoodDetectionResponse, 1)
	errCh := make(chan error, 1)

	go func() {
		req, err := helpers.CreateMultipartRequest(p.clientUrl, file, p.clientKey)
		if err != nil {
			errCh <- err
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errCh <- err
			return
		}
		defer resp.Body.Close()

		var mood entity.MoodDetectionResponse
		if err := json.NewDecoder(resp.Body).Decode(&mood); err != nil {
			errCh <- errors.New("failed to decode response")
			return
		}

		resultCh <- &mood
	}()

	select {
	case mood := <-resultCh:
		trackAttrs, _ := entity.MoodTrackAttributes[mood.Data]

		recommendations, err := p.service.GetRecommendations(10, &trackAttrs)
		if err != nil {
			return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
		}

		err = p.historyRepo.Transaction(func(tx *types.Repository[entity.History]) error {
			user, _ := ctx.Locals("user").(*entity2.User)

			history := CreateHistoryEntry(user, mood.Data, recommendations)

			if err = tx.Create(history); err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
		}

		return ctx.JSON(types.WebResponse[entity.PlaylistResponse]{
			Message:      "success",
			Success:      true,
			ShouldNotify: false,
			Data:         *recommendations,
		})
	case err := <-errCh:
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	case <-time.After(10 * time.Second):
		return helpers.ErrorResponse(ctx, fiber.StatusGatewayTimeout, true, errors.New("server timeout"))
	}
}
