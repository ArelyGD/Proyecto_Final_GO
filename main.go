package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	// Conexion a la base de datos MySQL
	db, err = sql.Open("mysql", "root:proyectogo@tcp(127.0.0.1:3306)/sistema_escolar")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	// Rutas para Calificaciones
	router.POST("/api/grades", createGrade)
	router.PUT("/api/grades/:grade_id", updateGrade)
	router.GET("/api/grades/:grade_id/student/:student_id", getGrade)

	router.Run(":8080")
}

// Función para crear una calificación
func createGrade(c *gin.Context) {
	var grade struct {
		StudentID int     `json:"student_id"`
		SubjectID int     `json:"subject_id"`
		Grade     float64 `json:"grade"`
	}
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("INSERT INTO grades (student_id, subject_id, grade) VALUES (?, ?, ?)", grade.StudentID, grade.SubjectID, grade.Grade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Calificacion creada exitosamente."})
}

// Función para actualizar una calificación por ID
func updateGrade(c *gin.Context) {
	id := c.Param("grade_id")
	var grade struct {
		StudentID int     `json:"student_id"`
		SubjectID int     `json:"subject_id"`
		Grade     float64 `json:"grade"`
	}
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE grades SET student_id = ?, subject_id = ?, grade = ? WHERE grade_id = ?", grade.StudentID, grade.SubjectID, grade.Grade, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la calificación."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Calificación actualizada exitosamente."})
}

// Función para obtener una calificación por ID y el ID del estudiante
func getGrade(c *gin.Context) {
	gradeID := c.Param("grade_id")
	studentID := c.Param("student_id")
	var grade struct {
		GradeID   int
		StudentID int
		SubjectID int
		Grade     float64
	}
	row := db.QueryRow("SELECT grade_id, student_id, subject_id, grade FROM grades WHERE grade_id = ? AND student_id = ?", gradeID, studentID)
	err := row.Scan(&grade.GradeID, &grade.StudentID, &grade.SubjectID, &grade.Grade)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Calificacion no encontrada."})
		return
	}
	c.JSON(http.StatusOK, grade)
}
