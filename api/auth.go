package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"starOJ-backend/config"
	"starOJ-backend/middleware/HmacSha256"
	"starOJ-backend/middleware/jwt"
	"starOJ-backend/model"
	"strconv"
	"time"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		type loginReq struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Captcha  string `json:"captcha"`
		}
		r := loginReq{}
		c.BindJSON(&r)

		db, rd := config.GetDB(), config.GetRD()
		user := model.User{}
		err := db.Where("username = ?", r.Username).Take(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "用户名不存在"})
		} else if user.Password != HmacSha256.HmacSha256ToBase64(r.Password) {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "密码错误"})
		} else {
			token, _ := jwt.GenerateToken(int64(user.ID), 7*time.Hour)
			rd.HSet("token", strconv.FormatInt(int64(user.ID), 10), token)
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "登录成功", "token": token})
		}
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		type registerReq struct {
			Username string `json:"username"`
			Nickname string `json:"nickname"`
			Password string `json:"password"`
			Captcha  string `json:"captcha"`
		}
		r := registerReq{}
		c.BindJSON(&r)

		db := config.GetDB()
		if err := db.Where("username = ?", r.Username).Take(&model.User{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(200, gin.H{"code": 0, "msg": "用户名已存在"})
		} else if err := db.Where("nickname = ?", r.Nickname).Take(&model.User{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(200, gin.H{"code": 0, "msg": "昵称已存在"})
		} else {
			c.JSON(200, gin.H{"code": 1, "msg": "注册成功"})
			db.Create(&model.User{Username: r.Username, Nickname: r.Nickname, Password: HmacSha256.HmacSha256ToBase64(r.Password)})
		}
	}
}
