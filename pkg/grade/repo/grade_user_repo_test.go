package repo

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGradeRepo_FetchAllStudentGrade(t *testing.T) {

	dsn := "host=localhost user=rootuser password=nosecret dbname=digidb port=5432 sslmode=disable"
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connect to db: %v", err)
	}

	t.Run("Success get student grades", func(t *testing.T) {

		userRepo := NewGradeRepo(dbConn)

		rtnGrades, err := userRepo.FetchAllStudentGrade()
		assert.NoError(t, err)
		assert.Len(t, rtnGrades, 12)
	})
}
