package test

import (
	"context"
	"itineraryplanner/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"itineraryplanner/dal/db"
	"itineraryplanner/dal"

)

func TestCreateRating(t *testing.T){
	ctx := context.Background()

	type arg struct {
		ctx context.Context
		rating *models.Rating
	}

	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantRating *models.Rating
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
				rating: &models.Rating{
					Score:5,
				},
			},
			wantRating: &models.Rating{
				Score:5,
			},
			wantErr: nil ,
		},
	}
	ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotRating, gotErr:= ratingDal.CreateRating(ctx, tt.arg.rating)
			gotRating.Id = ""
			assert.Equal(t, gotRating, tt.wantRating)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}

func TestGetRatingById(t *testing.T){
	ctx := context.Background()

	type arg struct {
		ctx context.Context
		ratingId string
	}

	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantRating *models.Rating
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
				ratingId: "6642102153edf9e8fb3dfcda",
			},
			wantRating: &models.Rating{
				Id:"6642102153edf9e8fb3dfcda",
				Score:5,
			},
			wantErr: nil ,
		},
	}
	ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotRating, gotErr:= ratingDal.GetRatingById(ctx, tt.arg.ratingId)
			assert.Equal(t, gotRating, tt.wantRating)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}

func TestGetRating(t *testing.T){
	ctx := context.Background()

	type arg struct {
		ctx context.Context
	}
	ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
	r, _ := ratingDal.GetRating(ctx)
	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantRating []*models.Rating
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
			},
			wantRating: r,
			wantErr: nil ,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotRating, gotErr:= ratingDal.GetRating(ctx)
			assert.Equal(t, gotRating, tt.wantRating)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}

func TestUpdateRating(t *testing.T){
	ctx := context.Background()

	type arg struct {
		ctx context.Context
		rating *models.Rating
	}

	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantRating *models.Rating
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
				rating :   &models.Rating{
					Id:"generated_id",
					Score:5,
				},
			},
			wantRating: &models.Rating{
				Id:"generated_id",
				Score:5,
			},
			wantErr: nil ,
		},
	}
	ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotRating, gotErr:= ratingDal.UpdateRating(ctx, tt.arg.rating)
			assert.Equal(t, gotRating, tt.wantRating)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}

func TestDeleteRating(t *testing.T){
	ctx := context.Background()

	type arg struct {
		ctx context.Context
		ratingId string
	}

	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantRating *models.Rating
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
				ratingId: "generated_id",
			},
			wantRating: &models.Rating{
				Id:"generated_id",
				Score:5,
			},
			wantErr: nil ,
		},
	}
	ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotRating, gotErr:= ratingDal.DeleteRating(ctx, tt.arg.ratingId)
			assert.Equal(t, gotRating, tt.wantRating)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}