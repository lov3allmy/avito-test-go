package infrastructure

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/lov3allmy/avito-test-go/internal/user"
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

	userRepository := user.NewRepository(postgres)

	userService := user.NewService(userRepository)

	user.NewHandler(app.Group("/api/users"), userService)

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
