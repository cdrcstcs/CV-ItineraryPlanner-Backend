package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"itineraryplanner/common/custom_errs"
	"itineraryplanner/constant"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	dal_inf "itineraryplanner/dal/inf"
	"itineraryplanner/models"
	"itineraryplanner/service/inf"
)

func NewCreateRatingService(cdal dal_inf.CreateRatingDal) inf.CreateRatingService {
	return &RatingService{
		CDal: cdal,
	}
}
func NewGetRatingByIdService(bdal dal_inf.GetRatingByIdDal) inf.GetRatingByIdService {
	return &RatingService{
		BDal: bdal,
	}
}
func NewGetRatingService(gdal dal_inf.GetRatingDal) inf.GetRatingService {
	return &RatingService{
		GDal: gdal,
	}
}
func NewUpdateRatingService(udal dal_inf.UpdateRatingDal) inf.UpdateRatingService {
	return &RatingService{
		UDal: udal,
	}
}
func NewDeleteRatingService(ddal dal_inf.DeleteRatingDal) inf.DeleteRatingService {
	return &RatingService{
		DDal: ddal,
	}
}
func NewRatingDTOService() inf.RatingDTOService {
	return &RatingService{}
}
type RatingService struct {
	CDal dal_inf.CreateRatingDal
	BDal dal_inf.GetRatingByIdDal
	GDal dal_inf.GetRatingDal
	UDal dal_inf.UpdateRatingDal
	DDal dal_inf.DeleteRatingDal
}

func (a *RatingService) CreateRating(ctx context.Context, req *models.CreateRatingReq) (*models.CreateRatingResp, error) {
	rating := &models.Rating{}
	err := copier.Copy(rating, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	rating, err = a.CDal.CreateRating(ctx, rating)
	if err != nil {
		// TODO logging
		return nil, err
	}

	dto, err := a.ConvertDBOToDTORating(ctx, rating)
	if err != nil {
		// TODO logging
		return nil, err
	}

	return &models.CreateRatingResp{Rating: dto}, nil
}

func (a *RatingService) ConvertDBOToDTORating(ctx context.Context, rat *models.Rating) (*models.RatingDTO, error) {
	if rat == nil {
		return nil, custom_errs.ServerError
	}
	facade := &Facade{
		Service: &FacadeField{
			US: &UserService{
				GDal: &dal.UserDal{
					MainDB: db.GetMemoMongo(constant.MainMongoDB),
				},
			},
		},
	}
	rating := &models.RatingDTO{}
	err := copier.Copy(rating, rat)
	if err != nil {
		// TODO logging
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	if rat.UserId != ""{
		req2 := &models.GetUserByIdReq{
			Id: rat.UserId,
		}
		resp2, err := facade.Execute(ctx, &models.ReqFacade{
			GURB: req2,
		}, "GUSB")
		if err != nil {
			// TODO logging
			return nil, err
		}
		rating.User = resp2.GURB.User
	}
	return rating, nil
}

func (r *RatingService) GetRatingById(ctx context.Context, req *models.GetRatingByIdReq) (*models.GetRatingByIdResp, error) {
	Rating, err := r.BDal.GetRatingById(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	Rating1 := &models.Rating{}
	ok := copier.Copy(Rating1, Rating)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dto, err := r.ConvertDBOToDTORating(ctx, Rating1)
	if err != nil {
		// TODO logging
		return nil, err
	}
	return &models.GetRatingByIdResp{Rating: dto}, nil
}

func (r *RatingService) GetRating(ctx context.Context, req *models.GetRatingReq) (*models.GetRatingResp, error) {
	Ratings, err := r.GDal.GetRating(ctx)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	Rating1 := []models.Rating{}
	ok := copier.Copy(Rating1, Ratings)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dtos := make([]*models.RatingDTO, 0)
	for _, v := range Rating1 {
		dto, err := r.ConvertDBOToDTORating(ctx, &v)
		if err != nil {
			// TODO logging
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
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	rating, err = r.UDal.UpdateRating(ctx, rating)
	if err != nil {
		// TODO logging
		return nil, err
	}
	dto, err := r.ConvertDBOToDTORating(ctx, rating)
	if err != nil {
		// TODO logging
		return nil, err
	}

	return &models.UpdateRatingResp{Rating: dto}, nil
}

func (r *RatingService) DeleteRating(ctx context.Context, req *models.DeleteRatingReq) (*models.DeleteRatingResp, error) {
	rating, err := r.DDal.DeleteRating(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	rating1 := &models.Rating{}
	ok := copier.Copy(rating1, rating)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dto, err := r.ConvertDBOToDTORating(ctx, rating1)
	if err != nil {
		// TODO logging
		return nil, err
	}
	return &models.DeleteRatingResp{Rating: dto}, nil
}