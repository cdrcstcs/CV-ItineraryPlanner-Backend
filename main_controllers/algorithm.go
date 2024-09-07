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



func NewMainRecommendItineraryController(ac con_inf.RecommendItineraryController) inf.MainRecommendItineraryController {
	return &MainAlgoController{
		RCon: ac,
	}
}
func NewMainBuildItineraryController(ac con_inf.BuildItineraryController) inf.MainBuildItineraryController {
	return &MainAlgoController{
		BCon: ac,
	}
}
type MainAlgoController struct {
	RCon con_inf.RecommendItineraryController
	BCon con_inf.BuildItineraryController
}

func (m *MainAlgoController) RecommendItinerary(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.RecommendItineraryReq](c, models.RecommendItineraryReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.RCon.RecommendItinerary(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainAlgoController) BuildItinerary(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.BuildItineraryReq](c, models.BuildItineraryReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.BCon.BuildItinerary(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}