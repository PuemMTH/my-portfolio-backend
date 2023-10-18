package api

import (
	"puem/puem/databases/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutesDatabase(app *fiber.App, db *gorm.DB) {
	app.Get("api/list", func(c *fiber.Ctx) error {
		var profiles []models.ProfileData                                    // คือ ตัวแปรที่เก็บข้อมูลที่ได้จากฐานข้อมูล
		if err := db.Preload("Websites").Find(&profiles).Error; err != nil { // คือ การดึงข้อมูลจากฐานข้อมูล และเก็บไว้ในตัวแปร profiles และ ถ้าเกิด error ให้แสดงข้อความ error
			return c.JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(profiles) // คือ การแสดงข้อมูลในรูปแบบ json
	})

	app.Get("api/add", func(c *fiber.Ctx) error {
		profile := models.ProfileData{
			Name:     "John Doe",
			Nickname: "johndoe",
			Phone:    "123-456-7890",
			Email:    "john@example.com",
			Github:   "https://github.com/johndoe",
			Linkedin: "https://linkedin.com/in/johndoe",
			Location: "New York",
		}
		websites := []models.Website{
			{Website: "https://website1.com", ProfileDataID: profile.ID},
			{Website: "https://website2.com", ProfileDataID: profile.ID},
		}
		profile.Websites = websites

		db.Create(&profile)

		return c.JSON(profile)
	})
}
