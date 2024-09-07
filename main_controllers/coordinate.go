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



func NewMainCreateCoordinateController(ac con_inf.CreateCoordinateController) inf.MainCreateCoordinateController {
	return &MainCoordinateController{
		CCon: ac,
	}
}
func NewMainGetCoordinateController(ac con_inf.GetCoordinateController) inf.MainGetCoordinateController {
	return &MainCoordinateController{
		GCon: ac,
	}
}
func NewMainGetCoordinateByIdController(ac con_inf.GetCoordinateByIdController) inf.MainGetCoordinateByIdController {
	return &MainCoordinateController{
		BCon: ac,
	}
}
func NewMainUpdateCoordinateController(ac con_inf.UpdateCoordinateController) inf.MainUpdateCoordinateController {
	return &MainCoordinateController{
		UCon: ac,
	}
}
func NewMainDeleteCoordinateController(ac con_inf.DeleteCoordinateController) inf.MainDeleteCoordinateController {
	return &MainCoordinateController{
		DCon: ac,
	}
}

type MainCoordinateController struct {
	CCon con_inf.CreateCoordinateController
	BCon con_inf.GetCoordinateByIdController
	GCon con_inf.GetCoordinateController
	UCon con_inf.UpdateCoordinateController
	DCon con_inf.DeleteCoordinateController
}

func (m *MainCoordinateController) CreateCoordinate(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.CreateCoordinateReq](c, models.CreateCoordinateReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.CCon.CreateCoordinate(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainCoordinateController) GetCoordinate(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetCoordinateReq](c, models.GetCoordinateReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.GCon.GetCoordinate(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainCoordinateController) GetCoordinateById(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetCoordinateByIdReq](c, models.GetCoordinateByIdReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.BCon.GetCoordinateById(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainCoordinateController) UpdateCoordinate(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.UpdateCoordinateReq](c, models.UpdateCoordinateReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.UCon.UpdateCoordinate(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainCoordinateController) DeleteCoordinate(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.DeleteCoordinateReq](c, models.DeleteCoordinateReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.DCon.DeleteCoordinate(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

