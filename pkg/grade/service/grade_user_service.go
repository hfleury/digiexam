package service

import (
	"fmt"
	"sync"

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
	studentsGrade, err := gs.gradeRepo.FetchAllStudentGrade()
	if err != nil {
		fmt.Printf("Error fetching students grade: %v", err)
		return grade.ResponseStudentGrade{}, nil
	}

	var rspStudentsGrade grade.ResponseStudentGrade

	rspStudentsGrade.StudentGrade = fetchStudentsGrade(studentsGrade)

	return rspStudentsGrade, nil
}

func fetchStudentsGrade(grades []grade.Grade) []grade.StudentGrade {
	var wg sync.WaitGroup
	gradesLen := len(grades)
	wg.Add(gradesLen)
	var rtnGrades = make([]grade.StudentGrade, 0)

	for i := 0; i < gradesLen; i++ {
		go func(i int) {
			defer wg.Done()
			rtnGrades = append(rtnGrades, grade.StudentGrade{
				StudentID: grades[i].StudentID,
				GPA:       grades[i].GradeGPA,
			})
		}(i)

	}

	wg.Wait()
	return rtnGrades
}
