package test
import (
	"context"
	"itineraryplanner/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"itineraryplanner/dal/db"
	"itineraryplanner/dal"

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
				Value:"test",
			},
			wantErr: nil ,
		},
	}
	tagDal := &dal.TagDal{MainDB: db.GetMemoMongo()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotTag, gotErr:= tagDal.CreateTag(ctx, tt.arg.tag)
			gotTag.Id = ""
			assert.Equal(t, gotTag, tt.wantTag)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}
func TestGetTagById(t *testing.T) {
    ctx := context.Background()
    tagDal := &dal.TagDal{MainDB: db.GetMemoMongo()}
    type arg struct {
        ctx    context.Context
        tagId  string
    }
    tests := []struct {
        name     string
        before   func(t *testing.T) string 
        arg      arg
        wantTag *models.Tag
        wantErr  error
    }{
        {
            name: "success",
            before: func(t *testing.T) string {
                tag := &models.Tag{
                    Value: "test", 
                }
                newTag, err := tagDal.CreateTag(ctx, tag) 
                if err != nil {
                    t.Fatalf("Failed to create tag: %v", err)
                }
                return newTag.Id 
            },
            arg: arg{
                ctx:    ctx,
                tagId:  "", 
            },
            wantTag: &models.Tag{
                Value: "test",
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tagId := tt.before(t)
            tt.arg.tagId = tagId
			tt.wantTag.Id = tagId 
            gotTag, gotErr := tagDal.GetTagById(tt.arg.ctx, tt.arg.tagId)
            if tt.wantErr != nil {
                assert.Equal(t, tt.wantErr, gotErr)
                assert.Nil(t, gotTag)
                return
            } 
			assert.Nil(t, gotErr)
            assert.Equal(t, tt.wantTag, gotTag)
        })
    }
}
func TestGetTag(t *testing.T) {
	ctx := context.Background()
	type arg struct {
		ctx    context.Context
	}
	tagDal := &dal.TagDal{MainDB: db.GetMemoMongo()}
	tags, err := tagDal.GetTag(ctx)
	if err != nil {
        t.Fatalf("Failed to retrieve expected tags: %v", err)
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
			wantTags: tags,
			wantErr: nil,
		},
	}
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
    tagDal := &dal.TagDal{MainDB: db.GetMemoMongo()}
    type arg struct {
        ctx    context.Context
        tag    *models.Tag
    }
    tests := []struct {
        name     string
        before   func(t *testing.T) string 
        arg      arg
        wantTag *models.Tag
        wantErr  error
    }{
        {
            name: "success",
            before: func(t *testing.T) string {
                tag := &models.Tag{
                    Value: "initial tag", 
                }
                newTag, err := tagDal.CreateTag(ctx, tag) 
                if err != nil {
                    t.Fatalf("Failed to create tag: %v", err)
                }
                return newTag.Id 
            },
            arg: arg{
                ctx: ctx,
                tag: &models.Tag{
                    Id:    "", 
                    Value: "updated tag", 
                },
            },
            wantTag: &models.Tag{
				Id: "",
                Value: "updated tag",
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tagId := tt.before(t)
            tt.arg.tag.Id = tagId 
			tt.wantTag.Id = tagId
            gotTag, gotErr := tagDal.UpdateTag(tt.arg.ctx, tt.arg.tag)
            if tt.wantErr != nil {
                assert.Equal(t, tt.wantErr, gotErr)
                assert.Nil(t, gotTag)
                return
            }
            assert.Equal(t, tt.wantTag, gotTag)
			assert.Nil(t, gotErr)
        })
    }
}
func TestDeleteTag(t *testing.T) {
    ctx := context.Background()
    tagDal := &dal.TagDal{MainDB: db.GetMemoMongo()}
    type arg struct {
        ctx    context.Context
        tagId  string
    }
    tests := []struct {
        name     string
        before   func(t *testing.T) string 
        arg      arg
        wantTag *models.Tag
        wantErr  error
    }{
        {
            name: "success",
            before: func(t *testing.T) string {
                tag := &models.Tag{
                    Value: "test tag", 
                }
                newTag, err := tagDal.CreateTag(ctx, tag)
                if err != nil {
                    t.Fatalf("Failed to create tag: %v", err)
                }
                return newTag.Id 
            },
            arg: arg{
                ctx:   ctx,
                tagId: "", 
            },
            wantTag: &models.Tag{
                Value: "test tag", 
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tagId := tt.before(t)
            tt.arg.tagId = tagId 
			tt.wantTag.Id = tagId
            gotTag, gotErr := tagDal.DeleteTag(tt.arg.ctx, tt.arg.tagId)
            if tt.wantErr != nil {
                assert.Equal(t, tt.wantErr, gotErr)
                assert.Nil(t, gotTag)
                return
            }
            assert.Equal(t, tt.wantTag, gotTag)
			assert.Nil(t, gotErr)
        })
    }
}