package api

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"starOJ-backend/config"
	"starOJ-backend/model"
	"strconv"
)

func GetContestList() gin.HandlerFunc {
	return func(c *gin.Context) {
		const PAGESIZE = 20

		page := c.Query("page")
		if len(page) == 0 {
			page = "0"
		}
		pageNumber, err := strconv.Atoi(page)
		if err != nil || pageNumber < 0 {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数错误"})
			return
		}

		db := config.GetDB()
		var results []model.Contest
		if db.Limit(PAGESIZE).Find(&results).Error != nil {
			c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "参数错误"})
			return
		}
		resultsJSON, _ := json.Marshal(results)
		c.JSON(http.StatusOK, gin.H{"results": json.RawMessage(resultsJSON), "code": 200, "msg": "查询完成"})
	}
}

func GetContest() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idNumber, err := strconv.Atoi(id)
		if err != nil || idNumber <= 0 {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数错误"})
			return
		}

		var results model.Contest
		db := config.GetDB()
		if db.Preload("Problems").Where("id = ?", idNumber).Take(&results).Error != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "记录不存在"})
			return
		}
		resultsJSON, _ := json.Marshal(results)
		c.JSON(http.StatusOK, gin.H{"results": json.RawMessage(resultsJSON), "code": 200, "msg": "查询完成"})
	}
}
