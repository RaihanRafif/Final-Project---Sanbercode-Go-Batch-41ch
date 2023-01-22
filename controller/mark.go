package controller

import (
	"finaltask/database"
	"finaltask/repository"
	"finaltask/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateMark(c *gin.Context) {
	var mark structs.Mark

	value, _ := strconv.Atoi(c.PostForm("mark"))
	mark.Mark = int64(value)

	studentID, _ := strconv.Atoi(c.PostForm("student_id"))
	mark.StudentID = int64(studentID)
	//id kelas
	id, _ := strconv.Atoi(c.Param("id"))

	mark.ClassID = int64(id)

	errs := repository.CreateMark(database.DbConnection, mark)
	if errs != nil {
		panic("ID tidak ditemukan")
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Insert Mark Success",
	})
}

func GetMarksByClassID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	errs, result := repository.GetMarksByID(database.DbConnection, id)
	if errs != nil {
		panic("ID tidak ditemukan")
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func UpdateMark(c *gin.Context) {
	var mark structs.Mark

	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&mark)
	if err != nil {
		panic(err)
	}

	mark.ID = int64(id)

	err = repository.UpdateMark(database.DbConnection, mark)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Mark",
	})
}

func DeleteMark(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := repository.DeleteMark(database.DbConnection, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Mark",
	})
}
