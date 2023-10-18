package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"puem/puem/api"
	"puem/puem/types"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Set up the web server
	app := fiber.New()

	// Set up the database
	db, err := initializeDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	// Set up middleware
	setupMiddlewares(app)

	// Set up routes
	api.SetupRoutesAPI(app)
	api.SetupRoutesDatabase(app, db)

	runServer(app, db)
}

func setupMiddlewares(app *fiber.App) {
	app.Use(logger.New())
	// app.Use(func(c *fiber.Ctx) error {
	// 	// log.Println(c.Method(), c.Path())
	// 	return c.Next()
	// })
}

func initializeDatabase() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME are required")
		log.Fatal("Error loading .env file")
	}

	Database := types.DB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Database.User,
		Database.Password,
		Database.Host,
		Database.Port,
		Database.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func runServer(app *fiber.App, db *gorm.DB) {
	go func() {
		fmt.Println("Listening on http://localhost:3000")
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt or terminate signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	// Gracefully shutdown the app
	if err := app.Shutdown(); err != nil {
		log.Printf("Error shutting down server: %v", err)
	}

	log.Println("Server exited")
}
