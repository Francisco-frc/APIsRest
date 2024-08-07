package handler

import (
	"net/http"

	"github.com/Francisco-frc/APIsRest/schemas"
	"github.com/gin-gonic/gin"
)

func CreatOpeningHandler(ctx *gin.Context) {
	request := CreateopeningRequest{}

	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrorF("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}
	sendSuccess(ctx, "create-opening", opening)
}
