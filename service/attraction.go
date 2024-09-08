package service
import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"
	"itineraryplanner/common/custom_errs"
	"itineraryplanner/constant"
	dal_inf "itineraryplanner/dal/inf"
	"itineraryplanner/models"
	"itineraryplanner/service/inf"
	"itineraryplanner/common/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)
func NewAttractionService(dal dal_inf.AttractionDal) inf.AttractionService {
	return &AttractionService{
		Dal: dal,
	}
}
type AttractionService struct {
	Dal dal_inf.AttractionDal
}
func (a *AttractionService) CreateAttraction(ctx context.Context, req *models.CreateAttractionReq) (*models.CreateAttractionResp, error) {
	attraction := &models.Attraction{}
	err := copier.Copy(attraction, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, custom_errs.InvalidInput
	}
	attraction, err = a.Dal.CreateAttraction(ctx, attraction)
	if err != nil {
		return nil, err
	}
	dto, err := a.ConvertDBOToDTOAttraction(ctx, attraction)
	if err != nil {
		return nil, err
	}
	return &models.CreateAttractionResp{Attraction: dto}, nil
}
func (a *AttractionService) ConvertDBOToDTOAttraction(ctx context.Context, att *models.Attraction) (*models.AttractionDTO, error) {
	log.Info().Msg("ConvertDBOToDTOAttraction is called")
	if att == nil {
		return nil, custom_errs.ServerError
	}
	if utils.IsEmpty(att.RatingId) {
		return nil, custom_errs.ErrDataNotFound
	}
	ratingcollection := a.Dal.GetDB().Collection(constant.RatingTable)
	ratingObjectID, err := primitive.ObjectIDFromHex(att.RatingId)
	if err != nil {
		return nil, custom_errs.DBErrIDConversion
	}
	ratingResult := ratingcollection.FindOne(ctx, bson.M{"_id": ratingObjectID})
	if ratingResult.Err() != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	var ratingDTO *models.RatingDTO
	if err := ratingResult.Decode(&ratingDTO); err != nil {
		return nil, custom_errs.DecodeErr
	}
	ret := &models.AttractionDTO{
		Id: att.Id,
		Name: att.Name,
		Address: att.Address,
		X: att.X,
		Y: att.Y,
		City: att.City,
	}
	ret.Rating = ratingDTO
	tagCollection := a.Dal.GetDB().Collection(constant.TagTable)
	for _, v := range att.TagIDs{
		if utils.IsEmpty(v) {
			return nil, custom_errs.ErrDataNotFound
		}
		tagObjectID, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return nil, custom_errs.DBErrIDConversion
		}
		tagResult := tagCollection.FindOne(ctx, bson.M{"_id": tagObjectID})
		if tagResult.Err() != nil {
			return nil, custom_errs.DBErrGetWithID
		}
		var tag *models.TagDTO
		if err := tagResult.Decode(&tag); err != nil {
			return nil, custom_errs.DecodeErr
		}
		ret.Tags = append(ret.Tags, tag)
	}
	return ret, nil
}
func (a *AttractionService) GetAttractionById(ctx context.Context, req *models.GetAttractionByIdReq) (*models.GetAttractionByIdResp, error) {
	attraction, err := a.Dal.GetAttractionById(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	attraction1 := &models.Attraction{}
	err = copier.Copy(attraction1, attraction)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, custom_errs.InvalidInput
	}
	dto, err := a.ConvertDBOToDTOAttraction(ctx, attraction1)
	if err != nil {
		return nil, err
	}
	return &models.GetAttractionByIdResp{Attraction: dto}, nil
}
func (a *AttractionService) GetAttraction(ctx context.Context, req *models.GetAttractionReq) (*models.GetAttractionResp, error) {
	attractions, err := a.Dal.GetAttraction(ctx)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	dtos := make([]*models.AttractionDTO, 0)
	for _, v := range attractions {
		dto, err := a.ConvertDBOToDTOAttraction(ctx, v)
		if err != nil {
			return nil, err
		}
		dtos = append(dtos, dto)
	}
	return &models.GetAttractionResp{Attractions: dtos}, nil
}
func (a *AttractionService) UpdateAttraction(ctx context.Context, req *models.UpdateAttractionReq) (*models.UpdateAttractionResp, error) {
	attraction := &models.Attraction{}
	err := copier.Copy(attraction, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, custom_errs.InvalidInput
	}
	attraction, err = a.Dal.UpdateAttraction(ctx, attraction)
	if err != nil {
		return nil, err
	}
	dto, err := a.ConvertDBOToDTOAttraction(ctx, attraction)
	if err != nil {
		return nil, err
	}
	return &models.UpdateAttractionResp{Attraction: dto}, nil
}
func (a *AttractionService) DeleteAttraction(ctx context.Context, req *models.DeleteAttractionReq) (*models.DeleteAttractionResp, error) {
	attraction, err := a.Dal.DeleteAttraction(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	attraction1 := &models.Attraction{}
	err = copier.Copy(attraction1, attraction)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, custom_errs.InvalidInput
	}
	dto, err := a.ConvertDBOToDTOAttraction(ctx, attraction1)
	if err != nil {
		return nil, err
	}
	return &models.DeleteAttractionResp{Attraction: dto}, nil
}