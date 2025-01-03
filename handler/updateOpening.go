package handler

import (
	"net/http"

	"github.com/PedroDalpa/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	r := UpdateOpeningRequest{}
	ctx.BindJSON(&r)
	if err := r.Validate(); err != nil {
		logger.Errf("validate error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "query parameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	if r.Role != "" {
		opening.Role = r.Role
	}

	if r.Company != "" {
		opening.Company = r.Company
	}

	if r.Location != "" {
		opening.Location = r.Location
	}
	if r.Remote != nil {
		opening.Remote = *r.Remote
	}
	if r.Salary >= 0 {
		opening.Salary = r.Salary
	}
	if r.Link != "" {
		opening.Link = r.Link
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errf("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "update opening", opening)
}
