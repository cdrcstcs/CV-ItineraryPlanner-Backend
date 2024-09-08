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
					Id    :"1",
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
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
    ctx := context.Background()
	type arg struct {
		ctx           context.Context
		itineraryId   string
	}    
	createItinerary := func(t *testing.T) string {
        itineraryDal := &dal.ItineraryDal{MainDB: db.GetMemoMongo()}
        itinerary := &models.Itinerary{
            StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
            EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
            CopiedId:     "1",
            UserId:       "1",
            EventIds:     []string{"1", "2"},
            EventCount:   2,
            RatingId:     "1",
        }
        createdItinerary, err := itineraryDal.CreateItinerary(ctx, itinerary) 
        if err != nil {
            t.Fatalf("Failed to create itinerary: %v", err)
        }
        return createdItinerary.Id
    }
    tests := []struct {
        name            string
        before          func(t *testing.T) string 
        arg
        wantItinerary *models.Itinerary
        wantErr         error
    }{
        {
            name: "success",
            before: createItinerary,
            arg: arg{
                ctx:           ctx,
                itineraryId:  "", 
            },
            wantItinerary: &models.Itinerary{
                StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
                EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
                Id:           "", 
                CopiedId:     "1",
                UserId:       "1",
                EventIds:     []string{"1", "2"},
                EventCount:   2,
                RatingId:     "1",
            },
            wantErr: nil,
        },
    }
    itineraryDal := &dal.ItineraryDal{MainDB: db.GetMemoMongo()}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            itineraryId := tt.before(t)
            tt.arg.itineraryId = itineraryId
            gotItinerary, err := itineraryDal.GetItineraryById(tt.arg.ctx, tt.arg.itineraryId)
            assert.Equal(t, tt.wantErr, err)
            if err != nil {
                return
            }
			tt.wantItinerary.Id = itineraryId
            assert.NotEmpty(t, gotItinerary)
            assert.Equal(t, tt.wantItinerary, gotItinerary)
        })
    }
}
func TestGetItinerary(t *testing.T) {
	ctx:= context.Background()
	type arg struct {
		ctx	context.Context
	}
	itineraryDal := &dal.ItineraryDal{MainDB: db.GetMemoMongo()}
	i, _ := itineraryDal.GetItinerary(ctx)
	tests := []struct {
		name            string
		before          func(t *testing.T)
		arg
		wantItinerary 	[]*models.Itinerary
		wantErr         error
	}{
		{
			name: "success",
			arg: arg{
				ctx:	ctx,
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
			assert.Equal(t, tt.wantItinerary, gotItinerary)
		})	
	}
}
func TestUpdateItinerary(t *testing.T) {
    ctx := context.Background()
	type arg struct {
		ctx           context.Context
		itinerary     *models.Itinerary
	}
    createItinerary := func(t *testing.T) (string) {
        itineraryDal := &dal.ItineraryDal{MainDB: db.GetMemoMongo()}
        itinerary := &models.Itinerary{
            StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
            EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
            CopiedId:     "1",
            UserId:       "1",
            EventIds:     []string{"1", "2"},
            EventCount:   2,
            RatingId:     "1",
        }
        createdItinerary, err := itineraryDal.CreateItinerary(ctx, itinerary) 
        if err != nil {
            t.Fatalf("Failed to create itinerary: %v", err)
        }
        return createdItinerary.Id
    }
    tests := []struct {
        name            string
        before          func(t *testing.T) (string) 
        arg
        wantItinerary *models.Itinerary
        wantErr       error
    }{
        {
            name: "success",
            before: createItinerary,
            arg: arg{
                ctx:           ctx,
                itinerary:  &models.Itinerary{
                    Id:          "",
                    StartTime:   time.Date(2024, 5, 13, 6, 0, 0, 0, time.UTC), 
                    EndTime:     time.Date(2024, 5, 13, 7, 0, 0, 0, time.UTC),
                    CopiedId:    "1",
                    UserId:      "1",
                    EventIds:    []string{"1", "2"},
                    EventCount:  2,
                    RatingId:    "1",
                },
            },
            wantItinerary: &models.Itinerary{
                Id:          "", 
                StartTime:   time.Date(2024, 5, 13, 6, 0, 0, 0, time.UTC),
                EndTime:     time.Date(2024, 5, 13, 7, 0, 0, 0, time.UTC),
                CopiedId:    "1",
                UserId:      "1",
                EventIds:    []string{"1", "2"},
                EventCount:  2,
                RatingId:    "1",
            },
            wantErr: nil,
        },
    }
    itineraryDal := &dal.ItineraryDal{MainDB: db.GetMemoMongo()}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            itineraryId := tt.before(t)
            tt.arg.itinerary.Id = itineraryId
            gotItinerary, err := itineraryDal.UpdateItinerary(tt.arg.ctx, tt.arg.itinerary)
            assert.Equal(t, tt.wantErr, err)
            if err != nil {
                return
            }
			tt.wantItinerary.Id = itineraryId
            assert.NotEmpty(t, gotItinerary)
            assert.Equal(t, tt.wantItinerary, gotItinerary)
        })
    }
}
func TestDeleteItinerary(t *testing.T) {
    ctx := context.Background()
	type arg struct {
		ctx           context.Context
		itineraryId   string
	}
    createItinerary := func(t *testing.T) string {
        itineraryDal := &dal.ItineraryDal{MainDB: db.GetMemoMongo()}
        itinerary := &models.Itinerary{
            StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
            EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
            CopiedId:     "1",
            UserId:       "1",
            EventIds:     []string{"1", "2"},
            EventCount:   2,
            RatingId:     "1",
        }
        createdItinerary, err := itineraryDal.CreateItinerary(ctx, itinerary) 
        if err != nil {
            t.Fatalf("Failed to create itinerary: %v", err)
        }
        return createdItinerary.Id
    }
    tests := []struct {
        name            string
        before          func(t *testing.T) string 
        arg
        wantItinerary   *models.Itinerary
        wantErr         error
    }{
        {
            name: "success",
            before: createItinerary,
            arg: arg{
                ctx:           ctx,
                itineraryId:  "", 
            },
            wantItinerary: &models.Itinerary{
				Id: "",
                StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
                EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
                CopiedId:     "1",
                UserId:       "1",
                EventIds:     []string{"1", "2"},
                EventCount:   2,
                RatingId:     "1",
            },
            wantErr: nil,
        },
        {
            name: "invalid id",
            before: func(t *testing.T) string {
                return "nonexistentID"
            },
            arg: arg{
                ctx:           ctx,
                itineraryId:  "", 
            },
            wantItinerary: nil, 
            wantErr:       custom_errs.DBErrIDConversion,
        },
    }
    itineraryDal := &dal.ItineraryDal{MainDB: db.GetMemoMongo()}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            itineraryId := tt.before(t)
            tt.arg.itineraryId = itineraryId
            gotItinerary, err := itineraryDal.DeleteItinerary(tt.arg.ctx, tt.arg.itineraryId)
            assert.Equal(t, tt.wantErr, err)
            if err != nil {
                assert.Equal(t, tt.wantItinerary, gotItinerary)
                return
            }
			tt.wantItinerary.Id = itineraryId
            assert.NotEmpty(t, gotItinerary)
            assert.Equal(t, tt.wantItinerary, gotItinerary)
        })
    }
}