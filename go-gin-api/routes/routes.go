package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tupizz/go-gin-api/controllers"
)

func HandleRequests() {
	router := gin.Default()
	router.GET("/:name", controllers.Greeting)

	// students
	router.GET("/students", controllers.GetAll)
	router.POST("/students", controllers.Create)
	router.GET("/students/:id", controllers.GetOne)
	router.DELETE("/students/:id", controllers.Delete)
	router.PUT("/students/:id", controllers.Update)
	router.GET("/students/cpf/:cpf", controllers.SearchByCPF)

	router.Run(":8585")
}
