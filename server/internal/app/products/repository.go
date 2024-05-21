package products

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type ProductRepository interface {
	FindAll(ctx context.Context) (*[]Product, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Product, error)
	Save(ctx context.Context, product Product) (Product, error)
	Update(ctx context.Context, id uuid.UUID, product Product) (Product, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type ProductRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) FindAll(ctx context.Context) (*[]Product, error) { // return pointer
	var products []Product
	SQL := "SELECT id, name, price, updated_at, created_at FROM products"
	res, err := r.db.QueryContext(ctx, SQL)

	if err != nil {
		panic(err)
	}

	defer res.Close()

	// Melakukan iterasi untuk mendapatkan semua data
	for res.Next() {
		product := &Product{}

		// Scan semua query sqlnya dan harus secara berurutan sesuai querynya
		err := res.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.UpdatedAt,
			&product.CreatedAt,
		)
		if err == sql.ErrNoRows && err != nil {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}

		//Append Menambahkan elemen pada slice
		products = append(products, *product)
	}

	return &products, nil
}

func (r *ProductRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*Product, error) {
	product := &Product{}

	SQL := "SELECT id, name, price, updated_at, created_at FROM products WHERE id = $1"
	row := r.db.QueryRowContext(ctx, SQL, id)

	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.UpdatedAt,
		&product.CreatedAt,
	)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err != nil {
		return nil, nil
	}

	return product, nil
}

func (r *ProductRepositoryImpl) Save(ctx context.Context, product Product) (Product, error) {
	id := uuid.New()
	updated_at := time.Now()
	created_at := time.Now()

	SQL := "INSERT INTO products(id, name, price, updated_at, created_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.ExecContext(ctx, SQL, id, product.Name, product.Price, updated_at, created_at)

	defer recover()

	if err != nil {
		panic(err)
	}

	product.ID = id
	product.UpdatedAt = time.Now()
	product.CreatedAt = time.Now()

	return product, nil
}

func (r *ProductRepositoryImpl) Update(ctx context.Context, id uuid.UUID, product Product) (Product, error) {

	updated_at := time.Now()
	SQL := "UPDATE products SET name = $1, price = $2, updated_at = $3 where id = $4"
	_, err := r.db.ExecContext(ctx, SQL, product.Name, product.Price, updated_at, id)

	defer recover()

	if err != nil {
		panic(err)
	}

	product.ID = id
	return product, nil
}

func (r *ProductRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	SQL := "DELETE FROM products where id = $1"
	_, err := r.db.ExecContext(ctx, SQL, id)

	if err != nil {
		return err
	}

	return nil
}
