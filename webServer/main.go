package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var auth = authService{Base: "http://localhost:8001"}

func main() {
	gin.SetMode(gin.ReleaseMode)
	s := gin.Default()

	s.POST("/login", login)
	s.GET("logout", logout)
	s.GET("/protedted-content", serveProtectedContent)

	s.Run(":8000")
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if response := auth.Login(username, password); response.Token != "" {
		c.SetCookie("username", username, 3600, "", "", false, true)
		c.SetCookie("token", response.Token, 3600, "", "", false, true)

		c.JSON(http.StatusOK, response)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func logout(c *gin.Context) {
	username, err1 := c.Cookie("username")
	token, err2 := c.Cookie("token")

	if err1 == nil && err2 == nil && auth.Logout(username, token) {
		// clear the Cookies
		c.SetCookie("username", "", -1, "", "", false, true)
		c.SetCookie("token", "", -1, "", "", false, true)

		c.JSON(http.StatusOK, nil)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func serveProtectedContent(c *gin.Context) {
	username, err1 := c.Cookie("username")
	token, err2 := c.Cookie("token")

	if err1 == nil && err2 == nil && auth.Authenticate(username, token) {
		c.JSON(http.StatusOK, gin.H{"content": "This should be visible to authenticated users only."})
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
