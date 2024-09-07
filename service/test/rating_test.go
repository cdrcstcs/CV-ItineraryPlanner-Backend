package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"itineraryplanner/dal/mock"
	"itineraryplanner/models"
	"itineraryplanner/service"
	ser_mock "itineraryplanner/service/mock"
)

func ConfigCreateRating(t *testing.T) (context.Context, *mock.MockCreateRatingDal, *ser_mock.MockRatingDTOService, *service.RatingService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockC := mock.NewMockCreateRatingDal(ctrl)
	mockDTO := ser_mock.NewMockRatingDTOService(ctrl)
	ratingService := &service.RatingService{CDal: mockC}
	return ctx, mockC, mockDTO, ratingService
}

func ConfigGetRating(t *testing.T) (context.Context, *mock.MockGetRatingDal, *ser_mock.MockRatingDTOService, *service.RatingService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockG := mock.NewMockGetRatingDal(ctrl)
	mockDTO := ser_mock.NewMockRatingDTOService(ctrl)
	ratingService := &service.RatingService{GDal: mockG}
	return ctx, mockG, mockDTO, ratingService
}

func ConfigGetRatingById(t *testing.T) (context.Context, *mock.MockGetRatingByIdDal, *ser_mock.MockRatingDTOService, *service.RatingService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockB := mock.NewMockGetRatingByIdDal(ctrl)
	mockDTO := ser_mock.NewMockRatingDTOService(ctrl)
	ratingService := &service.RatingService{BDal: mockB}
	return ctx, mockB, mockDTO, ratingService
}

func ConfigUpdateRating(t *testing.T) (context.Context, *mock.MockUpdateRatingDal, *ser_mock.MockRatingDTOService, *service.RatingService){
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockU := mock.NewMockUpdateRatingDal(ctrl)
	mockDTO := ser_mock.NewMockRatingDTOService(ctrl)
	ratingService := &service.RatingService{UDal: mockU}
	return ctx, mockU, mockDTO, ratingService
}
func ConfigDeleteRating(t *testing.T) (context.Context, *mock.MockDeleteRatingDal, *ser_mock.MockRatingDTOService, *service.RatingService){
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockD := mock.NewMockDeleteRatingDal(ctrl)
	mockDTO := ser_mock.NewMockRatingDTOService(ctrl)
	ratingService := &service.RatingService{DDal: mockD}
	return ctx, mockD, mockDTO, ratingService
}

func ConfigConvertToDTORating(t *testing.T) (context.Context, *ser_mock.MockFacadeDesignPatternService, *service.RatingService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockC := mock.NewMockCreateRatingDal(ctrl)
	mockG := mock.NewMockGetRatingDal(ctrl)
	mockF := ser_mock.NewMockFacadeDesignPatternService(ctrl)
	ratingService := &service.RatingService{CDal: mockC, GDal: mockG}
	return ctx, mockF, ratingService
}

func TestCreateRating(t *testing.T) {
	ctx, mockC, mockDTO, ratingService := ConfigCreateRating(t)

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
					Type:     models.TypeAttraction,
					UserId:   "test_user_id",
					ObjectId: "test_attraction_id",
					Score:    5,
				},
			},
			before: func(t *testing.T) {
				mockC.EXPECT().CreateRating(ctx, gomock.Any()).Return(
					&models.Rating{
						Id:       "test_rating_id",
						Type:     models.TypeAttraction,
						UserId:   "test_user_id",
						ObjectId: "test_attraction_id",
						Score:    5,
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTORating(
					ctx,
					&models.Rating{
						Id:       "test_rating_id",
						Type:     models.TypeAttraction,
						UserId:   "test_user_id",
						ObjectId: "test_attraction_id",
						Score:    5,
					},
				).Return(
					&models.RatingDTO{
						Id:   "test_rating_id",
						Type: "ATTRACTION",
						User: &models.UserDTO{
							Id: "test_user_id", 
							Name: "test_user_name",
						},
						ObjectId: "test_attraction_id",
						Score: 5,
					},
					nil,
				).Do(func(ctx context.Context, rating *models.Rating) {
					fmt.Println("ConvertDBOToDTORating called with", rating)
				})
			},
			wantResp: &models.CreateRatingResp{
				Rating: &models.RatingDTO{
					Id:   "test_rating_id",
					Type: "ATTRACTION",
					User: &models.UserDTO{
						Id: "test_user_id", 
						Name: "test_user_name",
					},
					ObjectId: "test_attraction_id",
					Score: 5,
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
	ctx, mockB, mockDTO, ratingService := ConfigGetRatingById(t)

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
					Id: "test_rating_id",
				},
			},
			before: func(t *testing.T) {
				mockB.EXPECT().GetRatingById(
					ctx,
					"test_rating_id",
				).Return(
					&models.Rating{
						Id:       "test_rating_id",
						Type:     models.TypeAttraction,
						UserId:   "test_user_id",
						ObjectId: "test_attraction_id",
						Score:    5,
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTORating(
					ctx,
					gomock.Any(),
				).Return(
					&models.RatingDTO{
						Id:   "test_rating_id",
						Type: "ATTRACTION",
						User: &models.UserDTO{
							Id: "test_user_id", 
							Name: "test_user_name",
						},
						ObjectId: "test_attraction_id",
						Score: 5,
					},
					nil,
				).Do(func(ctx context.Context, rating *models.Rating) {
					fmt.Println("ConvertDBOToDTORating called with", rating)
				})
			},
			wantResp: &models.GetRatingByIdResp{
				Rating: &models.RatingDTO{
					Id:   "test_rating_id",
					Type: "ATTRACTION",
					User: &models.UserDTO{
						Id: "test_user_id", 
						Name: "test_user_name",
					},
					ObjectId: "test_attraction_id",
					Score: 5,
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
	ctx, mockG, mockDTO, ratingService := ConfigGetRating(t)

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
				mockG.EXPECT().GetRating(ctx).Return(
					[]*models.Rating{
						{
							Id:       "test_rating_id",
							Type:     models.TypeAttraction,
							UserId:   "test_user_id",
							ObjectId: "test_attraction_id",
							Score:    5,
						},
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTORating(
					ctx,
					gomock.Any(),
				).Return(
					&models.RatingDTO{
						Id:   "test_rating_id",
						Type: "ATTRACTION",
						User: &models.UserDTO{
							Id: "test_user_id", 
							Name: "test_user_name",
						},
						ObjectId: "test_attraction_id",
						Score: 5,
					},
					nil,
				).Do(func(ctx context.Context, rating *models.Rating) {
					fmt.Println("ConvertDBOToDTORating called with", rating)
				})
			},
			wantResp: &models.GetRatingResp{
				Ratings: []*models.RatingDTO{
					{
						Id:   "test_rating_id",
						Type: "ATTRACTION",
						User: &models.UserDTO{
							Id: "test_user_id", 
							Name: "test_user_name",
						},
						ObjectId: "test_attraction_id",
						Score: 5,
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
	ctx, mockU, mockDTO, ratingService := ConfigUpdateRating(t)

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
					Type:     models.TypeAttraction,
					UserId:   "test_user_id",
					ObjectId: "test_attraction_id",
					Score:    5,
				},
			},
			before: func(t *testing.T) {
				mockU.EXPECT().UpdateRating(ctx, gomock.Any()).Return(
					&models.Rating{
						Id:       "test_rating_id",
						Type:     models.TypeAttraction,
						UserId:   "test_user_id",
						ObjectId: "test_attraction_id",
						Score:    5,
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTORating(
					ctx,
					&models.Rating{
						Id:       "test_rating_id",
						Type:     models.TypeAttraction,
						UserId:   "test_user_id",
						ObjectId: "test_attraction_id",
						Score:    5,
					},
				).Return(
					&models.RatingDTO{
						Id:   "test_rating_id",
						Type: "ATTRACTION",
						User: &models.UserDTO{
							Id: "test_user_id", 
							Name: "test_user_name",
						},
						ObjectId: "test_attraction_id",
						Score: 5,
					},
					nil,
				).Do(func(ctx context.Context, rating *models.Rating) {
					fmt.Println("ConvertDBOToDTORating called with", rating)
				})
			},
			wantResp: &models.UpdateRatingResp{
				Rating: &models.RatingDTO{
					Id:   "test_rating_id",
					Type: "ATTRACTION",
					User: &models.UserDTO{
						Id: "test_user_id", 
						Name: "test_user_name",
					},
					ObjectId: "test_attraction_id",
					Score: 5,
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
	ctx, mockD, mockDTO, ratingService := ConfigDeleteRating(t)

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
					Id: "test_rating_id",
				},
			},
			before: func(t *testing.T) {
				mockD.EXPECT().DeleteRating(
					ctx,
					"test_rating_id",
				).Return(
					&models.Rating{
						Id:       "test_rating_id",
						Type:     models.TypeAttraction,
						UserId:   "test_user_id",
						ObjectId: "test_attraction_id",
						Score:    5,
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTORating(
					ctx,
					&models.Rating{
						Id:       "test_rating_id",
						Type:     models.TypeAttraction,
						UserId:   "test_user_id",
						ObjectId: "test_attraction_id",
						Score:    5,
					},
				).Return(
					&models.RatingDTO{
						Id:   "test_rating_id",
						Type: "ATTRACTION",
						User: &models.UserDTO{
							Id: "test_user_id", 
							Name: "test_user_name",
						},
						ObjectId: "test_attraction_id",
						Score: 5,
					},
					nil,
				).Do(func(ctx context.Context, rating *models.Rating) {
					fmt.Println("ConvertDBOToDTORating called with", rating)
				})
			},
			wantResp: &models.DeleteRatingResp{
				Rating: &models.RatingDTO{
					Id:   "test_rating_id",
					Type: "ATTRACTION",
					User: &models.UserDTO{
						Id: "test_user_id", 
						Name: "test_user_name",
					},
					ObjectId: "test_attraction_id",
					Score: 5,
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


func TestConvertRatingToDTORating(t *testing.T) {
	ctx, mockF, ratingService := ConfigConvertToDTORating(t)

	type arg struct {
		req *models.Rating
		ctx context.Context
	}

	tests := []struct {
		name     string
		before   func(t *testing.T)
		arg      arg
		wantResp *models.RatingDTO
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				req: &models.Rating{
					Id:       "test_rating_id",
					Type:     models.TypeAttraction,
					UserId:   "test_user_id",
					ObjectId: "test_attraction_id",
					Score:    5,
				},
				ctx: ctx,
			},
			before: func(t *testing.T) {
				mockF.EXPECT().Execute(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					&models.RespFacade{
						GURB: &models.GetUserByIdResp{
							User: &models.UserDTO{Id: "test_user_id", Name: "test_user_name"},
						},
					},
					nil,
				)
			},
			wantResp: &models.RatingDTO{
				Id:   "test_rating_id",
				Type: "ATTRACTION",
				User: &models.UserDTO{
					Id: "test_user_id", 
					Name: "test_user_name",
				},
				ObjectId: "test_attraction_id",
				Score: 5,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := ratingService.ConvertDBOToDTORating(tt.arg.ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
