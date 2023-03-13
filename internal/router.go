package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hfleury/digiexam/internal/app"
)

type DigiRouter struct {
	appDigiRouter *app.AppDigiexam
}

func NewDigiRouter(
	appDigiRouter *app.AppDigiexam,
) *DigiRouter {
	return &DigiRouter{
		appDigiRouter: appDigiRouter,
	}
}

func (hr *DigiRouter) ConfigRouter() *gin.Engine {
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	hr.appDigiRouter.GinEngine.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	return hr.appDigiRouter.GinEngine
}
