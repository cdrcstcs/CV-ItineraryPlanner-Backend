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
func TestCreateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainEventController(ctrl)
	r := gin.Default()
	route.RouteE(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	validUUID := uuid.New().String()
	Event := &models.CreateEventReq{
		StartTime: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
		EndTime: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
		AttractionId: validUUID,
		AttractionName: "test",
		Description: "test",
	}
	body, err := json.Marshal(Event)
	assert.Nil(t, err)
	mockController.EXPECT().
		CreateEvent(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	resp, err := http.Post(ts.URL+"/event", "application/json", bytes.NewReader(body))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestGetEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainEventController(ctrl)
	r := gin.Default()
	route.RouteE(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	mockController.EXPECT().
		GetEvent(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	resp, err := http.Get(ts.URL + "/event")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestGetEventById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainEventController(ctrl)
	r := gin.Default()
	route.RouteE(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	mockController.EXPECT().
		GetEventById(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	resp, err := http.Get(ts.URL + "/event/123")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestUpdateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainEventController(ctrl)
	r := gin.Default()
	route.RouteE(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	validUUID := uuid.New().String()
	Event := &models.UpdateEventReq{
		Id:          validUUID,
		StartTime: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
		EndTime: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
		AttractionId: validUUID,
		AttractionName: "test",
		Description: "test",
	}
	body, err := json.Marshal(Event)
	assert.Nil(t, err)
	mockController.EXPECT().
		UpdateEvent(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	req, err := http.NewRequest(http.MethodPut, ts.URL+"/event", bytes.NewReader(body))
	assert.Nil(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestDeleteEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockController := mock.NewMockMainEventController(ctrl)
	r := gin.Default()
	route.RouteE(r, mockController)
	ts := httptest.NewServer(r)
	defer ts.Close()
	mockController.EXPECT().
		DeleteEvent(gomock.Any()).
		DoAndReturn(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}).MaxTimes(100)
	req, err := http.NewRequest(http.MethodDelete, ts.URL+"/event/123", nil)
	assert.Nil(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
