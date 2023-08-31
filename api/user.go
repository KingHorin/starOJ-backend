package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"starOJ-backend/config"
	"starOJ-backend/model"
)

func PostNickname() gin.HandlerFunc {
	return func(c *gin.Context) {
		type req struct {
			Username string `json:"username" validate:"required"`
			Nickname string `json:"nickname" validate:"required"`
		}
		r := req{}
		c.BindJSON(&r)

		validate := validator.New()
		if err := validate.Struct(&r); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数错误"})
			return
		}

		db := config.GetDB()
		user := model.User{}
		db.Where("username = ?", r.Username).Take(&user)
		tokenUserID, ok := c.Get("userID")
		if !ok || user.ID != tokenUserID {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "非法操作，账号不对应"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "昵称更新成功"})
			db.Model(&model.User{}).Where("username = ?", r.Username).Update("nickname", r.Nickname)
		}
	}
}

func PostAvatar() gin.HandlerFunc {
	return func(c *gin.Context) {
		type req struct {
			Username string `json:"username" validate:"required"`
			Avatar   string `json:"avatar" validate:"required"`
		}
		r := req{}
		c.BindJSON(&r)

		validate := validator.New()
		if err := validate.Struct(&r); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数错误"})
			return
		}

		db := config.GetDB()
		user := model.User{}
		db.Where("username = ?", r.Username).Take(&user)
		tokenUserID, ok := c.Get("userID")
		if !ok || user.ID != tokenUserID {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "非法操作，账号不对应"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "头像更新成功"})
			db.Model(&model.User{}).Where("username = ?", r.Username).Update("avatar", r.Avatar)
		}
	}
}

func PostPhone() gin.HandlerFunc {
	return func(c *gin.Context) {
		type req struct {
			Username string `json:"username" validate:"required"`
			Phone    string `json:"phone" validate:"required"`
		}
		r := req{}
		c.BindJSON(&r)

		validate := validator.New()
		if err := validate.Struct(&r); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数错误"})
			return
		}

		db := config.GetDB()
		user := model.User{}
		db.Where("username = ?", r.Username).Take(&user)
		tokenUserID, ok := c.Get("userID")
		if !ok || user.ID != tokenUserID {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "非法操作，账号不对应"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "手机号更新成功"})
			db.Model(&model.User{}).Where("username = ?", r.Username).Update("phone", r.Phone)
		}
	}
}

func PostEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		type req struct {
			Username string `json:"username"`
			Email    string `json:"email"`
		}
		r := req{}
		c.BindJSON(&r)

		validate := validator.New()
		if err := validate.Struct(&r); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数错误"})
			return
		}

		db := config.GetDB()
		user := model.User{}
		db.Where("username = ?", r.Username).Take(&user)
		tokenUserID, ok := c.Get("userID")
		if !ok || user.ID != tokenUserID {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "非法操作，账号不对应"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "邮箱更新成功"})
			db.Model(&model.User{}).Where("username = ?", r.Username).Update("email", r.Email)
		}
	}
}
