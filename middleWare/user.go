package middleware

import (
	"fmt"
	"net/http"
	"time"
	"webApp/Handller"
	"webApp/Helpers"
	"webApp/connections"
	"webApp/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserAuthentiaction(ctx *gin.Context) {
	fmt.Println("Processing user authentication...")

	token, ok := Helpers.GetToken(ctx, "user")

	if !ok {
		
		if ctx.Request.URL.Path == "/signup" {
			ctx.Next()
			return
		}
		//else case abort and go to login page
		ctx.Abort()
		Handller.LoginPage(ctx)
		return
	}

	//if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		//check its time is not over
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			fmt.Println(" checking exp not ok")
			ctx.Abort()
			ctx.Redirect(http.StatusSeeOther, "/")
			return
		}

		//get the user from database using token claims
		userId := uint(claims["userId"].(float64))

		var user models.User
		connections.Db.Find(&user, "id = ?", userId)
		fmt.Println(">>>>>>>>>>>>>>>>>>>>", userId)
		fmt.Println(">>>>>>>>>>>>>>>>>>>>", user)

		if user.ID == 0 { //user not found or user blocked by admin
			fmt.Println("user not ok")
			//check the path that user want to signup login after he is not a valid user or blocked by admin

			// if ctx.Request.URL.Path == "/signup" {
			// 	fmt.Println("if user tries signup")
			// 	ctx.Next()
			// 	return
			// }
			//any other path just show the login page
			ctx.Abort()
			fmt.Println("user not found but jwt is there admin deleted user")
			Handller.LoginPage(ctx)
			return
		}

		ctx.Set("userId", userId) //atach the user id in context if user is valid

		//if the user is valid and enter signup or login url show home page
		// if ctx.Request.URL.Path == "/" || ctx.Request.URL.Path == "/signup" {
		// 	fmt.Println("if user tries signup &login")
		// 	ctx.Abort()
		// 	ctx.Redirect(http.StatusSeeOther, "/home")
		// 	return
		// }

		//if all condition completed and the url is for home page
		if user.Ownstatus == "admin__" {
			Handller.AdminHome(ctx)
			return
		}
		Handller.UserHome(ctx, user.Name)
	} else {
		fmt.Println("if token invalid ")
		//if the token is invalid or cant claim then show login page
		ctx.Abort()
		ctx.Redirect(http.StatusSeeOther, "/")
	}

}
