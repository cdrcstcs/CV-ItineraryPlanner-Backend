package test

import (
	"context"
	"testing"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"itineraryplanner/dal/mock"
	"itineraryplanner/models"
	"itineraryplanner/service"
	ser_mock "itineraryplanner/service/mock"
)

func ConfigCreateTag(t *testing.T) (context.Context, *mock.MockCreateTagDal, *ser_mock.MockTagDTOService, *service.TagService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockC := mock.NewMockCreateTagDal(ctrl)
	mockDTO := ser_mock.NewMockTagDTOService(ctrl)
	tagService := &service.TagService{CDal: mockC}
	return ctx, mockC, mockDTO, tagService
}

func ConfigGetTag(t *testing.T) (context.Context, *mock.MockGetTagDal, *ser_mock.MockTagDTOService, *service.TagService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockG := mock.NewMockGetTagDal(ctrl)
	mockDTO := ser_mock.NewMockTagDTOService(ctrl)
	tagService := &service.TagService{GDal: mockG}
	return ctx, mockG, mockDTO, tagService
}

func ConfigGetTagById(t *testing.T) (context.Context, *mock.MockGetTagByIdDal, *ser_mock.MockTagDTOService, *service.TagService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockB := mock.NewMockGetTagByIdDal(ctrl)
	mockDTO := ser_mock.NewMockTagDTOService(ctrl)
	tagService := &service.TagService{BDal: mockB}
	return ctx, mockB, mockDTO, tagService
}
func ConfigUpdateTag(t *testing.T) (context.Context, *mock.MockUpdateTagDal, *ser_mock.MockTagDTOService, *service.TagService){
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockU := mock.NewMockUpdateTagDal(ctrl)
	mockDTO := ser_mock.NewMockTagDTOService(ctrl)
	tagService := &service.TagService{UDal: mockU}
	return ctx, mockU, mockDTO, tagService
}
func ConfigDeleteTag(t *testing.T) (context.Context, *mock.MockDeleteTagDal, *ser_mock.MockTagDTOService, *service.TagService){
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockD := mock.NewMockDeleteTagDal(ctrl)
	mockDTO := ser_mock.NewMockTagDTOService(ctrl)
	tagService := &service.TagService{DDal: mockD}
	return ctx, mockD, mockDTO, tagService
}
func ConfigTagDTOService(t *testing.T) (context.Context, *service.TagService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockC := mock.NewMockCreateTagDal(ctrl)
	mockG := mock.NewMockGetTagDal(ctrl)
	tagDTOService := &service.TagService{GDal: mockG, CDal: mockC}
	return ctx, tagDTOService
}


func TestCreateTag(t *testing.T) {
	ctx, mockC, mockDTO, tagService := ConfigCreateTag(t)

	type arg struct {
		req *models.CreateTagReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.CreateTagResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.CreateTagReq{
					Value: "test_value",
				},
			},
			before: func(t *testing.T) {
				mockC.EXPECT().CreateTag(
					ctx,
					gomock.Any(),
				).Return(
					&models.Tag{
						Id:    "test_tag_id",
						Value: "test_value",
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOTag(
					ctx,
					&models.Tag{
						Id:    "test_tag_id",
						Value: "test_value",
					},
				).Return(
					&models.TagDTO{
						Id:    "test_tag_id",
						Value: "test_value",
					},
					nil,
				).Do(func(ctx context.Context, tag *models.Tag) {
					fmt.Println("ConvertDBOToDTOTag called with", tag)
				})
			},
			wantResp: &models.CreateTagResp{
				Tag: &models.TagDTO{
					Id:    "test_tag_id",
					Value: "test_value",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := tagService.CreateTag(ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetTagById(t *testing.T) {
	ctx, mockB, mockDTO, tagService := ConfigGetTagById(t)

	type arg struct {
		req *models.GetTagByIdReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.GetTagByIdResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.GetTagByIdReq{
					Id: "test_tag_id",
				},
			},
			before: func(t *testing.T) {
				mockB.EXPECT().GetTagById(
					ctx,
					gomock.Any(),
				).Return(
					&models.Tag{
						Id:    "test_tag_id",
						Value: "test_value",
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOTag(
					ctx,
					gomock.Any(),
				).Return(
					&models.TagDTO{
						Id:    "test_tag_id",
						Value: "test_value",
					},
					nil,
				).Do(func(ctx context.Context, tag *models.Tag) {
					fmt.Println("ConvertDBOToDTOTag called with", tag)
				})
			},
			wantResp: &models.GetTagByIdResp{
				Tag: &models.TagDTO{
					Id:    "test_tag_id",
					Value: "test_value",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := tagService.GetTagById(ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetTag(t *testing.T) {
	ctx, mockG, mockDTO, tagService := ConfigGetTag(t)

	type arg struct {
		req *models.GetTagReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.GetTagResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.GetTagReq{
				},
			},
			before: func(t *testing.T) {
				mockG.EXPECT().GetTag(
					ctx,
				).Return(
					[]*models.Tag{
						{
							Id:    "test_tag_id",
							Value: "test_value",
						},
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOTag(
					ctx,
					&models.Tag{
						Id:    "test_tag_id",
						Value: "test_value",
					},
				).Return(
					&models.TagDTO{
						Id:    "test_tag_id",
						Value: "test_value",
					},
					nil,
				).Do(func(ctx context.Context, tag *models.Tag) {
					fmt.Println("ConvertDBOToDTOTag called with", tag)
				})
			},
			wantResp: &models.GetTagResp{
				Tags: []*models.TagDTO{
					{
						Id:    "test_tag_id",
						Value: "test_value",
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := tagService.GetTag(ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestUpdateTag(t *testing.T) {
	ctx, mockU, mockDTO, tagService := ConfigUpdateTag(t)

	type arg struct {
		req *models.UpdateTagReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.UpdateTagResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.UpdateTagReq{
					Value: "test_value",
				},
			},
			before: func(t *testing.T) {
				mockU.EXPECT().UpdateTag(
					ctx,
					gomock.Any(),
				).Return(
					&models.Tag{
						Id:    "test_tag_id",
						Value: "test_value",
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOTag(
					ctx,
					&models.Tag{
						Id:    "test_tag_id",
						Value: "test_value",
					},
				).Return(
					&models.TagDTO{
						Id:    "test_tag_id",
						Value: "test_value",
					},
					nil,
				).Do(func(ctx context.Context, tag *models.Tag) {
					fmt.Println("ConvertDBOToDTOTag called with", tag)
				})
			},
			wantResp: &models.UpdateTagResp{
				Tag: &models.TagDTO{
					Id:    "test_tag_id",
					Value: "test_value",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := tagService.UpdateTag(ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestDeleteTag(t *testing.T) {
	ctx, mockD, mockDTO, tagService := ConfigDeleteTag(t)

	type arg struct {
		req *models.DeleteTagReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.DeleteTagResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.DeleteTagReq{
					Id: "test_tag_id",
				},
			},
			before: func(t *testing.T) {
				mockD.EXPECT().DeleteTag(
					ctx,
					gomock.Any(),
				).Return(
					&models.Tag{
						Id:    "test_tag_id",
						Value: "test_value",
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOTag(
					ctx,
					&models.Tag{
						Id:    "test_tag_id",
						Value: "test_value",
					},
				).Return(
					&models.TagDTO{
						Id:    "test_tag_id",
						Value: "test_value",
					},
					nil,
				).Do(func(ctx context.Context, tag *models.Tag) {
					fmt.Println("ConvertDBOToDTOTag called with", tag)
				})
			},
			wantResp: &models.DeleteTagResp{
				Tag: &models.TagDTO{
					Id:    "test_tag_id",
					Value: "test_value",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := tagService.DeleteTag(ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
func TestConvertDBOToDTOTag(t *testing.T) {
	ctx, tagService := ConfigTagDTOService(t)

	type arg struct {
		tag *models.Tag
		ctx context.Context
	}

	tests := []struct {
		name     string
		before   func(t *testing.T)
		arg      arg
		wantResp *models.TagDTO
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				tag: &models.Tag{
					Id:    "test_tag_id",
					Value: "test_value",
				},
				ctx: ctx,
			},
			wantResp: &models.TagDTO{
				Id:    "test_tag_id",
				Value: "test_value",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := tagService.ConvertDBOToDTOTag(tt.arg.ctx, tt.arg.tag)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}