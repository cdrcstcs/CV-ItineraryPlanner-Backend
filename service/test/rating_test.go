package test
import (
	"context"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"itineraryplanner/dal/mock"
	"itineraryplanner/models"
	"itineraryplanner/service"
)
func ConfigRating(t *testing.T) (context.Context, *mock.MockRatingDal, *service.RatingService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mock := mock.NewMockRatingDal(ctrl)
	ratingService := &service.RatingService{Dal: mock}
	return ctx, mock, ratingService
}
func TestCreateRating(t *testing.T) {
	ctx, mock, ratingService := ConfigRating(t)
	type arg struct {
		req *models.CreateRatingReq
		ctx context.Context
	}
	tests := []struct {
		name     string
		arg
		before   func(t *testing.T)
		wantResp *models.CreateRatingResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.CreateRatingReq{
					Score:    5,
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().CreateRating(ctx, gomock.Any()).Return(
					&models.Rating{
						Id:       "test",
						Score:    5,
					},
					nil,
				)
			},
			wantResp: &models.CreateRatingResp{
				Rating: &models.RatingDTO{
					Id:       "test",
					Score:    5,
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := ratingService.CreateRating(tt.arg.ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
func TestGetRatingById(t *testing.T) {
	ctx, mock, ratingService := ConfigRating(t)
	type arg struct {
		req *models.GetRatingByIdReq
		ctx context.Context
	}
	tests := []struct {
		name     string
		arg
		before   func(t *testing.T)
		wantResp *models.GetRatingByIdResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.GetRatingByIdReq{
					Id: "test",
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().GetRatingById(
					ctx,
					"test",
				).Return(
					&models.Rating{
						Id:       "test",
						Score:    5,
					},
					nil,
				)
			},
			wantResp: &models.GetRatingByIdResp{
				Rating: &models.RatingDTO{
					Id:       "test",
					Score:    5,
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := ratingService.GetRatingById(tt.arg.ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
func TestGetRating(t *testing.T) {
	ctx, mock, ratingService := ConfigRating(t)
	type arg struct {
		req *models.GetRatingReq
		ctx context.Context
	}
	tests := []struct {
		name     string
		arg
		before   func(t *testing.T)
		wantResp *models.GetRatingResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.GetRatingReq{
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().GetRating(ctx).Return(
					[]*models.Rating{
						{
							Id:       "test",
							Score:    5,
						},
					},
					nil,
				)
			},
			wantResp: &models.GetRatingResp{
				Ratings: []*models.RatingDTO{
					{
						Id:       "test",
						Score:    5,
					},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := ratingService.GetRating(tt.arg.ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
func TestUpdateRating(t *testing.T) {
	ctx, mock, ratingService := ConfigRating(t)
	type arg struct {
		req *models.UpdateRatingReq
		ctx context.Context
	}
	tests := []struct {
		name     string
		arg
		before   func(t *testing.T)
		wantResp *models.UpdateRatingResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.UpdateRatingReq{
					Id:       "test",
					Score:    6,
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().UpdateRating(ctx, gomock.Any()).Return(
					&models.Rating{
						Id:       "test",
						Score:    6,
					},
					nil,
				)
			},
			wantResp: &models.UpdateRatingResp{
				Rating: &models.RatingDTO{
					Id:       "test",
					Score:    6,
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := ratingService.UpdateRating(tt.arg.ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
func TestDeleteRating(t *testing.T) {
	ctx, mock, ratingService := ConfigRating(t)
	type arg struct {
		req *models.DeleteRatingReq
		ctx context.Context
	}
	tests := []struct {
		name     string
		arg
		before   func(t *testing.T)
		wantResp *models.DeleteRatingResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.DeleteRatingReq{
					Id: "test",
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().DeleteRating(
					ctx,
					"test",
				).Return(
					&models.Rating{
						Id:       "test",
						Score:    5,
					},
					nil,
				)
			},
			wantResp: &models.DeleteRatingResp{
				Rating: &models.RatingDTO{
					Id:       "test",
					Score:    5,
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := ratingService.DeleteRating(tt.arg.ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}