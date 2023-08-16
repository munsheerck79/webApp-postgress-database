package Helpers

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

//-------------------------------------------------  email validation    ----------------------------------------

func IsValidEmail(email string) bool {
	// Regular expression to validate email address
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile regular expression
	regex := regexp.MustCompile(pattern)

	// Match email address against regular expression
	return regex.MatchString(email)
}

// JWT token & cookie setup for session handling
func JwtCookieSetup(c *gin.Context, name string, userId interface{}) bool {
	cookieTime := time.Now().Add(20 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId, // Store logged user info in token
		"exp":    float64(cookieTime),
	})

	// Generate signed JWT token using evn var of secret key
	if tokenString, err := token.SignedString([]byte(os.Getenv("Jwtkey"))); err == nil {

		// Set cookie with signed string if no error
		c.SetCookie(name, tokenString, 10*60, "", "", false, true)

		fmt.Println("JWT sign & set Cookie successful")
		return true
	}
	fmt.Println("Failed JWT setup")
	return false

}

func GetToken(ctx *gin.Context, name string) (*jwt.Token, bool) {

	// get cookie from client
	cookieval, ok := GetCookieVal(ctx, name)

	if !ok { // problem to get cookie so return false
		return nil, false
	}

	// Parse cookie to get JWT token
	token, err := jwt.Parse(cookieval, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("Jwtkey")), nil
	})
	if err != nil {
		fmt.Println("failed to parse the cookie to token")
		return nil, false
	}
	return token, true

}

// To get cookie from client
func GetCookieVal(ctx *gin.Context, name string) (string, bool) {

	if cookieVal, err := ctx.Cookie(name); err == nil {
		return cookieVal, true
	}

	fmt.Println("Failed to get cookie")
	return "", false
}
