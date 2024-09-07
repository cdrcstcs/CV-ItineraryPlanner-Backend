package service

import(
	"itineraryplanner/models"
	"context"
	"itineraryplanner/service/inf"
) 

type FacadeField struct {
	AS inf.AttractionService
	ES inf.EventService
	RS inf.RatingService
	IS inf.ItineraryService 
	US inf.UserService 
	CS inf.CoordinateService 
	TS inf.TagService 
}

type Facade struct {
	Service *FacadeField
}

func (f *Facade) Execute(ctx context.Context, req *models.ReqFacade, reqType string) (*models.RespFacade, error) {
	switch reqType {
		case "GASB":{
			result, err:= f.Service.AS.GetAttractionById(ctx, req.GARB)
			return &models.RespFacade{
				GARB: result,
			}, err
		} 
		case "GESB":{
			result, err:= f.Service.ES.GetEventById(ctx, req.GERB)
			return &models.RespFacade{
				GERB: result,
			}, err
		} 
		case "GRSB":{
			result, err:= f.Service.RS.GetRatingById(ctx, req.GRRB)
			return &models.RespFacade{
				GRRB: result,
			}, err
		} 
		case "GISB":{
			result, err:= f.Service.IS.GetItineraryById(ctx, req.GIRB)
			return &models.RespFacade{
				GIRB: result,
			}, err
		} 
		case "GUSB":{
			result, err:= f.Service.US.GetUserById(ctx, req.GURB)
			return &models.RespFacade{
				GURB: result,
			}, err
		} 
		case "GCSB":{
			result, err:= f.Service.CS.GetCoordinateById(ctx, req.GCRB)
			return &models.RespFacade{
				GCRB: result,
			}, err
		} 
		case "GTSB":{
			result, err:= f.Service.TS.GetTagById(ctx, req.GTRB)
			return &models.RespFacade{
				GTRB: result,
			}, err
		} 
	}
	return nil, nil
}
