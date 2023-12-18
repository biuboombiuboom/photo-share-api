package pkg

import (
	"github.com/gin-gonic/gin"
	"photo.share/pkg/model"
	"photo.share/pkg/service"
)

func newMessage(c *gin.Context) {
	message := &model.Message{}
	if err := c.ShouldBindJSON(message); err == nil {
		err := service.NewMessage(c.Request.Context(), *message)
		if err != nil {
			c.String(500, err.Error())
		} else {
			c.String(200, "success")
		}
	} else {
		c.String(400, err.Error())
	}
}

func readMessages(c *gin.Context) {
	messageIds := []int64{}
	if err := c.ShouldBind(messageIds); err != nil {
		if err := service.ReadMessages(c.Request.Context(), messageIds); err != nil {
			if err != nil {
				c.String(500, err.Error())
			} else {
				c.String(200, "success")
			}
		}
	} else {
		c.String(400, err.Error())
	}
}

func readAllMessage(c *gin.Context) {
	var userId int64
	if user, ok := c.Get(claimsKey); !ok {
		c.String(401, "请登录")
		c.Abort()
	} else {
		userId = user.(Claims).Id
	}
	if err := service.ReadAllsMessages(c.Request.Context(), userId); err != nil {
		if err != nil {
			c.String(500, err.Error())
		} else {
			c.String(200, "success")
		}
	}
}
