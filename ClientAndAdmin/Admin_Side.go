package ClientAndAdmin

import (
	"webApp/Handller"

	"github.com/gin-gonic/gin"
)

func AdminSide(gin *gin.Engine) {
	gin.GET("/admin/home", Handller.AdminHome)
	gin.GET("/admin/blockuser/:Status/:Email", Handller.BlockUser)
	gin.GET("/admin/logout", Handller.LogoutUser)
	gin.GET("/admin/delete/:Email", Handller.DeleteUser)
	//gin.GET("/admin/adduser",Handller.aAdminadduser)

}
