package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/smugg99/coding-night-2023/config"
	"github.com/smugg99/coding-night-2023/models"
	"github.com/smugg99/coding-night-2023/services"
)

func (base *Controller) Login(c *fiber.Ctx) error {
    var user models.UserLoginDTO
    if err := c.BodyParser(&user); err != nil {
        return err
    }

    authenticated, dbUser := services.GetUserByCredentials(base.Db, &user)
    if authenticated {
        claims := jwt.MapClaims{
            "email": dbUser.Password,
            "exp": time.Now().Add(time.Hour * 72).Unix(), 
        }

        token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

        t, err := token.SignedString(config.PrivateKey)
        if err != nil {
            return c.SendStatus(fiber.StatusInternalServerError)
        }

        return c.JSON(fiber.Map{"token": t})
    }

    return c.JSON(fiber.StatusUnauthorized)
}

func (base *Controller) Register(c *fiber.Ctx) error {
    var newUser models.UserRegisterDTO
    if err := c.BodyParser(&newUser); err != nil {
        return err
    }

    pass, err := services.HashPassword(newUser.Password)
    if err != nil {
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    user := models.User{
        Email: newUser.Email,
        Password: pass,
        Name: newUser.Name,
        Phone: newUser.Phone,
        Address: newUser.Address,
    }
    

    if err := services.CreateNewUser(base.Db, &user); err != nil {
        return c.SendStatus(fiber.StatusConflict)
    }

    return c.SendStatus(fiber.StatusCreated)
} 
