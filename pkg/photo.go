package pkg

import (
	"fmt"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"photo.share/pkg/model"
	"photo.share/pkg/service"
)

var phtotPath = "d:/"

func getPhotosByUserId(c *gin.Context) {

}

func newPhoto(c *gin.Context) {
	var userId int64
	if user, ok := c.Get(claimsKey); !ok {
		c.String(500, "请登录")
		c.Abort()
	} else {
		userId = user.(*Claims).Id
	}

	photo, _ := c.FormFile("photo")

	dst := path.Join(phtotPath, fmt.Sprintf("%d", userId), photo.Filename)
	err := c.SaveUploadedFile(photo, dst)

	desc, _ := c.GetPostForm("desc")
	title, _ := c.GetPostForm("title")
	isPublicStr, _ := c.GetPostForm("isPublic")
	ispublic, _ := strconv.ParseBool(isPublicStr)
	if err != nil {
		photoInfo := model.Photo{
			Path:        dst,
			UserId:      userId,
			CreatedAt:   time.Now(),
			Star:        0,
			Description: desc,
			Title:       title,
			IsPublic:    ispublic,
		}
		service.NewPhoto(c.Request.Context(), photoInfo)

	}

}
