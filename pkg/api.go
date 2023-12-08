package pkg

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	corsMiddleWare := cors.Default()
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(corsMiddleWare)
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "health is ok")
	})

	api := r.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/login", login)
	auth.POST("/register", register)

	api.GET("/user/info", authMiddleware, getUserInfo)
	api.GET("/user/:id/photos", authMiddleware, getPhotosByUserId)
	api.POST("/photo", authMiddleware, newPhoto)
	r.Run()
}
