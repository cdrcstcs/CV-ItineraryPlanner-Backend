package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"itineraryplanner/common/custom_errs"
	"itineraryplanner/constant"
	"itineraryplanner/dal/db"
	"itineraryplanner/models"
	"itineraryplanner/dal"
)

func TestCreateAttraction(t *testing.T) {
	ctx := context.Background()

	type arg struct {
		attraction *models.Attraction
		ctx        context.Context
	}

	tests := []struct {
		name   string
		before func(t *testing.T)
		arg
		wantErr        error
		wantAttraction *models.Attraction
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				attraction: &models.Attraction{
					Name: "test",
					CoordinateId: "1",
					Address: "123 Paya Lebar",
					TagIDs: []string{"1","2","3"},
				},
			},
			wantAttraction: &models.Attraction{
				Name: "test",
				CoordinateId: "1",
				Address: "123 Paya Lebar",
				TagIDs: []string{"1","2","3"},
			},
		},
		{
			name: "with id err",
			arg: arg{
				ctx: ctx,
				attraction: &models.Attraction{
					Id:   "1",
					Name: "test",
					CoordinateId: "1",
					Address: "123 Paya Lebar",
					TagIDs: []string{"1","2","3"},
				},
			},
			wantErr: custom_errs.DBErrCreateWithID,
		},
	}

	attractionDal := &dal.AttractionDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAttraction, err := attractionDal.CreateAttraction(tt.arg.ctx, tt.arg.attraction)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotAttraction.Id)
			gotAttraction.Id = ""
			assert.Equal(t, tt.wantAttraction, gotAttraction)
		})
	}
}
func TestGetAttractionById(t *testing.T) {
	ctx:= context.Background()
	type arg struct {
		ctx           context.Context
		attractionId string
	}
	
	tests := []struct {
		name            string
		before          func(t *testing.T)
		arg
		wantAttraction *models.Attraction
		wantErr         error
	}{
		{
			name: "success",
			arg: arg{
				ctx:           ctx,
				attractionId:  "1234",
			},
			wantAttraction: &models.Attraction{
				Id:         "1234",
				Name:       "water slide",
				Address:    "24 paya lebar",
				CoordinateId: "1",
				TagIDs:     []string{"1", "2", "3"},
			},
			wantErr: nil,
		},
		{
			name: "invalid id",
			arg: arg{
				ctx:           ctx,
				attractionId:  "12345",
			},
			wantAttraction: nil,
			wantErr: custom_errs.DBErrGetWithID,
		},
	}
	
	attractionDal := dal.AttractionDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAttraction, err := attractionDal.GetAttractionById(tt.arg.ctx, tt.arg.attractionId)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotAttraction)
			assert.ElementsMatch(t, tt.wantAttraction, gotAttraction)
		})	
	}
}
func TestGetAttraction(t *testing.T) {
	ctx:= context.Background()
	type arg struct {
		ctx           context.Context
	}
	
	tests := []struct {
		name            string
		before          func(t *testing.T)
		arg
		wantAttraction []*models.Attraction
		wantErr         error
	}{
		{
			name: "success",
			arg: arg{
				ctx:           ctx,
			},
			wantAttraction: []*models.Attraction{
				{
					Id:         "1234",
					Name:       "water slide",
					Address:    "24 paya lebar",
					CoordinateId: "1",
					TagIDs:     []string{"1", "2", "3"},
				},
			},
			wantErr: nil,
		},
	}
	
	attractionDal := dal.AttractionDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAttraction, err := attractionDal.GetAttraction(tt.arg.ctx)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotAttraction)
			assert.ElementsMatch(t, tt.wantAttraction, gotAttraction)
		})	
	}
}


func TestUpdateAttraction(t *testing.T) {
	ctx:= context.Background()
	type arg struct {
		ctx           context.Context
		attraction    *models.Attraction
	}
	
	tests := []struct {
		name            string
		before          func(t *testing.T)
		arg
		wantAttraction *models.Attraction
		wantErr         error
	}{
		{
			name: "success",
			arg: arg{
				ctx:           ctx,
				attraction:  &models.Attraction{
					Id:         "1234",
					Name:       "water slide",
					Address:    "24 paya lebar",
					CoordinateId: "1",
					TagIDs:     []string{"1", "2", "3"},
				},
			},
			wantAttraction: &models.Attraction{
				Id:         "1234",
				Name:       "water slide",
				Address:    "24 paya lebar",
				CoordinateId: "1",
				TagIDs:     []string{"1", "2", "3"},
			},
			wantErr: nil,
		},
	}
	
	attractionDal := dal.AttractionDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAttraction, err := attractionDal.UpdateAttraction(tt.arg.ctx, tt.arg.attraction)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotAttraction)
			assert.ElementsMatch(t, tt.wantAttraction, gotAttraction)
		})	
	}
}

func TestDeleteAttraction(t *testing.T) {
	ctx:= context.Background()
	type arg struct {
		ctx           context.Context
		attractionId    string
	}
	
	tests := []struct {
		name            string
		before          func(t *testing.T)
		arg
		wantAttraction *models.Attraction
		wantErr         error
	}{
		{
			name: "success",
			arg: arg{
				ctx:           ctx,
				attractionId:  "1234",
			},
			wantAttraction: &models.Attraction{
				Id:         "1234",
				Name:       "water slide",
				Address:    "24 paya lebar",
				CoordinateId: "1",
				TagIDs:     []string{"1", "2", "3"},
			},
			wantErr: nil,
		},
	}
	
	attractionDal := dal.AttractionDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAttraction, err := attractionDal.DeleteAttraction(tt.arg.ctx, tt.arg.attractionId)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotAttraction)
			assert.ElementsMatch(t, tt.wantAttraction, gotAttraction)
		})	
	}
}