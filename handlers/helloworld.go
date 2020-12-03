package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func (h *Handler) HelloWorld(ctx *fiber.Ctx) error {
	fmt.Println("hello world!")
	u, err := h.Client.User.Create().
		SetEmail("jason@aucklanders.co.nz").
		SetLastName("Shen").
		SetFirstName("Jason").
		SetPassword("1234").
		SetActivateCode(1234).
		SetJwttoken("1234556565").
		Save(ctx.Context())
	if err != nil {
		return fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)

	return nil
}
