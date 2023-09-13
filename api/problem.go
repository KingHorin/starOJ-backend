package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProblemList() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.Query("page")
		c.JSON(http.StatusOK, gin.H{"code": page, "msg": "开发中"})
	}
}

func GetProblem() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"code": id, "msg": "开发中"})
	}
}
