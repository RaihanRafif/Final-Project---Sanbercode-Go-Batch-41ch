package middleware

import (
	"finaltask/database"
	"finaltask/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var account structs.Account

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		data, err := repository.FindAccount(database.DbConnection, int(userID))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data doesn't exist",
			})
			return
		}

		if uint(data[0].ID) != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "your are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func TeacherAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData := c.MustGet("userData").(jwt.MapClaims)
		emailID := userData["email"].(string)

		err, data := repository.TeacherAuthorization(database.DbConnection, emailID)

		if err != nil || len(data) == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data doesn't exist",
			})
			return
		}

		if data[0].Role != "teacher" {
			fmt.Println("DDDDDDDDDd")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Only Teacher Can Make a class",
			})
			return
		}

		c.Next()
	}
}

func TeacherAccessAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		//id kelas
		classID, _ := strconv.Atoi(c.Param("id"))

		userData := c.MustGet("userData").(jwt.MapClaims)
		emailID := userData["email"].(string)

		err := repository.TeacherAccessAuthorization(database.DbConnection, emailID, classID)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data doesn't exist",
			})
			return
		}
		c.Next()
	}
}

func MemberAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		//id kelas
		classID, _ := strconv.Atoi(c.Param("id"))

		userData := c.MustGet("userData").(jwt.MapClaims)
		emailID := userData["email"].(string)

		err := repository.MemberAccessAuthorization(database.DbConnection, emailID, classID)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data doesn't exist",
			})
			return
		}
		c.Next()
	}
}
