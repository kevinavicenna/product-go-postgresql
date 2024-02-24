package main

import (
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
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
func (r *Repository) DeleteProduct() {

}

func (r *Repository) GetProductID() {

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

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
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
