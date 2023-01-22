package router

import (
	"finaltask/controller"
	"finaltask/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	//teacher controller
	teacherRouter := router.Group("/teacher")
	{
		teacherRouter.POST("/register", controller.CreateTeacher)
		teacherRouter.POST("/login", controller.LoginTeacher)
		teacherRouter.Use(middleware.Authentication())
		teacherRouter.PUT("/update", middleware.UserAuthorization(), controller.UpdateTeacher)
		teacherRouter.DELETE("/delete", middleware.UserAuthorization(), controller.DeleteTeacher)
	}

	//student controller
	studentRouter := router.Group("/student")
	{
		studentRouter.POST("/register", controller.CreateStudent)
		studentRouter.POST("/login", controller.LoginStudent)
		studentRouter.Use(middleware.Authentication())
		studentRouter.PUT("/update", middleware.UserAuthorization(), controller.UpdateStudent)
		studentRouter.DELETE("/delete", middleware.UserAuthorization(), controller.DeleteStudent)
	}

	//class controller
	classRouter := router.Group("/class")
	{
		classRouter.Use(middleware.Authentication())
		classRouter.POST("/", middleware.TeacherAuthorization(), controller.CreateClass)
		router.StaticFS("/upload", http.Dir("assets"))
		classRouter.GET("/", middleware.TeacherAuthorization(), controller.GetAllClassByTeacherID)
		classRouter.GET("/:id", middleware.TeacherAuthorization(), middleware.TeacherAccessAuthorization(), controller.GetClassByClassID)
		classRouter.PUT("/:id", middleware.TeacherAuthorization(), middleware.TeacherAccessAuthorization(), controller.UpdateClassByClassID)
		classRouter.DELETE("/:id", middleware.TeacherAuthorization(), controller.DeleteClass)
	}

	//marks controller
	markRouter := router.Group("/marks")
	{
		markRouter.GET("/:id", controller.GetMarksByClassID)
		markRouter.Use(middleware.Authentication())
		markRouter.POST("/", middleware.TeacherAuthorization(), middleware.TeacherAccessAuthorization(), controller.CreateMark)
		markRouter.PUT("/:id", middleware.TeacherAuthorization(), controller.UpdateMark)
		markRouter.DELETE("/:id", middleware.TeacherAuthorization(), controller.DeleteMark)
	}

	//member controller
	memberRouter := router.Group("/member")
	{
		memberRouter.POST("/", controller.CreateMember)
		memberRouter.DELETE("/:id", controller.DeleteMember)
	}

	//file controller
	fileRouter := router.Group("/file")
	{
		fileRouter.Use(middleware.Authentication())
		fileRouter.POST("/:id", middleware.MemberAuthorization(), controller.CreatFile)
	}

	return router
}
