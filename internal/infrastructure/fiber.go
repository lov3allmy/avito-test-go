package infrastructure

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	handler2 "github.com/lov3allmy/avito-test-go/internal/handler"
	"github.com/lov3allmy/avito-test-go/internal/repository"
	"github.com/lov3allmy/avito-test-go/internal/service"
	"log"
)

func Run() {
	postgres, err := ConnectToPostgres()
	if err != nil {
		log.Fatalf("Database connection error: %s", err.Error())
	}

	app := fiber.New(fiber.Config{
		AppName: "Avito Test Go",
	})

	userRepository := repository.NewRepository(postgres)

	userService := service.NewService(userRepository)

	userHandler := handler2.NewHandler(userService)

	api := app.Group("/api")

	handler2.Router(api, userHandler)

	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"message": errorMessage,
		})
	})

	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err.Error())
	}
}
