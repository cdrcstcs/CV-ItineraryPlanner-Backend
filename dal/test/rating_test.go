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
func TestGetRatingById(t *testing.T) {
    ctx := context.Background()
    type arg struct {
        ctx       context.Context
        ratingId  string
    }
    tests := []struct {
        name      string
        before    func(t *testing.T) (string, error) 
        arg
        wantRating *models.Rating
        wantErr    error
    }{
        {
            name: "success",
            before: func(t *testing.T) (string, error) {
                ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
                rating := &models.Rating{
                    Score: 5,
                }
                newRating, err := ratingDal.CreateRating(ctx, rating)
                if err != nil {
                    t.Fatalf("Failed to create rating: %v", err)
                    return "", err
                }
                return newRating.Id, nil
            },
            arg: arg{
                ctx: ctx,
            },
            wantRating: &models.Rating{
                Score: 5, 
            },
            wantErr: nil,
        },
    }
    ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ratingId, err := tt.before(t)
            if err != nil {
                t.Fatalf("Setup failed: %v", err)
            }
            tt.arg.ratingId = ratingId
            gotRating, gotErr := ratingDal.GetRatingById(tt.arg.ctx, tt.arg.ratingId)
			tt.wantRating.Id = ratingId
            assert.Equal(t, tt.wantRating, gotRating) 
            assert.Equal(t, tt.wantErr, gotErr)
        })
    }
}
func TestGetRating(t *testing.T){
	ctx := context.Background()
	type arg struct {
		ctx context.Context
	}
	ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
	r, err := ratingDal.GetRating(ctx)
	if err != nil {
        t.Fatalf("Failed to retrieve expected ratings: %v", err)
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
func TestUpdateRating(t *testing.T) {
    ctx := context.Background()
    ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
    type arg struct {
        ctx    context.Context
        rating *models.Rating
    }
    tests := []struct {
        name      string
        before    func(t *testing.T) string 
        arg       arg
        wantRating *models.Rating
        wantErr   error
    }{
        {
            name: "success",
            before: func(t *testing.T) string {
                rating := &models.Rating{
                    Score: 3, 
                }
                newRating, err := ratingDal.CreateRating(ctx, rating)
                if err != nil {
                    t.Fatalf("Failed to create rating: %v", err)
                }
                return newRating.Id
            },
            arg: arg{
                ctx: ctx,
                rating: &models.Rating{
                    Id:    "", 
                    Score: 5,  
                },
            },
            wantRating: &models.Rating{
				Id: "",
                Score: 5,
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ratingId := tt.before(t)
            tt.arg.rating.Id = ratingId 
            gotRating, gotErr := ratingDal.UpdateRating(ctx, tt.arg.rating)
            assert.Equal(t, tt.wantErr, gotErr)
			tt.wantRating.Id = ratingId
            assert.Equal(t, tt.wantRating, gotRating)
        })
    }
}
func TestDeleteRating(t *testing.T) {
    ctx := context.Background()
    ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
    type arg struct {
        ctx       context.Context
        ratingId  string
    }
    tests := []struct {
        name      string
        before    func(t *testing.T) string 
        arg       arg
        wantRating *models.Rating
        wantErr   error
    }{
        {
            name: "success",
            before: func(t *testing.T) string {
                rating := &models.Rating{
                    Score: 5, 
                }
                newRating, err := ratingDal.CreateRating(ctx, rating)
                if err != nil {
                    t.Fatalf("Failed to create rating: %v", err)
                }
                return newRating.Id 
            },
            arg: arg{
                ctx: ctx,
            },
            wantRating: &models.Rating{
				Id: "",
                Score: 5, 
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ratingId := tt.before(t)
            tt.arg.ratingId = ratingId 
			tt.wantRating.Id = ratingId
            gotRating, gotErr := ratingDal.DeleteRating(ctx, tt.arg.ratingId)
            assert.Equal(t, tt.wantErr, gotErr)
            assert.Equal(t, tt.wantRating, gotRating)
        })
    }
}