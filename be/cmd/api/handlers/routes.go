package handlers

import (
	"github.com/federico-paolillo/ssh-attempts/cmd/api/app"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	gin *gin.Engine,
	app *app.App,
) {
	registerStatRoutes(gin, app)
}

func registerStatRoutes(gin *gin.Engine, app *app.App) {
	gin.GET("/stats", getTop15LoginAttempts(app))
}
