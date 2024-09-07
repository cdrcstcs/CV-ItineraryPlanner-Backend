package test

import (
	"context"
	"itineraryplanner/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"itineraryplanner/dal/db"
	"itineraryplanner/constant"
	"itineraryplanner/dal"
)

func TestCreateCoordinate(t *testing.T){
	ctx := context.Background()

	type arg struct {
		ctx context.Context
		coordinate *models.Coordinate
	}

	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantCoordinate *models.Coordinate
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
				coordinate: &models.Coordinate{
					X :1,
					Y :1,
				},
			},
			wantCoordinate: &models.Coordinate{
				Id:"generated_id",
				X :1,
				Y :1,
			},
			wantErr: nil ,
		},
	}
	coordinateDal := &dal.CoordinateDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotCoordinate, gotErr:= coordinateDal.CreateCoordinate(ctx, tt.arg.coordinate)
			assert.Equal(t, gotCoordinate, tt.wantCoordinate)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}

func TestGetCoordinateById(t *testing.T){
	ctx := context.Background()

	type arg struct {
		ctx context.Context
		coordinateId string
	}

	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantCoordinate *models.Coordinate
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
				coordinateId: "generated_id",
			},
			wantCoordinate: &models.Coordinate{
				Id:"generated_id",
				X :1,
				Y :1,
			},
			wantErr: nil ,
		},
	}
	coordinateDal := &dal.CoordinateDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotCoordinate, gotErr:= coordinateDal.GetCoordinateById(ctx, tt.arg.coordinateId)
			assert.Equal(t, gotCoordinate, tt.wantCoordinate)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}
func TestGetCoordinate(t *testing.T){
	ctx := context.Background()

	type arg struct {
		ctx context.Context
	}

	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantCoordinate []*models.Coordinate
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
			},
			wantCoordinate: []*models.Coordinate{
				{
					Id:"generated_id",
					X :1,
					Y :1,
				},
			},
			wantErr: nil ,
		},
	}
	coordinateDal := &dal.CoordinateDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotCoordinate, gotErr:= coordinateDal.GetCoordinate(ctx)
			assert.Equal(t, gotCoordinate, tt.wantCoordinate)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}

func TestUpdateCoordinate(t *testing.T){
	ctx := context.Background()

	type arg struct {
		ctx context.Context
		coordinate *models.Coordinate
	}

	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantCoordinate *models.Coordinate
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
				coordinate: &models.Coordinate{
					Id:"generated_id",
					X :1,
					Y :1,
				},			
			},
			wantCoordinate: &models.Coordinate{
				Id:"generated_id",
				X :1,
				Y :1,
			},
			wantErr: nil ,
		},
	}
	coordinateDal := &dal.CoordinateDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotCoordinate, gotErr:= coordinateDal.UpdateCoordinate(ctx, tt.arg.coordinate)
			assert.Equal(t, gotCoordinate, tt.wantCoordinate)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}

func TestDeleteCoordinate(t *testing.T){
	ctx := context.Background()

	type arg struct {
		ctx context.Context
		coordinateId string
	}

	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantCoordinate *models.Coordinate
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
				coordinateId: "generated_id",
			},
			wantCoordinate: &models.Coordinate{
				Id:"generated_id",
				X :1,
				Y :1,
			},
			wantErr: nil ,
		},
	}
	coordinateDal := &dal.CoordinateDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotCoordinate, gotErr:= coordinateDal.DeleteCoordinate(ctx, tt.arg.coordinateId)
			assert.Equal(t, gotCoordinate, tt.wantCoordinate)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}