package service

import (
	"github.com/hfleury/digiexam/pkg/grade"
	"github.com/hfleury/digiexam/pkg/grade/repo"
)

type GradeService struct {
	gradeRepo *repo.GradeRepo
}

func NewGradeService(
	gradeRepo *repo.GradeRepo,
) *GradeService {
	return &GradeService{
		gradeRepo: gradeRepo,
	}
}

func (gs *GradeService) GetStudentGrade() (grade.ResponseStudentGrade, error) {
	return grade.ResponseStudentGrade{}, nil
}
