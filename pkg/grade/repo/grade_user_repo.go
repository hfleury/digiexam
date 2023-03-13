package repo

import (
	"github.com/hfleury/digiexam/pkg/grade"
	"gorm.io/gorm"
)

type GradeRepo struct {
	dbConn *gorm.DB
}

func NewGradeRepo(
	dbConn *gorm.DB,
) *GradeRepo {
	return &GradeRepo{
		dbConn: dbConn,
	}
}

func (gp *GradeRepo) FetchAllStudentGrade() ([]grade.Grade, error) {
	var rtnStudentGrade []grade.Grade
	result := gp.dbConn.Find(&rtnStudentGrade)
	if result.Error != nil {
		return []grade.Grade{}, result.Error
	}
	return rtnStudentGrade, nil
}
