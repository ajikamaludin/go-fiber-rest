package auth

import (
	"github.com/ajikamaludin/go-fiber-rest/app/models"
	userRepository "github.com/ajikamaludin/go-fiber-rest/app/repository/user"
	"github.com/ajikamaludin/go-fiber-rest/pkg/jwtmanager"
	"github.com/ajikamaludin/go-fiber-rest/pkg/utils/constants"
	"github.com/ajikamaludin/go-fiber-rest/pkg/utils/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	userRequest := new(models.UserReq)

	_ = c.BodyParser(&userRequest)

	errors := validator.ValidateRequest(userRequest)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "request body invalid",
			"error":   errors,
		})
	}

	user := &models.User{}
	userRepository.GetUserByEmail(userRequest.Email, user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "credentials invalid",
			"error":   err.Error(),
		})
	}

	accessToken := jwtmanager.CreateToken(user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":         user.ToUserRes(),
		"accessToken":  accessToken,
		"refreshToken": "",
	})
}

func Register(c *fiber.Ctx) error {
	userRequest := new(models.UserReq)

	_ = c.BodyParser(&userRequest)

	errors := validator.ValidateRequest(userRequest)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "request body invalid",
			"error":   errors,
		})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	user := &models.User{
		Email:    userRequest.Email,
		Password: string(hashedPassword),
	}
	err := userRepository.GetUserByEmail(userRequest.Email, user)
	if err == nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "email is already exists",
			"error":   "",
		})
	}

	userRepository.CreateUser(user)

	accessToken := jwtmanager.CreateToken(user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":         user.ToUserRes(),
		"accessToken":  accessToken,
		"refreshToken": "",
	})
}
