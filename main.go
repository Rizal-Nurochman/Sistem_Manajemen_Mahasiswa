package main

import (
	"belajar-api/database"
	"belajar-api/handling"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main()  {
	//SETUP DATABASE
	database.ConnectedDatabase()

	studentRepository:=database.NewRepository(database.DB)
	studentService:=handling.NewService(studentRepository)
	studentHandler:=handling.NewStudentHandler(studentService)

	router:=gin.Default()

		router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/students", studentHandler.GetStudentsHandler)
	router.GET("/students/:NIM", studentHandler.GetStudentsByNIMHandler)
	router.POST("/students", studentHandler.PostStudentsHandler)
	router.PUT("/students/:NIM", studentHandler.PutStudentsByNIMHandler)
	router.DELETE("/students/:NIM", studentHandler.DelStudentsByNIMHandler)

	router.Run()
}