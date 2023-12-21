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
	auth.POST("/logout", logout)
	auth.POST("/register", register)

	api.GET("/user/info", authMiddleware, getUserInfo)
	api.GET("/user/:id/photos", authMiddleware, getPhotosByUserId)

	api.POST("/photos", getAllPhotos)
	api.GET("photo/:id", getPhoto)
	api.GET("photo/:id/comments", getPhotoComments)
	api.POST("/photo/:id/comment", authMiddleware, newPhotoComment)
	api.POST("/photo", authMiddleware, newPhoto)
	api.POST("/photo/upload", authMiddleware, uploadPhoto)
	api.DELETE("/photo/:id", authMiddleware, deletePhoto)
	api.POST("/photo/:id/star", authMiddleware, starPhoto)
	api.POST("/photo/:id/collect", authMiddleware, collectPhoto)

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

func getUserIdAndName(c *gin.Context) (int64, string, bool) {
	found := false
	var userId int64
	var username string
	if user, ok := c.Get(claimsKey); !ok {
		c.String(500, "请登录")
		c.Abort()
	} else {
		userId = user.(Claims).Id
		username = user.(Claims).Username
		found = true
	}
	return userId, username, found

}
