package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hfleury/digiexam/internal/config"
)

type AppDigiexam struct {
	EnvVars   map[string]string
	GinEngine *gin.Engine
}

func NewAppDigiexam() (*AppDigiexam, error) {
	app := &AppDigiexam{}
	err := new(error)
	app.EnvVars, *err = config.SetEnv()
	app.GinEngine = gin.Default()

	return app, *err
}
