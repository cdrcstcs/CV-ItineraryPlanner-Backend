package test

import (
	"context"
	"fmt"
	"itineraryplanner/constant"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/dal/mock"
	"itineraryplanner/models"
	"itineraryplanner/service"
	mockser "itineraryplanner/service/mock"
	"testing"
    "go.mongodb.org/mongo-driver/bson"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)
func ConfigAttraction(t *testing.T) (context.Context, *mock.MockAttractionDal, *mockser.MockAttractionService, *service.AttractionService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockADal := mock.NewMockAttractionDal(ctrl)
	mockASer := mockser.NewMockAttractionService(ctrl)
	attractionService := &service.AttractionService{Dal: mockADal}
	return ctx, mockADal, mockASer, attractionService
}
func TestCreateAttraction(t *testing.T) {
    ctx, mockDal, mockService, attractionService := ConfigAttraction(t)
    type arg struct {
        req *models.CreateAttractionReq
        ctx context.Context
    }
    tagDal := &dal.TagDal{MainDB: db.GetMemoMongo()}
    ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
    newTag, _ := tagDal.CreateTag(ctx, &models.Tag{
        Value: "test",
    })
    newRating, _ := ratingDal.CreateRating(ctx, &models.Rating{
        Score: 5,
    })
    tests := []struct {
        name string
        arg
        before   func(t *testing.T)
        wantResp *models.CreateAttractionResp
        wantErr  error
    }{
        {
            name: "success",
            arg: arg{
                ctx: ctx,
                req: &models.CreateAttractionReq{
                    Name:          "test_name",
                    Address:       "test_address",
                    X:             1,
                    Y:             2,
                    TagIDs:        []string{newTag.Id},
                    RatingId:      newRating.Id,
                    City:          "test" ,
                },
            },
            before: func(t *testing.T) {
                mockDal.EXPECT().CreateAttraction(ctx, gomock.Any()).Return(
                    &models.Attraction{
                        Id:            "test_id",
                        Name:          "test_name",
                        Address:       "test_address",
                        X:             1,
                        Y:             2,
                        TagIDs:        []string{newTag.Id},
                        RatingId:      newRating.Id,
                        City:          "test" , 
                    },
                    nil,
                ).MaxTimes(100)
                mockDal.EXPECT().GetDB().Return(db.GetMemoMongo()).MaxTimes(100)
                mockService.EXPECT().ConvertDBOToDTOAttraction(ctx, &models.Attraction{
                    Id:            "test_id",
                    Name:          "test_name",
                    Address:       "test_address",
                    X:             1,
                    Y:             2,
                    TagIDs:        []string{newTag.Id},
                    RatingId:      newRating.Id,
                    City:          "test" ,
                }).Return(
                    &models.AttractionDTO{
                        Id:      "test_id",
                        Name:    "test_name",
                        Address: "test_address",
                        X:       1,
                        Y:       2,
                        Tags: []*models.TagDTO{
                            {
                                Id:    newTag.Id,
                                Value: "test",
                            },
                        },
                        Rating: &models.RatingDTO{
                            Id:    newRating.Id,
                            Score: 5,
                        },
                        City: "test",
                    },
                    nil,
                ).MaxTimes(100)
            },
            wantResp: &models.CreateAttractionResp{
                Attraction: &models.AttractionDTO{
                    Id:      "test_id",
                    Name:    "test_name",
                    Address: "test_address",
                    X:       1,
                    Y:       2,
                    Tags: []*models.TagDTO{
                        {
                            Id:    newTag.Id,
                            Value: "test",
                        },
                    },
                    Rating: &models.RatingDTO{
                        Id:    newRating.Id,
                        Score: 5,
                    },
                    City: "test",
                },
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.before(t)
            gotResp, err := attractionService.CreateAttraction(ctx, tt.arg.req)
            assert.Equal(t, tt.wantResp, gotResp)
            fmt.Printf("Expected: %+v\n", tt.wantResp)
            fmt.Printf("Actual: %+v\n", gotResp)
            assert.Equal(t, tt.wantErr, err)
        })
    }
}
func TestGetAttractionById(t *testing.T) {
    ctx, mockDal, mockService, attractionService := ConfigAttraction(t)
    type arg struct {
        req *models.GetAttractionByIdReq
        ctx context.Context
    }
    tagDal := &dal.TagDal{MainDB: db.GetMemoMongo()}
    ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
    attractionDal := &dal.AttractionDal{MainDB: db.GetMemoMongo()}
    newTag, _ := tagDal.CreateTag(ctx, &models.Tag{
        Value: "test",
    })
    newRating, _ := ratingDal.CreateRating(ctx, &models.Rating{
        Score: 5,
    })
    newAttraction, _ := attractionDal.CreateAttraction(ctx, &models.Attraction{
        Name:         "test",
        Address:      "123 Paya Lebar",
        X:            1,
        Y:            2,
        TagIDs:       []string{newTag.Id},
        RatingId:     newRating.Id,
        City:         "test",
    })
    tests := []struct {
        name string
        arg
        before   func(t *testing.T)
        wantResp *models.GetAttractionByIdResp
        wantErr  error
    }{
        {
            name: "success",
            arg: arg{
                ctx: ctx,
                req: &models.GetAttractionByIdReq{
                    Id: newAttraction.Id,
                },
            },
            before: func(t *testing.T) {
                mockDal.EXPECT().GetAttractionById(ctx, gomock.Any()).Return(
                    newAttraction,
                    nil,
                ).MaxTimes(100) 
                mockDal.EXPECT().GetDB().Return(db.GetMemoMongo()).MaxTimes(100)
                mockService.EXPECT().ConvertDBOToDTOAttraction(ctx, newAttraction).Return(
                    &models.AttractionDTO{
                        Id:           newAttraction.Id,
                        Name:         "test",
                        Address:      "123 Paya Lebar",
                        X:            1,
                        Y:            2,
                        Tags: []*models.TagDTO{
                            {
                                Id:    newTag.Id,
                                Value: "test",
                            },
                        },
                        Rating: &models.RatingDTO{
                            Id:    newRating.Id,
                            Score: 5,
                        },
                        City: "test",
                    },
                    nil,
                ).MaxTimes(100) 
            },
            wantResp: &models.GetAttractionByIdResp{
                Attraction: &models.AttractionDTO{
                    Id:            newAttraction.Id,
                    Name:         "test",
                    Address:      "123 Paya Lebar",
                    X:            1,
                    Y:            2,
                    Tags: []*models.TagDTO{
                        {
                            Id:    newTag.Id,
                            Value: "test",
                        },
                    },
                    Rating: &models.RatingDTO{
                        Id:    newRating.Id,
                        Score: 5,
                    },
                    City: "test",
                },
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.before(t)
            gotResp, err := attractionService.GetAttractionById(tt.arg.ctx, tt.arg.req)
            if !assert.Equal(t, tt.wantResp, gotResp) {
                fmt.Printf("Expected: %+v\n", tt.wantResp)
                fmt.Printf("Actual: %+v\n", gotResp)
            }
            assert.Equal(t, tt.wantErr, err)
        })
    }
}
func TestGetAttraction(t *testing.T) {
    ctx, mockDal, mockService, attractionService := ConfigAttraction(t)
    type arg struct {
        req *models.GetAttractionReq
        ctx context.Context
    }
    db.GetMemoMongo().Collection(constant.AttractionTable).DeleteMany(ctx, bson.M{})
    tagDal := &dal.TagDal{MainDB: db.GetMemoMongo()}
    ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
    attractionDal := &dal.AttractionDal{MainDB: db.GetMemoMongo()}
    newTag, _ := tagDal.CreateTag(ctx, &models.Tag{
        Value: "test",
    })
    newRating, _ := ratingDal.CreateRating(ctx, &models.Rating{
        Score: 5,
    })
    newAttraction, _ := attractionDal.CreateAttraction(ctx, &models.Attraction{
        Name:         "test",
        Address:      "123 Paya Lebar",
        X:            1,
        Y:            2,
        TagIDs:       []string{newTag.Id},
        RatingId:     newRating.Id,
        City:         "test",
    })
    tests := []struct {
        name string
        arg
        before   func(t *testing.T)
        wantResp *models.GetAttractionResp
        wantErr  error
    }{
        {
            name: "success",
            arg: arg{
                ctx: ctx,
                req: &models.GetAttractionReq{},
            },
            before: func(t *testing.T) {
                mockDal.EXPECT().GetAttraction(ctx).Return(
                    []*models.Attraction{
                        {
                            Id:           newAttraction.Id,
                            Name:         "test",
                            Address:      "123 Paya Lebar",
                            X:            1,
                            Y:            2,
                            TagIDs:       []string{newTag.Id},
                            RatingId:     newRating.Id,
                            City:         "test",
                        },
                    },
                    nil,
                ).MaxTimes(100)
                mockDal.EXPECT().GetDB().Return(db.GetMemoMongo()).MaxTimes(100)
                mockService.EXPECT().ConvertDBOToDTOAttraction(ctx, &models.Attraction{
                    Id:           newAttraction.Id,
                    Name:         "test",
                    Address:      "123 Paya Lebar",
                    X:            1,
                    Y:            2,
                    TagIDs:       []string{newTag.Id},
                    RatingId:     newRating.Id,
                    City:         "test",
                }).Return(
                    &models.AttractionDTO{
                        Id:           newAttraction.Id,
                        Name:         "test",
                        Address:      "123 Paya Lebar",
                        X:            1,
                        Y:            2,
                        Tags: []*models.TagDTO{
                            {
                                Id:    newTag.Id,
                                Value: "test",
                            },
                        },
                        Rating: &models.RatingDTO{
                            Id:    newRating.Id,
                            Score: 5,
                        },
                        City: "test",
                    },
                    nil,
                ).MaxTimes(100) 
            },
            wantResp: &models.GetAttractionResp{
                Attractions: []*models.AttractionDTO{
                    {
                        Id:            newAttraction.Id,
                        Name:         "test",
                        Address:      "123 Paya Lebar",
                        X:            1,
                        Y:            2,
                        Tags: []*models.TagDTO{
                            {
                                Id:    newTag.Id,
                                Value: "test",
                            },
                        },
                        Rating: &models.RatingDTO{
                            Id:    newRating.Id,
                            Score: 5,
                        },
                        City: "test",
                    },
                },
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.before(t)
            gotResp, err := attractionService.GetAttraction(tt.arg.ctx, tt.arg.req)
            if !assert.Equal(t, tt.wantResp, gotResp) {
                fmt.Printf("Expected: %+v\n", tt.wantResp)
                fmt.Printf("Actual: %+v\n", gotResp)
            }
            assert.Equal(t, tt.wantErr, err)
        })
    }
}
func TestUpdateAttraction(t *testing.T) {
    ctx, mockDal, mockService, attractionService := ConfigAttraction(t)
    type arg struct {
        req *models.UpdateAttractionReq
        ctx context.Context
    }
    db.GetMemoMongo().Collection(constant.AttractionTable).DeleteMany(ctx, bson.M{})
    tagDal := &dal.TagDal{MainDB: db.GetMemoMongo()}
    ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
    attractionDal := &dal.AttractionDal{MainDB: db.GetMemoMongo()}
    newTag, _ := tagDal.CreateTag(ctx, &models.Tag{Value: "test"})
    newRating, _ := ratingDal.CreateRating(ctx, &models.Rating{Score: 5})
    initialAttraction, _ := attractionDal.CreateAttraction(ctx, &models.Attraction{
        Name:         "test",
        Address:      "123 Paya Lebar",
        X:            1,
        Y:            2,
        TagIDs:       []string{newTag.Id},
        RatingId:     newRating.Id,
        City:         "test",
    })
    tests := []struct {
        name string
        arg
        before   func(t *testing.T)
        wantResp *models.UpdateAttractionResp
        wantErr  error
    }{
        {
            name: "success",
            arg: arg{
                ctx: ctx,
                req: &models.UpdateAttractionReq{
                    Id:           initialAttraction.Id,
                    Name:         "test2",
                    Address:      "123 Paya Lebar",
                    X:            1,
                    Y:            2,
                    TagIDs:       []string{newTag.Id},
                    RatingId:     newRating.Id,
                },
            },
            before: func(t *testing.T) {
                mockDal.EXPECT().UpdateAttraction(ctx, gomock.Any()).Return(
                    &models.Attraction{
                        Id:           initialAttraction.Id,
                        Name:         "test2",
                        Address:      "123 Paya Lebar",
                        X:            1,
                        Y:            2,
                        TagIDs:       []string{newTag.Id},
                        RatingId:     newRating.Id,
                        City:         "test",
                    },
                    nil,
                ).MaxTimes(100)
                mockDal.EXPECT().GetDB().Return(db.GetMemoMongo()).MaxTimes(100)
                mockService.EXPECT().ConvertDBOToDTOAttraction(ctx, &models.Attraction{
                    Id:           initialAttraction.Id,
                    Name:         "test2",
                    Address:      "123 Paya Lebar",
                    X:            1,
                    Y:            2,
                    TagIDs:       []string{newTag.Id},
                    RatingId:     newRating.Id,
                    City:         "test",
                }).Return(
                    &models.AttractionDTO{
                        Id:           initialAttraction.Id,
                        Name:         "test2",
                        Address:      "123 Paya Lebar",
                        X:            1,
                        Y:            2,
                        Tags: []*models.TagDTO{
                            {
                                Id:    newTag.Id,
                                Value: "test",
                            },
                        },
                        Rating: &models.RatingDTO{
                            Id:    newRating.Id,
                            Score: 5,
                        },
                        City: "test",
                    },
                    nil,
                ).MaxTimes(100)
            },
            wantResp: &models.UpdateAttractionResp{
                Attraction: &models.AttractionDTO{
                    Id:           initialAttraction.Id,
                    Name:         "test2",
                    Address:      "123 Paya Lebar",
                    X:            1,
                    Y:            2,
                    Tags: []*models.TagDTO{
                        {
                            Id:    newTag.Id,
                            Value: "test",
                        },
                    },
                    Rating: &models.RatingDTO{
                        Id:    newRating.Id,
                        Score: 5,
                    },
                    City: "test",
                },
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.before(t)
            gotResp, err := attractionService.UpdateAttraction(tt.arg.ctx, tt.arg.req)
            if !assert.Equal(t, tt.wantResp, gotResp) {
                fmt.Printf("Expected: %+v\n", tt.wantResp)
                fmt.Printf("Actual: %+v\n", gotResp)
            }
            assert.Equal(t, tt.wantErr, err)
        })
    }
}
func TestDeleteAttraction(t *testing.T) {
    ctx, mockDal, mockService, attractionService := ConfigAttraction(t)
    type arg struct {
        req *models.DeleteAttractionReq
        ctx context.Context
    }
    db.GetMemoMongo().Collection(constant.AttractionTable).DeleteMany(ctx, bson.M{})
    tagDal := &dal.TagDal{MainDB: db.GetMemoMongo()}
    ratingDal := &dal.RatingDal{MainDB: db.GetMemoMongo()}
    attractionDal := &dal.AttractionDal{MainDB: db.GetMemoMongo()}
    newTag, _ := tagDal.CreateTag(ctx, &models.Tag{Value: "test"})
    newRating, _ := ratingDal.CreateRating(ctx, &models.Rating{Score: 5})
    attraction, _ := attractionDal.CreateAttraction(ctx, &models.Attraction{
        Name:         "test2",
        Address:      "123 Paya Lebar",
        X:            1,
        Y:            2,
        TagIDs:       []string{newTag.Id},
        RatingId:     newRating.Id,
        City:         "test",
    })
    tests := []struct {
        name string
        arg
        before   func(t *testing.T)
        wantResp *models.DeleteAttractionResp
        wantErr  error
    }{
        {
            name: "success",
            arg: arg{
                ctx: ctx,
                req: &models.DeleteAttractionReq{
                    Id: attraction.Id,
                },
            },
            before: func(t *testing.T) {
                mockDal.EXPECT().DeleteAttraction(ctx, gomock.Any()).Return(
                    &models.Attraction{
                        Id:           attraction.Id,
                        Name:         "test2",
                        Address:      "123 Paya Lebar",
                        X:            1,
                        Y:            2,
                        TagIDs:       []string{newTag.Id},
                        RatingId:     newRating.Id,
                        City:         "test",
                    },
                    nil,
                ).MaxTimes(100)
                mockDal.EXPECT().GetDB().Return(db.GetMemoMongo()).MaxTimes(100)
                mockService.EXPECT().ConvertDBOToDTOAttraction(ctx, &models.Attraction{
                    Id:           attraction.Id,
                    Name:         "test2",
                    Address:      "123 Paya Lebar",
                    X:            1,
                    Y:            2,
                    TagIDs:       []string{newTag.Id},
                    RatingId:     newRating.Id,
                    City:         "test",
                }).Return(
                    &models.AttractionDTO{
                        Id:           attraction.Id,
                        Name:         "test2",
                        Address:      "123 Paya Lebar",
                        X:            1,
                        Y:            2,
                        Tags: []*models.TagDTO{
                            {
                                Id:    newTag.Id,
                                Value: "test",
                            },
                        },
                        Rating: &models.RatingDTO{
                            Id:    newRating.Id,
                            Score: 5,
                        },
                        City: "test",
                    },
                    nil,
                ).MaxTimes(100)
            },
            wantResp: &models.DeleteAttractionResp{
                Attraction: &models.AttractionDTO{
                    Id:           attraction.Id,
                    Name:         "test2",
                    Address:      "123 Paya Lebar",
                    X:            1,
                    Y:            2,
                    Tags: []*models.TagDTO{
                        {
                            Id:    newTag.Id,
                            Value: "test",
                        },
                    },
                    Rating: &models.RatingDTO{
                        Id:    newRating.Id,
                        Score: 5,
                    },
                    City: "test",
                },
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.before(t)
            gotResp, err := attractionService.DeleteAttraction(tt.arg.ctx, tt.arg.req)
            if !assert.Equal(t, tt.wantResp, gotResp) {
                fmt.Printf("Expected: %+v\n", tt.wantResp)
                fmt.Printf("Actual: %+v\n", gotResp)
            }
            assert.Equal(t, tt.wantErr, err)
        })
    }
}