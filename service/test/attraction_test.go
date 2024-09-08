package test
import(
	"context"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"itineraryplanner/dal/mock"
	"itineraryplanner/models"
	"itineraryplanner/service"
    "itineraryplanner/dal/db"
	mockser "itineraryplanner/service/mock"
    "itineraryplanner/dal"
    "fmt"
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
        Value: "Test",
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
                    City:          "Test" ,
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
                        City:          "Test" , 
                    },
                    nil,
                ).Times(100)
                mockDal.EXPECT().GetDB().Return(db.GetMemoMongo()).Times(100)
                mockService.EXPECT().ConvertDBOToDTOAttraction(ctx, &models.Attraction{
                    Id:            "test_id",
                    Name:          "test_name",
                    Address:       "test_address",
                    X:             1,
                    Y:             2,
                    TagIDs:        []string{newTag.Id},
                    RatingId:      newRating.Id,
                    City:          "Test" ,
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
                                Value: "Test",
                            },
                        },
                        Rating: &models.RatingDTO{
                            Id:    newRating.Id,
                            Score: 5,
                        },
                        City: "Test",
                    },
                    nil,
                ).Times(100)
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
                            Value: "Test",
                        },
                    },
                    Rating: &models.RatingDTO{
                        Id:    newRating.Id,
                        Score: 5,
                    },
                    City: "Test",
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
// func TestGetAttractionById(t *testing.T) {
// 	ctx, mock,  attractionService := ConfigAttraction(t)

// 	type arg struct {
// 		req *models.GetAttractionByIdReq
// 		ctx context.Context
// 	}

// 	tests := []struct {
// 		name string
// 		arg
// 		before   func(t *testing.T)
// 		wantResp *models.GetAttractionByIdResp
// 		wantErr  error
// 	}{
// 		{
// 			name: "success",
// 			arg: arg{
// 				ctx: ctx,
// 				req: &models.GetAttractionByIdReq{
// 					Id: "66423b801ce396edb58797e2",
// 				},
// 			},
// 			before: func(t *testing.T) {
// 				mock.EXPECT().GetAttractionById(
// 					ctx,
// 					gomock.Any(),
// 				).Return(
// 					&models.Attraction{
// 						Id:           "66423b801ce396edb58797e2",
// 						Name:         "test",
// 						Address:      "123 Paya Lebar",
// 						X:1,
// 						Y:2,
// 						TagIDs:       []string{"6641b096e8ee7b68f952efa8"},
// 						RatingId: "6642102153edf9e8fb3dfcda",
// 					},
// 					nil,
// 				)
// 			},
// 			wantResp: &models.GetAttractionByIdResp{Attraction: &models.AttractionDTO{
// 				Id:           "66423b801ce396edb58797e2",
// 				Name:         "test",
// 				Address:      "123 Paya Lebar",
// 				X:1,
// 				Y:2,
// 				Tags: []*models.TagDTO{
// 					{
// 						Id:    "6641b096e8ee7b68f952efa8",
// 						Value: "test",
// 					},
// 				},
// 				Rating: &models.RatingDTO{
// 					Id: "6642102153edf9e8fb3dfcda",
// 					Score: 5,
// 				},
// 			}},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.before(t)
// 			gotResp, err := attractionService.GetAttractionById(ctx, tt.arg.req)

// 			assert.Equal(t, gotResp, tt.wantResp)
// 			assert.Equal(t, err, tt.wantErr)
// 		})
// 	}
// }

// func TestGetAttraction(t *testing.T) {
// 	ctx, mock, attractionService := ConfigAttraction(t)

// 	type arg struct {
// 		req *models.GetAttractionReq
// 		ctx context.Context
// 	}

// 	tests := []struct {
// 		name string
// 		arg
// 		before   func(t *testing.T)
// 		wantResp *models.GetAttractionResp
// 		wantErr  error
// 	}{
// 		{
// 			name: "success",
// 			arg: arg{
// 				ctx: ctx,
// 				req: &models.GetAttractionReq{},
// 			},
// 			before: func(t *testing.T) {
// 				mock.EXPECT().GetAttraction(ctx).Return(
// 					[]*models.Attraction{
// 						{
// 							Id:           "66423b801ce396edb58797e2",
// 							Name:         "test",
// 							Address:      "123 Paya Lebar",
// 							X:1,
// 							Y:2,
// 							TagIDs:       []string{"6641b096e8ee7b68f952efa8"},
// 							RatingId: "6642102153edf9e8fb3dfcda",
// 						},
// 					},
// 					nil,
// 				)
// 			},
// 			wantResp: &models.GetAttractionResp{Attractions: []*models.AttractionDTO{
// 				{
// 					Id:           "66423b801ce396edb58797e2",
// 					Name:         "test",
// 					Address:      "123 Paya Lebar",
// 					X:1,
// 					Y:2,
// 					Tags: []*models.TagDTO{
// 						{
// 							Id:    "6641b096e8ee7b68f952efa8",
// 							Value: "test",
// 						},
// 					},
// 					Rating: &models.RatingDTO{
// 						Id: "6642102153edf9e8fb3dfcda",
// 						Score: 5,
// 					},
// 				},
// 			}},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.before(t)
// 			gotResp, err := attractionService.GetAttraction(ctx, tt.arg.req)

// 			assert.Equal(t, gotResp, tt.wantResp)
// 			assert.Equal(t, err, tt.wantErr)
// 		})
// 	}
// }

// func TestUpdateAttraction(t *testing.T) {
// 	ctx, mock, attractionService := ConfigAttraction(t)

// 	type arg struct {
// 		req *models.UpdateAttractionReq
// 		ctx context.Context
// 	}

// 	tests := []struct {
// 		name string
// 		arg
// 		before   func(t *testing.T)
// 		wantResp *models.UpdateAttractionResp
// 		wantErr  error
// 	}{
// 		{
// 			name: "success",
// 			arg: arg{
// 				ctx: ctx,
// 				req: &models.UpdateAttractionReq{
// 					Id:           "66423b801ce396edb58797e2",
// 					Name:         "test2",
// 					Address:      "123 Paya Lebar",
// 					X:1,
// 					Y:2,
// 					TagIDs:       []string{"6641b096e8ee7b68f952efa8"},
// 					RatingId: "6642102153edf9e8fb3dfcda",
// 				},
// 			},
// 			before: func(t *testing.T) {
// 				mock.EXPECT().UpdateAttraction(ctx, gomock.Any()).Return(
// 					&models.Attraction{
// 						Id:           "66423b801ce396edb58797e2",
// 						Name:         "test2",
// 						Address:      "123 Paya Lebar",
// 						X:1,
// 						Y:2,
// 						TagIDs:       []string{"6641b096e8ee7b68f952efa8"},
// 						RatingId: "6642102153edf9e8fb3dfcda",
// 					},
// 					nil,
// 				)
// 			},
// 			wantResp: &models.UpdateAttractionResp{Attraction: &models.AttractionDTO{
// 				Id:           "66423b801ce396edb58797e2",
// 				Name:         "test2",
// 				Address:      "123 Paya Lebar",
// 				X:1,
// 				Y:2,
// 				Tags: []*models.TagDTO{
// 					{
// 						Id:    "6641b096e8ee7b68f952efa8",
// 						Value: "test",
// 					},
// 				},
// 				Rating: &models.RatingDTO{
// 					Id: "6642102153edf9e8fb3dfcda",
// 					Score: 5,
// 				},
// 			}},
// 			wantErr: nil,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.before(t)
// 			gotResp, err := attractionService.UpdateAttraction(ctx, tt.arg.req)
// 			assert.Equal(t, gotResp, tt.wantResp)
// 			assert.Equal(t, err, tt.wantErr)
// 		})
// 	}
// }

// func TestDeleteAttraction(t *testing.T) {
// 	ctx, mock, attractionService := ConfigAttraction(t)

// 	type arg struct {
// 		req *models.DeleteAttractionReq
// 		ctx context.Context
// 	}

// 	tests := []struct {
// 		name string
// 		arg
// 		before   func(t *testing.T)
// 		wantResp *models.DeleteAttractionResp
// 		wantErr  error
// 	}{
// 		{
// 			name: "success",
// 			arg: arg{
// 				ctx: ctx,
// 				req: &models.DeleteAttractionReq{
// 					Id:           "66423b801ce396edb58797e2",
// 				},
// 			},
// 			before: func(t *testing.T) {
// 				mock.EXPECT().DeleteAttraction(
// 					ctx,
// 					gomock.Any(),
// 				).Return(
// 					&models.Attraction{
// 						Id:           "66423b801ce396edb58797e2",
// 						Name:         "test2",
// 						Address:      "123 Paya Lebar",
// 						X:1,
// 						Y:2,
// 						TagIDs:       []string{"6641b096e8ee7b68f952efa8"},
// 						RatingId: "6642102153edf9e8fb3dfcda",
// 					},
// 					nil,
// 				)
// 			},
// 			wantResp: &models.DeleteAttractionResp{
// 				Attraction: &models.AttractionDTO{
// 					Id:           "66423b801ce396edb58797e2",
// 					Name:         "test2",
// 					Address:      "123 Paya Lebar",
// 					X:1,
// 					Y:2,
// 					Tags: []*models.TagDTO{
// 						{
// 							Id:    "6641b096e8ee7b68f952efa8",
// 							Value: "test",
// 						},
// 					},
// 					Rating: &models.RatingDTO{
// 						Id: "6642102153edf9e8fb3dfcda",
// 						Score: 5,
// 					},
// 				},
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.before(t)
// 			gotResp, err := attractionService.DeleteAttraction(ctx, tt.arg.req)

// 			assert.Equal(t, gotResp, tt.wantResp)
// 			assert.Equal(t, err, tt.wantErr)
// 		})
// 	}
// }

// func TestConvertAttToDTOAttraction(t *testing.T) {
// 	ctx:= context.Background()
// 	attractionService := &service.AttractionService{}
// 	type arg struct {
// 		req *models.Attraction
// 		ctx context.Context
// 	}

// 	tests := []struct {
// 		name     string
// 		arg      arg
// 		wantResp *models.AttractionDTO
// 		wantErr  error
// 	}{
// 		{
// 			name: "success",
// 			arg: arg{
// 				req: &models.Attraction{
// 					Id:           "66423b801ce396edb58797e2",
// 					Name:         "test2",
// 					Address:      "123 Paya Lebar",
// 					X:1,
// 					Y:2,
// 					TagIDs:       []string{"6641b096e8ee7b68f952efa8"},
// 					RatingId: "6642102153edf9e8fb3dfcda",
// 				},
// 				ctx: ctx,
// 			},
// 			wantResp: &models.AttractionDTO{
// 				Id:           "66423b801ce396edb58797e2",
// 				Name:         "test2",
// 				Address:      "123 Paya Lebar",
// 				X:1,
// 				Y:2,
// 				Tags: []*models.TagDTO{
// 					{
// 						Id:    "6641b096e8ee7b68f952efa8",
// 						Value: "test",
// 					},
// 				},
// 				Rating: &models.RatingDTO{
// 					Id: "6642102153edf9e8fb3dfcda",
// 					Score: 5,
// 				},
// 			},
// 			wantErr: nil,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gotResp, err := attractionService.ConvertDBOToDTOAttraction(tt.arg.ctx, tt.arg.req)

// 			assert.Equal(t, tt.wantResp, gotResp)
// 			assert.Equal(t, tt.wantErr, err)
// 		})
// 	}
// }
