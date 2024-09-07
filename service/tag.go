package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"itineraryplanner/common/custom_errs"
	dal_inf "itineraryplanner/dal/inf"
	"itineraryplanner/models"
	"itineraryplanner/service/inf"
)

func NewCreateTagService(cdal dal_inf.CreateTagDal) inf.CreateTagService {
	return &TagService{
		CDal: cdal,
	}
}
func NewGetTagByIdService(bdal dal_inf.GetTagByIdDal) inf.GetTagByIdService {
	return &TagService{
		BDal: bdal,
	}
}
func NewGetTagService(gdal dal_inf.GetTagDal) inf.GetTagService {
	return &TagService{
		GDal: gdal,
	}
}
func NewUpdateTagService(udal dal_inf.UpdateTagDal) inf.UpdateTagService {
	return &TagService{
		UDal: udal,
	}
}
func NewDeleteTagService(ddal dal_inf.DeleteTagDal) inf.DeleteTagService {
	return &TagService{
		DDal: ddal,
	}
}
func NewTagDTOService() inf.TagDTOService {
	return &TagService{}
}
type TagService struct {
	CDal dal_inf.CreateTagDal
	BDal dal_inf.GetTagByIdDal
	GDal dal_inf.GetTagDal
	UDal dal_inf.UpdateTagDal
	DDal dal_inf.DeleteTagDal
}

func (t *TagService) CreateTag(ctx context.Context, req *models.CreateTagReq) (*models.CreateTagResp, error) {
	tag := &models.Tag{}
	err := copier.Copy(tag, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	tag, err = t.CDal.CreateTag(ctx, tag)
	if err != nil {
		// TODO logging
		return nil, err
	}

	dto, err := t.ConvertDBOToDTOTag(ctx, tag)
	if err != nil {
		// TODO logging
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
		// TODO logging
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	return tag1, nil
}

func (t *TagService) GetTagById(ctx context.Context, req *models.GetTagByIdReq) (*models.GetTagByIdResp, error) {
	Tag, err := t.BDal.GetTagById(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	Tag1 := &models.Tag{}
	ok := copier.Copy(Tag1, Tag)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dto, err := t.ConvertDBOToDTOTag(ctx, Tag1)
	if err != nil {
		// TODO logging
		return nil, err
	}
	return &models.GetTagByIdResp{Tag: dto}, nil
}

func (t *TagService) GetTag(ctx context.Context, req *models.GetTagReq) (*models.GetTagResp, error) {
	Tags, err := t.GDal.GetTag(ctx)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	Tag1 := []models.Tag{}
	ok := copier.Copy(Tag1, Tags)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dtos := make([]*models.TagDTO, 0)
	for _, v := range Tag1 {
		dto, err := t.ConvertDBOToDTOTag(ctx, &v)
		if err != nil {
			// TODO logging
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
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	tag, err = t.UDal.UpdateTag(ctx, tag)
	if err != nil {
		// TODO logging
		return nil, err
	}
	dto, err := t.ConvertDBOToDTOTag(ctx, tag)
	if err != nil {
		// TODO logging
		return nil, err
	}

	return &models.UpdateTagResp{Tag: dto}, nil
}

func (t *TagService) DeleteTag(ctx context.Context, req *models.DeleteTagReq) (*models.DeleteTagResp, error) {
	tag, err := t.DDal.DeleteTag(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	tag1 := &models.Tag{}
	ok := copier.Copy(tag1, tag)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dto, err := t.ConvertDBOToDTOTag(ctx, tag1)
	if err != nil {
		// TODO logging
		return nil, err
	}
	return &models.DeleteTagResp{Tag: dto}, nil
}