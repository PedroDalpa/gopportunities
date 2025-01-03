package handler

import (
	"fmt"
	"net/http"

	"github.com/PedroDalpa/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(
			ctx,
			http.StatusBadRequest,
			errParamIsRequired("id", "query parameter").Error(),
		)
		return
	}

	opening := schemas.Opening{}
	// find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("opening with id %s not found", id),
		)
		return
	}

	// delete opening
	if err := db.Delete(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id %s", id))
		return
	}

	sendSuccess(ctx, "delete opening", opening)

}
