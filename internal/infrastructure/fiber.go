package infrastructure

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	handler2 "github.com/lov3allmy/avito-test-go/internal/handler"
	"github.com/lov3allmy/avito-test-go/internal/repository"
	"github.com/lov3allmy/avito-test-go/internal/service"
	"github.com/spf13/viper"
	"log"
)

func Run() {
	if err := initConfig(); err != nil {
		log.Fatal("initializing viper config failed with error" + err.Error())
	}

	cfg := postgresConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		DBName:   "avito_test_go",
		SSLMode:  "disable",
	}

	postgres, err := ConnectToPostgres(cfg)
	if err != nil {
		log.Fatal("Connecting to db failed with error: " + err.Error())
	}

	app := fiber.New(fiber.Config{
		AppName: "Avito Test Go",
	})

	repos := repository.NewRepository(postgres)

	services := service.NewService(repos)

	handlers := handler2.NewHandler(services)

	api := app.Group("/api")

	handler2.Router(api, handlers)

	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"message": errorMessage,
		})
	})

	if err := app.Listen(":" + viper.GetString("port")); err != nil {
		log.Fatal("Launching server failed with error" + err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("main")
	return viper.ReadInConfig()
}
