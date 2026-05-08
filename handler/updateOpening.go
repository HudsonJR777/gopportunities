package handler

import (
	"net/http"

	"github.com/HudsonJR777/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Update opening
// @Description Update a job opening by ID
// @Tags Openings
// @Accept json
// @Produce json
// @Param id path int true "Opening ID"
// @Param request body UpdateOpeningRequest true "Update request body"
// @Success 200 {object} schemas.OpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening/{id} [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}
