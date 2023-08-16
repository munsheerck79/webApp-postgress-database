package Handller

import (
	"fmt"
	"net/http"
	"webApp/connections"
	"webApp/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func AdminHome(a *gin.Context) {
	fmt.Println(" admin home")
	var Data []models.User
	connections.Db.Find(&Data)
	type columns struct {
		ID     int
		Name   string
		Email  string
		Status bool
	}
	var ArrayOfField []columns

	for i, v := range Data {

		ArrayOfField = append(ArrayOfField, columns{
			ID:     i + 1,
			Name:   v.Name,
			Email:  v.Email,
			Status: v.Block,
		})

	}
	fmt.Println(ArrayOfField)
	a.HTML(http.StatusOK, "AdminPanel.html", ArrayOfField)

}

// func Block
// func DeleteUser

// Block user
func BlockUser(c *gin.Context) {
	fmt.Println("Admin tries to block")

	Email := c.Params.ByName("Email")
	fmt.Println("jjjjj", Email)

	if c.Params.ByName("Status") == "true" {
		connections.Db.Model(&models.User{}).Where("email = ?", Email).Update("Block", true)
	} else {
		connections.Db.Model(&models.User{}).Where("email = ?", Email).Update("Block", false)
	}
	c.Redirect(http.StatusSeeOther, "/admin/home")

}

// To delete
func DeleteUser(c *gin.Context) {

	usermail := c.Param("Email")

	connections.Db.Clauses(clause.OnConflict{DoNothing: true}).Delete(&models.User{}, "email= ?", usermail)

	c.Redirect(http.StatusSeeOther, "/admin/home")
}
