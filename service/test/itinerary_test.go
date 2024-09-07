// itinerary_test.go

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

func ConfigItinerary(t *testing.T) (context.Context, *mock.MockItineraryDal, *service.ItineraryService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mock := mock.NewMockItineraryDal(ctrl)
	itineraryService := &service.ItineraryService{Dal: mock}
	return ctx, mock, itineraryService
}
func TestCreateItinerary(t *testing.T) {
	ctx, mock, itineraryService := ConfigItinerary(t)

	type arg struct {
		req *models.CreateItineraryReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.CreateItineraryResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.CreateItineraryReq{
					CopiedId:  "6641b48d61634ce13f96c3f7",
					UserId:    "6641b48d61634ce13f96c3f7",
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					EventIds:  []string{
						"66420cf5a9df66a0a0737cb5",
					},
					EventCount: 1,
					RatingId:  "6642102153edf9e8fb3dfcda",
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().CreateItinerary(
					ctx,
					gomock.Any(),
				).Return(
					&models.Itinerary{
						CopiedId:  "6641b48d61634ce13f96c3f7",
						CopiedName: "test",
						Name: "test",
						UserId:    "6641b48d61634ce13f96c3f7",
						StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
						EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
						EventIds:  []string{
							"66420cf5a9df66a0a0737cb5",
						},
						EventCount: 1,
						RatingId:  "6642102153edf9e8fb3dfcda",
					},
					nil,
				)
			},
			wantResp: &models.CreateItineraryResp{
				Itinerary: &models.ItineraryDTO{
					CopiedId:  "6641b48d61634ce13f96c3f7",
					CopiedName: "test",
					Name: "test",
					UserId:    "6641b48d61634ce13f96c3f7",
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					EventIds:  []string{
						"66420cf5a9df66a0a0737cb5",
					},
					EventCount:1,
					Rating:  &models.RatingDTO{
						Id:"6642102153edf9e8fb3dfcda",
						Score:5,
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := itineraryService.CreateItinerary(ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetItineraryById(t *testing.T) {
	ctx, mock, itineraryService := ConfigItinerary(t)

	type arg struct {
		req *models.GetItineraryByIdReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.GetItineraryByIdResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.GetItineraryByIdReq{
					Id: "test_itinerary_id",
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().GetItineraryById(
					ctx,
					gomock.Any(),
				).Return(
					&models.Itinerary{
						CopiedId:  "6641b48d61634ce13f96c3f7",
						CopiedName: "test",
						Name: "test",
						UserId:    "6641b48d61634ce13f96c3f7",
						StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
						EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
						EventIds:  []string{
							"66420cf5a9df66a0a0737cb5",
						},
						EventCount: 1,
						RatingId:  "6642102153edf9e8fb3dfcda",
					},
					nil,
				)
			},
			wantResp: &models.GetItineraryByIdResp{
				Itinerary: &models.ItineraryDTO{
					Id:        "test_itinerary_id",
					CopiedId:  "6641b48d61634ce13f96c3f7",
					CopiedName: "test",
					Name: "test",
					UserId:    "6641b48d61634ce13f96c3f7",
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					EventIds:     []string{"66420cf5a9df66a0a0737cb5"},
					EventCount:   1,
					Rating:       &models.RatingDTO{Id: "6642102153edf9e8fb3dfcda", Score: 5},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := itineraryService.GetItineraryById(ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetItinerary(t *testing.T) {
	ctx, mock, itineraryService := ConfigItinerary(t)

	type arg struct {
		req *models.GetItineraryReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.GetItineraryResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.GetItineraryReq{
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().GetItinerary(ctx).Return(
					[]*models.Itinerary{
						{
							Id:        "test_itinerary_id",
							CopiedId:  "6641b48d61634ce13f96c3f7",
							CopiedName: "test",
							Name: "test",
							UserId:    "6641b48d61634ce13f96c3f7",
							StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
							EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
							EventIds:  []string{
								"66420cf5a9df66a0a0737cb5",
							},
							EventCount: 1,
							RatingId:  "6642102153edf9e8fb3dfcda",
						},
					},
					nil,
				)
			},	
			wantResp: &models.GetItineraryResp{
				Itineraries: []*models.ItineraryDTO{
					{
						Id:        "test_itinerary_id",
						CopiedId:  "6641b48d61634ce13f96c3f7",
						CopiedName: "test",
						Name: "test",
						UserId:    "6641b48d61634ce13f96c3f7",
						StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
						EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
						EventIds:     []string{"66420cf5a9df66a0a0737cb5"},
						EventCount:   1,
						Rating:       &models.RatingDTO{Id: "6642102153edf9e8fb3dfcda", Score: 5},
					},
				},
			},
			wantErr: nil,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := itineraryService.GetItinerary(ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestUpdateItinerary(t *testing.T) {
	ctx, mockU, itineraryService := ConfigItinerary(t)

	type arg struct {
		req *models.UpdateItineraryReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.UpdateItineraryResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.UpdateItineraryReq{
					Id:        "test_itinerary_id",
					CopiedId:  "6641b48d61634ce13f96c3f7",
					CopiedName: "test1",
					Name: "test1",
					UserId:    "6641b48d61634ce13f96c3f7",
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					EventIds:  []string{
						"66420cf5a9df66a0a0737cb5",
					},
					EventCount: 1,
					RatingId:  "6642102153edf9e8fb3dfcda",
				},
			},
			before: func(t *testing.T) {
				mockU.EXPECT().UpdateItinerary(
					ctx,
					gomock.Any(),
				).Return(
					&models.Itinerary{
						Id:        "test_itinerary_id",
						CopiedId:  "6641b48d61634ce13f96c3f7",
						CopiedName: "test1",
						Name: "test1",
						UserId:    "6641b48d61634ce13f96c3f7",
						StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
						EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
						EventIds:  []string{
							"66420cf5a9df66a0a0737cb5",
						},
						EventCount: 1,
						RatingId:  "6642102153edf9e8fb3dfcda",
					},
					nil,
				)
			},
			wantResp: &models.UpdateItineraryResp{
				Itinerary: &models.ItineraryDTO{
					Id:        "test_itinerary_id",
					CopiedId:  "6641b48d61634ce13f96c3f7",
					CopiedName: "test",
					Name: "test",
					UserId:    "6641b48d61634ce13f96c3f7",
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					EventIds:     []string{"66420cf5a9df66a0a0737cb5"},
					EventCount:   1,
					Rating:       &models.RatingDTO{Id: "6642102153edf9e8fb3dfcda", Score: 5},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := itineraryService.UpdateItinerary(ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestDeleteItinerary(t *testing.T) {
	ctx, mock, itineraryService := ConfigItinerary(t)

	type arg struct {
		req *models.DeleteItineraryReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.DeleteItineraryResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.DeleteItineraryReq{
					Id: "test_itinerary_id",
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().DeleteItinerary(
					ctx,
					gomock.Any(),
				).Return(
					&models.Itinerary{
						Id: "test_itinerary_id",
						CopiedId:  "6641b48d61634ce13f96c3f7",
						CopiedName: "test1",
						Name: "test1",
						UserId:    "6641b48d61634ce13f96c3f7",
						StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
						EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
						EventIds:  []string{
							"66420cf5a9df66a0a0737cb5",
						},
						EventCount: 1,
						RatingId:  "6642102153edf9e8fb3dfcda",
					},
					nil,
				)
			},
			wantResp: &models.DeleteItineraryResp{
				Itinerary: &models.ItineraryDTO{
					Id:        "test_itinerary_id",
					CopiedId:  "6641b48d61634ce13f96c3f7",
					CopiedName: "test",
					Name: "test",
					UserId:    "6641b48d61634ce13f96c3f7",
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					EventIds:     []string{"66420cf5a9df66a0a0737cb5"},
					EventCount:   1,
					Rating:       &models.RatingDTO{Id: "6642102153edf9e8fb3dfcda", Score: 5},
				},
			},
			wantErr: nil,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := itineraryService.DeleteItinerary(ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
func TestConvertDBOToDTOItinerary(t *testing.T) {
	ctx:= context.Background()
	itineraryService:= &service.ItineraryService{}
	type arg struct {
		req *models.Itinerary
		ctx context.Context
	}

	tests := []struct {
		name     string
		before   func(t *testing.T)
		arg      arg
		wantResp *models.ItineraryDTO
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				req: &models.Itinerary{
					Id: "test_itinerary_id",
					CopiedId:  "6641b48d61634ce13f96c3f7",
					CopiedName: "test1",
					Name: "test1",
					UserId:    "6641b48d61634ce13f96c3f7",
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					EventIds:  []string{
						"66420cf5a9df66a0a0737cb5",
					},
					EventCount: 1,
					RatingId:  "6642102153edf9e8fb3dfcda",
				},
				ctx: ctx,
			},
			wantResp: &models.ItineraryDTO{
				Id:        "test_itinerary_id",
				CopiedId:  "6641b48d61634ce13f96c3f7",
				CopiedName: "test",
				Name: "test",
				UserId:    "6641b48d61634ce13f96c3f7",
				StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
				EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
				EventIds:     []string{"66420cf5a9df66a0a0737cb5"},
				EventCount:   1,
				Rating:       &models.RatingDTO{Id: "6642102153edf9e8fb3dfcda", Score: 5},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := itineraryService.ConvertDBOToDTOItinerary(tt.arg.ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
	