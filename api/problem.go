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
		p := c.Param("p")
		c.JSON(http.StatusOK, gin.H{"code": p, "msg": "开发中"})
	}
}
