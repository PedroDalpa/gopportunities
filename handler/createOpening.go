package handler

import (
	"net/http"

	"github.com/PedroDalpa/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	r := CreateOpeningRequest{}

	ctx.BindJSON(&r)

	if err := r.Validate(); err != nil {
		logger.Errf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     r.Role,
		Company:  r.Company,
		Location: r.Location,
		Remote:   *r.Remote,
		Salary:   r.Salary,
		Link:     r.Link,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errf("Error creating opening request: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error creating database")
		return
	}

	sendSuccess(ctx, "create opening", opening)

}
