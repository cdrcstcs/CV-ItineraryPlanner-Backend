package test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"itineraryplanner/common/custom_errs"
	"itineraryplanner/dal/db"
	"itineraryplanner/models"
	"itineraryplanner/dal"

)

func TestCreateItinerary(t *testing.T) {
	ctx := context.Background()

	type arg struct {
		itinerary *models.Itinerary
		ctx   context.Context
	}

	tests := []struct {
		name          string
		before        func(t *testing.T)
		arg           arg
		wantErr       error
		wantItinerary     *models.Itinerary
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				itinerary: &models.Itinerary{
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					CopiedId :"1",
					UserId  : "1",
					EventIds  :  []string {"1","2"},
					EventCount: 2,
					RatingId:"1",
				},
			},
			wantItinerary: &models.Itinerary{
				StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
				EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
				CopiedId :"1",
				UserId  : "1",
				EventIds  :  []string {"1","2"},
				EventCount: 2,
				RatingId:"1",
			},
		},
		{
			name: "with id err",
			arg: arg{
				ctx: ctx,
				itinerary: &models.Itinerary{
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					Id    :"1",
					CopiedId :"1",
					UserId  : "1",
					EventIds  :  []string {"1","2"},
					EventCount: 2,
					RatingId:"1",
				},
			},
			wantErr: custom_errs.DBErrCreateWithID,
		},
	}
	itineraryDal := &dal.ItineraryDal{MainDB: db.GetMemoMongo()}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before(t)
			}
			gotItinerary, err := itineraryDal.CreateItinerary(tt.arg.ctx, tt.arg.itinerary)

			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotItinerary.Id)
			gotItinerary.Id = "" 
			assert.Equal(t, tt.wantItinerary, gotItinerary)

		})
	}
}

func TestGetItineraryById(t *testing.T) {
	ctx:= context.Background()
	type arg struct {
		ctx           context.Context
		itineraryId   string
	}
	
	tests := []struct {
		name            string
		before          func(t *testing.T)
		arg
		wantItinerary *models.Itinerary
		wantErr         error
	}{
		{
			name: "success",
			arg: arg{
				ctx:           ctx,
				itineraryId:  "66420fdc40732b8c0653b530",
			},
			wantItinerary: &models.Itinerary{
				StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
				EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
				Id    : 	"66420fdc40732b8c0653b530",
				CopiedId :"1",
				UserId  : "1",
				EventIds  :  []string {"1","2"},
				EventCount: 2,
				RatingId:"1",
			},
			wantErr: nil,
		},
	}
	
	itineraryDal := &dal.ItineraryDal{MainDB: db.GetMemoMongo()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItinerary, err := itineraryDal.GetItineraryById(tt.arg.ctx, tt.arg.itineraryId)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotItinerary)
			assert.Equal(t, tt.wantItinerary, gotItinerary)
		})	
	}
}

func TestGetItinerary(t *testing.T) {
	ctx:= context.Background()
	type arg struct {
		ctx           context.Context
	}
	itineraryDal := &dal.ItineraryDal{MainDB: db.GetMemoMongo()}
	i, _ := itineraryDal.GetItinerary(ctx)
	tests := []struct {
		name            string
		before          func(t *testing.T)
		arg
		wantItinerary []*models.Itinerary
		wantErr         error
	}{
		{
			name: "success",
			arg: arg{
				ctx:           ctx,
			},
			wantItinerary: i,
			wantErr: nil,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItinerary, err := itineraryDal.GetItinerary(tt.arg.ctx)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotItinerary)
			assert.ElementsMatch(t, tt.wantItinerary, gotItinerary)
		})	
	}
}

func TestUpdateItinerary(t *testing.T) {
	ctx:= context.Background()
	type arg struct {
		ctx           context.Context
		itinerary     *models.Itinerary
	}
	
	tests := []struct {
		name            string
		before          func(t *testing.T)
		arg
		wantItinerary *models.Itinerary
		wantErr         error
	}{
		{
			name: "success",
			arg: arg{
				ctx:           ctx,
				itinerary:  &models.Itinerary{
					StartTime:   time.Now(),
					EndTime:     time.Now().Add(1 * time.Hour),
					Id    :"1234",
					CopiedId :"1",
					UserId  : "1",
					EventIds  :  []string {"1","2"},
					EventCount: 2,
					RatingId:"1",
				},
			},
			wantItinerary: &models.Itinerary{
				StartTime:   time.Now(),
				EndTime:     time.Now().Add(1 * time.Hour),
				Id    :"1234",
				CopiedId :"1",
				UserId  : "1",
				EventIds  :  []string {"1","2"},
				EventCount: 2,
				RatingId:"1",
			},
			wantErr: nil,
		},
	}
	
	itineraryDal := &dal.ItineraryDal{MainDB: db.GetMemoMongo()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItinerary, err := itineraryDal.UpdateItinerary(tt.arg.ctx, tt.arg.itinerary)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotItinerary)
			assert.ElementsMatch(t, tt.wantItinerary, gotItinerary)
		})	
	}
}

func TestDeleteItinerary(t *testing.T) {
	ctx:= context.Background()
	type arg struct {
		ctx           context.Context
		itineraryId   string
	}
	
	tests := []struct {
		name            string
		before          func(t *testing.T)
		arg
		wantItinerary   *models.Itinerary
		wantErr         error
	}{
		{
			name: "success",
			arg: arg{
				ctx:           ctx,
				itineraryId:  "1234",
			},
			wantItinerary: &models.Itinerary{
				StartTime:   time.Now(),
				EndTime:     time.Now().Add(1 * time.Hour),
				Id    :"1234",
				CopiedId :"1",
				UserId  : "1",
				EventIds  :  []string {"1","2"},
				EventCount: 2,
				RatingId:"1",
			},
			wantErr: nil,
		},
		{
			name: "invalid id",
			arg: arg{
				ctx:           ctx,
				itineraryId:  "12345",
			},
			wantItinerary: nil,
			wantErr: custom_errs.DBErrDeleteWithID,
		},
	}
	
	itineraryDal := &dal.ItineraryDal{MainDB: db.GetMemoMongo()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItinerary, err := itineraryDal.DeleteItinerary(tt.arg.ctx, tt.arg.itineraryId)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotItinerary)
			assert.ElementsMatch(t, tt.wantItinerary, gotItinerary)
		})	
	}
}