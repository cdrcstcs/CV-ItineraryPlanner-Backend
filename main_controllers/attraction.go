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



func NewMainAttractionController(c con_inf.AttractionController) inf.MainAttractionController {
	return &MainAttractionController{
		Con: c,
	}
}
type MainAttractionController struct {
	Con con_inf.AttractionController
}

func (m *MainAttractionController) CreateAttraction(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.CreateAttractionReq](c, models.CreateAttractionReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.Con.CreateAttraction(ctx, &req)
	if err != nil {
        // Handle error returned from service
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating attraction: %s", err.Error())})
        return
    }
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainAttractionController) GetAttraction(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetAttractionReq](c, models.GetAttractionReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.Con.GetAttraction(ctx, &req)
	if err != nil {
        // Handle error returned from service
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting attraction: %s", err.Error())})
        return
    }
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainAttractionController) GetAttractionById(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetAttractionByIdReq](c, models.GetAttractionByIdReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.Con.GetAttractionById(ctx, &req)
	if err != nil {
        // Handle error returned from service
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting by id attraction: %s", err.Error())})
        return
    }
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainAttractionController) UpdateAttraction(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.UpdateAttractionReq](c, models.UpdateAttractionReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.Con.UpdateAttraction(ctx, &req)
	if err != nil {
        // Handle error returned from service
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating attraction: %s", err.Error())})
        return
    }
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainAttractionController) DeleteAttraction(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.DeleteAttractionReq](c, models.DeleteAttractionReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.Con.DeleteAttraction(ctx, &req)
	if err != nil {
        // Handle error returned from service
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting attraction: %s", err.Error())})
        return
    }
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

