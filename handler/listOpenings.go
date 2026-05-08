package handler

import (
	"net/http"

	"github.com/HudsonJR777/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary List openings
// @Description Get all job openings
// @Tags Openings
// @Produce json
// @Success 200 {array} schemas.OpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
