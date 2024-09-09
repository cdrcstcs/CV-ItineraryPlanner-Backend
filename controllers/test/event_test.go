package test
import (
	"context"
	"itineraryplanner/common/custom_errs"
	"itineraryplanner/controllers"
	"itineraryplanner/models"
	"github.com/google/uuid" 
	"itineraryplanner/service/mock"
	"testing"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
)
func TestCreateEvent_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockEventService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.EventController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.CreateEventReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.CreateEventReq{
				Name:       "test",
				Address:    "test",
				X:          1,
				Y:          1,
				TagIDs:     []string{validUUID, validUUID},
				RatingId:   validUUID,
				City:       "test",
			},
			wantErr: nil,
		},
		{
			name: "invalid request (missing Name)",
			req: &models.CreateEventReq{
				Address:    "test",
				X:          1,
				Y:          1,
				TagIDs:     []string{validUUID, validUUID},
				RatingId:   validUUID,
				City:       "test",
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().CreateEvent(gomock.Any(), tt.req).Return(&models.CreateEventResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().CreateEvent(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.CreateEvent(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestGetEventById_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockEventService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.EventController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.GetEventByIdReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.GetEventByIdReq{
				Id: validUUID,
			},
			wantErr: nil,
		},
		{
			name: "invalid request (empty Id)",
			req: &models.GetEventByIdReq{
				Id: "",
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().GetEventById(gomock.Any(), tt.req).Return(&models.GetEventByIdResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().GetEventById(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.GetEventById(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestUpdateEvent_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockEventService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.EventController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.UpdateEventReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.UpdateEventReq{
				Id: validUUID,
				Name: "test",
				Address: "test",
				X: 1,
				Y: 1,
				TagIDs: []string{validUUID, validUUID},
				RatingId: validUUID,
				City: "test",
			},
			wantErr: nil,
		},
		{
			name: "invalid request (empty Id)",
			req: &models.UpdateEventReq{
				Id: "",
				Name: "test",
				Address: "test",
				X: 1,
				Y: 1,
				TagIDs: []string{validUUID, validUUID},
				RatingId: validUUID,
				City: "test",
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().UpdateEvent(gomock.Any(), tt.req).Return(&models.UpdateEventResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().UpdateEvent(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.UpdateEvent(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestDeleteEvent_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockEventService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.EventController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.DeleteEventReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.DeleteEventReq{
				Id: validUUID,
			},
			wantErr: nil,
		},
		{
			name: "invalid request (empty Id)",
			req: &models.DeleteEventReq{
				Id: "",
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().DeleteEvent(gomock.Any(), tt.req).Return(&models.DeleteEventResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().DeleteEvent(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.DeleteEvent(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}