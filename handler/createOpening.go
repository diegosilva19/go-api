package handler

import (
	"net/http"

	"github.com/diegosilva19/go-api/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	/*
		request := struct {
			Role string `json:"role"`
		}{}*/

	request := CreateOpeningRequest{}

	ctx.Bind(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error")
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
		logger.Errorf("erro ao criar: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Infof("request received %+v", request)
	ctx.JSON(http.StatusOK, &opening)
}
