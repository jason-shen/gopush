package middlewares

import (
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/google/uuid"
	"github.com/jason-shen/gopush/config"
	"github.com/jason-shen/gopush/pkg/utils/logger"
	"strconv"
	"time"
)

func IsAuthenticated(config *config.Config) func(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Jwt.Secret),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error{
			ctx.Status(fiber.StatusUnauthorized).Send([]byte("Unauthorized"))
			return err
		},
	})
}

func GetUserIdFromContext(ctx *fiber.Ctx) (int64, error) {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	strId := claims["sub"].(string)
	id, err := strconv.Atoi(strId)
	if err != nil {
		return 0, fmt.Errorf("error while parsing the token id: %w", err)
	}

	return int64(id), nil
}

func ClaimToken(id uuid.UUID) (string, error) {
	 config := config.New()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30) // a month

	s, err := token.SignedString(config.Jwt.Secret)
	if err != nil {
		logger.Errorf("error", err)
		return "", err
	}
	return s, nil
}
