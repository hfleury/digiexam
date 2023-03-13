package grade

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hfleury/digiexam/pkg/grade/service"
)

type GradeHandler struct {
	gradeService *service.GradeService
}

func NewGradeHandler(
	gradeService *service.GradeService,
) *GradeHandler {
	return &GradeHandler{
		gradeService: gradeService,
	}
}

func (gh *GradeHandler) GetStudentsGpa(c *gin.Context) {

	stdGrades, err := gh.gradeService.GetStudentGrade()
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	c.IndentedJSON(http.StatusOK, stdGrades)
}
