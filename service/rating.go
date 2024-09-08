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
func NewRatingService(dal dal_inf.RatingDal) inf.RatingService {
	return &RatingService{
		Dal: dal,
	}
}
type RatingService struct {
	Dal dal_inf.RatingDal
}
func (a *RatingService) CreateRating(ctx context.Context, req *models.CreateRatingReq) (*models.CreateRatingResp, error) {
	rating := &models.Rating{}
	err := copier.Copy(rating, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, custom_errs.InvalidInput
	}
	rating, err = a.Dal.CreateRating(ctx, rating)
	if err != nil {
		return nil, err
	}
	dto, err := a.ConvertDBOToDTORating(ctx, rating)
	if err != nil {
		return nil, err
	}
	return &models.CreateRatingResp{Rating: dto}, nil
}
func (a *RatingService) ConvertDBOToDTORating(ctx context.Context, rat *models.Rating) (*models.RatingDTO, error) {
	if rat == nil {
		return nil, custom_errs.ServerError
	}
	ret := &models.RatingDTO{}
	err := copier.Copy(ret, rat)
	if err != nil {
		return nil, custom_errs.InvalidInput
	}
	return ret, nil
}
func (r *RatingService) GetRatingById(ctx context.Context, req *models.GetRatingByIdReq) (*models.GetRatingByIdResp, error) {
	Rating, err := r.Dal.GetRatingById(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	Rating1 := &models.Rating{}
	ok := copier.Copy(Rating1, Rating)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", ok.Error())
		return nil, custom_errs.InvalidInput
	}
	dto, err := r.ConvertDBOToDTORating(ctx, Rating1)
	if err != nil {
		return nil, err
	}
	return &models.GetRatingByIdResp{Rating: dto}, nil
}
func (r *RatingService) GetRating(ctx context.Context, req *models.GetRatingReq) (*models.GetRatingResp, error) {
	ratings, err := r.Dal.GetRating(ctx)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	dtos := make([]*models.RatingDTO, 0)
	for _, v := range ratings {
		dto, err := r.ConvertDBOToDTORating(ctx, v)
		if err != nil {
			return nil, err
		}
		dtos = append(dtos, dto)
	}
	return &models.GetRatingResp{Ratings: dtos}, nil
}
func (r *RatingService) UpdateRating(ctx context.Context, req *models.UpdateRatingReq) (*models.UpdateRatingResp, error) {
	rating := &models.Rating{}
	err := copier.Copy(rating, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, custom_errs.InvalidInput
	}
	rating, err = r.Dal.UpdateRating(ctx, rating)
	if err != nil {
		return nil, err
	}
	dto, err := r.ConvertDBOToDTORating(ctx, rating)
	if err != nil {
		return nil, err
	}
	return &models.UpdateRatingResp{Rating: dto}, nil
}
func (r *RatingService) DeleteRating(ctx context.Context, req *models.DeleteRatingReq) (*models.DeleteRatingResp, error) {
	rating, err := r.Dal.DeleteRating(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	rating1 := &models.Rating{}
	ok := copier.Copy(rating1, rating)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", ok.Error())
		return nil, custom_errs.InvalidInput
	}
	dto, err := r.ConvertDBOToDTORating(ctx, rating1)
	if err != nil {
		return nil, err
	}
	return &models.DeleteRatingResp{Rating: dto}, nil
}