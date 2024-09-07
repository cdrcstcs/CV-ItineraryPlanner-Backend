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
	ser_mock "itineraryplanner/service/mock"
)

func ConfigCreateItinerary(t *testing.T) (context.Context, *mock.MockCreateItineraryDal, *ser_mock.MockItineraryDTOService, *service.ItineraryService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockC := mock.NewMockCreateItineraryDal(ctrl)
	mockDTO := ser_mock.NewMockItineraryDTOService(ctrl)
	itineraryService := &service.ItineraryService{CDal: mockC}
	return ctx, mockC, mockDTO, itineraryService
}

func ConfigGetItinerary(t *testing.T) (context.Context, *mock.MockGetItineraryDal, *ser_mock.MockItineraryDTOService, *service.ItineraryService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockG := mock.NewMockGetItineraryDal(ctrl)
	mockDTO := ser_mock.NewMockItineraryDTOService(ctrl)
	itineraryService := &service.ItineraryService{GDal: mockG}
	return ctx, mockG, mockDTO, itineraryService
}

func ConfigGetItineraryById(t *testing.T) (context.Context, *mock.MockGetItineraryByIdDal, *ser_mock.MockItineraryDTOService, *service.ItineraryService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockB := mock.NewMockGetItineraryByIdDal(ctrl)
	mockDTO := ser_mock.NewMockItineraryDTOService(ctrl)
	itineraryService := &service.ItineraryService{BDal: mockB}
	return ctx, mockB, mockDTO, itineraryService
}

func ConfigUpdateItinerary(t *testing.T) (context.Context, *mock.MockUpdateItineraryDal, *ser_mock.MockItineraryDTOService, *service.ItineraryService){
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockU := mock.NewMockUpdateItineraryDal(ctrl)
	mockDTO := ser_mock.NewMockItineraryDTOService(ctrl)
	itineraryService := &service.ItineraryService{UDal: mockU}
	return ctx, mockU, mockDTO, itineraryService
}
func ConfigDeleteItinerary(t *testing.T) (context.Context, *mock.MockDeleteItineraryDal, *ser_mock.MockItineraryDTOService, *service.ItineraryService){
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockD := mock.NewMockDeleteItineraryDal(ctrl)
	mockDTO := ser_mock.NewMockItineraryDTOService(ctrl)
	itineraryService := &service.ItineraryService{DDal: mockD}
	return ctx, mockD, mockDTO, itineraryService
}

func ConfigConvertToDTOItinerary(t *testing.T) (context.Context, *ser_mock.MockFacadeDesignPatternService, *service.ItineraryService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockC := mock.NewMockCreateItineraryDal(ctrl)
	mockG := mock.NewMockGetItineraryDal(ctrl)
	mockF := ser_mock.NewMockFacadeDesignPatternService(ctrl)
	itineraryService := &service.ItineraryService{CDal: mockC, GDal: mockG}
	return ctx, mockF, itineraryService
}

func TestCreateItinerary(t *testing.T) {
	ctx, mockC, mockDTO, itineraryService := ConfigCreateItinerary(t)

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
					CopiedId:  "test_copied_id",
					UserId:    "test_user_id",
					StartTime: time.Now(),
					EndTime:   time.Now().Add(time.Hour),
					EventIds:  []string{
						"event-id-1", 
						"event-id-2",
					},
					RatingId:  "rating-id-1",
				},
			},
			before: func(t *testing.T) {
				mockC.EXPECT().CreateItinerary(
					ctx,
					gomock.Any(),
				).Return(
					&models.Itinerary{
						Id:        "test_itinerary_id",
						CopiedId:  "test_copied_id",
						UserId:    "test_user_id",
						StartTime: time.Now(),
						EndTime:   time.Now().Add(time.Hour),
						EventIds:  []string{
							"event-id-1",
							"event-id-2",
						},
						RatingId:  "rating-id-1",
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOItinerary(
					ctx,
					&models.Itinerary{
						Id:        "test_itinerary_id",
						CopiedId:  "test_copied_id",
						UserId:    "test_user_id",
						StartTime: time.Now(),
						EndTime:   time.Now().Add(time.Hour),
						EventIds:  []string{
							"event-id-1",
							"event-id-2",
						},
						RatingId:  "rating-id-1",
					},
				).Return(
					&models.ItineraryDTO{
						Id:        "test_itinerary_id",
						CopiedId:  "test_copied_id",
						User:    &models.UserDTO{
							Id:       "test_user_id",
							Name:     "test_name",
							Email:    "test_email@example.com",
							Password: "test_password",
						},
						StartTime: time.Now(),
						EndTime:   time.Now().Add(time.Hour),
						Events:    []*models.EventDTO{
							{
								Id:            "event-id-1",
								ItineraryId: "test_itinerary_id",
								StartTime:     time.Now(),
								EndTime:       time.Now().Add(time.Hour),
								Attraction:    &models.AttractionDTO{},
								Description:   "test_description",
							},
							{
								Id:            "event-id-2",
								ItineraryId: "test_itinerary_id",
								StartTime:     time.Now(),
								EndTime:       time.Now().Add(time.Hour),
								Attraction:    &models.AttractionDTO{},
								Description:   "test_description",
							},
						},
						Rating:    &models.RatingDTO{
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
					nil,
				)
			},
			wantResp: &models.CreateItineraryResp{
				Itinerary: &models.ItineraryDTO{
					Id:        "test_itinerary_id",
					CopiedId:  "test_copied_id",
					User:    &models.UserDTO{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
					StartTime: time.Now(),
					EndTime:   time.Now().Add(time.Hour),
					Events:    []*models.EventDTO{
						{
							Id:            "event-id-1",
							ItineraryId: "test_itinerary_id",
							StartTime:     time.Now(),
							EndTime:       time.Now().Add(time.Hour),
							Attraction:    &models.AttractionDTO{},
							Description:   "test_description",
						},
						{
							Id:            "event-id-2",
							ItineraryId: "test_itinerary_id",
							StartTime:     time.Now(),
							EndTime:       time.Now().Add(time.Hour),
							Attraction:    &models.AttractionDTO{},
							Description:   "test_description",
						},
					},
					Rating:    &models.RatingDTO{
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
			gotResp, err := itineraryService.CreateItinerary(ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetItineraryById(t *testing.T) {
	ctx, mockB, mockDTO, itineraryService := ConfigGetItineraryById(t)

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
				mockB.EXPECT().GetItineraryById(
					ctx,
					gomock.Any(),
				).Return(
					&models.Itinerary{
						Id:        "test_itinerary_id",
						CopiedId:  "test_copied_id",
						UserId:    "test_user_id",
						StartTime: time.Now(),
						EndTime:   time.Now().Add(time.Hour),
						EventIds:  []string{
							"event-id-1",
							"event-id-2",
						},
						RatingId:  "rating-id-1",
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOItinerary(
					ctx,
					gomock.Any()).Return(
						&models.ItineraryDTO{
							Id:        "test_itinerary_id",
							CopiedId:  "test_copied_id",
							User:    &models.UserDTO{
								Id:       "test_user_id",
								Name:     "test_name",
								Email:    "test_email@example.com",
								Password: "test_password",
							},
							StartTime: time.Now(),
							EndTime:   time.Now().Add(time.Hour),
							Events:    []*models.EventDTO{
								{
									Id:            "event-id-1",
									ItineraryId: "test_itinerary_id",
									StartTime:     time.Now(),
									EndTime:       time.Now().Add(time.Hour),
									Attraction:    &models.AttractionDTO{},
									Description:   "test_description",
								},
								{
									Id:            "event-id-2",
									ItineraryId: "test_itinerary_id",
									StartTime:     time.Now(),
									EndTime:       time.Now().Add(time.Hour),
									Attraction:    &models.AttractionDTO{},
									Description:   "test_description",
								},
							},
							Rating:    &models.RatingDTO{
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
					nil)
			},
			wantResp: &models.GetItineraryByIdResp{
				Itinerary: &models.ItineraryDTO{
					Id:        "test_itinerary_id",
					CopiedId:  "test_copied_id",
					User:    &models.UserDTO{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
					StartTime: time.Now(),
					EndTime:   time.Now().Add(time.Hour),
					Events:    []*models.EventDTO{
						{
							Id:            "event-id-1",
							ItineraryId: "test_itinerary_id",
							StartTime:     time.Now(),
							EndTime:       time.Now().Add(time.Hour),
							Attraction:    &models.AttractionDTO{},
							Description:   "test_description",
						},
						{
							Id:            "event-id-2",
							ItineraryId: "test_itinerary_id",
							StartTime:     time.Now(),
							EndTime:       time.Now().Add(time.Hour),
							Attraction:    &models.AttractionDTO{},
							Description:   "test_description",
						},
					},
					Rating:    &models.RatingDTO{
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
			gotResp, err := itineraryService.GetItineraryById(ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetItinerary(t *testing.T) {
	ctx, mockG, mockDTO, itineraryService := ConfigGetItinerary(t)

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
				mockG.EXPECT().GetItinerary(ctx).Return(
					[]*models.Itinerary{
						{
							Id:        "test_itinerary_id",
							CopiedId:  "test_copied_id",
							UserId:    "test_user_id",
							StartTime: time.Now(),
							EndTime:   time.Now().Add(time.Hour),
							EventIds:  []string{
								"event-id-1",
								"event-id-2",
							},
							RatingId:  "rating-id-1",
						},
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOItinerary(
					ctx,
					gomock.Any()).Return(
						&models.ItineraryDTO{
							Id:        "test_itinerary_id",
							CopiedId:  "test_copied_id",
							User:    &models.UserDTO{
								Id:       "test_user_id",
								Name:     "test_name",
								Email:    "test_email@example.com",
								Password: "test_password",
							},
							StartTime: time.Now(),
							EndTime:   time.Now().Add(time.Hour),
							Events:    []*models.EventDTO{
								{
									Id:            "event-id-1",
									ItineraryId: "test_itinerary_id",
									StartTime:     time.Now(),
									EndTime:       time.Now().Add(time.Hour),
									Attraction:    &models.AttractionDTO{},
									Description:   "test_description",
								},
								{
									Id:            "event-id-2",
									ItineraryId: "test_itinerary_id",
									StartTime:     time.Now(),
									EndTime:       time.Now().Add(time.Hour),
									Attraction:    &models.AttractionDTO{},
									Description:   "test_description",
								},
							},
							Rating:    &models.RatingDTO{
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
					nil)
			},
			wantResp: &models.GetItineraryResp{
				Itineraries: []*models.ItineraryDTO{
					{
						Id:        "test_itinerary_id",
						CopiedId:  "test_copied_id",
						User:    &models.UserDTO{
							Id:       "test_user_id",
							Name:     "test_name",
							Email:    "test_email@example.com",
							Password: "test_password",
						},
						StartTime: time.Now(),
						EndTime:   time.Now().Add(time.Hour),
						Events:    []*models.EventDTO{
							{
								Id:            "event-id-1",
								ItineraryId: "test_itinerary_id",
								StartTime:     time.Now(),
								EndTime:       time.Now().Add(time.Hour),
								Attraction:    &models.AttractionDTO{},
								Description:   "test_description",
							},
							{
								Id:            "event-id-2",
								ItineraryId: "test_itinerary_id",
								StartTime:     time.Now(),
								EndTime:       time.Now().Add(time.Hour),
								Attraction:    &models.AttractionDTO{},
								Description:   "test_description",
							},
						},
						Rating:    &models.RatingDTO{
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
	ctx, mockU, mockDTO, itineraryService := ConfigUpdateItinerary(t)

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
					CopiedId:  "test_copied_id",
					UserId:    "test_user_id",
					StartTime: time.Now(),
					EndTime:   time.Now().Add(time.Hour),
					EventIds:  []string{
						"event-id-1", 
						"event-id-2",
					},
					RatingId:  "rating-id-1",
				},
			},
			before: func(t *testing.T) {
				mockU.EXPECT().UpdateItinerary(
					ctx,
					gomock.Any(),
				).Return(
					&models.Itinerary{
						Id:        "test_itinerary_id",
						CopiedId:  "test_copied_id",
						UserId:    "test_user_id",
						StartTime: time.Now(),
						EndTime:   time.Now().Add(time.Hour),
						EventIds:  []string{
							"event-id-1",
							"event-id-2",
						},
						RatingId:  "rating-id-1",
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOItinerary(
					ctx,
					&models.Itinerary{
						Id:        "test_itinerary_id",
						CopiedId:  "test_copied_id",
						UserId:    "test_user_id",
						StartTime: time.Now(),
						EndTime:   time.Now().Add(time.Hour),
						EventIds:  []string{
							"event-id-1",
							"event-id-2",
						},
						RatingId:  "rating-id-1",
					},
				).Return(
					&models.ItineraryDTO{
						Id:        "test_itinerary_id",
						CopiedId:  "test_copied_id",
						User:    &models.UserDTO{
							Id:       "test_user_id",
							Name:     "test_name",
							Email:    "test_email@example.com",
							Password: "test_password",
						},
						StartTime: time.Now(),
						EndTime:   time.Now().Add(time.Hour),
						Events:    []*models.EventDTO{
							{
								Id:            "event-id-1",
								ItineraryId: "test_itinerary_id",
								StartTime:     time.Now(),
								EndTime:       time.Now().Add(time.Hour),
								Attraction:    &models.AttractionDTO{},
								Description:   "test_description",
							},
							{
								Id:            "event-id-2",
								ItineraryId: "test_itinerary_id",
								StartTime:     time.Now(),
								EndTime:       time.Now().Add(time.Hour),
								Attraction:    &models.AttractionDTO{},
								Description:   "test_description",
							},
						},
						Rating:    &models.RatingDTO{
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
					nil,
				)
			},
			wantResp: &models.UpdateItineraryResp{
				Itinerary: &models.ItineraryDTO{
					Id:        "test_itinerary_id",
					CopiedId:  "test_copied_id",
					User:    &models.UserDTO{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
					StartTime: time.Now(),
					EndTime:   time.Now().Add(time.Hour),
					Events:    []*models.EventDTO{
						{
							Id:            "event-id-1",
							ItineraryId: "test_itinerary_id",
							StartTime:     time.Now(),
							EndTime:       time.Now().Add(time.Hour),
							Attraction:    &models.AttractionDTO{},
							Description:   "test_description",
						},
						{
							Id:            "event-id-2",
							ItineraryId: "test_itinerary_id",
							StartTime:     time.Now(),
							EndTime:       time.Now().Add(time.Hour),
							Attraction:    &models.AttractionDTO{},
							Description:   "test_description",
						},
					},
					Rating:    &models.RatingDTO{
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
			gotResp, err := itineraryService.UpdateItinerary(ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestDeleteItinerary(t *testing.T) {
	ctx, mockD, mockDTO, itineraryService := ConfigDeleteItinerary(t)

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
				mockD.EXPECT().DeleteItinerary(
					ctx,
					gomock.Any(),
				).Return(
					&models.Itinerary{
						Id:        "test_itinerary_id",
						CopiedId:  "test_copied_id",
						UserId:    "test_user_id",
						StartTime: time.Now(),
						EndTime:   time.Now().Add(time.Hour),
						EventIds:  []string{
							"event-id-1",
							"event-id-2",
						},
						RatingId:  "rating-id-1",
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOItinerary(
					ctx,
					&models.Itinerary{
						Id:        "test_itinerary_id",
						CopiedId:  "test_copied_id",
						UserId:    "test_user_id",
						StartTime: time.Now(),
						EndTime:   time.Now().Add(time.Hour),
						EventIds:  []string{
							"event-id-1",
							"event-id-2",
						},
						RatingId:  "rating-id-1",
					}).Return(
						&models.ItineraryDTO{
							Id:        "test_itinerary_id",
							CopiedId:  "test_copied_id",
							User:    &models.UserDTO{
								Id:       "test_user_id",
								Name:     "test_name",
								Email:    "test_email@example.com",
								Password: "test_password",
							},
							StartTime: time.Now(),
							EndTime:   time.Now().Add(time.Hour),
							Events:    []*models.EventDTO{
								{
									Id:            "event-id-1",
									ItineraryId: "test_itinerary_id",
									StartTime:     time.Now(),
									EndTime:       time.Now().Add(time.Hour),
									Attraction:    &models.AttractionDTO{},
									Description:   "test_description",
								},
								{
									Id:            "event-id-2",
									ItineraryId: "test_itinerary_id",
									StartTime:     time.Now(),
									EndTime:       time.Now().Add(time.Hour),
									Attraction:    &models.AttractionDTO{},
									Description:   "test_description",
								},
							},
							Rating:    &models.RatingDTO{
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
					nil)
			},
			wantResp: &models.DeleteItineraryResp{
				Itinerary: &models.ItineraryDTO{
					Id:        "test_itinerary_id",
					CopiedId:  "test_copied_id",
					User:    &models.UserDTO{
						Id:       "test_user_id",
						Name:     "test_name",
						Email:    "test_email@example.com",
						Password: "test_password",
					},
					StartTime: time.Now(),
					EndTime:   time.Now().Add(time.Hour),
					Events:    []*models.EventDTO{
						{
							Id:            "event-id-1",
							ItineraryId: "test_itinerary_id",
							StartTime:     time.Now(),
							EndTime:       time.Now().Add(time.Hour),
							Attraction:    &models.AttractionDTO{},
							Description:   "test_description",
						},
						{
							Id:            "event-id-2",
							ItineraryId: "test_itinerary_id",
							StartTime:     time.Now(),
							EndTime:       time.Now().Add(time.Hour),
							Attraction:    &models.AttractionDTO{},
							Description:   "test_description",
						},
					},
					Rating:    &models.RatingDTO{
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
			gotResp, err := itineraryService.DeleteItinerary(ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
func TestConvertDBOToDTOItinerary(t *testing.T) {
	ctx, mockF, itineraryService := ConfigConvertToDTOItinerary(t)

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
					Id:        "test_itinerary_id",
					CopiedId:  "test_copied_id",
					UserId:    "test_user_id",
					StartTime: time.Now(),
					EndTime:   time.Now().Add(time.Hour),
					EventIds:  []string{
						"event-id-1",
					},
					RatingId:  "rating-id-1",
				},
				ctx: ctx,
			},
			before: func(t *testing.T) {
				mockF.EXPECT().Execute(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					&models.RespFacade{
						GURB: &models.GetUserByIdResp{
							User: &models.UserDTO{
								Id:       "test_user_id",
								Name:     "test_name",
								Email:    "test_email@example.com",
								Password: "test_password",
							},
						},
						GERB: &models.GetEventByIdResp{
							Event: &models.EventDTO{
								Id:            "event-id-1",
								ItineraryId: "test_itinerary_id",
								StartTime:     time.Now(),
								EndTime:       time.Now().Add(time.Hour),
								Attraction:    &models.AttractionDTO{},
								Description:   "test_description",
							},
						},
						GRRB: &models.GetRatingByIdResp{
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
					},
					nil,
				)
			},
			wantResp: &models.ItineraryDTO{
				Id:        "test_itinerary_id",
				CopiedId:  "test_copied_id",
				User:    &models.UserDTO{
					Id:       "test_user_id",
					Name:     "test_name",
					Email:    "test_email@example.com",
					Password: "test_password",
				},
				StartTime: time.Now(),
				EndTime:   time.Now().Add(time.Hour),
				Events:    []*models.EventDTO{
					{
						Id:            "event-id-1",
						ItineraryId: "test_itinerary_id",
						StartTime:     time.Now(),
						EndTime:       time.Now().Add(time.Hour),
						Attraction:    &models.AttractionDTO{},
						Description:   "test_description",
					},
				},
				Rating:    &models.RatingDTO{
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
			gotResp, err := itineraryService.ConvertDBOToDTOItinerary(tt.arg.ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
	