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
func TestCreateTag_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockTagService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.TagController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.CreateTagReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.CreateTagReq{
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
			req: &models.CreateTagReq{
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
				mockService.EXPECT().CreateTag(gomock.Any(), tt.req).Return(&models.CreateTagResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().CreateTag(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.CreateTag(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestGetTagById_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockTagService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.TagController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.GetTagByIdReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.GetTagByIdReq{
				Id: validUUID,
			},
			wantErr: nil,
		},
		{
			name: "invalid request (empty Id)",
			req: &models.GetTagByIdReq{
				Id: "",
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().GetTagById(gomock.Any(), tt.req).Return(&models.GetTagByIdResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().GetTagById(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.GetTagById(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestUpdateTag_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockTagService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.TagController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.UpdateTagReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.UpdateTagReq{
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
			req: &models.UpdateTagReq{
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
				mockService.EXPECT().UpdateTag(gomock.Any(), tt.req).Return(&models.UpdateTagResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().UpdateTag(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.UpdateTag(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestDeleteTag_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mock.NewMockTagService(ctrl)
	validUUID := uuid.New().String()
	controller := &controllers.TagController{
		Ser:       mockService,
		Validator: validator.New(),
	}
	tests := []struct {
		name    string
		req     *models.DeleteTagReq
		wantErr error
	}{
		{
			name: "valid request",
			req: &models.DeleteTagReq{
				Id: validUUID,
			},
			wantErr: nil,
		},
		{
			name: "invalid request (empty Id)",
			req: &models.DeleteTagReq{
				Id: "",
			},
			wantErr: custom_errs.ErrServiceError, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				mockService.EXPECT().DeleteTag(gomock.Any(), tt.req).Return(&models.DeleteTagResp{}, nil).MaxTimes(100)
			} else {
				mockService.EXPECT().DeleteTag(gomock.Any(), tt.req).Return(nil, custom_errs.ErrServiceError).MaxTimes(100)
			}
			_, err := controller.DeleteTag(context.Background(), tt.req)
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}