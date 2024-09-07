package test

import (
	"context"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"itineraryplanner/dal/mock"
	"itineraryplanner/models"
	"itineraryplanner/service"
)

func ConfigTag(t *testing.T) (context.Context, *mock.MockTagDal, *service.TagService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mock := mock.NewMockTagDal(ctrl)
	tagService := &service.TagService{Dal: mock}
	return ctx, mock, tagService
}
func TestCreateTag(t *testing.T) {
	ctx, mock, tagService := ConfigTag(t)

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
				mock.EXPECT().CreateTag(
					ctx,
					gomock.Any(),
				).Return(
					&models.Tag{
						Id:    "test_tag_id",
						Value: "test_value",
					},
					nil,
				)
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
	ctx, mock, tagService := ConfigTag(t)

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
				mock.EXPECT().GetTagById(
					ctx,
					gomock.Any(),
				).Return(
					&models.Tag{
						Id:    "test_tag_id",
						Value: "test_value",
					},
					nil,
				)
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
	ctx, mock, tagService := ConfigTag(t)

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
				mock.EXPECT().GetTag(
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
	ctx, mock, tagService := ConfigTag(t)

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
				mock.EXPECT().UpdateTag(
					ctx,
					gomock.Any(),
				).Return(
					&models.Tag{
						Id:    "test_tag_id",
						Value: "test_value",
					},
					nil,
				)
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
	ctx, mock, tagService := ConfigTag(t)

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
				mock.EXPECT().DeleteTag(
					ctx,
					gomock.Any(),
				).Return(
					&models.Tag{
						Id:    "test_tag_id",
						Value: "test_value",
					},
					nil,
				)
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
	ctx:= context.Background()
	tagService:= &service.TagService{}
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