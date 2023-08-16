package main

import (
	"webApp/ClientAndAdmin"
	"webApp/connections"

	"github.com/gin-gonic/gin"
)

func init() {
	connections.LoadENV()
	connections.ConnectTodb()

}

func main() {
	// Instance of gin frame work

	gin := gin.Default()

	//wrap html pages and load it in variable

	gin.LoadHTMLGlob("templates/*.html")

	ClientAndAdmin.ClientSide(gin)
	ClientAndAdmin.AdminSide(gin)

	gin.Run(":8080")

}
