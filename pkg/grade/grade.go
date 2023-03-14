package grade

type StudentGrade struct {
	StudentID int `json:"student_id"`
	GPA       int `json:"gpa"`
}

type ResponseStudentGrade struct {
	StudentGrade []StudentGrade
}

type Grade struct {
	GradeID   int    `gorm:"column:grade_id"`
	StudentID int    `gorm:"column:student_id"`
	CourseID  int    `gorm:"column:course_id"`
	ScaleID   string `gorm:"column:scale_id"`
	Grade     string `gorm:"column:grade_grade"`
	GradeMin  int    `gorm:"column:grade_min"`
	GradeGPA  int    `gorm:"column:grade_gpa"`
}

type GradeService interface {
	GetStudentGrade() (ResponseStudentGrade, error)
}

type GradeRepo interface {
	FetchAllStudentGrade() ([]Grade, error)
}

func (Grade) TableName() string {
	return "grade"
}
