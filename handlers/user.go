package handlers

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofiber/fiber/v2"
	"github.com/jason-shen/gopush/ent/user"
	"github.com/jason-shen/gopush/middlewares"
	"github.com/jason-shen/gopush/pkg/utils"
	"github.com/jason-shen/gopush/pkg/utils/generatecode"
	"github.com/jason-shen/gopush/pkg/utils/logger"
	"net/http"
)

type loginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (l loginRequest) validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Email, validation.Required, is.Email),
		validation.Field(&l.Password, validation.Required, validation.Length(6,12)),
	)
}

type registerRequest struct {
	Firstname 		string `json:"firstname"`
	Lastname 		string `json:"lastname"`
	Email 			string `json:"email"`
	Mobile 			string `json:"mobile"`
	Password 		string `json:"password"`
}

func (r registerRequest) validate() error  {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Firstname, validation.Required, validation.Length(2, 12)),
		validation.Field(&r.Lastname, validation.Required, validation.Length(2, 12)),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 12)),
	)
}

func (h *Handler) UserRegister(ctx *fiber.Ctx) error {
	var request registerRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		err = ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "Invalid Json",
		})
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		}
	}

	exist, _ := h.Client.User.Query().Where(user.Email(request.Email)).Only(ctx.Context())
	if exist != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "The Email is Already Taken!",
		})
		return nil
	}

	if err = request.validate(); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": err,
		})
		return nil
	}

	hashpassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return fmt.Errorf("failed hash user password: %v", err)
		return nil
	}

	_, err = h.Client.User.Create().
		SetEmail(request.Email).
		SetLastName(request.Lastname).
		SetFirstName(request.Firstname).
		SetPassword(hashpassword).
		SetActivateCode(generatecode.ActivationCode(4)).
		Save(ctx.Context())
	if err != nil {
		return fmt.Errorf("failed creating user: %v", err)
		return nil
	}


	_ = ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"error": false,
		"message": "registered successfully",
	})

	return nil
}

func (h *Handler) LoginUser(ctx *fiber.Ctx) error {
	var request loginRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		err = ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "Invalid Json",
		})
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		}
		return nil
	}

	if err = request.validate(); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": err,
		})
		return nil
	}

	u, err := h.Client.User.Query().Where(user.Email(request.Email)).Only(ctx.Context())
	// logger.Infof("user", u, err)
	if err != nil  {
		_ = ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "invalid credentials",
		})
		return nil
	}

	if u.Activated == false {
		_ = ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "you account is not activated!",
		})
		return nil
	}

	if u.Locked == true {
		_ = ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "you account is locked please contact the administrator!",
		})
		return nil
	}

	if err = utils.ComparePassword(request.Password, u.Password); err != nil {
		logger.Errorf("password compare error", err)
		_ = ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "invalid credentials",
		})
		return nil
	}

	token, err := middlewares.ClaimToken(u.ID)
	if err != nil {
		logger.Errorf("token generating error: ", err)
		return nil
	}

	// logger.Infof(token)

	_, err = h.Client.User.Update().Where(user.ID(u.ID)).SetJwttoken(token).Save(ctx.Context())

	if err != nil {
		logger.Errorf("failed to update jwttoken", err)
		return nil
	}

	response := map[string]interface{} {
		"firstname": u.FirstName,
		"lastname": u.LastName,
		"email": u.Email,
	}

	_ = ctx.Status(200).JSON(fiber.Map{
		"error": false,
		"data": response,
		"token": token,
		// "user":  user,
	})

	return nil
}
