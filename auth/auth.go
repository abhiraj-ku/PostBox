package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string
	Username string
	Password string
}

var mockuserDB = map[string]User{
	"user1": {ID: "1", Username: "Abhishek kumar", Password: "hello@dejfbjhfdu"},
	"user2": {ID: "2", Username: "Arko kumar", Password: "hello@dsds"},
}

func LoginHandler(c *gin.Context) {
	var loginjsonInfo struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindBodyWithJSON(&loginjsonInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not in required format"})
		return
	}

	user, exists := mockuserDB[loginjsonInfo.Username]

	if !exists {
		c.JSON(http.StatusNoContent, gin.H{"error": "user does not exists"})
	}

	if err:= bcrypt.CompareHashAndPassword([]byte(user.Password))

}
