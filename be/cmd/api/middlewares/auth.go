package middlewares

import (
	"fmt"
	"net/http"

	"github.com/federico-paolillo/ssh-attempts/cmd/api/app"
	"github.com/gin-gonic/gin"
)

func Auth(app *app.App) gin.HandlerFunc {
	headerKey := app.Cfg.Auth.HeaderKey
	headerValue := app.Cfg.Auth.HeaderValue

	return func(ctx *gin.Context) {
		req := fmt.Sprintf("%s %s", ctx.Request.Method, ctx.Request.URL)

		if _, ok := ctx.Request.Header[headerKey]; !ok {
			app.Log.Printf("auth: request %s denied. no auth header supplied", req)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		headerValueGiven := ctx.Request.Header.Get(headerKey)

		if headerValueGiven != headerValue {
			app.Log.Printf("auth: request %s denied. wrong auth header supplied", req)
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}
