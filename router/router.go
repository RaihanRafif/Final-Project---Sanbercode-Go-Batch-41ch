package router

import (
	"finaltask/controller"
	"finaltask/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	//user controller
	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", controller.CreateUser)
		userRouter.POST("/login", controller.LoginUser)
		userRouter.Use(middleware.Authentication())
		userRouter.PUT("/update", middleware.UserAuthorization(), controller.UpdateUser)
		userRouter.DELETE("/delete", middleware.UserAuthorization(), controller.DeleteUser)
	}

	// //student controller
	// studentRouter := router.Group("/student")
	// {
	// 	studentRouter.POST("/register", controller.CreateStudent)
	// 	studentRouter.POST("/login", controller.LoginStudent)
	// 	studentRouter.Use(middleware.Authentication())
	// 	studentRouter.PUT("/update", middleware.UserAuthorization(), controller.UpdateStudent)
	// 	studentRouter.DELETE("/delete", middleware.UserAuthorization(), controller.DeleteStudent)
	// 	// CLASS SECTION
	// }

	//class controller
	classRouter := router.Group("/class")
	{
		classRouter.Use(middleware.Authentication())
		classRouter.POST("/", middleware.TeacherAuthorization(), controller.CreateClass)
		router.StaticFS("/upload", http.Dir("assets"))
		classRouter.GET("/", middleware.TeacherAuthorization(), controller.GetAllClassByTeacherID)
		classRouter.GET("/:id", middleware.TeacherAuthorization(), middleware.TeacherAccessAuthorization(), controller.TeacherGetClassByClassID)
		classRouter.PUT("/:id", middleware.TeacherAuthorization(), middleware.TeacherAccessAuthorization(), controller.UpdateClassByClassID)
		classRouter.DELETE("/:id", middleware.TeacherAuthorization(), controller.DeleteClass)

		// //STUDENT  CLASS
		classRouter.GET("/student/", middleware.StudentAuthorization(), controller.StudentGetAllClassByStudentID)
		classRouter.POST("/:id/student/", middleware.StudentAuthorization(), middleware.MemberAuthorization(), controller.StudentPostFileByClassID)
		classRouter.PUT("/:id/student/:fileid", middleware.StudentAuthorization(), middleware.MemberAuthorization(), controller.StudentUpdateFileByID)
		classRouter.DELETE("/:id/student/:fileid", middleware.StudentAuthorization(), middleware.MemberAuthorization(), controller.StudentDeleteFileByID)
	}

	//marks controller
	markRouter := router.Group("/marks")
	{
		// markRouter.GET("/:id", controller.GetMarksByClassID)
		markRouter.Use(middleware.Authentication())
		markRouter.POST("/:id", middleware.TeacherAuthorization(), middleware.TeacherAccessAuthorization(), controller.CreateMark)
		markRouter.PUT("/:id", middleware.TeacherAuthorization(), middleware.TeacherAccessAuthorization(), controller.UpdateMark)
		// markRouter.DELETE("/:id", middleware.TeacherAuthorization(), controller.DeleteMark)
	}

	//member controller
	memberRouter := router.Group("/member")
	{
		memberRouter.POST("/", controller.CreateMember)
		memberRouter.DELETE("/:id", controller.DeleteMember)
	}

	return router
}
