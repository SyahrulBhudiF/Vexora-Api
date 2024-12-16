package history

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/history/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/history/repository"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/history/service"
	entity2 "github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/helpers"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/services"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"time"
)

type Handler struct {
	service         *services.SpotifyService
	historyRepo     *repository.HistoryRepository
	tokenRepository *types.RedisRepository
	moodService     *service.Service
}

func NewPlaylistHandler(services *services.SpotifyService, clientUri string, ClientKey string, repo *repository.HistoryRepository, token *types.RedisRepository) *Handler {
	moodService := service.NewService(clientUri, ClientKey)
	return &Handler{service: services, historyRepo: repo, moodService: moodService, tokenRepository: token}
}

func (p *Handler) GetRecommendations(ctx *fiber.Ctx) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomGenre := availableGenres[r.Intn(len(availableGenres))]

	recommendations, err := p.service.SearchTracks(randomGenre, 10)
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

func (p *Handler) GetSearch(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	if search == "" {
		return helpers.ErrorResponse(ctx, fiber.StatusBadRequest, true, errors.New("search query is required"))
	}

	searchQuery := "genre:" + search
	result, err := p.service.SearchTracks(searchQuery, 10)
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

	mood, err := p.moodService.DetectMood(file)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	}

	trackAttrs, _ := GenreMoodTrackAttributes[mood.Data]

	rand.Shuffle(len(trackAttrs), func(i, j int) {
		trackAttrs[i], trackAttrs[j] = trackAttrs[j], trackAttrs[i]
	})

	existingTrackIds := make(map[string]bool)
	finalPlaylist := &entity.MoodResponse{
		Music: []entity.RandomMusic{},
	}

	trackCh := make(chan []entity.RandomMusic, len(trackAttrs))
	errCh := make(chan error, len(trackAttrs))
	maxGoroutines := 5
	semaphore := make(chan struct{}, maxGoroutines)

	targetTracks := 10
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, genre := range trackAttrs {
		semaphore <- struct{}{}
		go func(genre string) {
			defer func() { <-semaphore }()

			searchQuery := "genre:" + genre
			recommendations, err := p.service.SearchTracks(searchQuery, 3)
			if err != nil {
				select {
				case errCh <- err:
				case <-ctxWithTimeout.Done():
				}
				return
			}

			select {
			case trackCh <- recommendations.Music:
			case <-ctxWithTimeout.Done():
			}
		}(genre)
	}

	go func() {
		for i := 0; i < len(trackAttrs); i++ {
			select {
			case tracks := <-trackCh:
				for _, track := range tracks {
					if len(finalPlaylist.Music) >= targetTracks {
						cancel()
						return
					}
					if !existingTrackIds[track.ID] {
						existingTrackIds[track.ID] = true
						finalPlaylist.Music = append(finalPlaylist.Music, track)
					}
				}
			case <-ctxWithTimeout.Done():
				return
			}
		}
		close(trackCh)
		close(errCh)
	}()

	select {
	case <-ctxWithTimeout.Done():
		if len(finalPlaylist.Music) < targetTracks {
			return helpers.ErrorResponse(ctx, fiber.StatusGatewayTimeout, true, errors.New("timeout while fetching tracks"))
		}
	}

	err = p.historyRepo.Transaction(func(tx *types.Repository[entity.History]) error {
		user, _ := ctx.Locals("user").(*entity2.User)
		history := CreateHistoryEntry(user, mood.Data, finalPlaylist)
		if err := tx.Create(history); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	}

	user, _ := ctx.Locals("user").(*entity2.User)
	redisKey := fmt.Sprintf("user:%s:history", user.UUID)
	if err := p.tokenRepository.Delete(redisKey); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to delete history cache"))
	}

	finalPlaylist.Mood = mood.Data
	finalPlaylist.CreatedAt = time.Now().Format(time.RFC3339)

	return ctx.JSON(types.WebResponse[entity.MoodResponse]{
		Message:      "success",
		Success:      true,
		ShouldNotify: false,
		Data:         *finalPlaylist,
	})
}

func (p *Handler) GetHistory(ctx *fiber.Ctx) error {
	user, _ := ctx.Locals("user").(*entity2.User)

	redisKey := fmt.Sprintf("user:%s:history", user.UUID)
	cachedHistory, err := p.tokenRepository.Get(redisKey)
	if err == nil && cachedHistory != "" {
		var history []entity.History
		if err := json.Unmarshal([]byte(cachedHistory), &history); err == nil {
			return ctx.JSON(types.WebResponse[[]entity.History]{Message: "success", Success: true, ShouldNotify: false, Data: history})
		}
	}

	history, err := p.historyRepo.FindByColumnValue("user_uuid", user.UUID)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	}

	historyJSON, err := json.Marshal(history)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to cache history"))
	}

	if err := p.tokenRepository.Set(redisKey, string(historyJSON), 60*time.Minute); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to save history data to cache"))
	}
	return ctx.JSON(types.WebResponse[[]entity.History]{Message: "success", Success: true, ShouldNotify: false, Data: history})
}

func (p *Handler) GetMostFrequentMood(ctx *fiber.Ctx) error {
	user, _ := ctx.Locals("user").(*entity2.User)
	mood, err := p.historyRepo.GetMostFrequentMoodByUserUUID(user.UUID)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, err)
	}

	return ctx.JSON(types.WebResponse[entity.MostMood]{Message: "success", Success: true, ShouldNotify: false, Data: entity.MostMood{Mood: mood}})
}
