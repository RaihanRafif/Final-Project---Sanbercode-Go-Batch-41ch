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

var (
	appJson = "application/json"
)

func CreatFile(c *gin.Context) {
	var dataFile structs.File
	// db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	// Photo := models.Photo{}
	userID := uint(userData["id"].(float64))
	dt := time.Now()

	classID, _ := strconv.Atoi(c.Param("id"))

	currentTime := dt.Format("01-02-2006")
	timeInSecond := dt.Format("15-04-05")

	if contentType == appJson {
		c.ShouldBindJSON(&dataFile)
	} else {
		c.ShouldBind(&dataFile)
	}

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
