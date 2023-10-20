package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"puem/puem/api"
	"puem/puem/databases/connection"
	"puem/puem/databases/models"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func main() {
	// Set up the web server
	app := fiber.New()

	// Set up the database
	db, err := connection.GetDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	db.AutoMigrate(
		&models.ProfileData{},
		&models.Projects{},
		&models.Images{},
		&models.Technologies{},
		&models.Skills{},
		&models.Github{},
		&models.Website{},
		&models.Experience{},
		&models.Certificate{},
		&models.EducationData{},
		&models.Languages{},
		&models.Interests{},
	)

	// Set up middleware
	setupMiddlewares(app)

	// Set up routes
	api.SetupRoutesDatabase(app, db)

	runServer(app, db)
}

func setupMiddlewares(app *fiber.App) {
	app.Use(logger.New())
}

func runServer(app *fiber.App, db *gorm.DB) {
	go func() {
		fmt.Println("Listening on http://localhost:3000")
		if err := app.Listen(getPort()); err != nil {
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
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}
