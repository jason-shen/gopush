package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jason-shen/gopush/handlers"
)

func SetupUserRoutes(grp fiber.Router, handlers *handlers.Handler) {
	// conf := config.New()
	userRoute := grp.Group("/user")
	userRoute.Post("/register", handlers.UserRegister)
	userRoute.Post("/login", handlers.LoginUser)
}
