package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)

func SetupMiddleware(app *fiber.App) {
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
}
