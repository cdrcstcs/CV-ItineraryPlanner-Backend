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

func NewCreateCoordinateService(cdal dal_inf.CreateCoordinateDal) inf.CreateCoordinateService {
	return &CoordinateService{
		CDal: cdal,
	}
}
func NewGetCoordinateByIdService(bdal dal_inf.GetCoordinateByIdDal) inf.GetCoordinateByIdService {
	return &CoordinateService{
		BDal: bdal,
	}
}
func NewGetCoordinateService(gdal dal_inf.GetCoordinateDal) inf.GetCoordinateService {
	return &CoordinateService{
		GDal: gdal,
	}
}
func NewUpdateCoordinateService(udal dal_inf.UpdateCoordinateDal) inf.UpdateCoordinateService {
	return &CoordinateService{
		UDal: udal,
	}
}
func NewDeleteCoordinateService(ddal dal_inf.DeleteCoordinateDal) inf.DeleteCoordinateService {
	return &CoordinateService{
		DDal: ddal,
	}
}
func NewCoordinateDTOService() inf.CoordinateDTOService {
	return &CoordinateService{}
}

type CoordinateService struct {
	CDal dal_inf.CreateCoordinateDal
	BDal dal_inf.GetCoordinateByIdDal
	GDal dal_inf.GetCoordinateDal
	UDal dal_inf.UpdateCoordinateDal
	DDal dal_inf.DeleteCoordinateDal
}

func (c *CoordinateService) CreateCoordinate(ctx context.Context, req *models.CreateCoordinateReq) (*models.CreateCoordinateResp, error) {
	Coordinate := &models.Coordinate{}
	err := copier.Copy(Coordinate, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	Coordinate, err = c.CDal.CreateCoordinate(ctx, Coordinate)
	if err != nil {
		// TODO logging
		return nil, err
	}
	dto, err := c.ConvertDBOToDTOCoordinate(ctx, Coordinate)
	if err != nil {
		// TODO logging
		return nil, err
	}

	return &models.CreateCoordinateResp{Coordinate: dto}, nil
}

func (c *CoordinateService) ConvertDBOToDTOCoordinate(ctx context.Context, att *models.Coordinate) (*models.CoordinateDTO, error) {
	log.Info().Msg("ConvertDBOToDTOCoordinate is called") // Add this line for debugging
	if att == nil {
		return nil, custom_errs.ServerError
	}
	Coordinate := &models.CoordinateDTO{}
	err := copier.Copy(Coordinate, att)
	if err != nil {
		// TODO logging
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	return Coordinate, nil
}

func (c *CoordinateService) GetCoordinateById(ctx context.Context, req *models.GetCoordinateByIdReq) (*models.GetCoordinateByIdResp, error) {
	Coordinate, err := c.BDal.GetCoordinateById(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	Coordinate1 := &models.Coordinate{}
	ok := copier.Copy(Coordinate1, Coordinate)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dto, err := c.ConvertDBOToDTOCoordinate(ctx, Coordinate1)
	if err != nil {
		// TODO logging
		return nil, err
	}
	return &models.GetCoordinateByIdResp{Coordinate: dto}, nil
}

func (c *CoordinateService) GetCoordinate(ctx context.Context, req *models.GetCoordinateReq) (*models.GetCoordinateResp, error) {
	Coordinates, err := c.GDal.GetCoordinate(ctx)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	Coordinate1 := []models.Coordinate{}
	ok := copier.Copy(Coordinate1, Coordinates)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dtos := make([]*models.CoordinateDTO, 0)
	for _, v := range Coordinate1 {
		dto, err := c.ConvertDBOToDTOCoordinate(ctx, &v)
		if err != nil {
			// TODO logging
			return nil, err
		}
		dtos = append(dtos, dto)
	}
	return &models.GetCoordinateResp{Coordinates: dtos}, nil
}

func (c *CoordinateService) UpdateCoordinate(ctx context.Context, req *models.UpdateCoordinateReq) (*models.UpdateCoordinateResp, error) {
	Coordinate := &models.Coordinate{}
	err := copier.Copy(Coordinate, req)
	if err != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	Coordinate, err = c.UDal.UpdateCoordinate(ctx, Coordinate)
	if err != nil {
		// TODO logging
		return nil, err
	}
	dto, err := c.ConvertDBOToDTOCoordinate(ctx, Coordinate)
	if err != nil {
		// TODO logging
		return nil, err
	}

	return &models.UpdateCoordinateResp{Coordinate: dto}, nil
}

func (c *CoordinateService) DeleteCoordinate(ctx context.Context, req *models.DeleteCoordinateReq) (*models.DeleteCoordinateResp, error) {
	Coordinate, err := c.DDal.DeleteCoordinate(ctx, req.Id)
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}

	Coordinate1 := &models.Coordinate{}
	ok := copier.Copy(Coordinate1, Coordinate)
	if ok != nil {
		log.Error().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}
	dto, err := c.ConvertDBOToDTOCoordinate(ctx, Coordinate1)
	if err != nil {
		// TODO logging
		return nil, err
	}
	return &models.DeleteCoordinateResp{Coordinate: dto}, nil
}
