package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/HudsonJR777/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if idParam == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "pathParameter").Error())
		return
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		sendError(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %d not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %d", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)
}
