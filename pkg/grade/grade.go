package grade

type StudentGrade struct {
	StudentName string `json:"student_grade"`
	CourseName  string `json:"course_name"`
	ScaleType   string `json:"scale_name"`
	Grade       string `json:"grade"`
}

type ResponseStudentGrade struct {
	StudentGrade []StudentGrade
}

type Grade struct {
	GradeID      int    `gorm:"column:grade_id"`
	StudentID    int    `gorm:"column:student_id"`
	CourseID     int    `gorm:"column:course_id"`
	ScaleID      string `gorm:"column:scale_id"`
	GradeScaleID int    `gorm:"grade_scale"`
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
