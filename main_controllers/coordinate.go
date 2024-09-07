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



func NewMainCoordinateController(c con_inf.CoordinateController) inf.MainCoordinateController {
	return &MainCoordinateController{
		Con: c,
	}
}

type MainCoordinateController struct {
	Con con_inf.CoordinateController
}

func (m *MainCoordinateController) CreateCoordinate(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.CreateCoordinateReq](c, models.CreateCoordinateReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.CreateCoordinate(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating coordinate: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (m *MainCoordinateController) GetCoordinate(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetCoordinateReq](c, models.GetCoordinateReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.GetCoordinate(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting coordinate: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (m *MainCoordinateController) GetCoordinateById(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetCoordinateByIdReq](c, models.GetCoordinateByIdReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.GetCoordinateById(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting coordinate by ID: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (m *MainCoordinateController) UpdateCoordinate(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.UpdateCoordinateReq](c, models.UpdateCoordinateReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.UpdateCoordinate(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating coordinate: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (m *MainCoordinateController) DeleteCoordinate(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.DeleteCoordinateReq](c, models.DeleteCoordinateReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.DeleteCoordinate(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting coordinate: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}
