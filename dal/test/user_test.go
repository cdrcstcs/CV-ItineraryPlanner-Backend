package test
import (
	"context"
	"itineraryplanner/models"
	"testing"
	"github.com/stretchr/testify/assert"
	"itineraryplanner/dal/db"
	"itineraryplanner/dal"

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
				Name    :"test",
				Password   :"test",
				Email  :"test",
				EmailPassword :"test",
			},
			wantErr: nil ,
		},
	}
	userDal := &dal.UserDal{MainDB: db.GetMemoMongo()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotUser, gotErr:= userDal.CreateUser(ctx, tt.arg.user)
			gotUser.Id = ""
			assert.Equal(t, gotUser, tt.wantUser)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}
func TestGetUserById(t *testing.T) {
    ctx := context.Background()
    userDal := &dal.UserDal{MainDB: db.GetMemoMongo()}
    type arg struct {
        ctx     context.Context
        userId  string
    }
    tests := []struct {
        name     string
        before   func(t *testing.T) string 
        arg      arg
        wantUser *models.User
        wantErr  error
    }{
        {
            name: "success",
            before: func(t *testing.T) string {
                user := &models.User{
                    Name:          "test",
                    Password:      "test",
                    Email:         "test",
                    EmailPassword: "test",
                }
                newUser, err := userDal.CreateUser(ctx, user) 
                if err != nil {
                    t.Fatalf("Failed to create user: %v", err)
                }
                return newUser.Id 
            },
            arg: arg{
                ctx:    ctx,
                userId: "", 
            },
            wantUser: &models.User{
                Name:          "test",
                Password:      "test",
                Email:         "test",
                EmailPassword: "test",
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            userId := tt.before(t)
            tt.arg.userId = userId
            gotUser, gotErr := userDal.GetUserById(tt.arg.ctx, tt.arg.userId)
            if tt.wantErr != nil {
                assert.Equal(t, tt.wantErr, gotErr)
                assert.Nil(t, gotUser)
                return
            }
			tt.wantUser.Id = userId
            assert.Equal(t, tt.wantUser, gotUser)
            assert.Nil(t, gotErr)
        })
    }
}
func TestGetUser(t *testing.T){
	ctx := context.Background()
	type arg struct {
		ctx context.Context
	}
	userDal := &dal.UserDal{MainDB: db.GetMemoMongo()}
	us, err := userDal.GetUser(ctx)
	if err != nil {
        t.Fatalf("Failed to retrieve expected users: %v", err)
    }
	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantUsers []*models.User
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
			},
			wantUsers: us, 
			wantErr: nil ,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotUser, gotErr:= userDal.GetUser(ctx)
			assert.Equal(t, gotUser, tt.wantUsers)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}
func TestUpdateUser(t *testing.T) {
    ctx := context.Background()
    userDal := &dal.UserDal{MainDB: db.GetMemoMongo()}
    type arg struct {
        ctx  context.Context
        user *models.User
    }
    tests := []struct {
        name     string
        before   func(t *testing.T) string 
        arg      arg
        wantUser *models.User
        wantErr  error
    }{
        {
            name: "success",
            before: func(t *testing.T) string {
                user := &models.User{
                    Name:          "test",
                    Password:      "test",
                    Email:         "test",
                    EmailPassword: "test",
                }
                newUser, err := userDal.CreateUser(ctx, user) 
                if err != nil {
                    t.Fatalf("Failed to create user: %v", err)
                }
                return newUser.Id
            },
            arg: arg{
                ctx: ctx,
                user: &models.User{
                    Id:    "", 
                    Name:  "updated_test", 
                    Password:      "updated_test",
                    Email:         "updated_test",
                    EmailPassword: "updated_test",
                },
            },
            wantUser: &models.User{
                Name:          "updated_test",
                Password:      "updated_test",
                Email:         "updated_test",
                EmailPassword: "updated_test",
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            userId := tt.before(t)
            tt.arg.user.Id = userId 
            gotUser, gotErr := userDal.UpdateUser(tt.arg.ctx, tt.arg.user)
            if tt.wantErr != nil {
                assert.Equal(t, tt.wantErr, gotErr)
                assert.Nil(t, gotUser)
                return
            }
			tt.wantUser.Id = userId
            assert.Equal(t, tt.wantUser, gotUser)
            assert.Nil(t, gotErr)
        })
    }
}
func TestDeleteUser(t *testing.T) {
    ctx := context.Background()
    userDal := &dal.UserDal{MainDB: db.GetMemoMongo()}
    type arg struct {
        ctx    context.Context
        userId string
    }
    tests := []struct {
        name     string
        before   func(t *testing.T) string
        arg      arg
        wantUser *models.User
        wantErr  error
    }{
        {
            name: "success",
            before: func(t *testing.T) string {
                user := &models.User{
                    Name:          "test",
                    Password:      "test",
                    Email:         "test",
                    EmailPassword: "test",
                }
                newUser, err := userDal.CreateUser(ctx, user) 
                if err != nil {
                    t.Fatalf("Failed to create user: %v", err)
                }
                return newUser.Id 
            },
            arg: arg{
                ctx:    ctx,
                userId: "", 
            },
            wantUser: &models.User{
                Name:          "test",
                Password:      "test",
                Email:         "test",
                EmailPassword: "test",
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            userId := tt.before(t)
            tt.arg.userId = userId 
            gotUser, gotErr := userDal.DeleteUser(tt.arg.ctx, tt.arg.userId)
            if tt.wantErr != nil {
                assert.Equal(t, tt.wantErr, gotErr)
                assert.Nil(t, gotUser)
                return
            }
			tt.wantUser.Id = userId
            assert.Equal(t, tt.wantUser, gotUser)
			assert.Nil(t, gotErr)
        })
    }
}