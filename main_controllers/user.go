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



func NewMainCreateUserController(ac con_inf.CreateUserController) inf.MainCreateUserController {
	return &MainUserController{
		CCon: ac,
	}
}
func NewMainGetUserController(ac con_inf.GetUserController) inf.MainGetUserController {
	return &MainUserController{
		GCon: ac,
	}
}
func NewMainGetUserByIdController(ac con_inf.GetUserByIdController) inf.MainGetUserByIdController {
	return &MainUserController{
		BCon: ac,
	}
}
func NewMainUpdateUserController(ac con_inf.UpdateUserController) inf.MainUpdateUserController {
	return &MainUserController{
		UCon: ac,
	}
}
func NewMainDeleteUserController(ac con_inf.DeleteUserController) inf.MainDeleteUserController {
	return &MainUserController{
		DCon: ac,
	}
}
func NewMainLoginUserController(ac con_inf.LoginUserController) inf.MainLoginUserController {
	return &MainUserController{
		LCon: ac,
	}
}

type MainUserController struct {
	CCon con_inf.CreateUserController
	BCon con_inf.GetUserByIdController
	GCon con_inf.GetUserController
	UCon con_inf.UpdateUserController
	DCon con_inf.DeleteUserController
	LCon con_inf.LoginUserController
}

func (m *MainUserController) CreateUser(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.CreateUserReq](c, models.CreateUserReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.CCon.CreateUser(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainUserController) GetUser(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetUserReq](c, models.GetUserReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.GCon.GetUser(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainUserController) GetUserById(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.GetUserByIdReq](c, models.GetUserByIdReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.BCon.GetUserById(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainUserController) UpdateUser(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.UpdateUserReq](c, models.UpdateUserReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.UCon.UpdateUser(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainUserController) DeleteUser(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.DeleteUserReq](c, models.DeleteUserReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.DCon.DeleteUser(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

func (m *MainUserController) LoginUser(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.LoginUserReq](c, models.LoginUserReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.LCon.LoginUser(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}

