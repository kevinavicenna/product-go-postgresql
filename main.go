package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kevinavicenna/product-go-postgresql/models"
	"github.com/kevinavicenna/product-go-postgresql/storage"
	"gorm.io/gorm"
)

type Product struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateProduct(context *fiber.Ctx) error {
	product := Product{}
	err := context.BodyParser(&product)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	err = r.DB.Create(&product).Error
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
			&fiber.Map{"message": "could not get any product"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "product fetch",
			"data":    ProductModels,
		})
	return nil
}

func (r *Repository) DeleteProduct(context *fiber.Ctx) error {
	ProductModels := &[]models.Products{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cant empty",
		})
		return nil
	}

	err := r.DB.Delete(ProductModels, id)
	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not delete product"})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "delete product successfully",
	})
	return nil
}

func (r *Repository) GetProductID(context *fiber.Ctx) error {
	id := context.Params("id")
	ProductModels := &models.Products{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "could not get product by id"})
	}

	fmt.Println("this id :", id)

	err := r.DB.Where("id = ?", id).First(ProductModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could nnot get product by id"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "get id",
			"data":    ProductModels,
		})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_product", func(c *fiber.Ctx) error { return r.CreateProduct(c) })
	api.Delete("/delete_product/:id", func(c *fiber.Ctx) error { return r.DeleteProduct(c) })
	api.Get("/get_product/:id", func(c *fiber.Ctx) error { return r.GetProductID(c) })
	api.Get("/products", func(c *fiber.Ctx) error { return r.GetAllProduct(c) })
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
		SSLMode:  os.Getenv("DB_SSL"),
		DB:       os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Cant load database")
	}

	err = models.MigrateProduct(db)
	if err != nil {
		log.Fatal("cant migrate db")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
