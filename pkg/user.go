package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"photo.share/pkg/model"
	"photo.share/pkg/service"
)

var secretKey = "20231208"
var claimsKey = "claims"

type Claims struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func updatePassword(c *gin.Context) {
	if userId, found := getUserId(c); found {
		password := model.PasswordUpdate{}
		if err := c.ShouldBindJSON(&password); err == nil {
			if err := service.UpdatePassword(c.Request.Context(), userId, password.OldPassword, password.NewPassword); err != nil {
				c.String(500, err.Error())
			} else {
				c.String(200, "success")
			}
		} else {
			c.String(500, err.Error())
		}
	}
}

func updateSetting(c *gin.Context) {
	userSetting := model.UserSetting{}
	if err := c.ShouldBindJSON(&userSetting); err == nil {
		if err := service.UpdateUserSetting(c.Request.Context(), userSetting); err != nil {
			c.String(500, err.Error())
		} else {
			c.String(200, "success")
		}
	} else {
		c.String(500, err.Error())
	}
}

func getUserInfo(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"result": gin.H{"username": "aaa", "role": gin.H{"permissions": []gin.H{{
		"roleId": "admin", "permissionId": "support", "permissionName": "",
	}}}}})
}

func logout(c *gin.Context) {
	c.String(200, "logout")
}

func login(c *gin.Context) {
	loginInfo := model.LoginInfo{}
	if c.ShouldBindJSON(&loginInfo) == nil {
		if user, err := service.Login(c.Request.Context(), loginInfo.Login, loginInfo.Password); err == nil {
			token, err := createToken(user.Id, user.UserName)
			if err != nil {
				c.JSON(500, gin.H{"error": "Internal Server Error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"result": gin.H{"token": token, "user": user}})
		} else {
			c.String(500, err.Error())
		}
	} else {
		c.String(400, "非法的登录信息")
	}
}

func register(c *gin.Context) {
	user := model.User{}
	if c.ShouldBindJSON(&user) == nil {
		if userId, err := service.NewUser(c.Request.Context(), user); err != nil {
			c.String(500, err.Error())
		} else {
			c.JSON(200, userId)
		}
	} else {
		c.String(400, "非法的注册信息")
	}
}

func authMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Access-Token")
	if tokenString == "" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not found auth header"})
		c.Abort()
		return
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized invalid token"})
		c.Abort()
		return
	}

	// Pass the claims to the next handler
	c.Set(claimsKey, *claims)
}

func createToken(userid int64, username string) (string, error) {
	claims := Claims{
		Username: username,
		Id:       userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 0,
			Issuer:    "202312071026",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
