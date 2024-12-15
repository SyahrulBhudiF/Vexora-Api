package music

import (
	"encoding/json"
	"fmt"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/music/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/music/repository"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/helpers"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Handler struct {
	musicRepo *repository.MusicRepository
	tokenRepo *types.RedisRepository
}

func NewMusicHandler(musicRepo *repository.MusicRepository, token *types.RedisRepository) *Handler {
	return &Handler{musicRepo: musicRepo, tokenRepo: token}
}

func (m Handler) GetMusic(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	redisKey := fmt.Sprintf("music:%s", id)
	cachedMusic, err := m.tokenRepo.Get(redisKey)
	if err == nil && cachedMusic != "" {
		var music []entity.Music
		if err := json.Unmarshal([]byte(cachedMusic), &music); err == nil {
			return ctx.JSON(types.WebResponse[[]entity.Music]{
				Data:         music,
				Message:      "success",
				ShouldNotify: false,
				Success:      true,
			})
		}
	}

	music, err := m.musicRepo.FindByColumnValue("history_uuid", id)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	}

	musicJSON, err := json.Marshal(music)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to cache music"))
	}

	if err := m.tokenRepo.Set(redisKey, string(musicJSON), 60*time.Minute); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to save music data to cache"))
	}

	return ctx.JSON(types.WebResponse[[]entity.Music]{
		Data:         music,
		Message:      "success",
		ShouldNotify: false,
		Success:      true,
	})
}
