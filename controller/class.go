package controller

import (
	"finaltask/database"
	"finaltask/helpers"
	"finaltask/repository"
	"finaltask/structs"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateClass(c *gin.Context) {
	var class structs.Class
	var dataFile structs.File

	class.Topic = c.PostForm("topic")
	maxMarks, _ := strconv.Atoi(c.PostForm("max_marks"))
	class.MaxMarks = int64(maxMarks)
	class.Description = c.PostForm("description")

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	repository.CreateClass(database.DbConnection, class, int(userID))
	data, err := repository.GetAllClass(database.DbConnection)
	lengthData := len(data)

	if err != nil {
		panic(err)
	}

	contentType := helpers.GetContentType(c)
	file, header, _ := c.Request.FormFile("file_url")

	if header != nil {
		dt := time.Now()
		classID := data[lengthData-1].ID
		dataFile.ClassID = classID

		currentTime := dt.Format("01-02-2006")
		timeInSecond := dt.Format("15-04-05")

		if contentType == appJson {
			c.ShouldBindJSON(&dataFile)
		} else {
			c.ShouldBind(&dataFile)
		}

		// GET Format File
		sourceFile := header.Filename
		splitedFile := strings.Split(sourceFile, ".")
		formatFile := splitedFile[1]
		fmt.Println("sourceFile", sourceFile)
		fmt.Println("formatFile", formatFile)

		filename := strconv.FormatUint(uint64(userID), 10) + "_" + currentTime + "_" + timeInSecond + "." + formatFile
		fmt.Println("filename", filename)
		out, err := os.Create("assets/" + filename)
		if err != nil {
			log.Fatal(err)
		}

		dataFile.Filename = filename
		dataFile.CreatedAt = dt
		dataFile.UpdatedAt = dt
		dataFile.UserID = int64(userID)
		dataFile.ClassID = int64(classID)

		// err = db.Debug().Create(&dataFile).Error

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err":     "Bad Request",
				"message": err.Error(),
			})
			return
		}

		repository.CreateFiles(database.DbConnection, dataFile)
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Success Create Class",
			"result":  dataFile,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Create Class",
	})
}

func GetAllClassByTeacherID(c *gin.Context) {
	var (
		result gin.H
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	class, err := repository.GetAllClassByTeacherID(database.DbConnection, int(userID))

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": class,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetClassByClassID(c *gin.Context) {
	var (
		result gin.H
	)

	//id kelas
	id, _ := strconv.Atoi(c.Param("id"))

	class, err := repository.GetClassByClassID(database.DbConnection, id)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": class,
		}
	}

	c.JSON(http.StatusOK, result)
}

// func GetClassByID(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	errs, result := repository.GetClassByID(database.DbConnection, id)
// 	if errs != nil {
// 		panic("ID tidak ditemukan")
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"result": result,
// 	})
// }

func UpdateClassByClassID(c *gin.Context) {
	var class structs.Class
	var dataFile structs.File

	class.Topic = c.PostForm("topic")
	maxMarks, _ := strconv.Atoi(c.PostForm("max_marks"))
	class.MaxMarks = int64(maxMarks)
	class.Description = c.PostForm("description")

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	id, _ := strconv.Atoi(c.Param("id"))

	class.ID = int64(id)

	err := repository.UpdateClass(database.DbConnection, class)
	if err != nil {
		panic(err)
	}
	contentType := helpers.GetContentType(c)
	file, header, _ := c.Request.FormFile("file_url")

	if header != nil {
		dt := time.Now()
		// classID := data[lengthData-1].ID

		currentTime := dt.Format("01-02-2006")
		timeInSecond := dt.Format("15-04-05")

		if contentType == appJson {
			c.ShouldBindJSON(&dataFile)
		} else {
			c.ShouldBind(&dataFile)
		}

		// GET Format File
		sourceFile := header.Filename
		splitedFile := strings.Split(sourceFile, ".")
		formatFile := splitedFile[1]
		fmt.Println("sourceFile", sourceFile)
		fmt.Println("formatFile", formatFile)

		filename := strconv.FormatUint(uint64(userID), 10) + "_" + currentTime + "_" + timeInSecond + "." + formatFile
		fmt.Println("filename", filename)
		out, err := os.Create("assets/" + filename)
		fmt.Println("444444")

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": err,
			})
		}

		dataFile.Filename = filename
		dataFile.CreatedAt = dt
		dataFile.UpdatedAt = dt
		dataFile.UserID = int64(userID)
		dataFile.ClassID = int64(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err":     "Bad Request",
				"message": err.Error(),
			})
			return
		}

		repository.UpdateFiles(database.DbConnection, dataFile, id)
		defer out.Close()
		_, err = io.Copy(out, file)
		fmt.Println("hhhhhh")

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": err,
			})
		}

		fmt.Println("llllll")
		c.JSON(http.StatusCreated, gin.H{
			"message":   "Success Update Class",
			"data file": dataFile,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Class",
	})
}

func DeleteClass(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := repository.DeleteClass(database.DbConnection, id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete class",
	})
}
