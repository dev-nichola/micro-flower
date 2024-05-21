package helper

import "github.com/gofiber/fiber/v2"

func Response(ctx *fiber.Ctx, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success",
		"data":    data,
	})
}
