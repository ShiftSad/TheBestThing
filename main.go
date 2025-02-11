package main

import (
	"TheBestThing/endpoints"
	"TheBestThing/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func main() {
	// Connect to database
	dsn := os.Getenv("DSN")

	log.Println("Connecting to database")

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database with error: ", err)
	}

	err = db.AutoMigrate(&models.Thing{}, &models.Image{})
	if err != nil {
		log.Fatal("Failed to migrate schema with error: ", err)
	}

	// Initialize the router
	r := gin.Default()

	r.POST("/admin/thing", func(c *gin.Context) { endpoints.ThingRequest(c, db) })
	r.POST("/admin/image", func(c *gin.Context) { endpoints.ImageRequest(c, db) })

	log.Println("Server started, running on port 8080")
	log.Fatal(r.Run("0.0.0.0:8080"))
}
