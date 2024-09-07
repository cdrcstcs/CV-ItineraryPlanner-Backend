package test

import (
	"context"
	"itineraryplanner/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"itineraryplanner/dal/db"
	"itineraryplanner/dal"

	"itineraryplanner/constant"
)

func TestCreateTag(t *testing.T){
	ctx := context.Background()

	type arg struct {
		ctx context.Context
		tag *models.Tag
	}

	tests:=[]struct {
		name string 
		before func(t *testing.T)
		arg
		wantTag *models.Tag
		wantErr error
	}{
		{
			name: "success",
			arg: arg{
				ctx:ctx,
				tag: &models.Tag{
					Value: "test",
				},
			},
			wantTag: &models.Tag{
				Id:"generated_id",
				Value:"test",
			},
			wantErr: nil ,
		},
	}
	tagDal := &dal.TagDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotRating, gotErr:= tagDal.CreateTag(ctx, tt.arg.tag)
			assert.Equal(t, gotRating, tt.wantTag)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}

func TestGetTagById(t *testing.T) {
	ctx := context.Background()

	type arg struct {
		ctx    context.Context
		tagId   string
	}

	tests := []struct {
		name     string
		before   func(t *testing.T)
		arg      arg
		wantErr  error
		wantTags *models.Tag
	}{
		{
			name: "success",
			arg: arg{
				ctx:    ctx,
				tagId: "65a6158d254f3843c2a3d42b",
			},
			wantTags: &models.Tag{
				Id:    "65a6158d254f3843c2a3d42b",
				Value: "test tag",
			},
			wantErr: nil,
		},
	}
	tagDal := &dal.TagDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before(t)
			}

			gotTags, err := tagDal.GetTagById(tt.arg.ctx, tt.arg.tagId)

			if tt.wantErr != nil {
				assert.Equal(t, err, tt.wantErr)
				assert.Nil(t, gotTags)
				return
			}
			assert.Equal(t, tt.wantTags, gotTags)
		})
	}
}
func TestGetTag(t *testing.T) {
	ctx := context.Background()

	type arg struct {
		ctx    context.Context
	}

	tests := []struct {
		name     string
		before   func(t *testing.T)
		arg      arg
		wantErr  error
		wantTags []*models.Tag
	}{
		{
			name: "success",
			arg: arg{
				ctx:    ctx,
			},
			wantTags: []*models.Tag{
				{
					Id:    "65a6158d254f3843c2a3d42b",
					Value: "test tag",
				},
			},
			wantErr: nil,
		},
	}
	tagDal := &dal.TagDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before(t)
			}

			gotTags, err := tagDal.GetTag(tt.arg.ctx)

			if tt.wantErr != nil {
				assert.Equal(t, err, tt.wantErr)
				assert.Nil(t, gotTags)
				return
			}
			assert.Equal(t, tt.wantTags, gotTags)
		})
	}
}

func TestUpdateTag(t *testing.T) {
	ctx := context.Background()

	type arg struct {
		ctx    context.Context
		tag 	*models.Tag
	}

	tests := []struct {
		name     string
		before   func(t *testing.T)
		arg      arg
		wantErr  error
		wantTags *models.Tag
	}{
		{
			name: "success",
			arg: arg{
				ctx:    ctx,
				tag:   &models.Tag{
					Id:    "65a6158d254f3843c2a3d42b",
					Value: "test tag",
				},
			},
			wantTags: &models.Tag{
				Id:    "65a6158d254f3843c2a3d42b",
				Value: "test tag",
			},
			wantErr: nil,
		},
	}
	tagDal := &dal.TagDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before(t)
			}

			gotTags, err := tagDal.UpdateTag(tt.arg.ctx, tt.arg.tag)

			if tt.wantErr != nil {
				assert.Equal(t, err, tt.wantErr)
				assert.Nil(t, gotTags)
				return
			}
			assert.Equal(t, tt.wantTags, gotTags)
		})
	}
}

func TestDeleteTag(t *testing.T) {
	ctx := context.Background()

	type arg struct {
		ctx    context.Context
		tagId string
	}

	tests := []struct {
		name     string
		before   func(t *testing.T)
		arg      arg
		wantErr  error
		wantTags *models.Tag
	}{
		{
			name: "success",
			arg: arg{
				ctx:    ctx,
				tagId: "65a6158d254f3843c2a3d42b",
			},
			wantTags: &models.Tag{
				Id:    "65a6158d254f3843c2a3d42b",
				Value: "test tag",
			},
			wantErr: nil,
		},
	}
	tagDal := &dal.TagDal{MainDB: db.GetMemoMongo(constant.MainMongoDB)}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before(t)
			}

			gotTags, err := tagDal.DeleteTag(tt.arg.ctx, tt.arg.tagId)

			if tt.wantErr != nil {
				assert.Equal(t, err, tt.wantErr)
				assert.Nil(t, gotTags)
				return
			}
			assert.Equal(t, tt.wantTags, gotTags)
		})
	}
}

