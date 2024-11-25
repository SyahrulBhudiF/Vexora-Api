package playlist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/playlist/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/helpers"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/services"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/gofiber/fiber/v2"
	"io"
	"mime/multipart"
	"net/http"
)

type Handler struct {
	service   *services.SpotifyService
	clientUrl string
	clientKey string
}

func NewPlaylistHandler(service *services.SpotifyService, clientUrl string, clientKey string) *Handler {
	return &Handler{service: service, clientUrl: clientUrl, clientKey: clientKey}
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

	if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" {
		return helpers.ErrorResponse(ctx, fiber.StatusBadRequest, true, errors.New("file must be an image"))
	}

	src, err := file.Open()
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, errors.New("failed to open file"))
	}
	defer src.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("image", file.Filename)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, errors.New("failed to create form file"))
	}

	if _, err = io.Copy(part, src); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, errors.New("failed to copy file"))
	}

	if err = writer.Close(); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, errors.New("failed to close writer"))
	}

	req, err := http.NewRequest("POST", p.clientUrl, body)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, errors.New("failed to create request"))
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Secret-Key", p.clientKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	}

	fmt.Printf("Response: %v\n", resp)

	defer resp.Body.Close()

	var mood entity.MoodDetectionResponse
	if err := json.NewDecoder(resp.Body).Decode(&mood); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, errors.New("failed to decode response"))
	}

	trackAttrs, _ := entity.MoodTrackAttributes[mood.Data]

	recommendations, err := p.service.GetRecommendations(10, &trackAttrs)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	}

	return ctx.JSON(types.WebResponse[entity.PlaylistResponse]{
		Message:      "success",
		Success:      true,
		ShouldNotify: false,
		Data:         *recommendations,
	})
}
