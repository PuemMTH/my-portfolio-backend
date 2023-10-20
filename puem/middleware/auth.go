package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Create Middleware to check if the user is authorized to access the API
func Auth(c *fiber.Ctx) (err error) {
	// if c.Get("Authorization") != "" {
	// 	log.Println("Unauthorized access")
	// 	c.Status(401).JSON(fiber.Map{
	// 		"message": "Unauthorized access",
	// 	})
	// }

	// ใน Header จะมี Authorization อยู่ ถ้าไม่มีจะเข้าเงื่อนไข
	if c.Get("Authorization") == "" {
		log.Println("Unauthorized access")
		c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized access",
		})
		return nil
	}

	// ถ้ามี Authorization อยู่ใน Header จะเข้าเงื่อนไขนี้
	if c.Get("Authorization") != "" {
		log.Println("Authorized access")
		return c.Next()
	}

	return nil
}
