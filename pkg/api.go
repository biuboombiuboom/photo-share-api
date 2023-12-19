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

	api.POST("/photos", getAllPhotos)
	api.POST("/photo", authMiddleware, newPhoto)
	api.POST("/photo/upload", authMiddleware, uploadPhoto)
	api.DELETE("/photo/:id", authMiddleware, deletePhoto)

	api.POST("/message", authMiddleware, newMessage)
	r.Run()
}

func getUserId(c *gin.Context) (int64, bool) {
	var userId int64
	found := false
	if user, ok := c.Get(claimsKey); !ok {
		c.String(500, "请登录")
		c.Abort()
	} else {
		userId = user.(Claims).Id
		found = true
	}
	return userId, found
}
