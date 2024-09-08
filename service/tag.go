package service
import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"
	"itineraryplanner/common/custom_errs"
	dal_inf "itineraryplanner/dal/inf"
	"itineraryplanner/models"
	"itineraryplanner/service/inf"
)
func NewTagService(dal dal_inf.TagDal) inf.TagService {
	return &TagService{
		Dal: dal,
	}
}
type TagService struct {
	Dal dal_inf.TagDal
}
func (t *TagService) CreateTag(ctx context.Context, req *models.CreateTagReq) (*models.CreateTagResp, error) {
	tag := &models.Tag{}
	err := copier.Copy(tag, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, custom_errs.InvalidInput
	}
	tag, err = t.Dal.CreateTag(ctx, tag)
	if err != nil {
		return nil, err
	}
	dto, err := t.ConvertDBOToDTOTag(ctx, tag)
	if err != nil {
		return nil, err
	}
	return &models.CreateTagResp{Tag: dto}, nil
}
func (t *TagService) ConvertDBOToDTOTag(ctx context.Context, tag *models.Tag) (*models.TagDTO, error) {
	if tag == nil {
		return nil, custom_errs.ServerError
	}
	tag1 := &models.TagDTO{}
	err:= copier.Copy(tag1, tag)
	if err != nil {
		return nil, custom_errs.InvalidInput
	}
	return tag1, nil
}
func (t *TagService) GetTagById(ctx context.Context, req *models.GetTagByIdReq) (*models.GetTagByIdResp, error) {
	Tag, err := t.Dal.GetTagById(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	Tag1 := &models.Tag{}
	ok := copier.Copy(Tag1, Tag)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", ok.Error())
		return nil, custom_errs.InvalidInput
	}
	dto, err := t.ConvertDBOToDTOTag(ctx, Tag1)
	if err != nil {
		return nil, err
	}
	return &models.GetTagByIdResp{Tag: dto}, nil
}
func (t *TagService) GetTag(ctx context.Context, req *models.GetTagReq) (*models.GetTagResp, error) {
	Tags, err := t.Dal.GetTag(ctx)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	Tag1 := []models.Tag{}
	ok := copier.Copy(Tag1, Tags)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", ok.Error())
		return nil, custom_errs.InvalidInput
	}
	dtos := make([]*models.TagDTO, 0)
	for _, v := range Tag1 {
		dto, err := t.ConvertDBOToDTOTag(ctx, &v)
		if err != nil {
			return nil, err
		}
		dtos = append(dtos, dto)
	}
	return &models.GetTagResp{Tags: dtos}, nil
}
func (t *TagService) UpdateTag(ctx context.Context, req *models.UpdateTagReq) (*models.UpdateTagResp, error) {
	tag := &models.Tag{}
	err := copier.Copy(tag, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, custom_errs.InvalidInput
	}
	tag, err = t.Dal.UpdateTag(ctx, tag)
	if err != nil {
		return nil, err
	}
	dto, err := t.ConvertDBOToDTOTag(ctx, tag)
	if err != nil {
		return nil, err
	}
	return &models.UpdateTagResp{Tag: dto}, nil
}
func (t *TagService) DeleteTag(ctx context.Context, req *models.DeleteTagReq) (*models.DeleteTagResp, error) {
	tag, err := t.Dal.DeleteTag(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	tag1 := &models.Tag{}
	ok := copier.Copy(tag1, tag)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", ok.Error())
		return nil, custom_errs.InvalidInput
	}
	dto, err := t.ConvertDBOToDTOTag(ctx, tag1)
	if err != nil {
		return nil, err
	}
	return &models.DeleteTagResp{Tag: dto}, nil
}