package ClientAndAdmin

import (
	"webApp/Handller"
	middleware "webApp/middleWare"

	"github.com/gin-gonic/gin"
)

func ClientSide(gin *gin.Engine) {

	gin.GET("/", middleware.UserAuthentiaction)
	gin.POST("/", Handller.LoginSubmission)

	gin.GET("/signup", middleware.UserAuthentiaction, Handller.SignUpUser)
	gin.POST("/signup", Handller.SignUpSubmission)

	gin.GET("/home", middleware.UserAuthentiaction)
	gin.GET("/logout", Handller.LogoutUser)

}
