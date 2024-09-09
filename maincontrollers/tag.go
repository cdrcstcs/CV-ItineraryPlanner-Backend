package maincontrollers
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"itineraryplanner/maincontrollers/inf"
	"itineraryplanner/common/gin_ctx"
	con_inf "itineraryplanner/controllers/inf"
	"itineraryplanner/models"
)
func NewMainTagController(c con_inf.TagController) inf.MainTagController {
	return &MainTagController{
		Con: c,
	}
}
type MainTagController struct {
	Con con_inf.TagController
}
func (m *MainTagController) CreateTag(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.CreateTagReq](c, models.CreateTagReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    resp, err := m.Con.CreateTag(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating Tag: %s", err.Error())})
        return
    }
    c.JSON(http.StatusOK, resp)
}
func (m *MainTagController) GetTag(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetTagReq](c, models.GetTagReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    resp, err := m.Con.GetTag(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting Tag: %s", err.Error())})
        return
    }
    c.JSON(http.StatusOK, resp)
}
func (m *MainTagController) GetTagById(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetTagByIdReq](c, models.GetTagByIdReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    resp, err := m.Con.GetTagById(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting Tag by ID: %s", err.Error())})
        return
    }
    c.JSON(http.StatusOK, resp)
}
func (m *MainTagController) UpdateTag(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.UpdateTagReq](c, models.UpdateTagReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    resp, err := m.Con.UpdateTag(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating Tag: %s", err.Error())})
        return
    }
    c.JSON(http.StatusOK, resp)
}
func (m *MainTagController) DeleteTag(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.DeleteTagReq](c, models.DeleteTagReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    resp, err := m.Con.DeleteTag(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting Tag: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, resp)
}