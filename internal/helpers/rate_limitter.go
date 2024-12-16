package helpers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func RateLimiterConfig(expiration time.Duration, maxRequests int, errorMsg string) fiber.Handler {
	return limiter.New(limiter.Config{
		Expiration: expiration,
		Max:        maxRequests,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return ErrorResponse(c, fiber.StatusTooManyRequests, true, errors.New(errorMsg))
		},
	})
}
