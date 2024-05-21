package products

import (
	"context"

	"github.com/google/uuid"
)

type ProductService interface {
	FindAll(ctx context.Context) (*[]Product, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Product, error)
	Save(ctx context.Context, product Product) (Product, error)
	Update(ctx context.Context, id uuid.UUID, product Product) (Product, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type ProductServiceImpl struct {
	Repository ProductRepository
}

func NewProductService(repository ProductRepository) ProductService {
	return &ProductServiceImpl{
		Repository: repository,
	}
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) (*[]Product, error) {
	products, err := service.Repository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (service *ProductServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*Product, error) {
	product, err := service.Repository.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *ProductServiceImpl) Save(ctx context.Context, product Product) (Product, error) {
	product, err := service.Repository.Save(ctx, product)

	defer recover()

	if err != nil {
		panic(err)
	}

	return product, err
}

func (service *ProductServiceImpl) Update(ctx context.Context, id uuid.UUID, product Product) (Product, error) {
	result, err := service.Repository.Update(ctx, id, product)

	defer recover()

	if err != nil {
		panic(err)
	}

	return result, err
}

func (service *ProductServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	err := service.Repository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
