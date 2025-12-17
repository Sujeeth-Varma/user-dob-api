package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RequestLogger(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()

		duration := time.Since(start)
		requestID, _ := c.Locals(RequestIDKey).(string)

		logger.Info("http request",
			zap.String("method", c.Method()),
			zap.String("request_id", requestID),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
		)
		return err
	}
}
