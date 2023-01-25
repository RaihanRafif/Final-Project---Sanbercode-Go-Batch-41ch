package controller

import (
	"finaltask/database"
	"finaltask/helpers"
	"finaltask/repository"
	"finaltask/structs"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var User structs.User

	User.Username = c.PostForm("username")
	User.Email = c.PostForm("email")
	User.Password = helpers.HashPass(c.PostForm("password"))
	User.Role = c.PostForm("role")
	phone, _ := strconv.Atoi(c.PostForm("phone"))
	User.Phone = int64(phone)

	err := repository.CreateUser(database.DbConnection, User)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Register User Account",
	})
}

func LoginUser(c *gin.Context) {
	var User structs.User

	User.Email = c.PostForm("email")
	User.Password = c.PostForm("password")
	User.Role = c.PostForm("role")

	errs, data := repository.LoginUser(database.DbConnection, User)
	if errs != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password/role",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(data[0].Password), []byte(User.Password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password/role",
		})
		return
	}

	if errs != nil {
		panic(errs)
	}
	User.ID = data[0].ID

	token := helpers.GenerateToken(uint(User.ID), User.Email)
	c.JSON(http.StatusOK, gin.H{
		"result": token,
	})
}

func UpdateUser(c *gin.Context) {
	var User structs.User
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	User.Username = c.PostForm("username")
	User.Email = c.PostForm("email")
	User.Password = helpers.HashPass(c.PostForm("password"))

	err := repository.UpdateUser(database.DbConnection, int(userID), User)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Updated Account",
	})
}

func DeleteUser(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := repository.DeleteUser(database.DbConnection, int(userID))
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Account",
	})
}
