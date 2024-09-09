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
)
func TestCreateTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainTagController(ctrl)
	r := gin.Default()
	route.RouteT(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	Tag := &models.CreateTagReq{
		Value: "test",
	}
	body, err := json.Marshal(Tag)
	assert.Nil(t, err)
	mockController.EXPECT().
		CreateTag(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	resp, err := http.Post(ts.URL+"/tag", "application/json", bytes.NewReader(body))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestGetTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainTagController(ctrl)
	r := gin.Default()
	route.RouteT(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	mockController.EXPECT().
		GetTag(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	resp, err := http.Get(ts.URL + "/tag")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestGetTagById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainTagController(ctrl)
	r := gin.Default()
	route.RouteT(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	mockController.EXPECT().
		GetTagById(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	resp, err := http.Get(ts.URL + "/tag/123")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestUpdateTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainTagController(ctrl)
	r := gin.Default()
	route.RouteT(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	validUUID := uuid.New().String()
	Tag := &models.UpdateTagReq{
		Id:          validUUID,
		Value: "test",
	}
	body, err := json.Marshal(Tag)
	assert.Nil(t, err)
	mockController.EXPECT().
		UpdateTag(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	req, err := http.NewRequest(http.MethodPut, ts.URL+"/tag", bytes.NewReader(body))
	assert.Nil(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestDeleteTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainTagController(ctrl)
	r := gin.Default()
	route.RouteT(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	mockController.EXPECT().
		DeleteTag(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	req, err := http.NewRequest(http.MethodDelete, ts.URL+"/tag/123", nil)
	assert.Nil(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
