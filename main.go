package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lenna-ai/azureOneSmile.git/config"
	appconfig "github.com/lenna-ai/azureOneSmile.git/config/appConfig"
	"github.com/lenna-ai/azureOneSmile.git/helpers"
	"github.com/lenna-ai/azureOneSmile.git/routes"
)

func main()  {
	defer helpers.RecoverPanicContext(&fiber.Ctx{})
	appconfig.InitApplication()
	app := fiber.New()
	app.Use(cors.New())
	config.Logger(app)
	routes.Router(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err.Error())
	}
}