package api

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"starOJ-backend/config"
	"starOJ-backend/model"
	"strconv"
)

func GetProblemList() gin.HandlerFunc {
	return func(c *gin.Context) {
		const PAGESIZE = 20
		type resp struct {
			ID         int32       `json:"ID"`
			Name       string      `json:"name"`
			Difficulty int32       `json:"difficulty"`
			CreatedBy  int32       `json:"createdBy"`
			Tags       []model.Tag `json:"tags" gorm:"many2many:problem_tag"`
		}

		page := c.Query("page")
		if len(page) == 0 {
			page = "0"
		}
		pageNumber, err := strconv.Atoi(page)
		if err != nil || pageNumber < 0 {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数错误"})
			return
		}

		var tmp []model.Problem
		db := config.GetDB()
		db.Model(model.Problem{}).Preload("Tags").Limit(PAGESIZE).Offset(PAGESIZE * pageNumber).Find(&tmp)
		var results []resp
		for i := 0; i < len(tmp); i++ {
			results = append(results, resp{tmp[i].ID, tmp[i].Name, tmp[i].Difficulty, tmp[i].CreatedBy, tmp[i].Tags})
		}
		resultsJSON, _ := json.Marshal(results)
		c.JSON(http.StatusOK, gin.H{"results": json.RawMessage(resultsJSON), "code": 200, "msg": "查询完成"})
	}
}

func GetProblem() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idNumber, err := strconv.Atoi(id)
		if err != nil || idNumber <= 0 {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数错误"})
			return
		}

		var results model.Problem
		db := config.GetDB()
		if db.Preload("Tags").Where("id = ?", idNumber).Take(&results).Error != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "记录不存在"})
			return
		}
		resultsJSON, _ := json.Marshal(results)
		c.JSON(http.StatusOK, gin.H{"results": json.RawMessage(resultsJSON), "code": 200, "msg": "查询完成"})
	}
}
