package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func (h *Handler) HelloWorld(ctx *fiber.Ctx) error {
	fmt.Println("hello world!")
	u, err := h.Client.User.Create().
		SetEmail("jason_2000_nz@hotamil.com").
		SetLastName("Shen").
		SetFirstName("Jason").
		SetPassword("1234").
		SetActivateCode(12).
		Save(ctx.Context())
	if err != nil {
		return fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)

	return nil
}
