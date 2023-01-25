package controller

import (
	"finaltask/database"
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

	class.Topic = c.PostForm("topic")
	maxMarks, _ := strconv.Atoi(c.PostForm("max_marks"))
	class.MaxMarks = int64(maxMarks)
	class.Description = c.PostForm("description")

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	class.TeacherID = int64(userID)

	_, err := repository.GetAllClass(database.DbConnection)

	if err != nil {
		panic(err)
	}

	file, header, _ := c.Request.FormFile("file_url")

	if header != nil {
		dt := time.Now()

		currentTime := dt.Format("01-02-2006")
		timeInSecond := dt.Format("15-04-05")

		// GET Format File
		sourceFile := header.Filename
		splitedFile := strings.Split(sourceFile, ".")
		formatFile := splitedFile[1]
		fmt.Println("sourceFile", sourceFile)
		fmt.Println("formatFile", formatFile)

		filename := strconv.FormatUint(uint64(userID), 10) + "_" + currentTime + "_" + timeInSecond + "." + formatFile

		out, err := os.Create("assets/" + filename)
		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err":     "Bad Request",
				"message": err.Error(),
			})
			return
		}

		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}

		class.Filename = filename

		repository.CreateClass(database.DbConnection, class)

		c.JSON(http.StatusCreated, gin.H{
			"message":    "Success Create Class",
			"data class": class,
			"file":       class.Filename,
		})

		return
	}

	class.Filename = " "
	// fmt.Println("CCCCCCCC", class)
	repository.CreateClass(database.DbConnection, class)

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

func TeacherGetClassByClassID(c *gin.Context) {
	var (
		result gin.H
	)

	//id kelas
	id, _ := strconv.Atoi(c.Param("id"))

	// userData := c.MustGet("userData").(jwt.MapClaims)
	// userID := uint(userData["id"].(float64))

	err, class, members := repository.TeacherGetClassByClassID(database.DbConnection, id)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	}

	// fmt.Println("MMMMMMMM", err)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"data class": class,
			"member":     members,
		}
	}
	fmt.Println("LLLLL", result)

	c.JSON(http.StatusOK, result)
}

func UpdateClassByClassID(c *gin.Context) {
	var class structs.Class
	class.Topic = c.PostForm("topic")
	maxMarks, _ := strconv.Atoi(c.PostForm("max_marks"))
	class.MaxMarks = int64(maxMarks)
	class.Description = c.PostForm("description")

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	id, _ := strconv.Atoi(c.Param("id"))

	class.ID = int64(id)

	file, header, _ := c.Request.FormFile("file_url")

	if header != nil {
		dt := time.Now()
		// classID := data[lengthData-1].ID

		currentTime := dt.Format("01-02-2006")
		timeInSecond := dt.Format("15-04-05")

		// GET Format File
		sourceFile := header.Filename
		splitedFile := strings.Split(sourceFile, ".")
		formatFile := splitedFile[1]
		fmt.Println("sourceFile", sourceFile)
		fmt.Println("formatFile", formatFile)

		filename := strconv.FormatUint(uint64(userID), 10) + "_" + currentTime + "_" + timeInSecond + "." + formatFile
		out, err := os.Create("assets/" + filename)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": err,
			})
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err":     "Bad Request",
				"message": err.Error(),
			})
			return
		}

		defer out.Close()
		_, err = io.Copy(out, file)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": err,
			})
		}

		class.Filename = filename

		repository.UpdateClass(database.DbConnection, class)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Success Update Class",
		})
	}

	repository.UpdateClass(database.DbConnection, class)

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

func StudentGetAllClassByStudentID(c *gin.Context) {
	var (
		result gin.H
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	_, err := repository.FindAccount(database.DbConnection, int(userID))

	if err != nil {
		result = gin.H{
			"result": "ID Not Found",
		}
	}

	class, err := repository.StudentGetAllClassByStudentID(database.DbConnection, int(userID))

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

func StudentPostFileByClassID(c *gin.Context) {
	var dataFile structs.File
	// db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	dt := time.Now()

	classID, _ := strconv.Atoi(c.Param("id"))

	currentTime := dt.Format("01-02-2006")
	timeInSecond := dt.Format("15-04-05")

	file, header, _ := c.Request.FormFile("file_url")

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

	c.JSON(http.StatusCreated, dataFile)
}

func StudentUpdateFileByID(c *gin.Context) {
	var dataFile structs.File
	// db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	dt := time.Now()

	classID, _ := strconv.Atoi(c.Param("id"))
	fileID, _ := strconv.Atoi(c.Param("fileid"))

	currentTime := dt.Format("01-02-2006")
	timeInSecond := dt.Format("15-04-05")

	file, header, _ := c.Request.FormFile("file_url")

	// GET Format File
	sourceFile := header.Filename
	splitedFile := strings.Split(sourceFile, ".")
	formatFile := splitedFile[1]

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

	repository.UpdateFiles(database.DbConnection, dataFile, classID, fileID)

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, dataFile)
}

func StudentDeleteFileByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fileId, _ := strconv.Atoi(c.Param("fileid"))

	err := repository.DeleteFile(database.DbConnection, id, fileId)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete File",
	})
}
