package grade

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hfleury/digiexam/internal"
	"github.com/hfleury/digiexam/internal/app"
	"github.com/hfleury/digiexam/pkg/grade/repo"
	"github.com/hfleury/digiexam/pkg/grade/service"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ginEngine *gin.Engine
var dbConn *gorm.DB

func TestMain(m *testing.M) {
	os.Setenv("DIGI_ENVIRONMENT", "TEST")
	appDigi, errapp := app.NewAppDigiexam()
	if errapp != nil {
		log.Fatalf("error initialing the App billogram %v", errapp)
	}

	var err error
	dsn := "host=localhost user=rootuser password=nosecret dbname=digidb port=5432 sslmode=disable"
	dbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connect to db: %v", err)
	}

	digiRouter := internal.NewDigiRouter(appDigi)
	ginEngine = digiRouter.ConfigRouter()
	digiRepo := repo.NewGradeRepo(dbConn)
	digiService := service.NewGradeService(digiRepo)
	digiHandler := NewGradeHandler(digiService)
	gradeRouter := NewGradeRouter(appDigi, digiHandler)

	gradeRouter.SetGradeRouters()

	m.Run()
}

func TestGradeRouter(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/students/gpa", nil)
		ginEngine.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, 200, w.Code)
	})

}
