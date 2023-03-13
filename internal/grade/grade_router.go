package grade

import (
	"github.com/hfleury/digiexam/internal/app"
)

type GradeRouter struct {
	appDigiRouter *app.AppDigiexam
	GradeHandler  *GradeHandler
}

func NewGradeRouter(
	appDig *app.AppDigiexam,
	GradeHandler *GradeHandler,
) *GradeRouter {
	return &GradeRouter{
		appDigiRouter: appDig,
		GradeHandler:  GradeHandler,
	}
}

func (gr *GradeRouter) SetGradeRouters() {
	studentGroup := gr.appDigiRouter.GinEngine.Group("/students")

	studentGroup.Use()
	{
		studentGroup.GET("/gpa", gr.GradeHandler.GetStudentsGpa)
	}
}
