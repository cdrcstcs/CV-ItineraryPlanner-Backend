package test
import (
	"context"
	"testing"
	"time"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"itineraryplanner/dal/mock"
	"itineraryplanner/models"
	"itineraryplanner/service"
)
func ConfigEvent(t *testing.T) (context.Context, *mock.MockEventDal, *service.EventService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mock := mock.NewMockEventDal(ctrl)
	eventService := &service.EventService{Dal: mock}
	return ctx, mock, eventService
}
func TestCreateEvent(t *testing.T) {
	ctx, mock, eventService := ConfigEvent(t)
	type arg struct {
		req *models.CreateEventReq
		ctx context.Context
	}
	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.CreateEventResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.CreateEventReq{
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					AttractionId:  "test",
					AttractionName: "test",
					Description: "test",
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().CreateEvent(
					ctx,
					gomock.Any(),
				).Return(
					&models.Event{
						StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
						EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
						AttractionId:  "test",
						AttractionName: "test",
						Description: "test",
					},
					nil,
				)
			},
			wantResp: &models.CreateEventResp{
				Event: &models.EventDTO{
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					AttractionId:  "test",
					AttractionName: "test",
					Description: "test",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := eventService.CreateEvent(ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
func TestGetEventById(t *testing.T) {
	ctx, mock, eventService := ConfigEvent(t)
	type arg struct {
		req *models.GetEventByIdReq
		ctx context.Context
	}
	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.GetEventByIdResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.GetEventByIdReq{
					Id: "test",
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().GetEventById(
					ctx,
					gomock.Any(),
				).Return(
					&models.Event{
						Id: "test",
						StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
						EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
						AttractionId:  "test",
						AttractionName: "test",
						Description: "test",
					},
					nil,
				)
			},
			wantResp: &models.GetEventByIdResp{
				Event: &models.EventDTO{
					Id: "test",
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					AttractionId:  "test",
					AttractionName: "test",
					Description: "test",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := eventService.GetEventById(ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
func TestGetEvent(t *testing.T) {
	ctx, mock, eventService := ConfigEvent(t)
	type arg struct {
		req *models.GetEventReq
		ctx context.Context
	}
	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.GetEventResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.GetEventReq{
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().GetEvent(ctx).Return(
					[]*models.Event{
						{
							Id: "test",
							StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
							EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
							AttractionId:  "test",
							AttractionName: "test",
							Description: "test",
						},
					},
					nil,
				)
			},
			wantResp: &models.GetEventResp{
				Events: []*models.EventDTO{
					{
						Id: "test",
						StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
						EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
						AttractionId:  "test",
						AttractionName: "test",
						Description: "test",
					},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := eventService.GetEvent(ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
func TestUpdateEvent(t *testing.T) {
	ctx, mock, eventService := ConfigEvent(t)
	type arg struct {
		req *models.UpdateEventReq
		ctx context.Context
	}
	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.UpdateEventResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.UpdateEventReq{
					Id: "test",
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					AttractionId:  "test",
					AttractionName: "test1",
					Description: "test",
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().UpdateEvent(
					ctx,
					gomock.Any(),
				).Return(
					&models.Event{
						Id: "test",
						StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
						EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
						AttractionId:  "test",
						AttractionName: "test1",
						Description: "test",
					},
					nil,
				)
			},
			wantResp: &models.UpdateEventResp{
				Event: &models.EventDTO{
					Id: "test",
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					AttractionId:  "test",
					AttractionName: "test1",
					Description: "test",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := eventService.UpdateEvent(ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
func TestDeleteEvent(t *testing.T) {
	ctx, mock, eventService := ConfigEvent(t)
	type arg struct {
		req *models.DeleteEventReq
		ctx context.Context
	}
	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.DeleteEventResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.DeleteEventReq{
					Id: "test",
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().DeleteEvent(
					ctx,
					gomock.Any(),
				).Return(
					&models.Event{
						Id: "test",
						StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
						EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
						AttractionId:  "test",
						AttractionName: "test",
						Description: "test",
					},
					nil,
				)
			},
			wantResp: &models.DeleteEventResp{
				Event: &models.EventDTO{
					Id: "test",
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					AttractionId:  "test",
					AttractionName: "test",
					Description: "test",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := eventService.DeleteEvent(ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}