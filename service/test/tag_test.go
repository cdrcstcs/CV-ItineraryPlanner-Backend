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
					Value: "test",
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().CreateTag(
					ctx,
					gomock.Any(),
				).Return(
					&models.Tag{
						Id:    "test",
						Value: "test",
					},
					nil,
				)
			},
			wantResp: &models.CreateTagResp{
				Tag: &models.TagDTO{
					Id:    "test",
					Value: "test",
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
					Id: "test",
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().GetTagById(
					ctx,
					gomock.Any(),
				).Return(
					&models.Tag{
						Id:    "test",
						Value: "test",
					},
					nil,
				)
			},
			wantResp: &models.GetTagByIdResp{
				Tag: &models.TagDTO{
					Id:    "test",
					Value: "test",
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
							Id:    "test",
							Value: "test",
						},
					},
					nil,
				)
			},
			wantResp: &models.GetTagResp{
				Tags: []*models.TagDTO{
					{
						Id:    "test",
						Value: "test",
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
					Value: "test",
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().UpdateTag(
					ctx,
					gomock.Any(),
				).Return(
					&models.Tag{
						Id:    "test",
						Value: "test",
					},
					nil,
				)
			},
			wantResp: &models.UpdateTagResp{
				Tag: &models.TagDTO{
					Id:    "test",
					Value: "test",
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
					Id: "test",
				},
			},
			before: func(t *testing.T) {
				mock.EXPECT().DeleteTag(
					ctx,
					gomock.Any(),
				).Return(
					&models.Tag{
						Id:    "test",
						Value: "test",
					},
					nil,
				)
			},
			wantResp: &models.DeleteTagResp{
				Tag: &models.TagDTO{
					Id:    "test",
					Value: "test",
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