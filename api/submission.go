package api

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"starOJ-backend/config"
	"starOJ-backend/model"
	"strconv"
	"time"
)

type resp struct {
	ID        int32     `json:"ID"`
	ProblemID int32     `json:"problemID"`
	UserID    int32     `json:"userID"`
	Status    int32     `json:"status"`
	Language  int32     `json:"language"`
	Time      int32     `json:"time"`
	Memory    int32     `json:"memory"`
	CreatedAt time.Time `json:"createdAt"`
}

func GetSubmissionList() gin.HandlerFunc {
	return func(c *gin.Context) {
		const PAGESIZE = 20
		problemID := c.Query("problemid")
		userID := c.Query("userid")
		status := c.Query("status")
		language := c.Query("language")
		page := c.Query("page")
		order := c.Query("order")

		db := config.GetDB().Model(&model.Submission{})
		if len(problemID) != 0 {
			db = db.Where("problemid = ?", problemID)
		}
		if len(userID) != 0 {
			db = db.Where("userid = ?", userID)
		}
		if len(status) != 0 {
			db = db.Where("status = ?", status)
		}
		if len(language) != 0 {
			db = db.Where("language = ?", language)
		}
		if len(page) != 0 {
			pageNumber, err := strconv.Atoi(page)
			if err != nil || pageNumber < 0 {
				c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数错误"})
				return
			}
			db = db.Offset(PAGESIZE * pageNumber)
		}
		if len(order) != 0 {
			db = db.Order(order + " asc")
		}
		db = db.Order("created_at desc")

		var results []resp
		if db.Limit(PAGESIZE).Find(&results).Error != nil {
			c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "参数错误"})
			return
		}
		resultsJSON, _ := json.Marshal(results)
		c.JSON(http.StatusOK, gin.H{"results": json.RawMessage(resultsJSON), "code": 200, "msg": "查询完成"})
	}
}

func GetSubmission() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idNumber, err := strconv.Atoi(id)
		if err != nil || idNumber <= 0 {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数错误"})
			return
		}

		var results resp
		db := config.GetDB()
		if db.Model(model.Submission{}).Where("id = ?", idNumber).Take(&results).Error != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "记录不存在"})
			return
		}
		resultsJSON, _ := json.Marshal(results)
		c.JSON(http.StatusOK, gin.H{"results": json.RawMessage(resultsJSON), "code": 200, "msg": "查询完成"})
	}
}
