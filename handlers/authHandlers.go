package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/Tirivashe/go-fiber-jwt/auth"
	"github.com/Tirivashe/go-fiber-jwt/db"
	"github.com/Tirivashe/go-fiber-jwt/sqlc"
	"github.com/Tirivashe/go-fiber-jwt/types"
	"github.com/Tirivashe/go-fiber-jwt/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var userDto types.UserLoginDto
	if err := c.BodyParser(&userDto); err != nil {
		return err
	}

	queries := db.GetQueries()
	user, err := queries.GetUserByEmail(context.Background(), userDto.Email)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Credentials"})
	}
	if user.ID == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid credentials"})
	}
	if err = utils.CheckPassword(userDto.Password, user.Password); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid credentials"})
	}
	token, err := auth.CreateToken(userDto.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"token_error": err})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "access_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 24)
	c.Cookie(cookie)
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "login success"})

}

func SignUp(c *fiber.Ctx) error {
	var userDto types.UserDto
	if err := c.BodyParser(&userDto); err != nil {
		return err
	}

	queries := db.GetQueries()
	_, err := queries.GetUserByEmail(context.Background(), userDto.Email)
	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "already exists"})
	}
	hashedPassword, err := utils.HashPassword(userDto.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"hash_message": err})
	}

	err = queries.CreateUser(context.Background(), sqlc.CreateUserParams{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: hashedPassword,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"save_message": err})
	}

	token, err := auth.CreateToken(userDto.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"t_error": err})
	}
	cookie := new(fiber.Cookie)
	cookie.Name = "access_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 24)
	c.Cookie(cookie)
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "created"})
}

func HealthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}
