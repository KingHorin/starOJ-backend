package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/nu7hatch/gouuid"
	"net/http"
	"path"
	"starOJ-backend/config"
	"starOJ-backend/model"
	"strconv"
)

func GetInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		type resp struct {
			ID       int32  `json:"id"`
			Nickname string `json:"nickname"`
			Avatar   string `json:"avatar"`
		}
		id := c.Query("id")
		idNumber, err := strconv.Atoi(id)
		if err != nil || len(id) == 0 {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数错误"})
			return
		}

		var results resp
		db := config.GetDB()
		if db.Model(&model.User{}).Where("id = ?", idNumber).Take(&results).Error != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数错误"})
			return
		}
		resultsJSON, _ := json.Marshal(results)
		c.JSON(http.StatusOK, gin.H{"results": json.RawMessage(resultsJSON), "code": 1, "msg": "查询完成"})
	}
}

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
		extName := path.Ext(file.Filename) //获取后缀名
		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".jpeg": true,
		}
		if _, ok := allowExtMap[extName]; !ok {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "文件类型错误"})
			return
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
