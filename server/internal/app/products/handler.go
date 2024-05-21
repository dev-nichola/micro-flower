package products

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductHandler interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Edit(ctx *fiber.Ctx) error
	Destroy(ctx *fiber.Ctx) error
}

type ProductHandlerImpl struct {
	Service ProductService
}

func NewProductHanlder(service ProductService) ProductHandler {
	return &ProductHandlerImpl{
		Service: service,
	}
}

func (handler *ProductHandlerImpl) FindAll(ctx *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())

	defer cancel()

	products, err := handler.Service.FindAll(customContext)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed to get products",
			"message": err.Error(),
		})
	}

	// return resultnya
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": products,
	})

}

func (handler *ProductHandlerImpl) FindByID(ctx *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	productId, err := uuid.Parse(ctx.Params("productId"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "bad request",
			"message": err.Error(),
		})
	}

	product, err := handler.Service.FindByID(customContext, productId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed to get product",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success",
		"data":    product,
	})
}

func (handler *ProductHandlerImpl) Create(ctx *fiber.Ctx) error {
	var product Product
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "bad request",
			"message": err.Error(),
		})
	}

	// Create one product
	product, err := handler.Service.Save(customContext, product)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed to create product",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "successfully created product",
		"message": product,
	})
}

func (handler *ProductHandlerImpl) Edit(ctx *fiber.Ctx) error {
	return nil
}

func (handler *ProductHandlerImpl) Destroy(ctx *fiber.Ctx) error {
	return nil
}
