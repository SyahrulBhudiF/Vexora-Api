package core

import (
	"fmt"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/delivery/middleware"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/delivery/routes"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/history"
	repository2 "github.com/SyahrulBhudiF/Vexora-Api/internal/domains/history/repository"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/music"
	repository3 "github.com/SyahrulBhudiF/Vexora-Api/internal/domains/music/repository"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user/repository"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/services"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Vexora struct {
	Viper    *viper.Viper
	DB       *gorm.DB
	App      *fiber.App
	Redis    *redis.Client
	JWT      *services.JWTService
	ImageKit *services.ImageKitService
	Spotify  *services.SpotifyService
	Mail     *services.MailService
}

func Init(vexora *Vexora) {
	userRepo := repository.NewUserRepository(vexora.DB)
	tokenRepo := types.NewRedisRepository(vexora.Redis, "token")
	userHandler := user.NewUserHandler(userRepo, tokenRepo, vexora.JWT, vexora.ImageKit, vexora.Viper, vexora.Mail)

	historyRepo := repository2.NewHistoryRepository(vexora.DB)
	playlistHandler := history.NewPlaylistHandler(vexora.Spotify, vexora.Viper.GetString("auth.client_url"), vexora.Viper.GetString("auth.client_key"), historyRepo)

	musicHandler := music.NewMusicHandler(repository3.NewMusicRepository(vexora.DB))

	authMiddleware := middleware.NewAuthMiddleware(userRepo, tokenRepo, vexora.JWT)

	route := routes.Route{
		App:             vexora.App,
		UserHandler:     userHandler,
		AuthMiddleware:  authMiddleware,
		PlaylistHandler: playlistHandler,
		MusicHandler:    musicHandler,
	}

	route.InitV1()
}

func (a *Vexora) Start() {
	err := a.App.Listen(fmt.Sprintf("%s:%s", a.Viper.GetString("app.host"), a.Viper.GetString("app.port")))
	if err != nil {
		logrus.Fatal(err)
	}
}
