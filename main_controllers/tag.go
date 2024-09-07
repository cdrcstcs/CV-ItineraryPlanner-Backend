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



func NewMainCreateTagController(ac con_inf.CreateTagController) inf.MainCreateTagController {
	return &MainTagController{
		CCon: ac,
	}
}
func NewMainGetTagController(ac con_inf.GetTagController) inf.MainGetTagController {
	return &MainTagController{
		GCon: ac,
	}
}
func NewMainGetTagByIdController(ac con_inf.GetTagByIdController) inf.MainGetTagByIdController {
	return &MainTagController{
		BCon: ac,
	}
}
func NewMainUpdateTagController(ac con_inf.UpdateTagController) inf.MainUpdateTagController {
	return &MainTagController{
		UCon: ac,
	}
}
func NewMainDeleteTagController(ac con_inf.DeleteTagController) inf.MainDeleteTagController {
	return &MainTagController{
		DCon: ac,
	}
}

type MainTagController struct {
	CCon con_inf.CreateTagController
	BCon con_inf.GetTagByIdController
	GCon con_inf.GetTagController
	UCon con_inf.UpdateTagController
	DCon con_inf.DeleteTagController
}

func (m *MainTagController) CreateTag(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.CreateTagReq](c, models.CreateTagReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.CCon.CreateTag(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainTagController) GetTag(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetTagReq](c, models.GetTagReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.GCon.GetTag(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainTagController) GetTagById(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetTagByIdReq](c, models.GetTagByIdReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.BCon.GetTagById(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainTagController) UpdateTag(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.UpdateTagReq](c, models.UpdateTagReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.UCon.UpdateTag(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainTagController) DeleteTag(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.DeleteTagReq](c, models.DeleteTagReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.DCon.DeleteTag(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

