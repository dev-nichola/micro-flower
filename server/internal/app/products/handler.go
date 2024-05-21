package products

import (
	"context"

	"github.com/dev-nichola/todo-go/internal/exception"
	"github.com/dev-nichola/todo-go/internal/pkg/helper"
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
		exception.ResponseError(ctx, fiber.StatusNotFound, "Not Found Product", err)
	}

	product, err := handler.Service.FindByID(customContext, productId)
	if err != nil {
		exception.ResponseError(ctx, fiber.StatusInternalServerError, "Failed to find Product", err)
	}

	return helper.Response(ctx, product)
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
	var product Product
	customContext, cancel := context.WithCancel(context.Background())

	defer cancel()

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "bad request",
			"message": err.Error(),
		})
	}

	productId, err := uuid.Parse(ctx.Params("productId"))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Internal server error",
			"message": err.Error(),
		})
	}

	result, err := handler.Service.Update(customContext, productId, product)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed to update product",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": result,
	})

}

func (handler *ProductHandlerImpl) Destroy(ctx *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())

	defer cancel()

	productId, err := uuid.Parse(ctx.Params("productId"))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "ID Not Found",
			"message": err.Error(),
		})
	}

	err = handler.Service.Delete(customContext, productId)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed to delete product",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "successfully deleted product",
	})
}
