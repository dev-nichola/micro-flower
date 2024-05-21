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
		exception.ResponseError(ctx, fiber.StatusNotFound, fiber.ErrNotFound.Message, err)
	}

	return helper.Response(ctx, "successfully findall", products)

}

func (handler *ProductHandlerImpl) FindByID(ctx *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	productId, err := uuid.Parse(ctx.Params("productId"))

	if err != nil {
		return exception.ResponseError(ctx, fiber.StatusNotFound, "Not Found Product", err)
	}

	product, err := handler.Service.FindByID(customContext, productId)
	if err != nil {
		return exception.ResponseError(ctx, fiber.StatusInternalServerError, "Failed to find Product", err)
	}

	return helper.Response(ctx, "successfully find id", product)
}

func (handler *ProductHandlerImpl) Create(ctx *fiber.Ctx) error {
	var product Product
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := ctx.BodyParser(&product); err != nil {
		return exception.ResponseError(ctx, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, err)
	}

	// Create one product
	product, err := handler.Service.Save(customContext, product)
	if err != nil {
		return exception.ResponseError(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message, err)

	}

	return helper.Response(ctx, "Successfully created product", product)
}

func (handler *ProductHandlerImpl) Edit(ctx *fiber.Ctx) error {
	var product Product
	customContext, cancel := context.WithCancel(context.Background())

	defer cancel()

	if err := ctx.BodyParser(&product); err != nil {
		return exception.ResponseError(ctx, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, err)
	}

	productId, err := uuid.Parse(ctx.Params("productId"))

	if err != nil {
		return exception.ResponseError(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message, err)
	}

	result, err := handler.Service.Update(customContext, productId, product)

	if err != nil {
		return exception.ResponseError(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message, err)
	}

	return helper.Response(ctx, "Successfully created product", result)

}

func (handler *ProductHandlerImpl) Destroy(ctx *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Parsing UUID dari parameter URL
	productId, err := uuid.Parse(ctx.Params("productId"))
	if err != nil || productId == uuid.Nil {
		return exception.ResponseError(ctx, fiber.StatusNotFound, fiber.ErrNotFound.Message, err)
	}

	// Menghapus produk
	err = handler.Service.Delete(customContext, productId)
	if err != nil {
		return exception.ResponseError(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "successfully deleted product",
	})
}
