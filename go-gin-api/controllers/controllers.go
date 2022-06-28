package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tupizz/go-gin-api/database"
	"github.com/tupizz/go-gin-api/models"
)

func GetAll(context *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	context.JSON(200, students)
}

func Greeting(context *gin.Context) {
	name := context.Params.ByName("name")
	context.JSON(200, gin.H{
		"user": name,
	})
}

func Create(context *gin.Context) {
	var student models.Student
	if err := context.ShouldBindJSON(&student); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid payload submitted",
			"details": err.Error(),
		})
		return
	}
	database.DB.Create(&student)
	context.JSON(http.StatusCreated, student)
}

func GetOne(context *gin.Context) {
	studentId := context.Params.ByName("id")
	var student models.Student
	database.DB.First(&student, studentId)

	if student.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "student was not found"})
		return
	}

	context.JSON(http.StatusOK, student)
}

func Delete(context *gin.Context) {
	studentId := context.Params.ByName("id")
	var student models.Student
	database.DB.Delete(&student, studentId)
	context.JSON(http.StatusNoContent, gin.H{})
}

func Update(context *gin.Context) {
	studentId := context.Params.ByName("id")
	var student models.Student
	database.DB.First(&student, studentId)

	if err := context.ShouldBindJSON(&student); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid payload submitted",
			"details": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	context.JSON(http.StatusOK, student)
}

func SearchByCPF(context *gin.Context) {
	cpf := context.Param("cpf")
	var student models.Student
	database.DB.Where(&models.Student{
		CPF: cpf,
	}).First(&student)

	if student.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "student was not found"})
		return
	}

	context.JSON(http.StatusOK, student)
}
