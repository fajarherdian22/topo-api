package controller

import (
	"net/http"

	"github.com/fajarherdian22/topo-api/exception"
	"github.com/fajarherdian22/topo-api/helper"
	"github.com/fajarherdian22/topo-api/repository"
	"github.com/fajarherdian22/topo-api/service"
	"github.com/fajarherdian22/topo-api/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SpatialController struct {
	SpatialService *service.KabKotaServiceImpl
	Validate       *validator.Validate
}

func NewSpatialController(SpatialService *service.KabKotaServiceImpl, validate *validator.Validate) *SpatialController {
	return &SpatialController{
		SpatialService: SpatialService,
		Validate:       validate,
	}
}

func (controller *SpatialController) GetAllSpatial(c *gin.Context) {
	payload, err := controller.SpatialService.GetAllSpatial(c.Request.Context())
	if err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	WebResponse := web.WebResponse{
		Code:   http.StatusOK,
		Data:   payload,
		Status: "OK",
	}

	helper.HandleEncodeWriteJson(c, WebResponse)
}

func (controller *SpatialController) GetSpatialByFilter(c *gin.Context) {
	var req repository.GetSpatialLv

	if err := c.ShouldBindJSON(&req); err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	if err := controller.Validate.Struct(req); err != nil {
		exception.ErrorHandler(c, err)
		return
	}
	ctx := c.Request.Context()

	payload, err := controller.SpatialService.GetSpatialByFilter(ctx, req)
	if err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	WebResponse := web.WebResponse{
		Code:   http.StatusOK,
		Data:   payload,
		Status: "OK",
	}

	helper.HandleEncodeWriteJson(c, WebResponse)
}
