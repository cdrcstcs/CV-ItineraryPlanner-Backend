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



func NewMainItineraryController(c con_inf.ItineraryController) inf.MainItineraryController {
	return &MainItineraryController{
		Con: c,
	}
}

type MainItineraryController struct {
	Con con_inf.ItineraryController
}

func (m *MainItineraryController) CreateItinerary(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.CreateItineraryReq](c, models.CreateItineraryReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.CreateItinerary(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating Itinerary: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (m *MainItineraryController) GetItinerary(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetItineraryReq](c, models.GetItineraryReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.GetItinerary(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting Itinerary: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (m *MainItineraryController) GetItineraryById(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetItineraryByIdReq](c, models.GetItineraryByIdReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.GetItineraryById(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting Itinerary by ID: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (m *MainItineraryController) UpdateItinerary(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.UpdateItineraryReq](c, models.UpdateItineraryReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.UpdateItinerary(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating Itinerary: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (m *MainItineraryController) DeleteItinerary(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.DeleteItineraryReq](c, models.DeleteItineraryReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.DeleteItinerary(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting Itinerary: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}
