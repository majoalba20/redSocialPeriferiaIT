package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/majoalba20/redSocialPeriferiaIT/cmd/initializers"
	"github.com/majoalba20/redSocialPeriferiaIT/cmd/internal/models"
)

func CreateProfile(c *gin.Context) {

	var req struct {
		Name      string    `json:"name" binding:"required"`
		LastName  string    `json:"last_name"`
		Alias     string    `json:"alias" binding:"required"`
		BirthDate time.Time `json:"birth_date"`
		Bio       string    `json:"bio"`
		AvatarURL string    `json:"avatar_url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	profile := models.Profile{
		UserID:    userID,
		Name:      req.Name,
		LastName:  req.LastName,
		Alias:     req.Alias,
		BirthDate: req.BirthDate,
		Bio:       req.Bio,
		AvatarURL: req.AvatarURL,
	}

	if err := initializers.DB.Create(&profile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, profile)
}

func GetProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	var profile models.Profile
	err := initializers.DB.Where("user_id = ?", userID).First(&profile).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "profile not found"})
		return
	}

	c.JSON(200, profile)
}
