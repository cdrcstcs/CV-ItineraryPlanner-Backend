package test
import (
	"context"
	"testing"
	"time"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"itineraryplanner/dal/mock"
	"itineraryplanner/models"
	"itineraryplanner/service"
	"fmt"
	"itineraryplanner/dal/db"
	mockser "itineraryplanner/service/mock"
	"itineraryplanner/dal"
	"go.mongodb.org/mongo-driver/bson"
	"itineraryplanner/constant"
)
func ConfigItinerary(t *testing.T) (context.Context, *mock.MockItineraryDal, *mockser.MockItineraryService, *service.ItineraryService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockDal := mock.NewMockItineraryDal(ctrl)
	mockSer := mockser.NewMockItineraryService(ctrl)
	itineraryService := &service.ItineraryService{Dal: mockDal}
	return ctx, mockDal, mockSer, itineraryService
}
func TestCreateItinerary(t *testing.T) {
    ctx, mockDal, mockSer, itineraryService := ConfigItinerary(t)
    type arg struct {
        req *models.CreateItineraryReq
        ctx context.Context
    }
    ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
    newRating, _ := ratingDal.CreateRating(ctx, &models.Rating{
        Score: 5,
    })
    tests := []struct {
        name     string
        arg      arg
        before   func(t *testing.T)
        wantResp *models.CreateItineraryResp
        wantErr  error
    }{
        {
            name: "success",
            arg: arg{
                ctx: ctx,
                req: &models.CreateItineraryReq{
                    CopiedId:  "test",
					CopiedName: "test",
					Name: "test",
                    UserId:    "test",
                    StartTime: time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
                    EndTime:   time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
                    EventIds:  []string{"test"},
                    EventCount: 1,
                    RatingId:  newRating.Id,
                },
            },
            before: func(t *testing.T) {
                mockDal.EXPECT().CreateItinerary(ctx, gomock.Any()).Return(
                    &models.Itinerary{
                        CopiedId:    "test",
                        CopiedName:  "test",
                        Name:        "test",
                        UserId:      "test",
                        StartTime:   time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
                        EndTime:     time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
                        EventIds:    []string{"test"},
                        EventCount:  1,
                        RatingId:    newRating.Id,
                    },
                    nil,
                ).MaxTimes(100)
                mockDal.EXPECT().GetDB().Return(db.GetMemoMongo()).MaxTimes(100)
                mockSer.EXPECT().ConvertDBOToDTOItinerary(ctx, &models.Itinerary{
                    CopiedId:    "test",
                    CopiedName:  "test",
                    Name:        "test",
                    UserId:      "test",
                    StartTime:   time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
                    EndTime:     time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
                    EventIds:    []string{"test"},
                    EventCount:  1,
                    RatingId:    newRating.Id,
                }).Return(
                    &models.ItineraryDTO{
                        CopiedId:    "test",
                        CopiedName:  "test",
                        Name:        "test",
                        UserId:      "test",
                        StartTime:   time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
                        EndTime:     time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
                        EventIds:    []string{"test"},
                        EventCount:  1,
                        Rating: &models.RatingDTO{
                            Id:    newRating.Id,
                            Score: 5,
                        },
                    },
                    nil,
                ).MaxTimes(100)
            },
            wantResp: &models.CreateItineraryResp{
                Itinerary: &models.ItineraryDTO{
                    CopiedId:    "test",
                    CopiedName:  "test",
                    Name:        "test",
                    UserId:      "test",
                    StartTime:   time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
                    EndTime:     time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
                    EventIds:    []string{"test"},
                    EventCount:  1,
                    Rating: &models.RatingDTO{
                        Id:    newRating.Id,
                        Score: 5,
                    },
                },
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.before(t)
            gotResp, err := itineraryService.CreateItinerary(ctx, tt.arg.req)
            assert.Equal(t, tt.wantResp, gotResp)
            fmt.Printf("Expected: %+v\n", tt.wantResp)
            fmt.Printf("Actual: %+v\n", gotResp)
            assert.Equal(t, tt.wantErr, err)
        })
    }
}
func TestGetItineraryById(t *testing.T) {
    ctx, mockDal, mockSer, itineraryService := ConfigItinerary(t)
    type arg struct {
        req *models.GetItineraryByIdReq
        ctx context.Context
    }
    ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
    newRating, _ := ratingDal.CreateRating(ctx, &models.Rating{
        Score: 5,
    })
    itineraryDal := &dal.ItineraryDal{MainDB: db.GetMemoMongo()}
    newItinerary, _ := itineraryDal.CreateItinerary(ctx, &models.Itinerary{
        CopiedId:    "test",
        CopiedName:  "test",
        Name:        "test",
        UserId:      "test",
        StartTime:   time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
        EndTime:     time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
        EventIds:    []string{"test"},
        EventCount:  1,
        RatingId:    newRating.Id,
    })
    tests := []struct {
        name     string
        arg      arg
        before   func(t *testing.T)
        wantResp *models.GetItineraryByIdResp
        wantErr  error
    }{
        {
            name: "success",
            arg: arg{
                ctx: ctx,
                req: &models.GetItineraryByIdReq{
                    Id: newItinerary.Id,
                },
            },
            before: func(t *testing.T) {
                mockDal.EXPECT().GetItineraryById(ctx, gomock.Any()).Return(
                    newItinerary,
                    nil,
                ).MaxTimes(100)
                mockDal.EXPECT().GetDB().Return(db.GetMemoMongo()).MaxTimes(100)
                mockSer.EXPECT().ConvertDBOToDTOItinerary(ctx, newItinerary).Return(
                    &models.ItineraryDTO{
                        Id:        newItinerary.Id,
                        CopiedId:  newItinerary.CopiedId,
                        CopiedName: newItinerary.CopiedName,
                        Name:      newItinerary.Name,
                        UserId:    newItinerary.UserId,
                        StartTime: newItinerary.StartTime,
                        EndTime:   newItinerary.EndTime,
                        EventIds:  newItinerary.EventIds,
                        EventCount: newItinerary.EventCount,
                        Rating: &models.RatingDTO{
                            Id:    newRating.Id,
                            Score: 5,
                        },
                    },
                    nil,
                ).MaxTimes(100)
            },
            wantResp: &models.GetItineraryByIdResp{
                Itinerary: &models.ItineraryDTO{
                    Id:        newItinerary.Id,
                    CopiedId:  newItinerary.CopiedId,
                    CopiedName: newItinerary.CopiedName,
                    Name:      newItinerary.Name,
                    UserId:    newItinerary.UserId,
                    StartTime: newItinerary.StartTime,
                    EndTime:   newItinerary.EndTime,
                    EventIds:  newItinerary.EventIds,
                    EventCount: newItinerary.EventCount,
                    Rating: &models.RatingDTO{
                        Id:    newRating.Id,
                        Score: 5,
                    },
                },
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.before(t)
            gotResp, err := itineraryService.GetItineraryById(tt.arg.ctx, tt.arg.req)
            if !assert.Equal(t, tt.wantResp, gotResp) {
                fmt.Printf("Expected: %+v\n", tt.wantResp)
                fmt.Printf("Actual: %+v\n", gotResp)
            }
            assert.Equal(t, tt.wantErr, err)
        })
    }
}
func TestGetItinerary(t *testing.T) {
    ctx, mockDal, mockService, itineraryService := ConfigItinerary(t)
    type arg struct {
        req *models.GetItineraryReq
        ctx context.Context
    }
    db.GetMemoMongo().Collection(constant.ItineraryTable).DeleteMany(ctx, bson.M{})
    ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
    newRating, _ := ratingDal.CreateRating(ctx, &models.Rating{
        Score: 5,
    })
    itineraryDal := &dal.ItineraryDal{MainDB: db.GetMemoMongo()}
    newItinerary, _ := itineraryDal.CreateItinerary(ctx, &models.Itinerary{
        CopiedId:    "test",
        CopiedName:  "test",
        Name:        "test",
        UserId:      "test",
        StartTime:   time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
        EndTime:     time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
        EventIds:    []string{"test"},
        EventCount:  1,
        RatingId:    newRating.Id,
    })
    tests := []struct {
        name     string
        arg      arg
        before   func(t *testing.T)
        wantResp *models.GetItineraryResp
        wantErr  error
    }{
        {
            name: "success",
            arg: arg{
                ctx: ctx,
                req: &models.GetItineraryReq{},
            },
            before: func(t *testing.T) {
                mockDal.EXPECT().GetItinerary(ctx).Return(
                    []*models.Itinerary{
                        newItinerary,
                    },
                    nil,
                ).MaxTimes(100)
                mockDal.EXPECT().GetDB().Return(db.GetMemoMongo()).MaxTimes(100)
                mockService.EXPECT().ConvertDBOToDTOItinerary(ctx, newItinerary).Return(
                    &models.ItineraryDTO{
                        Id:        newItinerary.Id,
                        CopiedId:  newItinerary.CopiedId,
                        CopiedName: newItinerary.CopiedName,
                        Name:      newItinerary.Name,
                        UserId:    newItinerary.UserId,
                        StartTime: newItinerary.StartTime,
                        EndTime:   newItinerary.EndTime,
                        EventIds:  newItinerary.EventIds,
                        EventCount: newItinerary.EventCount,
                        Rating: &models.RatingDTO{
                            Id:    newRating.Id,
                            Score: 5,
                        },
                    },
                    nil,
                ).MaxTimes(100)
            },
            wantResp: &models.GetItineraryResp{
                Itineraries: []*models.ItineraryDTO{
                    {
                        Id:        newItinerary.Id,
                        CopiedId:  newItinerary.CopiedId,
                        CopiedName: newItinerary.CopiedName,
                        Name:      newItinerary.Name,
                        UserId:    newItinerary.UserId,
                        StartTime: newItinerary.StartTime,
                        EndTime:   newItinerary.EndTime,
                        EventIds:  newItinerary.EventIds,
                        EventCount: newItinerary.EventCount,
                        Rating: &models.RatingDTO{
                            Id:    newRating.Id,
                            Score: 5,
                        },
                    },
                },
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.before(t)
            gotResp, err := itineraryService.GetItinerary(tt.arg.ctx, tt.arg.req)
            if !assert.Equal(t, tt.wantResp, gotResp) {
                fmt.Printf("Expected: %+v\n", tt.wantResp)
                fmt.Printf("Actual: %+v\n", gotResp)
            }
            assert.Equal(t, tt.wantErr, err)
        })
    }
}
func TestUpdateItinerary(t *testing.T) {
    ctx, mockDal, mockService, itineraryService := ConfigItinerary(t)
    type arg struct {
        req *models.UpdateItineraryReq
        ctx context.Context
    }
    db.GetMemoMongo().Collection(constant.ItineraryTable).DeleteMany(ctx, bson.M{})
    ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
    newRating, _ := ratingDal.CreateRating(ctx, &models.Rating{Score: 5})
	itineraryDal := &dal.ItineraryDal{MainDB: db.GetMemoMongo()}
    initialItinerary, _ := itineraryDal.CreateItinerary(ctx, &models.Itinerary{
        CopiedId:    "test",
        CopiedName:  "test",
        Name:        "test",
        UserId:      "test",
        StartTime:   time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
        EndTime:     time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
        EventIds:    []string{"test"},
        EventCount:  1,
        RatingId:    newRating.Id,
    })
    tests := []struct {
        name     string
        arg      arg
        before   func(t *testing.T)
        wantResp *models.UpdateItineraryResp
        wantErr  error
    }{
        {
            name: "success",
            arg: arg{
                ctx: ctx,
                req: &models.UpdateItineraryReq{
                    Id:        initialItinerary.Id,
                    CopiedId:  initialItinerary.CopiedId,
                    CopiedName: initialItinerary.CopiedName,
                    Name: initialItinerary.Name,
                    UserId:    initialItinerary.UserId,
                    StartTime:   initialItinerary.StartTime,
                    EndTime:      initialItinerary.EndTime,
                    EventIds:  initialItinerary.EventIds,
                    EventCount: 2,
                    RatingId:  initialItinerary.RatingId,
                },
            },
            before: func(t *testing.T) {
                mockDal.EXPECT().UpdateItinerary(ctx, gomock.Any()).Return(
                    &models.Itinerary{
                        Id:        initialItinerary.Id,
						CopiedId:  initialItinerary.CopiedId,
						CopiedName: initialItinerary.CopiedName,
						Name: initialItinerary.Name,
						UserId:    initialItinerary.UserId,
						StartTime:   initialItinerary.StartTime,
						EndTime:      initialItinerary.EndTime,
						EventIds:  initialItinerary.EventIds,
						EventCount: 2,
						RatingId:  initialItinerary.RatingId,
                    },
                    nil,
                ).MaxTimes(100)
                mockDal.EXPECT().GetDB().Return(db.GetMemoMongo()).MaxTimes(100)
                mockService.EXPECT().ConvertDBOToDTOItinerary(ctx, &models.Itinerary{
                    Id:        initialItinerary.Id,
                    CopiedId:  initialItinerary.CopiedId,
                    CopiedName: initialItinerary.CopiedName,
                    Name: initialItinerary.Name,
                    UserId:    initialItinerary.UserId,
                    StartTime:   initialItinerary.StartTime,
                    EndTime:      initialItinerary.EndTime,
                    EventIds:  initialItinerary.EventIds,
                    EventCount: 2,
                    RatingId:  initialItinerary.RatingId,
                }).Return(
                    &models.ItineraryDTO{
                        Id:        initialItinerary.Id,
						CopiedId:  initialItinerary.CopiedId,
						CopiedName: initialItinerary.CopiedName,
						Name: initialItinerary.Name,
						UserId:    initialItinerary.UserId,
						StartTime:   initialItinerary.StartTime,
						EndTime:      initialItinerary.EndTime,
						EventIds:  initialItinerary.EventIds,
						EventCount: 2,
                        Rating:     &models.RatingDTO{Id: initialItinerary.RatingId, Score: 5},
                    },
                    nil,
                ).MaxTimes(100)
            },
            wantResp: &models.UpdateItineraryResp{
                Itinerary: &models.ItineraryDTO{
                    Id:        initialItinerary.Id,
					CopiedId:  initialItinerary.CopiedId,
					CopiedName: initialItinerary.CopiedName,
					Name: initialItinerary.Name,
					UserId:    initialItinerary.UserId,
					StartTime:   initialItinerary.StartTime,
					EndTime:      initialItinerary.EndTime,
					EventIds:  initialItinerary.EventIds,
					EventCount: 2,
					Rating:     &models.RatingDTO{Id: initialItinerary.RatingId, Score: 5},
                },
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.before(t)
            gotResp, err := itineraryService.UpdateItinerary(tt.arg.ctx, tt.arg.req)
            if !assert.Equal(t, tt.wantResp, gotResp) {
                fmt.Printf("Expected: %+v\n", tt.wantResp)
                fmt.Printf("Actual: %+v\n", gotResp)
            }
            assert.Equal(t, tt.wantErr, err)
        })
    }
}
func TestDeleteItinerary(t *testing.T) {
    ctx, mockDal, mockService, itineraryService := ConfigItinerary(t)
    type arg struct {
        req *models.DeleteItineraryReq
        ctx context.Context
    }
    db.GetMemoMongo().Collection(constant.ItineraryTable).DeleteMany(ctx, bson.M{})
    ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
	itineraryDal := &dal.ItineraryDal{MainDB: db.GetMemoMongo()}
    newRating, _ := ratingDal.CreateRating(ctx, &models.Rating{Score: 5})
    toDeleteItinerary, _ := itineraryDal.CreateItinerary(ctx, &models.Itinerary{
        CopiedId:    "test",
        CopiedName:  "test",
        Name:        "test",
        UserId:      "test",
        StartTime:   time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
        EndTime:     time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
        EventIds:    []string{"test"},
        EventCount:  1,
        RatingId:    newRating.Id,
    })
    tests := []struct {
        name     string
        arg      arg
        before   func(t *testing.T)
        wantResp *models.DeleteItineraryResp
        wantErr  error
    }{
        {
            name: "success",
            arg: arg{
                ctx: ctx,
                req: &models.DeleteItineraryReq{
                    Id: toDeleteItinerary.Id,
                },
            },
            before: func(t *testing.T) {
                mockDal.EXPECT().DeleteItinerary(ctx, gomock.Any()).Return(
                    &models.Itinerary{
                        Id:        toDeleteItinerary.Id,
                        CopiedId:  toDeleteItinerary.CopiedId,
                        CopiedName: toDeleteItinerary.CopiedName,
                        Name: toDeleteItinerary.Name,
                        UserId:    toDeleteItinerary.UserId,
                        StartTime:    toDeleteItinerary.StartTime,
                        EndTime:      toDeleteItinerary.EndTime,
                        EventIds:  toDeleteItinerary.EventIds,
                        EventCount: toDeleteItinerary.EventCount,
                        RatingId:  newRating.Id,
                    },
                    nil,
                ).MaxTimes(100)
                mockDal.EXPECT().GetDB().Return(db.GetMemoMongo()).MaxTimes(100)
                mockService.EXPECT().ConvertDBOToDTOItinerary(ctx, &models.Itinerary{
                    Id:        toDeleteItinerary.Id,
					CopiedId:  toDeleteItinerary.CopiedId,
					CopiedName: toDeleteItinerary.CopiedName,
					Name: toDeleteItinerary.Name,
					UserId:    toDeleteItinerary.UserId,
					StartTime:    toDeleteItinerary.StartTime,
					EndTime:      toDeleteItinerary.EndTime,
					EventIds:  toDeleteItinerary.EventIds,
					EventCount: toDeleteItinerary.EventCount,
					RatingId:  newRating.Id,
                }).Return(
                    &models.ItineraryDTO{
                        Id:        toDeleteItinerary.Id,
                        CopiedId:  toDeleteItinerary.CopiedId,
                        CopiedName: toDeleteItinerary.CopiedName,
                        Name: toDeleteItinerary.Name,
                        UserId:    toDeleteItinerary.UserId,
                        StartTime:    toDeleteItinerary.StartTime,
                        EndTime:      toDeleteItinerary.EndTime,
                        EventIds:  toDeleteItinerary.EventIds,
                        EventCount: toDeleteItinerary.EventCount,
                        Rating:      &models.RatingDTO{Id: newRating.Id, Score: 5},
                    },
                    nil,
                ).MaxTimes(100)
            },
            wantResp: &models.DeleteItineraryResp{
                Itinerary: &models.ItineraryDTO{
                    Id:        toDeleteItinerary.Id,
					CopiedId:  toDeleteItinerary.CopiedId,
					CopiedName: toDeleteItinerary.CopiedName,
					Name: toDeleteItinerary.Name,
					UserId:    toDeleteItinerary.UserId,
					StartTime:    toDeleteItinerary.StartTime,
					EndTime:      toDeleteItinerary.EndTime,
					EventIds:  toDeleteItinerary.EventIds,
					EventCount: toDeleteItinerary.EventCount,
                    Rating:       &models.RatingDTO{Id: newRating.Id, Score: 5},
                },
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.before(t)
            gotResp, err := itineraryService.DeleteItinerary(tt.arg.ctx, tt.arg.req)
            if !assert.Equal(t, tt.wantResp, gotResp) {
                fmt.Printf("Expected: %+v\n", tt.wantResp)
                fmt.Printf("Actual: %+v\n", gotResp)
            }
            assert.Equal(t, tt.wantErr, err)
        })
    }
}