package main

import (
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
