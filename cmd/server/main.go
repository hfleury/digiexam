package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/hfleury/digiexam/internal"
	"github.com/hfleury/digiexam/internal/app"
	"github.com/hfleury/digiexam/internal/grade"
	"github.com/hfleury/digiexam/pkg/grade/repo"
	"github.com/hfleury/digiexam/pkg/grade/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// ctx := context.Background()
	appDigi, err := app.NewAppDigiexam()
	if err != nil {
		log.Fatalln("Error initializing App")
	}

	// Graceful shutdown the server, wait remaings connections to finish
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - 15s")
	flag.Parse()

	// Init DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=CET",
		appDigi.EnvVars["DIGI_POSTGRESQL_HOST"],
		appDigi.EnvVars["DIGI_POSTGRESQL_USER"],
		appDigi.EnvVars["DIGI_POSTGRESQL_PASS"],
		appDigi.EnvVars["DIGI_POSTGRESQL_DB"],
		appDigi.EnvVars["DIGI_POSTGRESQL_PORT"],
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connectio to DB - error: %v", err)
	}

	// Init Repo
	digiRepo := repo.NewGradeRepo(db)

	// Ini Services
	digiService := service.NewGradeService(digiRepo)

	// Init Handlers
	gradeHandler := grade.NewGradeHandler(digiService)

	// Init and config router
	digiRouter := internal.NewDigiRouter(appDigi)
	ginEngine := digiRouter.ConfigRouter()
	gradeRouter := grade.NewGradeRouter(appDigi, gradeHandler)
	gradeRouter.SetGradeRouters()

	// Start the server
	hostUrl := fmt.Sprintf("%v:%v", appDigi.EnvVars["DIGI_HOST_ADDRESS"], appDigi.EnvVars["DIGI_HOST_PORT"])
	ginEngine.Run(hostUrl)
}
