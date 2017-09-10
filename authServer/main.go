package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type user struct {
	Username string `json:"usernamer"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// This map will store the username/token key value pairs
var users = make(map[string]string)

var seedUsers = []user{
	user{
		Username: "user1",
		Password: "pass1",
	},
	user{
		Username: "user2",
		Password: "pass2",
	},
	user{
		Username: "user3",
		Password: "pass3",
	},
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	s := gin.Default()

	s.POST("/login", login)
	s.POST("/authenticate", authenticate)
	s.POST("/logout", logout)

	s.Run(":8001")
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if token := validateUser(username, password); token == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	} else {
		users[username] = token
		// this sets response with success status code and token
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func authenticate(c *gin.Context) {
	username := c.PostForm("username")
	token := c.PostForm("token")

	if v, ok := users[username]; ok && v == token {
		// If the username/token pair is found in the users map, respond with HTTP success status
		c.JSON(http.StatusOK, nil)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func logout(c *gin.Context) {
	username := c.PostForm("username")
	token := c.PostForm("token")

	if v, ok := users[username]; ok && v == token {
		delete(users, username)
		c.JSON(http.StatusOK, nil)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

// validateUser validate username/password against seed value.
func validateUser(username, password string) string {
	for _, u := range seedUsers {
		if username == u.Username {
			if u.Password == password {
				return generateSessionToken()
			}
			return ""
		}
	}
	return ""
}
