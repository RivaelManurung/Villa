package middleware

import (
	"strings"
	"time"

	database "Backend/database"
	"Backend/models/entity"
	utils "Backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func checkLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("jwt")
		if cookie == "" {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		tokenString := strings.Replace(cookie, "jwt", "", 1)
		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(utils.Secret_Key), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		// Refresh token expiration
		claims.ExpiresAt = time.Now().Add(time.Hour * 1).Unix()

		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err = newToken.SignedString([]byte(utils.Secret_Key))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error refreshing token",
			})
		}

		// Set the new token as a cookie
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    tokenString,
			Expires:  time.Now().Add(time.Hour * 1),
			HTTPOnly: true,
		})

		var admin entity.Admin
		if err := database.DB.Where("username = ?", claims.Issuer).First(&admin).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Admin not found",
			})
		}
		c.Locals("admin", admin)

		return c.Next()
	}
}
