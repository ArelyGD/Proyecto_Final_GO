package models

type Grade struct {
	GradeID   int     `json:"grade_id"`
	StudentID int     `json:"student_id"`
	SubjectID int     `json:"subject_id"`
	Grade     float64 `json:"grade"`
}
