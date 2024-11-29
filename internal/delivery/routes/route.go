package routes

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/delivery/middleware"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/history"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/music"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"time"
)

type Route struct {
	App             *fiber.App
	UserHandler     *user.Handler
	AuthMiddleware  *middleware.AuthMiddleware
	PlaylistHandler *history.Handler
	MusicHandler    *music.Handler
}

func (r *Route) InitV1() {
	r.App.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))

	api := r.App.Group("/api")
	v1 := api.Group("/v1")
	r.initializeUserRoutes(v1)
	r.initializeSpotifyRoutes(v1)
	r.initializeHistoryRoutes(v1)
	r.initializeMusicRoutes(v1)
}

func (r *Route) initializeUserRoutes(router fiber.Router) {
	router.Post("/register", middleware.EnsureJsonValidRequest[user.RegisterRequest], r.UserHandler.Register)
	router.Post("/login", middleware.EnsureJsonValidRequest[user.LoginRequest], r.UserHandler.Login)
	router.Post("/logout", r.AuthMiddleware.EnsureAuthenticated, middleware.EnsureJsonValidRequest[user.LogoutRequest], r.UserHandler.Logout)
	router.Post("/refresh", middleware.EnsureJsonValidRequest[user.RefreshTokenRequest], r.UserHandler.RefreshToken)
	router.Post("/send-otp", middleware.EnsureJsonValidRequest[user.VerifyEmailRequest], r.UserHandler.SendOtp)
	router.Post("/verify-email", middleware.EnsureJsonValidRequest[user.VerifyOtpRequest], r.UserHandler.VerifyEmail)
	router.Post("/reset-password", middleware.EnsureJsonValidRequest[user.ResetPasswordRequest], r.UserHandler.ResetPassword)

	router.Post("/mood-detection", r.AuthMiddleware.EnsureAuthenticated, r.PlaylistHandler.MoodDetect)

	userRoute := router.Group("/user")
	userRoute.Get("/", r.AuthMiddleware.EnsureAuthenticated, r.UserHandler.GetProfile)
	userRoute.Put("/", r.AuthMiddleware.EnsureAuthenticated, middleware.EnsureJsonValidRequest[user.UpdateProfileRequest], r.UserHandler.UpdateProfile)
	userRoute.Put("/profile-picture", r.AuthMiddleware.EnsureAuthenticated, r.UserHandler.UploadProfilePicture)
	userRoute.Put("/change-password", r.AuthMiddleware.EnsureAuthenticated, middleware.EnsureJsonValidRequest[user.ChangePasswordRequest], r.UserHandler.ChangePassword)
}

func (r *Route) initializeSpotifyRoutes(router fiber.Router) {
	spotify := router.Group("/spotify")
	spotify.Get("/random-playlist", r.AuthMiddleware.EnsureAuthenticated, r.PlaylistHandler.GetRecommendations)
	spotify.Get("/search", r.AuthMiddleware.EnsureAuthenticated, r.PlaylistHandler.GetSearch)
	spotify.Get("/:id", r.AuthMiddleware.EnsureAuthenticated, r.PlaylistHandler.GetTrackByID)
}

func (r *Route) initializeHistoryRoutes(router fiber.Router) {
	historyRoute := router.Group("/history")
	historyRoute.Get("/", r.AuthMiddleware.EnsureAuthenticated, r.PlaylistHandler.GetHistory)
}

func (r *Route) initializeMusicRoutes(router fiber.Router) {
	musicRoute := router.Group("/music")
	musicRoute.Get("/:id", r.AuthMiddleware.EnsureAuthenticated, r.MusicHandler.GetMusic)
}
