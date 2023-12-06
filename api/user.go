package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nu7hatch/gouuid"
	"net/http"
	"path"
	"starOJ-backend/config"
	"starOJ-backend/model"
)

func PostNickname() gin.HandlerFunc {
	return func(c *gin.Context) {
		type req struct {
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
		userID, _ := c.Get("userID")
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "昵称更新成功"})
		db.Model(&model.User{}).Where("id = ?", userID).Update("nickname", r.Nickname)
	}
}

func PostAvatar() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "上传错误"})
		}
		filename, _ := uuid.NewV4()
		dest := path.Join("./user/avatar", filename.String()+".jpg")
		c.SaveUploadedFile(file, dest)

		db := config.GetDB()
		userID, _ := c.Get("userID")
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "头像更新成功"})
		db.Model(&model.User{}).Where("id = ?", userID).Update("avatar", dest)
	}
}

func PostPhone() gin.HandlerFunc {
	return func(c *gin.Context) {
		type req struct {
			Phone string `json:"phone" validate:"required"`
		}
		r := req{}
		c.BindJSON(&r)

		validate := validator.New()
		if err := validate.Struct(&r); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数错误"})
			return
		}

		db := config.GetDB()
		userID, _ := c.Get("userID")
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "手机号更新成功"})
		db.Model(&model.User{}).Where("id = ?", userID).Update("phone", r.Phone)
	}
}

func PostEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		type req struct {
			Email string `json:"email"`
		}
		r := req{}
		c.BindJSON(&r)

		validate := validator.New()
		if err := validate.Struct(&r); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数错误"})
			return
		}

		db := config.GetDB()
		userID, _ := c.Get("userID")
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "邮箱更新成功"})
		db.Model(&model.User{}).Where("id = ?", userID).Update("email", r.Email)
	}
}
