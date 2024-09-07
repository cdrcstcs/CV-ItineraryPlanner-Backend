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

func ConfigCreateCoordinate(t *testing.T) (context.Context, *mock.MockCreateCoordinateDal, *ser_mock.MockCoordinateDTOService, *service.CoordinateService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockC := mock.NewMockCreateCoordinateDal(ctrl)
	mockDTO := ser_mock.NewMockCoordinateDTOService(ctrl)
	coordinateService := &service.CoordinateService{CDal: mockC}
	return ctx, mockC, mockDTO, coordinateService
}


func ConfigGetCoordinate(t *testing.T) (context.Context, *mock.MockGetCoordinateDal, *ser_mock.MockCoordinateDTOService, *service.CoordinateService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockG := mock.NewMockGetCoordinateDal(ctrl)
	mockDTO := ser_mock.NewMockCoordinateDTOService(ctrl)
	coordinateService := &service.CoordinateService{GDal: mockG}
	return ctx, mockG, mockDTO, coordinateService
}

func ConfigGetCoordinateById(t *testing.T) (context.Context, *mock.MockGetCoordinateByIdDal, *ser_mock.MockCoordinateDTOService, *service.CoordinateService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockB := mock.NewMockGetCoordinateByIdDal(ctrl)
	mockDTO := ser_mock.NewMockCoordinateDTOService(ctrl)
	coordinateService := &service.CoordinateService{BDal: mockB}
	return ctx, mockB, mockDTO, coordinateService
}

func ConfigUpdateCoordinate(t *testing.T) (context.Context, *mock.MockUpdateCoordinateDal, *ser_mock.MockCoordinateDTOService, *service.CoordinateService){
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockU := mock.NewMockUpdateCoordinateDal(ctrl)
	mockDTO := ser_mock.NewMockCoordinateDTOService(ctrl)
	coordinateService := &service.CoordinateService{UDal: mockU}
	return ctx, mockU, mockDTO, coordinateService
}
func ConfigDeleteCoordinate(t *testing.T) (context.Context, *mock.MockDeleteCoordinateDal, *ser_mock.MockCoordinateDTOService, *service.CoordinateService){
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockD := mock.NewMockDeleteCoordinateDal(ctrl)
	mockDTO := ser_mock.NewMockCoordinateDTOService(ctrl)
	coordinateService := &service.CoordinateService{DDal: mockD}
	return ctx, mockD, mockDTO, coordinateService
}

func ConfigConvertToDTOCoordinate(t *testing.T) (context.Context, *ser_mock.MockFacadeDesignPatternService, *service.CoordinateService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockC := mock.NewMockCreateCoordinateDal(ctrl)
	mockG := mock.NewMockGetCoordinateDal(ctrl)
	mockF := ser_mock.NewMockFacadeDesignPatternService(ctrl)
	coordinateService := &service.CoordinateService{CDal: mockC, GDal: mockG}
	return ctx, mockF, coordinateService
}

func TestCreateCoordinate(t *testing.T) {
	ctx, mockC, mockDTO, coordinateService := ConfigCreateCoordinate(t)

	type arg struct {
		req *models.CreateCoordinateReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg
		before   func(t *testing.T)
		wantResp *models.CreateCoordinateResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.CreateCoordinateReq{
					X: 1,
					Y: 2,
				},
			},
			before: func(t *testing.T) {
				mockC.EXPECT().CreateCoordinate(ctx, gomock.Any()).Return(
					&models.Coordinate{
						Id: "test_coordinate_id",
						X:  1,
						Y:  2,
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOCoordinate(
					ctx,
					&models.Coordinate{
						Id: "test_coordinate_id",
						X:  1,
						Y:  2,
					},
				).Return(
					&models.CoordinateDTO{
						Id: "test_coordinate_id",
						X:  1,
						Y:  2,
					},
					nil,
				).Do(func(ctx context.Context, coordinate *models.Coordinate) {
					fmt.Println("ConvertDBOToDTOCoordinate called with", coordinate)
				})
			},
			wantResp: &models.CreateCoordinateResp{
				Coordinate: &models.CoordinateDTO{
					Id: "test_coordinate_id",
					X:  1,
					Y:  2,
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := coordinateService.CreateCoordinate(tt.arg.ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetCoordinateById(t *testing.T) {
	ctx, mockB, mockDTO, coordinateService := ConfigGetCoordinateById(t)

	type arg struct {
		req *models.GetCoordinateByIdReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg
		before   func(t *testing.T)
		wantResp *models.GetCoordinateByIdResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.GetCoordinateByIdReq{
					Id: "test_coordinate_id",
				},
			},
			before: func(t *testing.T) {
				mockB.EXPECT().GetCoordinateById(
					ctx,
					gomock.Any(),
				).Return(
					&models.Coordinate{
						Id: "test_coordinate_id",
						X:  1,
						Y:  2,
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOCoordinate(
					ctx,
					&models.Coordinate{
						Id: "test_coordinate_id",
						X:  1,
						Y:  2,
					},
				).Return(
					&models.CoordinateDTO{
						Id: "test_coordinate_id",
						X:  1,
						Y:  2,
					},
					nil,
				).Do(func(ctx context.Context, coordinate *models.Coordinate) {
					fmt.Println("ConvertDBOToDTOCoordinate called with", coordinate)
				})
			},
			wantResp: &models.GetCoordinateByIdResp{
				Coordinate: &models.CoordinateDTO{
					Id: "test_coordinate_id",
					X:  1,
					Y:  2,
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := coordinateService.GetCoordinateById(tt.arg.ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetCoordinate(t *testing.T) {
	ctx, mockG, mockDTO, coordinateService := ConfigGetCoordinate(t)

	type arg struct {
		req *models.GetCoordinateReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg
		before   func(t *testing.T)
		wantResp *models.GetCoordinateResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.GetCoordinateReq{
				},
			},
			before: func(t *testing.T) {
				mockG.EXPECT().GetCoordinate(ctx).Return(
					[]*models.Coordinate{
						{
							Id: "test_coordinate_id",
							X:  1,
							Y:  2,
						},
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOCoordinate(
					ctx,
					&models.Coordinate{
						Id: "test_coordinate_id",
						X:  1,
						Y:  2,
					},
				).Return(
					&models.CoordinateDTO{
						Id: "test_coordinate_id",
						X:  1,
						Y:  2,
					},
					nil,
				).Do(func(ctx context.Context, coordinate *models.Coordinate) {
					fmt.Println("ConvertDBOToDTOCoordinate called with", coordinate)
				})
			},
			wantResp: &models.GetCoordinateResp{
				Coordinates: []*models.CoordinateDTO{
					{
						Id: "test_coordinate_id",
						X:  1,
						Y:  2,
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := coordinateService.GetCoordinate(tt.arg.ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestUpdateCoordinate(t *testing.T) {
	ctx, mockU, mockDTO, coordinateService := ConfigUpdateCoordinate(t)

	type arg struct {
		req *models.UpdateCoordinateReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg
		before   func(t *testing.T)
		wantResp *models.UpdateCoordinateResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.UpdateCoordinateReq{
					Id: "test_coordinate_id",
					X: 1,
					Y: 2,
				},
			},
			before: func(t *testing.T) {
				mockU.EXPECT().UpdateCoordinate(ctx, gomock.Any()).Return(
					&models.Coordinate{
						Id: "test_coordinate_id",
						X:  1,
						Y:  2,
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOCoordinate(
					ctx,
					&models.Coordinate{
						Id: "test_coordinate_id",
						X:  1,
						Y:  2,
					},
				).Return(
					&models.CoordinateDTO{
						Id: "test_coordinate_id",
						X:  1,
						Y:  2,
					},
					nil,
				).Do(func(ctx context.Context, coordinate *models.Coordinate) {
					fmt.Println("ConvertDBOToDTOCoordinate called with", coordinate)
				})
			},
			wantResp: &models.UpdateCoordinateResp{
				Coordinate: &models.CoordinateDTO{
					Id: "test_coordinate_id",
					X:  1,
					Y:  2,
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := coordinateService.UpdateCoordinate(tt.arg.ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestDeleteCoordinate(t *testing.T) {
	ctx, mockD, mockDTO, coordinateService := ConfigDeleteCoordinate(t)

	type arg struct {
		req *models.DeleteCoordinateReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg
		before   func(t *testing.T)
		wantResp *models.DeleteCoordinateResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.DeleteCoordinateReq{
					Id: "test_coordinate_id",
				},
			},
			before: func(t *testing.T) {
				mockD.EXPECT().DeleteCoordinate(
					ctx,
					gomock.Any(),
				).Return(
					&models.Coordinate{
						Id: "test_coordinate_id",
						X:  1,
						Y:  2,
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOCoordinate(
					ctx,
					&models.Coordinate{
						Id: "test_coordinate_id",
						X:  1,
						Y:  2,
					},
				).Return(
					&models.CoordinateDTO{
						Id: "test_coordinate_id",
						X:  1,
						Y:  2,
					},
					nil,
				).Do(func(ctx context.Context, coordinate *models.Coordinate) {
					fmt.Println("ConvertDBOToDTOCoordinate called with", coordinate)
				})
			},
			wantResp: &models.DeleteCoordinateResp{
				Coordinate: &models.CoordinateDTO{
					Id: "test_coordinate_id",
					X:  1,
					Y:  2,
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := coordinateService.DeleteCoordinate(tt.arg.ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestConvertCoordinateToDTOCoordinate(t *testing.T) {
	ctx, mockF, coordinateService := ConfigConvertToDTOCoordinate(t)

	type arg struct {
		req *models.Coordinate
		ctx context.Context
	}

	tests := []struct {
		name     string
		before   func(t *testing.T)
		arg      arg
		wantResp *models.CoordinateDTO
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				req: &models.Coordinate{
					Id: "test_coordinate_id",
					X:  1,
					Y:  2,
				},
				ctx: ctx,
			},
			before: func(t *testing.T) {
				mockF.EXPECT().Execute(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					&models.RespFacade{
						GCRB: &models.GetCoordinateByIdResp{
							Coordinate: &models.CoordinateDTO{
								Id: "test_coordinate_id",
								X:  1,
								Y:  2,
							},
						},
					},
					nil,
				)
			},
			wantResp: &models.CoordinateDTO{
				Id: "test_coordinate_id",
				X:  1,
				Y:  2,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := coordinateService.ConvertDBOToDTOCoordinate(tt.arg.ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
