package infrastructure

import (
	"log"
	"time"

	"github.com/dev-nichola/todo-go/internal/app/products"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func Run() {

	DB, err := NewDB()

	if err != nil {
		log.Println("Ada Yang Masalah Ini Master :)")
	}

	productRepository := products.NewProductRepository(DB)
	productService := products.NewProductService(productRepository)
	productHandler := products.NewProductHanlder(productService)

	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
	})

	app.Get("/products", productHandler.FindAll)
	app.Get("/product/:productId", productHandler.FindByID)
	app.Post("/product", productHandler.Create)
	app.Put("/product/:productId", productHandler.Edit)
	app.Delete("/product/:productId", productHandler.Destroy)

	app.Listen("localhost:8080")
}
