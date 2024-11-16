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
			return c.Params("refresh") == "T"
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

	userRoute := router.Group("/user")
	userRoute.Get("/", r.AuthMiddleware.EnsureAuthenticated, r.UserHandler.GetProfile)
	userRoute.Put("/", r.AuthMiddleware.EnsureAuthenticated, middleware.EnsureJsonValidRequest[user.UpdateProfileRequest], r.UserHandler.UpdateProfile)
	userRoute.Put("/user/profile-picture", r.AuthMiddleware.EnsureAuthenticated, r.UserHandler.UploadProfilePicture)
	userRoute.Put("/user/change-password", r.AuthMiddleware.EnsureAuthenticated, middleware.EnsureJsonValidRequest[user.ChangePasswordRequest], r.UserHandler.ChangePassword)

	router.Get("/random-playlist", r.AuthMiddleware.EnsureAuthenticated, r.PlaylistHandler.GetRecommendations)
	router.Get("/available-genres", r.AuthMiddleware.EnsureAuthenticated, r.PlaylistHandler.GetAvailableGenres)
	router.Get("/spotify/search", r.AuthMiddleware.EnsureAuthenticated, r.PlaylistHandler.GetSearch)
}
