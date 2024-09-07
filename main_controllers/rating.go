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



func NewMainRatingController(c con_inf.RatingController) inf.MainRatingController {
	return &MainRatingController{
		Con: c,
	}
}

type MainRatingController struct {
	Con con_inf.RatingController
}

func (m *MainRatingController) CreateRating(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.CreateRatingReq](c, models.CreateRatingReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.CreateRating(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating Rating: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (m *MainRatingController) GetRating(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetRatingReq](c, models.GetRatingReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.GetRating(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting Rating: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (m *MainRatingController) GetRatingById(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetRatingByIdReq](c, models.GetRatingByIdReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.GetRatingById(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting Rating by ID: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (m *MainRatingController) UpdateRating(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.UpdateRatingReq](c, models.UpdateRatingReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.UpdateRating(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating Rating: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (m *MainRatingController) DeleteRating(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.DeleteRatingReq](c, models.DeleteRatingReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    
    resp, err := m.Con.DeleteRating(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting Rating: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}
