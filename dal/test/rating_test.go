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
					Type:1,
					UserId:"user_id",
					ObjectId:"object_id",
					Score:5,
				},
			},
			wantRating: &models.Rating{
				Id:"generated_id",
				Type:1,
				UserId:"user_id",
				ObjectId:"object_id",
				Score:5,
			},
			wantErr: nil ,
		},
	}
	ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotRating, gotErr:= ratingDal.CreateRating(ctx, tt.arg.rating)
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
				ratingId: "generated_id",
			},
			wantRating: &models.Rating{
				Id:"generated_id",
				Type:1,
				UserId:"user_id",
				ObjectId:"object_id",
				Score:5,
			},
			wantErr: nil ,
		},
	}
	ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
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
			wantRating: []*models.Rating{
				{
					Id:"generated_id",
					Type:1,
					UserId:"user_id",
					ObjectId:"object_id",
					Score:5,
				},
			},
			wantErr: nil ,
		},
	}
	ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
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
					Type:1,
					UserId:"user_id",
					ObjectId:"object_id",
					Score:5,
				},
			},
			wantRating: &models.Rating{
				Id:"generated_id",
				Type:1,
				UserId:"user_id",
				ObjectId:"object_id",
				Score:5,
			},
			wantErr: nil ,
		},
	}
	ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
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
				Type:1,
				UserId:"user_id",
				ObjectId:"object_id",
				Score:5,
			},
			wantErr: nil ,
		},
	}
	ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotRating, gotErr:= ratingDal.DeleteRating(ctx, tt.arg.ratingId)
			assert.Equal(t, gotRating, tt.wantRating)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}