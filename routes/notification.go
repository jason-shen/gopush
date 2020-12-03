package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jason-shen/gopush/handlers"
)

func SetupNotificationRoutes(grp fiber.Router, handlers *handlers.Handler) {
	// conf := config.New()
	userRoute := grp.Group("/notification")
	userRoute.Post("/send", handlers.SendNotification)
}
