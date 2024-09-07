package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	// "itineraryplanner/common/custom_errs"
	"itineraryplanner/dal/mock"

	"itineraryplanner/models"
	"itineraryplanner/service"
	ser_mock "itineraryplanner/service/mock"
)

func ConfigCreateAttraction(t *testing.T) (context.Context, *mock.MockCreateAttractionDal, *ser_mock.MockAttractionDTOService, *service.AttractionService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockC := mock.NewMockCreateAttractionDal(ctrl)
	mockDTO := ser_mock.NewMockAttractionDTOService(ctrl)
	attractionService := &service.AttractionService{CDal: mockC}
	return ctx, mockC, mockDTO, attractionService
}
func ConfigGetAttraction(t *testing.T) (context.Context, *mock.MockGetAttractionDal, *ser_mock.MockAttractionDTOService, *service.AttractionService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockG := mock.NewMockGetAttractionDal(ctrl)
	mockDTO := ser_mock.NewMockAttractionDTOService(ctrl)
	attractionService := &service.AttractionService{GDal: mockG}
	return ctx, mockG, mockDTO, attractionService
}
func ConfigGetAttractionById(t *testing.T) (context.Context, *mock.MockGetAttractionByIdDal, *ser_mock.MockAttractionDTOService, *service.AttractionService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockB := mock.NewMockGetAttractionByIdDal(ctrl)
	mockDTO := ser_mock.NewMockAttractionDTOService(ctrl)
	attractionService := &service.AttractionService{BDal: mockB}
	return ctx, mockB, mockDTO, attractionService
}

func ConfigUpdateAttraction(t *testing.T) (context.Context, *mock.MockUpdateAttractionDal, *ser_mock.MockAttractionDTOService, *service.AttractionService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockU := mock.NewMockUpdateAttractionDal(ctrl)
	mockDTO := ser_mock.NewMockAttractionDTOService(ctrl)
	attractionService := &service.AttractionService{UDal: mockU}
	return ctx, mockU, mockDTO, attractionService
}
func ConfigDeleteAttraction(t *testing.T) (context.Context, *mock.MockDeleteAttractionDal, *ser_mock.MockAttractionDTOService, *service.AttractionService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockD := mock.NewMockDeleteAttractionDal(ctrl)
	mockDTO := ser_mock.NewMockAttractionDTOService(ctrl)
	attractionService := &service.AttractionService{DDal: mockD}
	return ctx, mockD, mockDTO, attractionService
}

func ConfigConvertToDTOAttraction(t *testing.T) (context.Context, *ser_mock.MockFacadeDesignPatternService, *service.AttractionService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockC := mock.NewMockCreateAttractionDal(ctrl)
	mockG := mock.NewMockGetAttractionDal(ctrl)
	mockF := ser_mock.NewMockFacadeDesignPatternService(ctrl)
	attractionService := &service.AttractionService{CDal: mockC, GDal: mockG}
	return ctx, mockF, attractionService
}

func TestCreateAttraction(t *testing.T) {
	ctx, mockC, mockDTO, attractionService := ConfigCreateAttraction(t)

	type arg struct {
		req *models.CreateAttractionReq
		ctx context.Context
	}

	tests := []struct {
		name string
		arg
		before   func(t *testing.T)
		wantResp *models.CreateAttractionResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.CreateAttractionReq{
					Name:         "test_name",
					Address:      "test_address",
					CoordinateId: "1",
					TagIDs:       []string{"1", "2"},
				},
			},
			before: func(t *testing.T) {
				mockC.EXPECT().CreateAttraction(ctx, gomock.Any()).Return(
					&models.Attraction{
						Id:           "test_id",
						Name:         "test_name",
						Address:      "test_address",
						CoordinateId: "1",
						TagIDs:       []string{"1", "2"},
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOAttraction(
					ctx,
					&models.Attraction{
						Id:           "test_id",
						Name:         "test_name",
						Address:      "test_address",
						CoordinateId: "1",
						TagIDs:       []string{"1", "2"},
					},
				).Return(
					&models.AttractionDTO{
						Id:      "test_id",
						Name:    "test_name",
						Address: "test_address",
						Coordinate: &models.CoordinateDTO{
							Id: "1",
							X:  1,
							Y:  1,
						},
						Tags: []*models.TagDTO{
							{
								Id:    "1",
								Value: "test 1",
							},
							{
								Id:    "2",
								Value: "test 2",
							},
						},
					},
					nil,
				).Do(func(ctx context.Context, att *models.Attraction) {
					fmt.Println("ConvertDBOToDTOAttractionAttraction called with", att)
				})
			},
			wantResp: &models.CreateAttractionResp{Attraction: &models.AttractionDTO{
				Id:      "test_id",
				Name:    "test_name",
				Address: "test_address",
				Coordinate: &models.CoordinateDTO{
					Id: "1",
					X:  1,
					Y:  1,
				},
				Tags: []*models.TagDTO{
					{
						Id:    "1",
						Value: "test 1",
					},
					{
						Id:    "2",
						Value: "test 2",
					},
				},
			}},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := attractionService.CreateAttraction(ctx, tt.arg.req)
			assert.Equal(t, gotResp, tt.wantResp)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}

func TestGetAttractionById(t *testing.T) {
	ctx, mockB, mockDTO, attractionService := ConfigGetAttractionById(t)

	type arg struct {
		req *models.GetAttractionByIdReq
		ctx context.Context
	}

	tests := []struct {
		name string
		arg
		before   func(t *testing.T)
		wantResp *models.GetAttractionByIdResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.GetAttractionByIdReq{
					Id: "test_id",
				},
			},
			before: func(t *testing.T) {
				mockB.EXPECT().GetAttractionById(
					ctx,
					gomock.Any(),
				).Return(
					&models.Attraction{
						Id:           "test_id",
						Name:         "test_name",
						Address:      "test_address",
						CoordinateId: "1",
						TagIDs:       []string{"1", "2"},
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOAttraction(
					ctx,
					&models.Attraction{
						Id:           "test_id",
						Name:         "test_name",
						Address:      "test_address",
						CoordinateId: "1",
						TagIDs:       []string{"1", "2"},
					},
				).Return(
					&models.AttractionDTO{
						Id:      "test_id",
						Name:    "test_name",
						Address: "test_address",
						Coordinate: &models.CoordinateDTO{
							Id: "1",
							X:  1,
							Y:  1,
						},
						Tags: []*models.TagDTO{
							{
								Id:    "1",
								Value: "test 1",
							},
							{
								Id:    "2",
								Value: "test 2",
							},
						},
					},
					nil,
				).Do(func(ctx context.Context, att *models.Attraction) {
					fmt.Println("ConvertDBOToDTOAttraction called with", att)
				})
			},
			wantResp: &models.GetAttractionByIdResp{Attraction: &models.AttractionDTO{
				Id:      "test_id",
				Name:    "test_name",
				Address: "test_address",
				Coordinate: &models.CoordinateDTO{
					Id: "1",
					X:  1,
					Y:  1,
				},
				Tags: []*models.TagDTO{
					{
						Id:    "1",
						Value: "test 1",
					},
					{
						Id:    "2",
						Value: "test 2",
					},
				},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := attractionService.GetAttractionById(ctx, tt.arg.req)

			assert.Equal(t, gotResp, tt.wantResp)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}

func TestGetAttraction(t *testing.T) {
	ctx, mockG, mockDTO, attractionService := ConfigGetAttraction(t)

	type arg struct {
		req *models.GetAttractionReq
		ctx context.Context
	}

	tests := []struct {
		name string
		arg
		before   func(t *testing.T)
		wantResp *models.GetAttractionResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.GetAttractionReq{},
			},
			before: func(t *testing.T) {
				mockG.EXPECT().GetAttraction(ctx).Return(
					[]*models.Attraction{
						{
							Id:           "test_id",
							Name:         "test_name",
							Address:      "test_address",
							CoordinateId: "1",
							TagIDs:       []string{"1", "2"},
						},
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOAttraction(
					ctx,
					&models.Attraction{
						Id:           "test_id",
						Name:         "test_name",
						Address:      "test_address",
						CoordinateId: "1",
						TagIDs:       []string{"1", "2"},
					},
				).Return(
					&models.AttractionDTO{
						Id:      "test_id",
						Name:    "test_name",
						Address: "test_address",
						Coordinate: &models.CoordinateDTO{
							Id: "1",
							X:  1,
							Y:  1,
						},
						Tags: []*models.TagDTO{
							{
								Id:    "1",
								Value: "test 1",
							},
							{
								Id:    "2",
								Value: "test 2",
							},
						},
					},
					nil,
				).Do(func(ctx context.Context, att *models.Attraction) {
					fmt.Println("ConvertDBOToDTOAttraction called with", att)
				})
			},
			wantResp: &models.GetAttractionResp{Attractions: []*models.AttractionDTO{
				{
					Id:      "test_id",
					Name:    "test_name",
					Address: "test_address",
					Coordinate: &models.CoordinateDTO{
						Id: "1",
						X:  1,
						Y:  1,
					},
					Tags: []*models.TagDTO{
						{
							Id:    "1",
							Value: "test 1",
						},
						{
							Id:    "2",
							Value: "test 2",
						},
					},
				},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := attractionService.GetAttraction(ctx, tt.arg.req)

			assert.Equal(t, gotResp, tt.wantResp)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}

func TestUpdateAttraction(t *testing.T) {
	ctx, mockU, mockDTO, attractionService := ConfigUpdateAttraction(t)

	type arg struct {
		req *models.UpdateAttractionReq
		ctx context.Context
	}

	tests := []struct {
		name string
		arg
		before   func(t *testing.T)
		wantResp *models.UpdateAttractionResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.UpdateAttractionReq{
					Id:           "test_id",
					Name:         "test_name",
					Address:      "test_address",
					CoordinateId: "1",
					TagIDs:       []string{"1", "2"},
				},
			},
			before: func(t *testing.T) {
				mockU.EXPECT().UpdateAttraction(ctx, gomock.Any()).Return(
					&models.Attraction{
						Id:           "test_id",
						Name:         "test_name",
						Address:      "test_address",
						CoordinateId: "1",
						TagIDs:       []string{"1", "2"},
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOAttraction(
					ctx,
					&models.Attraction{
						Id:           "test_id",
						Name:         "test_name",
						Address:      "test_address",
						CoordinateId: "1",
						TagIDs:       []string{"1", "2"},
					},
				).Return(
					&models.AttractionDTO{
						Id:      "test_id",
						Name:    "test_name",
						Address: "test_address",
						Coordinate: &models.CoordinateDTO{
							Id: "1",
							X:  1,
							Y:  1,
						},
						Tags: []*models.TagDTO{
							{
								Id:    "1",
								Value: "test 1",
							},
							{
								Id:    "2",
								Value: "test 2",
							},
						},
					},
					nil,
				).Do(func(ctx context.Context, att *models.Attraction) {
					fmt.Println("ConvertDBOToDTOAttractionAttraction called with", att)
				})
			},
			wantResp: &models.UpdateAttractionResp{Attraction: &models.AttractionDTO{
				Id:      "test_id",
				Name:    "test_name",
				Address: "test_address",
				Coordinate: &models.CoordinateDTO{
					Id: "1",
					X:  1,
					Y:  1,
				},
				Tags: []*models.TagDTO{
					{
						Id:    "1",
						Value: "test 1",
					},
					{
						Id:    "2",
						Value: "test 2",
					},
				},
			}},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := attractionService.UpdateAttraction(ctx, tt.arg.req)
			assert.Equal(t, gotResp, tt.wantResp)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}

func TestDeleteAttraction(t *testing.T) {
	ctx, mockD, mockDTO, attractionService := ConfigDeleteAttraction(t)

	type arg struct {
		req *models.DeleteAttractionReq
		ctx context.Context
	}

	tests := []struct {
		name string
		arg
		before   func(t *testing.T)
		wantResp *models.DeleteAttractionResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.DeleteAttractionReq{
					Id: "test_id",
				},
			},
			before: func(t *testing.T) {
				mockD.EXPECT().DeleteAttraction(
					ctx,
					gomock.Any(),
				).Return(
					&models.Attraction{
						Id:           "test_id",
						Name:         "test_name",
						Address:      "test_address",
						CoordinateId: "1",
						TagIDs:       []string{"1", "2"},
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOAttraction(
					ctx,
					&models.Attraction{
						Id:           "test_id",
						Name:         "test_name",
						Address:      "test_address",
						CoordinateId: "1",
						TagIDs:       []string{"1", "2"},
					},
				).Return(
					&models.AttractionDTO{
						Id:      "test_id",
						Name:    "test_name",
						Address: "test_address",
						Coordinate: &models.CoordinateDTO{
							Id: "1",
							X:  1,
							Y:  1,
						},
						Tags: []*models.TagDTO{
							{
								Id:    "1",
								Value: "test 1",
							},
							{
								Id:    "2",
								Value: "test 2",
							},
						},
					},
					nil,
				).Do(func(ctx context.Context, att *models.Attraction) {
					fmt.Println("ConvertDBOToDTOAttraction called with", att)
				})
			},
			wantResp: &models.DeleteAttractionResp{
				Attraction: &models.AttractionDTO{
					Id:      "test_id",
					Name:    "test_name",
					Address: "test_address",
					Coordinate: &models.CoordinateDTO{
						Id: "1",
						X:  1,
						Y:  1,
					},
					Tags: []*models.TagDTO{
						{
							Id:    "1",
							Value: "test 1",
						},
						{
							Id:    "2",
							Value: "test 2",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := attractionService.DeleteAttraction(ctx, tt.arg.req)

			assert.Equal(t, gotResp, tt.wantResp)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}

func TestConvertAttToDTOAttraction(t *testing.T) {
	ctx, mockF, attractionService := ConfigConvertToDTOAttraction(t)

	type arg struct {
		req *models.Attraction
		ctx context.Context
	}

	tests := []struct {
		name     string
		before   func(t *testing.T)
		arg      arg
		wantResp *models.AttractionDTO
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				req: &models.Attraction{
					Id:           "test_id",
					Name:         "test_name",
					Address:      "test_address",
					CoordinateId: "1",
					TagIDs:       []string{"1"},
				},
				ctx: ctx,
			},
			before: func(t *testing.T) {
				mockF.EXPECT().Execute(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					&models.RespFacade{
						GTRB: &models.GetTagByIdResp{
							Tag: &models.TagDTO{
								Id:    "1",
								Value: "test 1",
							},
						},
					},
					nil,
				)
				mockF.EXPECT().Execute(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					&models.RespFacade{
						GCRB: &models.GetCoordinateByIdResp{
							Coordinate: &models.CoordinateDTO{
								Id: "1",
								X:  1,
								Y:  1,
							},
						},
					},
					nil,
				)
			},
			wantResp: &models.AttractionDTO{
				Id:      "test_id",
				Name:    "test_name",
				Address: "test_address",
				Coordinate: &models.CoordinateDTO{
					X: 1,
					Y: 2,
				},
				Tags: []*models.TagDTO{
					{
						Id:    "1",
						Value: "test 1",
					},
					{
						Id:    "2",
						Value: "test 2",
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := attractionService.ConvertDBOToDTOAttraction(tt.arg.ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
