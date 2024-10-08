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
func NewItineraryService(dal dal_inf.ItineraryDal) inf.ItineraryService {
	return &ItineraryService{
		Dal: dal,
	}
}
type ItineraryService struct {
	Dal dal_inf.ItineraryDal
}
func (i *ItineraryService) CreateItinerary(ctx context.Context, req *models.CreateItineraryReq) (*models.CreateItineraryResp, error) {
	itinerary := &models.Itinerary{}
	err := copier.Copy(itinerary, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, custom_errs.InvalidInput
	}
	itinerary, err = i.Dal.CreateItinerary(ctx, itinerary)
	if err != nil {
		return nil, err
	}
	dto, err := i.ConvertDBOToDTOItinerary(ctx, itinerary)
	if err != nil {
		return nil, err
	}
	return &models.CreateItineraryResp{Itinerary: dto}, nil
}
func (i *ItineraryService) ConvertDBOToDTOItinerary(ctx context.Context, iti *models.Itinerary) (*models.ItineraryDTO, error) {
	if iti == nil {
		return nil, custom_errs.ServerError
	}
	ret := &models.ItineraryDTO{
		Id: iti.Id,
		CopiedId: iti.CopiedId,
		CopiedName: iti.CopiedName,
		Name: iti.Name,
		UserId: iti.UserId,
		StartTime: iti.StartTime,
		EndTime: iti.EndTime,
		EventIds: iti.EventIds,
		EventCount: iti.EventCount,
	}
	if utils.IsEmpty(iti.RatingId) {
		return nil, custom_errs.DBErrGetWithID
	}
	Icollection := i.Dal.GetDB().Collection(constant.RatingTable)
	IObjectID, err := primitive.ObjectIDFromHex(iti.RatingId)
	if err != nil {
		return nil, custom_errs.DBErrIDConversion
	}
	Iresult := Icollection.FindOne(ctx, bson.M{"_id": IObjectID})
	if Iresult.Err() != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	var rating *models.RatingDTO
	if err := Iresult.Decode(&rating); err != nil {
		return nil, custom_errs.DecodeErr
	}
	ret.Rating = rating
	return ret, nil
}
func (i *ItineraryService) GetItineraryById(ctx context.Context, req *models.GetItineraryByIdReq) (*models.GetItineraryByIdResp, error) {
	Itinerary, err := i.Dal.GetItineraryById(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	Itinerary1 := &models.Itinerary{}
	ok := copier.Copy(Itinerary1, Itinerary)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", ok.Error())
		return nil, custom_errs.InvalidInput
	}
	dto, err := i.ConvertDBOToDTOItinerary(ctx, Itinerary1)
	if err != nil {
		return nil, err
	}
	return &models.GetItineraryByIdResp{Itinerary: dto}, nil
}
func (i *ItineraryService) GetItinerary(ctx context.Context, req *models.GetItineraryReq) (*models.GetItineraryResp, error) {
	itineraries, err := i.Dal.GetItinerary(ctx)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	dtos := make([]*models.ItineraryDTO, 0)
	for _, v := range itineraries {
		dto, err := i.ConvertDBOToDTOItinerary(ctx, v)
		if err != nil {
			return nil, err
		}
		dtos = append(dtos, dto)
	}
	return &models.GetItineraryResp{Itineraries: dtos}, nil
}
func (i *ItineraryService) UpdateItinerary(ctx context.Context, req *models.UpdateItineraryReq) (*models.UpdateItineraryResp, error) {
	itinerary := &models.Itinerary{}
	err := copier.Copy(itinerary, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, custom_errs.InvalidInput
	}
	itinerary, err = i.Dal.UpdateItinerary(ctx, itinerary)
	if err != nil {
		return nil, err
	}
	dto, err := i.ConvertDBOToDTOItinerary(ctx, itinerary)
	if err != nil {
		return nil, err
	}
	return &models.UpdateItineraryResp{Itinerary: dto}, nil
}
func (i *ItineraryService) DeleteItinerary(ctx context.Context, req *models.DeleteItineraryReq) (*models.DeleteItineraryResp, error) {
	itinerary, err := i.Dal.DeleteItinerary(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	itinerary1 := &models.Itinerary{}
	ok := copier.Copy(itinerary1, itinerary)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", ok.Error())
		return nil, custom_errs.InvalidInput
	}
	dto, err := i.ConvertDBOToDTOItinerary(ctx, itinerary1)
	if err != nil {
		return nil, err
	}
	return &models.DeleteItineraryResp{Itinerary: dto}, nil
}