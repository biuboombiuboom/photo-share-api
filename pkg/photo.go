package pkg

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"photo.share/pkg/model"
	"photo.share/pkg/service"
)

// var phtotPath = "d:\\"
var photoPath = "E:\\workspace\\node\\photo-share\\src\\assets\\photo"

func getPhotosByUserId(c *gin.Context) {
	if userId, found := getUserId(c); found {
		photos, err := service.GetPhotosByUserId(c.Request.Context(), userId)
		if err != nil {
			c.String(500, err.Error())
		} else {
			c.JSON(200, photos)
		}
	}
}

func uploadPhoto(c *gin.Context) {
	var userId int64
	if user, ok := c.Get(claimsKey); !ok {
		c.String(500, "请登录")
		c.Abort()
	} else {
		userId = user.(Claims).Id
	}

	photo, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{"result": err})
		c.Abort()
	}

	dst := path.Join(photoPath, fmt.Sprintf("%d", userId), photo.Filename)
	err = c.SaveUploadedFile(photo, dst)
	if err != nil {
		c.JSON(500, gin.H{"result": err})
	}
	c.JSON(200, gin.H{"result": dst})
}

func newPhoto(c *gin.Context) {
	var userId int64
	if user, ok := c.Get(claimsKey); !ok {
		c.String(500, "请登录")
		c.Abort()
		return
	} else {
		userId = user.(Claims).Id
	}

	photo, _ := c.FormFile("file")

	dst := path.Join(photoPath, fmt.Sprintf("%d", userId), photo.Filename)
	err := c.SaveUploadedFile(photo, dst)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	desc, _ := c.GetPostForm("desc")
	title, _ := c.GetPostForm("title")
	isPublicStr, _ := c.GetPostForm("isPublic")
	ispublic, _ := strconv.ParseBool(isPublicStr)
	if err == nil {
		photoInfo := model.Photo{
			Path:        dst,
			UserId:      userId,
			CreatedAt:   time.Now(),
			Description: desc,
			Title:       title,
			IsPublic:    ispublic,
		}
		_, err := service.NewPhoto(c.Request.Context(), photoInfo)
		if err == nil {
			c.String(200, "success")
		} else {
			os.Remove(dst)
			c.String(500, err.Error())
		}

	}

}
