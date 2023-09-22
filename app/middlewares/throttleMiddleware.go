package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/spf13/viper"
	"time"
)

type ThrottleMiddleware struct {
}

func (throttle *ThrottleMiddleware) Limit(ctx *fiber.Ctx) error {
	// Throttling middleware for specific routes
	limiterConfig := limiter.Config{
		Max:        viper.GetInt("throttle.max"),                     // Maximum number of requests allowed within the duration
		Expiration: viper.GetDuration("throttle.time") * time.Minute, // Duration within which the maximum number of requests is allowed
		LimitReached: func(c *fiber.Ctx) error { // Custom handler when the limit is reached
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too Many Requests",
			})
		},
	}
	limiterMiddleware := limiter.New(limiterConfig)
	return limiterMiddleware(ctx)
}
