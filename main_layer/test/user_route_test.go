package test
import (
	"bytes"
	"encoding/json"
	"itineraryplanner/main_layer/route"
	"itineraryplanner/maincontrollers/mock"
	"itineraryplanner/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"time"
)
func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainUserController(ctrl)
	r := gin.Default()
	route.RouteU(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	User := &models.CreateUserReq{
		Name: "test", 
		Password: "test",    
		Email: "test@gmail.com",       
		EmailPassword: "test",
		Phone: "test",         
		Token: "test",         
		UserType: "USER",    
		RefreshToken: "test",
		CreatedAt: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),   
		UpdatedAt: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
	}
	body, err := json.Marshal(User)
	assert.Nil(t, err)
	mockController.EXPECT().
		CreateUser(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	resp, err := http.Post(ts.URL+"/user", "application/json", bytes.NewReader(body))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainUserController(ctrl)
	r := gin.Default()
	route.RouteU(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	mockController.EXPECT().
		GetUser(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	resp, err := http.Get(ts.URL + "/user")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestGetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainUserController(ctrl)
	r := gin.Default()
	route.RouteU(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	mockController.EXPECT().
		GetUserById(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	resp, err := http.Get(ts.URL + "/user/123")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainUserController(ctrl)
	r := gin.Default()
	route.RouteU(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	validUUID := uuid.New().String()
	User := &models.UpdateUserReq{
		Id:          validUUID,
		Name: "test", 
		Password: "test",    
		Email: "test@gmail.com",       
		EmailPassword: "test",
		Phone: "test",         
		Token: "test",         
		UserType: "USER",    
		RefreshToken: "test",
		CreatedAt: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),   
		UpdatedAt: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
	}
	body, err := json.Marshal(User)
	assert.Nil(t, err)
	mockController.EXPECT().
		UpdateUser(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	req, err := http.NewRequest(http.MethodPut, ts.URL+"/user", bytes.NewReader(body))
	assert.Nil(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainUserController(ctrl)
	r := gin.Default()
	route.RouteU(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	mockController.EXPECT().
		DeleteUser(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	req, err := http.NewRequest(http.MethodDelete, ts.URL+"/user/123", nil)
	assert.Nil(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
