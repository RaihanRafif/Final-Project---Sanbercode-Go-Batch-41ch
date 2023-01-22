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

func CreateTeacher(c *gin.Context) {
	var Teacher structs.Teacher

	Teacher.Username = c.PostForm("username")
	Teacher.Email = c.PostForm("email")
	Teacher.Password = helpers.HashPass(c.PostForm("password"))
	Teacher.Role = "teacher"

	phone, _ := strconv.Atoi(c.PostForm("phone"))
	Teacher.Phone = int64(phone)

	err := repository.CreateTeacher(database.DbConnection, Teacher)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Register Teacher Account",
	})
}

func LoginTeacher(c *gin.Context) {
	var Teacher structs.Teacher

	Teacher.Email = c.PostForm("email")
	Teacher.Password = c.PostForm("password")

	errs, data := repository.LoginTeacher(database.DbConnection, Teacher)
	comparePass := helpers.ComparePass([]byte(data[0].Password), []byte(Teacher.Password))

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
	Teacher.ID = data[0].ID

	token := helpers.GenerateToken(uint(Teacher.ID), Teacher.Email)
	c.JSON(http.StatusOK, gin.H{
		"result": token,
	})
}

func UpdateTeacher(c *gin.Context) {
	var Teacher structs.Teacher
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	Teacher.Username = c.PostForm("username")
	Teacher.Email = c.PostForm("email")
	Teacher.Password = helpers.HashPass(c.PostForm("password"))
	Teacher.Role = "teacher"

	err := repository.UpdateTeacher(database.DbConnection, int(userID), Teacher)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Updated Account",
	})
}

func DeleteTeacher(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := repository.DeleteTeacher(database.DbConnection, int(userID))
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Account",
	})
}
