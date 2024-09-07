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



func NewMainCreateEventController(ac con_inf.CreateEventController) inf.MainCreateEventController {
	return &MainEventController{
		CCon: ac,
	}
}
func NewMainGetEventController(ac con_inf.GetEventController) inf.MainGetEventController {
	return &MainEventController{
		GCon: ac,
	}
}
func NewMainGetEventByIdController(ac con_inf.GetEventByIdController) inf.MainGetEventByIdController {
	return &MainEventController{
		BCon: ac,
	}
}
func NewMainUpdateEventController(ac con_inf.UpdateEventController) inf.MainUpdateEventController {
	return &MainEventController{
		UCon: ac,
	}
}
func NewMainDeleteEventController(ac con_inf.DeleteEventController) inf.MainDeleteEventController {
	return &MainEventController{
		DCon: ac,
	}
}

type MainEventController struct {
	CCon con_inf.CreateEventController
	BCon con_inf.GetEventByIdController
	GCon con_inf.GetEventController
	UCon con_inf.UpdateEventController
	DCon con_inf.DeleteEventController
}

func (m *MainEventController) CreateEvent(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.CreateEventReq](c, models.CreateEventReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.CCon.CreateEvent(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainEventController) GetEvent(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetEventReq](c, models.GetEventReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.GCon.GetEvent(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainEventController) GetEventById(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetEventByIdReq](c, models.GetEventByIdReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.BCon.GetEventById(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainEventController) UpdateEvent(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.UpdateEventReq](c, models.UpdateEventReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.UCon.UpdateEvent(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainEventController) DeleteEvent(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.DeleteEventReq](c, models.DeleteEventReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.DCon.DeleteEvent(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

