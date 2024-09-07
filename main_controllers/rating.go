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



func NewMainCreateRatingController(ac con_inf.CreateRatingController) inf.MainCreateRatingController {
	return &MainRatingController{
		CCon: ac,
	}
}
func NewMainGetRatingController(ac con_inf.GetRatingController) inf.MainGetRatingController {
	return &MainRatingController{
		GCon: ac,
	}
}
func NewMainGetRatingByIdController(ac con_inf.GetRatingByIdController) inf.MainGetRatingByIdController {
	return &MainRatingController{
		BCon: ac,
	}
}
func NewMainUpdateRatingController(ac con_inf.UpdateRatingController) inf.MainUpdateRatingController {
	return &MainRatingController{
		UCon: ac,
	}
}
func NewMainDeleteRatingController(ac con_inf.DeleteRatingController) inf.MainDeleteRatingController {
	return &MainRatingController{
		DCon: ac,
	}
}

type MainRatingController struct {
	CCon con_inf.CreateRatingController
	BCon con_inf.GetRatingByIdController
	GCon con_inf.GetRatingController
	UCon con_inf.UpdateRatingController
	DCon con_inf.DeleteRatingController
}

func (m *MainRatingController) CreateRating(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.CreateRatingReq](c, models.CreateRatingReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.CCon.CreateRating(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainRatingController) GetRating(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetRatingReq](c, models.GetRatingReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.GCon.GetRating(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainRatingController) GetRatingById(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetRatingByIdReq](c, models.GetRatingByIdReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.BCon.GetRatingById(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainRatingController) UpdateRating(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.UpdateRatingReq](c, models.UpdateRatingReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.UCon.UpdateRating(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainRatingController) DeleteRating(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.DeleteRatingReq](c, models.DeleteRatingReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.DCon.DeleteRating(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

