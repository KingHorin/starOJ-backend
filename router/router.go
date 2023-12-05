package router

import (
	"github.com/gin-gonic/gin"
	"starOJ-backend/api"
	"starOJ-backend/middleware/jwt"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	indexGroup := r.Group("/index")
	{
		indexGroup.GET("/getDailyProblem", api.GetDailyProblem())
		indexGroup.GET("/getUserSheet", api.GetUserSheet())
		indexGroup.GET("/getRecentContest", api.GetRecentContest())
	}

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", api.Login())
		authGroup.POST("/register", api.Register())
	}

	userGroup := r.Group("/user").Use(jwt.JwtMiddleware())
	{
		//userGroup.GET("/getProfile/:username", api.getProfile())
		userGroup.POST("/postNickname", api.PostNickname())
		userGroup.POST("/postAvatar", api.PostAvatar())
		userGroup.POST("/postPhone", api.PostPhone())
		userGroup.POST("/postEmail", api.PostEmail())
	}

	problemGroup := r.Group("/problem")
	{
		problemGroup.GET("/list", api.GetProblemList())
		problemGroup.GET("/:id", api.GetProblem())
	}

	submissionGroup := r.Group("/submission")
	{
		submissionGroup.GET("/list", api.GetSubmissionList())
		submissionGroup.GET("/:id", api.GetSubmission())
	}

	sheetGroup := r.Group("/sheet")
	{
		sheetGroup.GET("/list", api.GetSheetList())
		sheetGroup.GET("/:id", api.GetSheet())
	}

	return r
}
