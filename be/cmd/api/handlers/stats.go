package handlers

import (
	"net/http"
	"slices"

	"github.com/federico-paolillo/ssh-attempts/cmd/api/app"
	"github.com/federico-paolillo/ssh-attempts/cmd/api/dtos"
	"github.com/gin-gonic/gin"
)

var ammissibleNodenames = []string{
	"controlplane-1",
	"worker-1",
}

func getTop15LoginAttempts(app *app.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		nodename := ctx.Param("nodename")

		if !slices.Contains(ammissibleNodenames, nodename) {
			ctx.Status(http.StatusBadRequest)
			return
		}

		attempts, err := app.Provider.Top15LoginAttempts(nodename)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		attemptsDto := dtos.MapAttemptsToDto(attempts)

		ctx.JSON(http.StatusOK, attemptsDto)
	}
}
