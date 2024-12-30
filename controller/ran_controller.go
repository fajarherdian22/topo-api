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

type RanController struct {
	RanService *service.RanServiceImpl
	Validate   *validator.Validate
}

func NewRanController(RanService *service.RanServiceImpl, validate *validator.Validate) *RanController {
	return &RanController{
		RanService: RanService,
		Validate:   validate,
	}
}

func (controller *RanController) ListLevel(c *gin.Context) {
	type LevelReq struct {
		Level string `json:"level" binding:"required,LevelName"`
	}
	var req LevelReq

	if err := c.ShouldBindJSON(&req); err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	if err := controller.Validate.Struct(req); err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	ctx := c.Request.Context()

	payload, err := controller.RanService.ListLevel(ctx, req.Level)
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

func (controller *RanController) GetAllData(c *gin.Context) {
	type LevelReq struct {
		Level string `json:"level" binding:"required,LevelName"`
	}
	var req LevelReq

	if err := c.ShouldBindJSON(&req); err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	if err := controller.Validate.Struct(req); err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	ctx := c.Request.Context()

	payload, err := controller.RanService.GetAllData(ctx, req.Level)
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

func (controller *RanController) GetByLevel(c *gin.Context) {
	var req repository.GetByLevelAndNameParams

	if err := c.ShouldBindJSON(&req); err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	if err := controller.Validate.Struct(req); err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	ctx := c.Request.Context()

	payload, err := controller.RanService.GetByLevel(ctx, req)
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

func (controller *RanController) GetByReference(c *gin.Context) {
	var req repository.GetByLevelAndReferenceParams

	if err := c.ShouldBindJSON(&req); err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	if err := controller.Validate.Struct(req); err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	ctx := c.Request.Context()

	payload, err := controller.RanService.GetByReference(ctx, req)
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
