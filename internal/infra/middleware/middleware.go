package middleware

import "github.com/gofiber/fiber/v2"

func DefaultMiddleware(config ...fiber.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}

var GetUserInfo = func(c fiber.Ctx) error {
	return c.Next()
}

var InjectLogger = func(c fiber.Ctx) error {
	return c.Next()
}
