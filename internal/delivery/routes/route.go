package routes

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/delivery/middleware"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/playlist"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"time"
)

type Route struct {
	App             *fiber.App
	UserHandler     *user.UserHandler
	AuthMiddleware  *middleware.AuthMiddleware
	PlaylistHandler *playlist.PlaylistHandler
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
	r.InitializeUserRoutes(v1)
}

func (r *Route) InitializeUserRoutes(router fiber.Router) {
	router.Post("/register", middleware.EnsureJsonValidRequest[user.RegisterRequest], r.UserHandler.Register)
	router.Post("/login", middleware.EnsureJsonValidRequest[user.LoginRequest], r.UserHandler.Login)
	router.Post("/logout", r.AuthMiddleware.EnsureAuthenticated, middleware.EnsureJsonValidRequest[user.LogoutRequest], r.UserHandler.Logout)
	router.Post("/refresh", middleware.EnsureJsonValidRequest[user.RefreshTokenRequest], r.UserHandler.RefreshToken)

	router.Get("/user", r.AuthMiddleware.EnsureAuthenticated, r.UserHandler.GetProfile)
	router.Put("/user", r.AuthMiddleware.EnsureAuthenticated, middleware.EnsureJsonValidRequest[user.UpdateProfileRequest], r.UserHandler.UpdateProfile)
	router.Put("/user/profile-picture", r.AuthMiddleware.EnsureAuthenticated, r.UserHandler.UploadProfilePicture)
	router.Put("/user/change-password", r.AuthMiddleware.EnsureAuthenticated, middleware.EnsureJsonValidRequest[user.ChangePasswordRequest], r.UserHandler.ChangePassword)

	router.Get("/random-playlist", r.AuthMiddleware.EnsureAuthenticated, r.PlaylistHandler.GetRecommendations)
}
