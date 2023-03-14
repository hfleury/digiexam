package repo

import (
	"log"
	"testing"

	"github.com/hfleury/digiexam/pkg/grade"
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
		assert.Len(t, rtnGrades, 10)
	})

	t.Run("Fail", func(t *testing.T) {
		dsnFail := "host=localhost user=rootuser password=nosecret dbname=digidb port=5432 sslmode=disable"
		dbConnFail, errConn := gorm.Open(postgres.Open(dsnFail), &gorm.Config{})
		if errConn != nil {
			log.Fatalf("error connect to db: %v", err)
		}

		dbConnFail.Migrator().DropTable(grade.Grade{})
		userRepo := NewGradeRepo(dbConnFail)

		_, err := userRepo.FetchAllStudentGrade()
		assert.Error(t, err, "ERROR: relation \"grade\" does not exist (SQLSTATE 42P01)")
	})
}
