package router

import (
	"github.com/gin-gonic/gin"
	"starOJ-backend/api"
	"starOJ-backend/middleware/jwt"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", api.Login())
		authGroup.POST("/register", api.Register())
	}

	userGroup := r.Group("/user").Use(jwt.JwtMiddleware())
	{
		userGroup.POST("/postNickname", api.PostNickname())
		userGroup.POST("/postAvatar", api.PostAvatar())
		userGroup.POST("/postPhone", api.PostPhone())
		userGroup.POST("/postEmail", api.PostEmail())
	}

	problemGroup := r.Group("/problem")
	{
		problemGroup.GET("/list", api.GetProblemList())
		problemGroup.GET("/:p", api.GetProblem())
	}

	return r
}
