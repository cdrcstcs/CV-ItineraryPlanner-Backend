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



func NewMainSearchEngineController(s con_inf.SearchEngineController) inf.MainSearchEngineController {
	return &MainSearchEngineController{
		Con: s,
	}
}
type MainSearchEngineController struct {
	Con con_inf.SearchEngineController
}

func (m *MainSearchEngineController) SearchEngine(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.SearchEngineReq](c, models.SearchEngineReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.Con.SearchEngine(ctx, &req)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error search engine: %s", err.Error())})
        return
    }
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}
