package endpoints

import (
	"TheBestThing/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func ThingRequest(c *gin.Context, db *gorm.DB) {
	var thing models.Thing
	if err := c.ShouldBindJSON(&thing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
