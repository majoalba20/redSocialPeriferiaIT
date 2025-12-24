package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/majoalba20/redSocialPeriferiaIT/cmd/initializers"
	"github.com/majoalba20/redSocialPeriferiaIT/cmd/internal/controllers"
	"github.com/majoalba20/redSocialPeriferiaIT/cmd/internal/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/profile", middleware.RequireAuth, controllers.CreateProfile)
	r.GET("/profile", middleware.RequireAuth, controllers.GetProfile)
	r.POST("/post", middleware.RequireAuth, controllers.CreatePost)
	r.GET("/get/all", middleware.RequireAuth, controllers.GetPostsFeed)
	r.POST("/posts/:id/like", middleware.RequireAuth, controllers.LikePost)
	r.Run()
}
