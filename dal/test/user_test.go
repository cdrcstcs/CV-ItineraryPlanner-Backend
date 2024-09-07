package test

import (
	"context"
	"itineraryplanner/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"itineraryplanner/dal/db"
	"itineraryplanner/dal"

	"itineraryplanner/constant"
)

func TestCreateUser(t *testing.T){
	ctx := context.Background()

	type arg struct {
		ctx context.Context
		user *models.User
	}

	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantUser *models.User
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
				user: &models.User{
					Name    :"test",
					Password   :"test",
					Email  :"test",
					EmailPassword :"test",
				},
			},
			wantUser: &models.User{
				Id  :"generatedId",
				Name    :"test",
				Password   :"test",
				Email  :"test",
				EmailPassword :"test",
			},
			wantErr: nil ,
		},
	}
	userDal := &dal.UserDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotRating, gotErr:= userDal.CreateUser(ctx, tt.arg.user)
			assert.Equal(t, gotRating, tt.wantUser)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}

func TestGetUserById(t *testing.T){
	ctx := context.Background()

	type arg struct {
		ctx context.Context
		userId string
	}

	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantUser *models.User
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
				userId: "generated_id",
			},
			wantUser: &models.User{
				Id  :"generatedId",
				Name    :"test",
				Password   :"test",
				Email  :"test",
				EmailPassword :"test",
			},
			wantErr: nil ,
		},
	}
	userDal := &dal.UserDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotUser, gotErr:= userDal.GetUserById(ctx, tt.arg.userId)
			assert.Equal(t, gotUser, tt.wantUser)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}
func TestGetUser(t *testing.T){
	ctx := context.Background()

	type arg struct {
		ctx context.Context
	}

	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantUser []*models.User
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
			},
			wantUser: []*models.User{
				{
					Id  :"generatedId",
					Name    :"test",
					Password   :"test",
					Email  :"test",
					EmailPassword :"test",
				},
			},
			wantErr: nil ,
		},
	}
	userDal := &dal.UserDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotUser, gotErr:= userDal.GetUser(ctx)
			assert.Equal(t, gotUser, tt.wantUser)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}

func TestUpdateUser(t *testing.T){
	ctx := context.Background()

	type arg struct {
		ctx context.Context
		user *models.User
	}

	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantUser *models.User
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
				user: &models.User{
					Id  :"generatedId",
					Name    :"test",
					Password   :"test",
					Email  :"test",
					EmailPassword :"test",
				},
			},
			wantUser: &models.User{
				Id  :"generatedId",
				Name    :"test",
				Password   :"test",
				Email  :"test",
				EmailPassword :"test",
			},
			wantErr: nil ,
		},
	}
	userDal := &dal.UserDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotRating, gotErr:= userDal.UpdateUser(ctx, tt.arg.user)
			assert.Equal(t, gotRating, tt.wantUser)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}

func TestDeleteUser(t *testing.T){
	ctx := context.Background()

	type arg struct {
		ctx context.Context
		userId string
	}

	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantUser *models.User
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
				userId: "generated_id",
			},
			wantUser: &models.User{
				Id  :"generatedId",
				Name    :"test",
				Password   :"test",
				Email  :"test",
				EmailPassword :"test",
			},
			wantErr: nil ,
		},
	}
	userDal := &dal.UserDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotUser, gotErr:= userDal.DeleteUser(ctx, tt.arg.userId)
			assert.Equal(t, gotUser, tt.wantUser)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}