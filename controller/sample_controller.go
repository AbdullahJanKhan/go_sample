package controller

import (
	"net/http"

	"github.com/abdullahjankhan/go_sample/models"
	"github.com/abdullahjankhan/go_sample/service"
	"github.com/abdullahjankhan/go_sample/utils"
	"github.com/gin-gonic/gin"
)

type SampleController interface {
	SampleFunc(ctx *gin.Context)
}

type sampleController struct {
	sampleService service.SampleService
	logger        utils.Logger
}

func NewSampleController(sampleService service.SampleService, logger utils.Logger) SampleController {
	return &sampleController{
		sampleService: sampleService,
		logger:        logger,
	}
}

func (sc *sampleController) SampleFunc(ctx *gin.Context) {
	data := GetBody[models.SampleRequest](ctx)
	err := data.Validate()
	if err != nil {
		sc.logger.Info("/api/sample/create was called with invalid body")
		ctx.JSON(http.StatusOK, models.NewStandardResponse(false, models.INVALID_INPUT, models.INVALID_INPUT_MESSAGE, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewStandardResponse(true, models.SUCCESS, models.SAMPLE_REQUEST_COMPLETE_MESSAGE, data))
}
