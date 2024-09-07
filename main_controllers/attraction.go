package main_controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"itineraryplanner/main_controllers/inf"
	"itineraryplanner/common/gin_ctx"
	con_inf "itineraryplanner/controllers/inf"
	"itineraryplanner/models"
)



func NewMainCreateAttractionController(ac con_inf.CreateAttractionController) inf.MainCreateAttractionController {
	return &MainAttractionController{
		CCon: ac,
	}
}
func NewMainGetAttractionController(ac con_inf.GetAttractionController) inf.MainGetAttractionController {
	return &MainAttractionController{
		GCon: ac,
	}
}
func NewMainGetAttractionByIdController(ac con_inf.GetAttractionByIdController) inf.MainGetAttractionByIdController {
	return &MainAttractionController{
		BCon: ac,
	}
}
func NewMainUpdateAttractionController(ac con_inf.UpdateAttractionController) inf.MainUpdateAttractionController {
	return &MainAttractionController{
		UCon: ac,
	}
}
func NewMainDeleteAttractionController(ac con_inf.DeleteAttractionController) inf.MainDeleteAttractionController {
	return &MainAttractionController{
		DCon: ac,
	}
}

type MainAttractionController struct {
	CCon con_inf.CreateAttractionController
	BCon con_inf.GetAttractionByIdController
	GCon con_inf.GetAttractionController
	UCon con_inf.UpdateAttractionController
	DCon con_inf.DeleteAttractionController
}

func (m *MainAttractionController) CreateAttraction(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.CreateAttractionReq](c, models.CreateAttractionReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.CCon.CreateAttraction(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainAttractionController) GetAttraction(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetAttractionReq](c, models.GetAttractionReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.GCon.GetAttraction(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainAttractionController) GetAttractionById(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetAttractionByIdReq](c, models.GetAttractionByIdReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.BCon.GetAttractionById(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainAttractionController) UpdateAttraction(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.UpdateAttractionReq](c, models.UpdateAttractionReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.UCon.UpdateAttraction(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainAttractionController) DeleteAttraction(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.DeleteAttractionReq](c, models.DeleteAttractionReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.DCon.DeleteAttraction(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

