package controller

import (
	"finaltask/database"
	"finaltask/repository"
	"finaltask/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateMember(c *gin.Context) {
	var member structs.Member

	class_id, _ := strconv.Atoi(c.PostForm("class_id"))
	user_id, _ := strconv.Atoi(c.PostForm("user_id"))

	member.ClassID = int64(class_id)
	member.UserID = int64(user_id)

	errs := repository.CreateMember(database.DbConnection, member)
	if errs != nil {
		panic("ID tidak ditemukan")
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Insert Member Success",
	})
}

func DeleteMember(c *gin.Context) {
	// var member structs.Member

	id, _ := strconv.Atoi(c.Param("id"))

	// member.ClassID = int64(class_id)
	// member.UserID = int64(user_id)

	errs := repository.DeleteMember(database.DbConnection, id)
	if errs != nil {
		panic("ID tidak ditemukan")
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Delete Member Success",
	})
}
