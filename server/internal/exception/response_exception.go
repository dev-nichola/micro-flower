package exception

import "github.com/gofiber/fiber/v2"

func ResponseError(ctx *fiber.Ctx, statusCode any, status string, err error) error {
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"status":  "Not Found",
		"message": err.Error(),
	})
}
