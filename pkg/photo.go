package pkg

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"photo.share/pkg/model"
	"photo.share/pkg/service"
)

// var phtotPath = "d:\\"

var viewPath = "/photo"

var photoPath = "G:/workspace/node/photo-share-web/public/photo"

// var photoPath = "E:/workspace/node/photo-share/public/photo"

func newPhotoComment(c *gin.Context) {
	if userId, userName, f := getUserIdAndName(c); f {
		comment := model.PhotoComment{}
		if err := c.ShouldBind(&comment); err != nil {
			c.String(400, "bad request:"+err.Error())
			c.Abort()
		} else {
			comment.UserId = userId
			comment.UserName = userName
			comment.CreatedAt = time.Now()
			if err := service.NewComment(c.Request.Context(), comment); err != nil {
				c.String(500, err.Error())
			} else {
				c.String(200, "success")
			}
		}
	}
}

func getPhotoComments(c *gin.Context) {
	if pidStr, got := c.Params.Get("id"); !got {
		c.String(500, "miss photo id")
	} else {
		if photoId, err := strconv.ParseInt(pidStr, 0, 64); err != nil {
			c.String(500, fmt.Sprintf("invaild photoid %s", pidStr))
		} else {
			comments, err := service.GetComments(c.Request.Context(), photoId)
			if err != nil {
				c.String(500, err.Error())
			} else {
				c.JSON(200, comments)
			}
		}
	}
}

func getPhoto(c *gin.Context) {
	if pidStr, got := c.Params.Get("id"); !got {
		c.String(500, "miss photo id")
	} else {
		if photoId, err := strconv.ParseInt(pidStr, 0, 64); err != nil {
			c.String(500, fmt.Sprintf("invaild photoid %s", pidStr))
		} else {
			photo, err := service.GetPhoto(c.Request.Context(), photoId)
			if err != nil {
				c.String(500, err.Error())
			} else {
				c.JSON(200, photo)
			}
		}
	}
}

func collectPhoto(c *gin.Context) {
	if userId, username, found := getUserIdAndName(c); !found {
		return
	} else {
		if pidStr, got := c.Params.Get("id"); !got {
			c.String(500, "miss photo id")
		} else {
			photoId, _ := strconv.ParseInt(pidStr, 0, 64)
			photo_collect := model.PhotoCollect{
				PhotoId:   photoId,
				UserId:    userId,
				UserName:  username,
				CreatedAt: time.Now(),
			}
			if rows, err := service.CollectPhoto(c.Request.Context(), photo_collect); err != nil {
				c.String(500, err.Error())
			} else {
				c.JSON(200, gin.H{"message": "success", "rows": rows})
			}
		}
	}
}

func starPhoto(c *gin.Context) {
	if userId, username, found := getUserIdAndName(c); !found {
		return
	} else {
		if pidStr, got := c.Params.Get("id"); !got {
			c.String(500, "miss photo id")
		} else {
			photoId, _ := strconv.ParseInt(pidStr, 0, 64)
			photoStar := model.PhotoStar{
				PhotoId:   photoId,
				UserId:    userId,
				UserName:  username,
				CreatedAt: time.Now(),
			}
			if rows, err := service.StarPhoto(c.Request.Context(), photoStar); err != nil {
				c.String(500, err.Error())
			} else {
				c.JSON(200, gin.H{"message": "success", "rows": rows})
			}
		}
	}
}

func getAllPhotos(c *gin.Context) {
	query := model.PageQuery{
		OrderBy:  "created_at",
		Page:     1,
		PageSize: 6,
	}
	c.ShouldBind(&query)
	if photos, total, err := service.GetPublicPhotos(c.Request.Context(), query.OrderBy, (query.Page-1)*query.PageSize, query.PageSize); err != nil {
		c.String(500, err.Error())
	} else {
		c.JSON(200, gin.H{"total": total, "data": photos})
	}
}

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
		c.String(401, "请登录")
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
	newPhotoName := uuid.NewString() + path.Ext(photo.Filename)
	dst := path.Join(photoPath, fmt.Sprintf("%d", userId), newPhotoName)
	htmpPath := path.Join(viewPath, fmt.Sprintf("%d", userId), newPhotoName)
	err := c.SaveUploadedFile(photo, dst)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	desc, _ := c.GetPostForm("desc")
	fmt.Println(desc)
	title, _ := c.GetPostForm("title")
	isPublicStr, _ := c.GetPostForm("isPublic")
	ispublic, _ := strconv.ParseBool(isPublicStr)
	if err == nil {
		photoInfo := model.Photo{
			Path:        htmpPath,
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

func deletePhoto(c *gin.Context) {
	if photoIdStr, got := c.Params.Get("id"); got {
		photoId, _ := strconv.ParseInt(photoIdStr, 0, 64)
		if err := service.DeletePhoto(c.Request.Context(), photoId); err != nil {
			if err != nil {
				c.String(500, err.Error())
			} else {
				c.String(200, "success")
			}
		}
	} else {
		c.String(500, "missing phtoto id")
	}

}
