package music

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/music/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/music/repository"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/helpers"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	musicRepo *repository.MusicRepository
}

func NewMusicHandler(musicRepo *repository.MusicRepository) *Handler {
	return &Handler{musicRepo: musicRepo}
}

func (m Handler) GetMusic(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	music, err := m.musicRepo.FindByColumnValue("history_uuid", id)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	}

	return ctx.JSON(types.WebResponse[[]entity.Music]{
		Data:         music,
		Message:      "success",
		ShouldNotify: false,
		Success:      true,
	})
}
