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
func TestCreateItinerary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainItineraryController(ctrl)
	r := gin.Default()
	route.RouteI(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	validUUID := uuid.New().String()
	Itinerary := &models.CreateItineraryReq{
		CopiedId: validUUID,
		CopiedName: "test",
		Name: "test",
		UserId: validUUID,
		StartTime: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
		EndTime: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
		EventIds: []string{validUUID,validUUID},
		EventCount: 1,
		RatingId: validUUID,
	}
	body, err := json.Marshal(Itinerary)
	assert.Nil(t, err)
	mockController.EXPECT().
		CreateItinerary(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	resp, err := http.Post(ts.URL+"/itinerary", "application/json", bytes.NewReader(body))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestGetItinerary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainItineraryController(ctrl)
	r := gin.Default()
	route.RouteI(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	mockController.EXPECT().
		GetItinerary(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	resp, err := http.Get(ts.URL + "/itinerary")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestGetItineraryById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainItineraryController(ctrl)
	r := gin.Default()
	route.RouteI(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	mockController.EXPECT().
		GetItineraryById(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	resp, err := http.Get(ts.URL + "/itinerary/123")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestUpdateItinerary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainItineraryController(ctrl)
	r := gin.Default()
	route.RouteI(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	validUUID := uuid.New().String()
	Itinerary := &models.UpdateItineraryReq{
		Id:          validUUID,
		CopiedId: validUUID,
		CopiedName: "test",
		Name: "test",
		UserId: validUUID,
		StartTime: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
		EndTime: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
		EventIds: []string{validUUID,validUUID},
		EventCount: 1,
		RatingId: validUUID,
	}
	body, err := json.Marshal(Itinerary)
	assert.Nil(t, err)
	mockController.EXPECT().
		UpdateItinerary(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	req, err := http.NewRequest(http.MethodPut, ts.URL+"/itinerary", bytes.NewReader(body))
	assert.Nil(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestDeleteItinerary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainItineraryController(ctrl)
	r := gin.Default()
	route.RouteI(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	mockController.EXPECT().
		DeleteItinerary(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	req, err := http.NewRequest(http.MethodDelete, ts.URL+"/itinerary/123", nil)
	assert.Nil(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
