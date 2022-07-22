package auth

import (
	"github.com/ajikamaludin/go-fiber-rest/app/models"
	userRepository "github.com/ajikamaludin/go-fiber-rest/app/repository/user"
	"github.com/ajikamaludin/go-fiber-rest/pkg/utils/constants"
	"github.com/ajikamaludin/go-fiber-rest/pkg/utils/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	userRequest := new(models.User)

	_ = c.BodyParser(&userRequest)

	errors := validator.ValidateRequest(userRequest)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "request body mismatch",
			"error":   errors,
		})
	}

	user := &models.User{}
	userRepository.GetUserByEmail(userRequest.Email, user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "credentials mismatch",
			"error":   err.Error(),
		})
	}
	// generate jwt token
	// return token and user detail
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":         user,
		"accessToken":  "",
		"refreshToken": "",
	})
}
