package routes

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/delivery/middleware"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"time"
)

type Route struct {
	App            *fiber.App
	UserHandler    *user.UserHandler
	AuthMiddleware *middleware.AuthMiddleware
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
}
