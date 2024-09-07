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



func NewMainEventController(c con_inf.EventController) inf.MainEventController {
	return &MainEventController{
		Con: c,
	}
}
type MainEventController struct {
	Con con_inf.EventController
}

func (m *MainEventController) CreateEvent(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.CreateEventReq](c, models.CreateEventReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.CreateEvent(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating event: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (m *MainEventController) GetEvent(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetEventReq](c, models.GetEventReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.GetEvent(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting event: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (m *MainEventController) GetEventById(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetEventByIdReq](c, models.GetEventByIdReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.GetEventById(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting event by ID: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (m *MainEventController) UpdateEvent(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.UpdateEventReq](c, models.UpdateEventReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.UpdateEvent(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating event: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (m *MainEventController) DeleteEvent(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.DeleteEventReq](c, models.DeleteEventReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.DeleteEvent(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting event: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}
