package service

import (
	"context"
	"github.com/rs/zerolog/log"

	"itineraryplanner/common/custom_errs"
	"itineraryplanner/constant"
	dal_inf "itineraryplanner/dal/inf"
	"itineraryplanner/models"
	"itineraryplanner/service/inf"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"itineraryplanner/common/utils"
)


func NewSearchEngineService(dal dal_inf.SearchEngineDal) inf.SearchEngineService {
	return &SearchEngineService{
		Dal: dal,
	}
}

type SearchEngineService struct {
	Dal dal_inf.SearchEngineDal
}
func (s *SearchEngineService) SearchEngine(ctx context.Context, req *models.SearchEngineReq) (*models.SearchEngineResp, error) {
	attractions, err := s.Dal.SearchEngine(ctx, req.Query)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	dtos := make([]*models.AttractionDTO, 0)
	for _, v := range attractions {
		dto, err := s.ConvertDBOToDTOAttraction(ctx, v)
		if err != nil {
			return nil, err
		}
		dtos = append(dtos, dto)
	}
	return &models.SearchEngineResp{Attractions: dtos}, nil
}

func (s *SearchEngineService) ConvertDBOToDTOAttraction(ctx context.Context, att *models.Attraction) (*models.AttractionDTO, error) {
	log.Info().Msg("ConvertDBOToDTOAttraction is called") // Add this line for debugging
	if att == nil {
		return nil, custom_errs.ServerError
	}
	ret := &models.AttractionDTO{}
	if utils.IsEmpty(att.RatingId) {
		// TODO logging here
		return nil, custom_errs.DBErrGetWithID
	}
	Rcollection := s.Dal.GetDB().Collection(constant.RatingTable)
	RObjectID, err := primitive.ObjectIDFromHex(att.RatingId)
	if err != nil {
		return nil, custom_errs.DBErrIDConversion
	}
	Rresult := Rcollection.FindOne(ctx, bson.M{"_id": RObjectID})
	if Rresult.Err() != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	var rating *models.RatingDTO
	if err := Rresult.Decode(&rating); err != nil {
		return nil, custom_errs.DecodeErr
	}
	ret.Rating = rating

	for _, v := range att.TagIDs{
		if utils.IsEmpty(v) {
			// TODO logging here
			return nil, custom_errs.DBErrGetWithID
		}
		Tcollection := s.Dal.GetDB().Collection(constant.TagTable)
		TObjectID, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return nil, custom_errs.DBErrIDConversion
		}
		Tresult := Tcollection.FindOne(ctx, bson.M{"_id": TObjectID})
		if Tresult.Err() != nil {
			return nil, custom_errs.DBErrGetWithID
		}
		var tag *models.TagDTO
		if err := Tresult.Decode(&tag); err != nil {
			return nil, custom_errs.DecodeErr
		}
		ret.Tags = append(ret.Tags, tag)
	}

	return ret, nil
}
