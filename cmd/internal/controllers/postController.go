package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/majoalba20/redSocialPeriferiaIT/cmd/initializers"
	"github.com/majoalba20/redSocialPeriferiaIT/cmd/internal/models"
	"gorm.io/gorm"
)

func CreatePost(c *gin.Context) {
	var req struct {
		Message string `json:"message" binding:"required,min=1,max=500"`
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

	var profile models.Profile
	if err := initializers.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "profile not found"})
		return
	}

	post := models.Post{
		Message:   req.Message,
		ProfileID: profile.ID,
	}

	if err := initializers.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create post"})
		return
	}

	c.JSON(http.StatusCreated, post)

}

func GetPostsFeed(c *gin.Context) {
	var posts []struct {
		UserID    string `json:"user_id"`
		Alias     string `json:"alias"`
		Message   string `json:"message"`
		Likes     int    `json:"likes"`
		AvatarURL string `json:"avatar_url"`
	}

	query := `
		SELECT 
			pr.user_id,
			pr.alias,
			p.message,
			p.likes,
			pr.avatar_url
		FROM posts p
		JOIN profiles pr ON pr.id = p.profile_id
		ORDER BY p.created_at DESC
	`

	if err := initializers.DB.Raw(query).Scan(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func LikePost(c *gin.Context) {
	postIDParam := c.Param("id")

	postID, err := strconv.ParseUint(postIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	result := initializers.DB.Model(&models.Post{}).
		Where("id = ?", postID).
		Update("likes", gorm.Expr("likes + 1"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post liked"})
}
