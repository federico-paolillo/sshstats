package middlewares

import (
	"github.com/federico-paolillo/ssh-attempts/cmd/api/app"
	"github.com/gin-gonic/gin"
)

func Logger(app *app.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		app.Log.Printf(
			"gin: completed request %s %s (%d)",
			ctx.Request.Method,
			ctx.Request.URL,
			ctx.Writer.Status(),
		)
	}
}
