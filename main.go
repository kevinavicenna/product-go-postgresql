package main

import (
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"github.com/kevinavicenna/product-go-postgresql/models"
	"github.com/kevinavicenna/product-go-postgresql/storage"
	"gorm.io/gorm"

	"log"
	"net/http"
	"os"
)

type Product struct {
	Name        string `json:"author"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_product", r.CreateProduct)
	api.Delete("delete_product/:id", r.DeleteProduct)
	api.Get("/get_product/:id", r.GetProductID)
	api.Get("/product", r.GetAllProduct)
}

func (r *Repository) CreateProduct(context *fiber.Ctx) error {
	product := Product{}
	err := context.BodyParser(&product)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create product"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "well done you create product:v"})
	return nil
}

func (r *Repository) GetAllProduct(context *fiber.Ctx) error {
	ProductModels := &[]models.Products{}

	err := r.DB.Find(ProductModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.MAP{"message": "invd"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "get all product",
			"data": ProductModels})
	return nil
}

func (r *Repository) DeleteProduct(context *fiber.Ctx) {
	ProductModels := &[]models.Products{}

}

func (r *Repository) GetProductID() {
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		db:       os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Cant load database")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
