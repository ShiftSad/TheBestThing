package endpoints

import (
	"TheBestThing/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"net/http"
)

func ThingRequest(c *gin.Context, db *gorm.DB) {
	var thing models.Thing
	if err := c.ShouldBindJSON(&thing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	thing.ID = 0

	if err := db.Create(&thing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create thing"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": thing.ID})
}

func ImageRequest(c *gin.Context, db *gorm.DB) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image uploaded"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer f.Close()

	imageData, err := io.ReadAll(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	image := models.Image{
		Name: file.Filename,
		Data: imageData,
	}

	if err := db.Create(&image).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": image.ID})
}
