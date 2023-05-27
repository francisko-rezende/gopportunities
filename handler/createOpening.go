package handler

import (
	"net/http"

	"github.com/francisko-rezende/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	err := request.Validate()

	if err != nil {
		logger.ErrF("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// required because the type is used by gorm to find the table where to save the data so if you try to pass something else (like the request), gorm tries to save in a table that doesn't exist and thus it won't work
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrF("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening in the database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
