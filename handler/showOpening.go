package handler

import (
	"net/http"

	"github.com/HudsonJR777/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Show opening
// @Description Get one job opening by ID
// @Tags Openings
// @Produce json
// @Param id path int true "Opening ID"
// @Success 200 {object} schemas.OpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening/{id} [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "pathParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
