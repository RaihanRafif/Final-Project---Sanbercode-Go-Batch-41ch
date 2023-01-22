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

func CreateStudent(c *gin.Context) {
	var Student structs.Student

	Student.Username = c.PostForm("username")
	Student.Email = c.PostForm("email")
	Student.Password = helpers.HashPass(c.PostForm("password"))
	Student.Role = "student"
	phone, _ := strconv.Atoi(c.PostForm("phone"))
	Student.Phone = int64(phone)

	err := repository.CreateStudent(database.DbConnection, Student)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Register Student Account",
	})
}

func LoginStudent(c *gin.Context) {
	var Student structs.Student

	Student.Email = c.PostForm("email")
	Student.Password = c.PostForm("password")

	errs, data := repository.LoginStudent(database.DbConnection, Student)
	comparePass := helpers.ComparePass([]byte(data[0].Password), []byte(Student.Password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	if errs != nil {
		panic(errs)
	}
	Student.ID = data[0].ID

	token := helpers.GenerateToken(uint(Student.ID), Student.Email)
	c.JSON(http.StatusOK, gin.H{
		"result": token,
	})
}

func UpdateStudent(c *gin.Context) {
	var Student structs.Student

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	Student.Username = c.PostForm("username")
	Student.Email = c.PostForm("email")
	Student.Password = helpers.HashPass(c.PostForm("password"))

	err := repository.UpdateStudent(database.DbConnection, int(userID), Student)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Updated Account",
	})
}

func DeleteStudent(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := repository.DeleteStudent(database.DbConnection, int(userID))
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Account",
	})
}
