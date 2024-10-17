package handlers

import (
	"net/http"
	"sistema_escolar/database"
	"sistema_escolar/models"

	"github.com/gin-gonic/gin"
)

// Crear una nueva calificación
func CreateGrade(c *gin.Context) {
	var grade models.Grade
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.Exec("INSERT INTO grades (student_id, subject_id, grade) VALUES (?, ?, ?)", grade.StudentID, grade.SubjectID, grade.Grade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Grade created successfully"})
}

// Obtener una calificación por su ID y el ID del estudiante
func GetGrade(c *gin.Context) {
	var grade models.Grade
	gradeID := c.Param("grade_id")
	studentID := c.Param("student_id")

	row := database.DB.QueryRow("SELECT grade_id, student_id, subject_id, grade FROM grades WHERE grade_id = ? AND student_id = ?", gradeID, studentID)
	err := row.Scan(&grade.GradeID, &grade.StudentID, &grade.SubjectID, &grade.Grade)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Calificacion no encontrada."})
		return
	}
	c.JSON(http.StatusOK, grade)
}
