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



func NewMainCreateItineraryController(ac con_inf.CreateItineraryController) inf.MainCreateItineraryController {
	return &MainItineraryController{
		CCon: ac,
	}
}
func NewMainGetItineraryController(ac con_inf.GetItineraryController) inf.MainGetItineraryController {
	return &MainItineraryController{
		GCon: ac,
	}
}
func NewMainGetItineraryByIdController(ac con_inf.GetItineraryByIdController) inf.MainGetItineraryByIdController {
	return &MainItineraryController{
		BCon: ac,
	}
}
func NewMainUpdateItineraryController(ac con_inf.UpdateItineraryController) inf.MainUpdateItineraryController {
	return &MainItineraryController{
		UCon: ac,
	}
}
func NewMainDeleteItineraryController(ac con_inf.DeleteItineraryController) inf.MainDeleteItineraryController {
	return &MainItineraryController{
		DCon: ac,
	}
}

type MainItineraryController struct {
	CCon con_inf.CreateItineraryController
	BCon con_inf.GetItineraryByIdController
	GCon con_inf.GetItineraryController
	UCon con_inf.UpdateItineraryController
	DCon con_inf.DeleteItineraryController
}

func (m *MainItineraryController) CreateItinerary(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.CreateItineraryReq](c, models.CreateItineraryReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.CCon.CreateItinerary(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainItineraryController) GetItinerary(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetItineraryReq](c, models.GetItineraryReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.GCon.GetItinerary(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainItineraryController) GetItineraryById(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetItineraryByIdReq](c, models.GetItineraryByIdReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.BCon.GetItineraryById(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainItineraryController) UpdateItinerary(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.UpdateItineraryReq](c, models.UpdateItineraryReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.UCon.UpdateItinerary(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainItineraryController) DeleteItinerary(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.DeleteItineraryReq](c, models.DeleteItineraryReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.DCon.DeleteItinerary(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

