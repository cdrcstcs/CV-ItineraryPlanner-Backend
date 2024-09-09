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
func TestCreateRating(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainRatingController(ctrl)
	r := gin.Default()
	route.RouteR(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	Rating := &models.CreateRatingReq{
		Score: 1,

	}
	body, err := json.Marshal(Rating)
	assert.Nil(t, err)
	mockController.EXPECT().
		CreateRating(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	resp, err := http.Post(ts.URL+"/rating", "application/json", bytes.NewReader(body))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestGetRating(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainRatingController(ctrl)
	r := gin.Default()
	route.RouteR(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	mockController.EXPECT().
		GetRating(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	resp, err := http.Get(ts.URL + "/rating")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestGetRatingById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainRatingController(ctrl)
	r := gin.Default()
	route.RouteR(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	mockController.EXPECT().
		GetRatingById(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	resp, err := http.Get(ts.URL + "/rating/123")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestUpdateRating(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainRatingController(ctrl)
	r := gin.Default()
	route.RouteR(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	validUUID := uuid.New().String()
	Rating := &models.UpdateRatingReq{
		Id:          validUUID,
		Score: 1,
	}
	body, err := json.Marshal(Rating)
	assert.Nil(t, err)
	mockController.EXPECT().
		UpdateRating(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	req, err := http.NewRequest(http.MethodPut, ts.URL+"/rating", bytes.NewReader(body))
	assert.Nil(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestDeleteRating(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainRatingController(ctrl)
	r := gin.Default()
	route.RouteR(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	mockController.EXPECT().
		DeleteRating(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	req, err := http.NewRequest(http.MethodDelete, ts.URL+"/rating/123", nil)
	assert.Nil(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
