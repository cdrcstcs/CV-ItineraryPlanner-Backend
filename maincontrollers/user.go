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
func NewMainUserController(c con_inf.UserController) inf.MainUserController {
	return &MainUserController{
		Con: c,
	}
}
type MainUserController struct {
	Con con_inf.UserController
}
func (m *MainUserController) CreateUser(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.CreateUserReq](c, models.CreateUserReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    resp, err := m.Con.CreateUser(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating User: %s", err.Error())})
        return
    }
    c.JSON(http.StatusOK, resp)
}
func (m *MainUserController) GetUser(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetUserReq](c, models.GetUserReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    resp, err := m.Con.GetUser(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting User: %s", err.Error())})
        return
    }
    c.JSON(http.StatusOK, resp)
}
func (m *MainUserController) GetUserById(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetUserByIdReq](c, models.GetUserByIdReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    resp, err := m.Con.GetUserById(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting User by ID: %s", err.Error())})
        return
    }
    c.JSON(http.StatusOK, resp)
}
func (m *MainUserController) UpdateUser(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.UpdateUserReq](c, models.UpdateUserReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    resp, err := m.Con.UpdateUser(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating User: %s", err.Error())})
        return
    }
    c.JSON(http.StatusOK, resp)
}
func (m *MainUserController) DeleteUser(c *gin.Context) {
    ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.DeleteUserReq](c, models.DeleteUserReq{})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
        return
    }
    resp, err := m.Con.DeleteUser(ctx, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting User: %s", err.Error())})
        return
    }
    c.JSON(http.StatusOK, resp)
}