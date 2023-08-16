package Handller

import (
	"fmt"
	"net/http"
	"webApp/Helpers"
	"webApp/connections"

	"github.com/gin-gonic/gin"
)

// ----------------------------------------login page get-------------------------------------------------------

func LoginPage(d *gin.Context) {

	// display login page
	fmt.Println("login user")

	d.Header("Cache-Control", "no-cache, no-store, must-revalidate")

	d.HTML(http.StatusOK, "loginPage.html", nil)

}

//-------------------------------------------signup page get----------------------------------------------------

func SignUpUser(d *gin.Context) {
	fmt.Println("call get method of signup")

	//view signup page
	d.Header("Cache-Control", "no-cache, no-store, must-revalidate")

	d.HTML(http.StatusOK, "signup.html", nil)

}

//---------------------------------------------user home page get---------------------------------------------------

func UserHome(u *gin.Context, Name string) {

	//Cookie validaton

	// display login page
	fmt.Println(" user home")

	u.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	res := Name
	u.HTML(http.StatusOK, "homePage.html", res)

}

//-----------------------------------------------login page post------------------------------------------------------

func LoginSubmission(r *gin.Context) {

	// validation (of nill or not)

	email := r.Request.PostFormValue("Email")
	password := r.Request.PostFormValue("Password")
	if email == "" || password == "" {
		fmt.Println("enter values")
		Response := "*Missing inputs,Pleas fill in the column"
		r.HTML(http.StatusBadRequest, "loginPage.html", Response)
		return
	}

	//validation of email structure

	if !Helpers.IsValidEmail(email) {
		fmt.Println("enter values")
		Response := "*Please enter the valid email & password"
		r.HTML(http.StatusBadRequest, "loginPage.html", Response)
		return
	}

	// check user exist in database
	condition, id, sts, err := connections.FindUser(email, password)
	if err == nil {
		fmt.Println("enter values")
		Response := "*Please enter the valid email & password"
		r.HTML(http.StatusBadRequest, "loginPage.html", Response)
		return
	}
	if !condition && err != nil {
		fmt.Println("blocked by admin")
		Response := "* you are blocked by admin"
		r.HTML(http.StatusBadRequest, "loginPage.html", Response)
		return

	}

	// generate jwt tokens
	if !Helpers.JwtCookieSetup(r, "user", id) {
		r.Redirect(http.StatusSeeOther, "/")
		return
	}
	//===================================================================================

	if err != nil && sts == "user__" {
		fmt.Println("login and display home ")
		Response := err
		r.HTML(http.StatusBadRequest, "homePage.html", Response)
		return
	}

	if err != nil && sts == "admin__" {
		fmt.Println("login and display admin home ")

		// r.HTML(http.StatusBadRequest, "AdminPanel.html")
		AdminHome(r)
		return
	}

	// ======else {
	// 	// set cookie and jwt
	// 	r.HTML(http.StatusBadRequest, "loginPage.html", nil)
	// }=======

	// check admin blocked

	// redirect to homepage if success else login page

}

//-------------------------------------------------  user signup post   -----------------------------------------------

func SignUpSubmission(s *gin.Context) {

	//get data from form vlues and
	//store data
	Firstname := s.Request.PostFormValue("FirstName")
	Emailuser := s.Request.PostFormValue("Email")
	passworduser := s.Request.PostFormValue("Password")
	//validate data & check data allready in db
	if Firstname == "" || Emailuser == "" || passworduser == "" {
		fmt.Println("fill in the column")
		Response := "*Please fill in the all column "
		s.HTML(http.StatusBadRequest, "signup.html", Response)
		return
	}
	if !Helpers.IsValidEmail(Emailuser) {
		fmt.Println("invalid email")
		Response := "*Please enter the valid email "
		s.HTML(http.StatusBadRequest, "signup.html", Response)
		return
	}

	err := connections.AddNewUser(Firstname, Emailuser, passworduser)
	if err != nil {
		s.HTML(400, "signup.html", err)
	} else if err == nil {
		fmt.Println("signup compleated")
		//renter login page
		responce := "Login your account"
		s.HTML(http.StatusOK, "loginPage.html", responce)
	}
}

// To logout user
func LogoutUser(c *gin.Context) {
	fmt.Println("User Logged out")

	_, ok := Helpers.GetCookieVal(c, "user")

	if !ok {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	c.SetCookie("user", "", -1, "", "", false, true)

	//atlast redirect to login page
	c.Redirect(http.StatusSeeOther, "/")
}
