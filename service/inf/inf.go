package inf

import (
	"context"

	"itineraryplanner/models"
)

//go:generate mockgen -source=./inf.go -destination=../mock/service_inf_mock.go -package=mock .
type CreateAttractionService interface {
	CreateAttraction(ctx context.Context, req *models.CreateAttractionReq) (*models.CreateAttractionResp, error)
}
type GetAttractionByIdService interface {
	GetAttractionById(ctx context.Context, req *models.GetAttractionByIdReq) (*models.GetAttractionByIdResp, error)
}
type GetAttractionService interface {
	GetAttraction(ctx context.Context, req *models.GetAttractionReq) (*models.GetAttractionResp, error)
}
type UpdateAttractionService interface {
	UpdateAttraction(ctx context.Context, req *models.UpdateAttractionReq) (*models.UpdateAttractionResp, error)
}
type DeleteAttractionService interface {
	DeleteAttraction(ctx context.Context, req *models.DeleteAttractionReq) (*models.DeleteAttractionResp, error)
}
type AttractionService interface {
	CreateAttraction(ctx context.Context, req *models.CreateAttractionReq) (*models.CreateAttractionResp, error)
	GetAttractionById(ctx context.Context, req *models.GetAttractionByIdReq) (*models.GetAttractionByIdResp, error)
	GetAttraction(ctx context.Context, req *models.GetAttractionReq) (*models.GetAttractionResp, error)
	UpdateAttraction(ctx context.Context, req *models.UpdateAttractionReq) (*models.UpdateAttractionResp, error)
	DeleteAttraction(ctx context.Context, req *models.DeleteAttractionReq) (*models.DeleteAttractionResp, error)
}
type AttractionDTOService interface {
	ConvertDBOToDTOAttraction(ctx context.Context, att *models.Attraction) (*models.AttractionDTO, error)
}
type CreateEventService interface {
	CreateEvent(ctx context.Context, req *models.CreateEventReq)(*models.CreateEventResp,error)
}
type GetEventByIdService interface {
	GetEventById(ctx context.Context,req *models.GetEventByIdReq) (*models.GetEventByIdResp, error)
}
type GetEventService interface {
	GetEvent(ctx context.Context,req *models.GetEventReq) (*models.GetEventResp, error)
}
type UpdateEventService interface {
	UpdateEvent(ctx context.Context, req *models.UpdateEventReq)(*models.UpdateEventResp,error)
}
type DeleteEventService interface {
	DeleteEvent(ctx context.Context,req *models.DeleteEventReq) (*models.DeleteEventResp, error)
}
type EventService interface {
	CreateEvent(ctx context.Context, req *models.CreateEventReq) (*models.CreateEventResp, error)
	GetEventById(ctx context.Context,req *models.GetEventByIdReq) (*models.GetEventByIdResp, error)
	GetEvent(ctx context.Context, req *models.GetEventReq) (*models.GetEventResp, error)
	UpdateEvent(ctx context.Context, req *models.UpdateEventReq)(*models.UpdateEventResp,error)
	DeleteEvent(ctx context.Context,req *models.DeleteEventReq) (*models.DeleteEventResp, error)
}
type EventDTOService interface {
	ConvertDBOToDTOEvent(ctx context.Context, eve *models.Event) (*models.EventDTO, error)
}
type GetRatingByIdService interface {
	GetRatingById(ctx context.Context, req *models.GetRatingByIdReq) (*models.GetRatingByIdResp, error)
}
type GetRatingService interface {
	GetRating(ctx context.Context, req *models.GetRatingReq) (*models.GetRatingResp, error)
}
type CreateRatingService interface {
	CreateRating(ctx context.Context, req *models.CreateRatingReq) (*models.CreateRatingResp, error)
}
type UpdateRatingService interface {
	UpdateRating(ctx context.Context, req *models.UpdateRatingReq)(*models.UpdateRatingResp,error)
}
type DeleteRatingService interface {
	DeleteRating(ctx context.Context,req *models.DeleteRatingReq) (*models.DeleteRatingResp, error)
}
type RatingService interface {
	CreateRating(ctx context.Context, req *models.CreateRatingReq) (*models.CreateRatingResp, error)
	GetRatingById(ctx context.Context, req *models.GetRatingByIdReq) (*models.GetRatingByIdResp, error)
	GetRating(ctx context.Context, req *models.GetRatingReq) (*models.GetRatingResp, error)
	UpdateRating(ctx context.Context, req *models.UpdateRatingReq)(*models.UpdateRatingResp,error)
	DeleteRating(ctx context.Context,req *models.DeleteRatingReq) (*models.DeleteRatingResp, error)
}
type RatingDTOService interface {
	ConvertDBOToDTORating(ctx context.Context, rat *models.Rating) (*models.RatingDTO, error)
}
type GetItineraryByIdService interface {
	GetItineraryById(ctx context.Context, req *models.GetItineraryByIdReq) (*models.GetItineraryByIdResp, error)
}
type GetItineraryService interface {
	GetItinerary(ctx context.Context, req *models.GetItineraryReq) (*models.GetItineraryResp, error)
}
type CreateItineraryService interface {
	CreateItinerary(ctx context.Context, req *models.CreateItineraryReq) (*models.CreateItineraryResp, error)
}
type UpdateItineraryService interface {
	UpdateItinerary(ctx context.Context, req *models.UpdateItineraryReq)(*models.UpdateItineraryResp,error)
}
type DeleteItineraryService interface {
	DeleteItinerary(ctx context.Context,req *models.DeleteItineraryReq) (*models.DeleteItineraryResp, error)
}
type ItineraryService interface {
	CreateItinerary(ctx context.Context, req *models.CreateItineraryReq) (*models.CreateItineraryResp, error)
	GetItineraryById(ctx context.Context, req *models.GetItineraryByIdReq) (*models.GetItineraryByIdResp, error)
	GetItinerary(ctx context.Context, req *models.GetItineraryReq) (*models.GetItineraryResp, error)
	UpdateItinerary(ctx context.Context, req *models.UpdateItineraryReq)(*models.UpdateItineraryResp,error)
	DeleteItinerary(ctx context.Context,req *models.DeleteItineraryReq) (*models.DeleteItineraryResp, error)
}
type ItineraryDTOService interface {
	ConvertDBOToDTOItinerary(ctx context.Context, att *models.Itinerary) (*models.ItineraryDTO, error)
}
type GetCoordinateByIdService interface {
	GetCoordinateById(ctx context.Context, req *models.GetCoordinateByIdReq) (*models.GetCoordinateByIdResp, error)
}
type GetCoordinateService interface {
	GetCoordinate(ctx context.Context, req *models.GetCoordinateReq) (*models.GetCoordinateResp, error)
}
type CreateCoordinateService interface {
	CreateCoordinate(ctx context.Context, req *models.CreateCoordinateReq) (*models.CreateCoordinateResp, error)
}
type UpdateCoordinateService interface {
	UpdateCoordinate(ctx context.Context, req *models.UpdateCoordinateReq)(*models.UpdateCoordinateResp,error)
}
type DeleteCoordinateService interface {
	DeleteCoordinate(ctx context.Context,req *models.DeleteCoordinateReq) (*models.DeleteCoordinateResp, error)
}
type CoordinateService interface {
	CreateCoordinate(ctx context.Context, req *models.CreateCoordinateReq) (*models.CreateCoordinateResp, error)
	GetCoordinateById(ctx context.Context, req *models.GetCoordinateByIdReq) (*models.GetCoordinateByIdResp, error)
	GetCoordinate(ctx context.Context, req *models.GetCoordinateReq) (*models.GetCoordinateResp, error)
	UpdateCoordinate(ctx context.Context, req *models.UpdateCoordinateReq)(*models.UpdateCoordinateResp,error)
	DeleteCoordinate(ctx context.Context,req *models.DeleteCoordinateReq) (*models.DeleteCoordinateResp, error)
}
type CoordinateDTOService interface {
	ConvertDBOToDTOCoordinate(ctx context.Context, cor *models.Coordinate) (*models.CoordinateDTO, error)
}
type GetTagByIdService interface {
	GetTagById(ctx context.Context, req *models.GetTagByIdReq) (*models.GetTagByIdResp, error)
}
type GetTagService interface {
	GetTag(ctx context.Context, req *models.GetTagReq) (*models.GetTagResp, error)
}
type CreateTagService interface {
	CreateTag(ctx context.Context, req *models.CreateTagReq) (*models.CreateTagResp, error)
}
type UpdateTagService interface {
	UpdateTag(ctx context.Context, req *models.UpdateTagReq)(*models.UpdateTagResp,error)
}
type DeleteTagService interface {
	DeleteTag(ctx context.Context,req *models.DeleteTagReq) (*models.DeleteTagResp, error)
}
type TagService interface {
	CreateTag(ctx context.Context, req *models.CreateTagReq) (*models.CreateTagResp, error)
	GetTagById(ctx context.Context, req *models.GetTagByIdReq) (*models.GetTagByIdResp, error)
	GetTag(ctx context.Context, req *models.GetTagReq) (*models.GetTagResp, error)
	UpdateTag(ctx context.Context, req *models.UpdateTagReq)(*models.UpdateTagResp,error)
	DeleteTag(ctx context.Context,req *models.DeleteTagReq) (*models.DeleteTagResp, error)
}
type TagDTOService interface {
	ConvertDBOToDTOTag(ctx context.Context, eve *models.Tag) (*models.TagDTO, error)
}
type GetUserByIdService interface {
	GetUserById(ctx context.Context, req *models.GetUserByIdReq) (*models.GetUserByIdResp, error)
}
type GetUserService interface {
	GetUser(ctx context.Context, req *models.GetUserReq) (*models.GetUserResp, error)
}
type CreateUserService interface {
	CreateUser(ctx context.Context, req *models.CreateUserReq) (*models.CreateUserResp, error)
}
type UpdateUserService interface {
	UpdateUser(ctx context.Context, req *models.UpdateUserReq)(*models.UpdateUserResp,error)
}
type DeleteUserService interface {
	DeleteUser(ctx context.Context,req *models.DeleteUserReq) (*models.DeleteUserResp, error)
}
type LoginUserService interface {
	LoginUser(ctx context.Context, req *models.LoginUserReq) (*models.LoginUserResp, error)
}
type UserService interface {
	CreateUser(ctx context.Context, req *models.CreateUserReq) (*models.CreateUserResp, error)
	GetUserById(ctx context.Context, req *models.GetUserByIdReq) (*models.GetUserByIdResp, error)
	GetUser(ctx context.Context, req *models.GetUserReq) (*models.GetUserResp, error)
	UpdateUser(ctx context.Context, req *models.UpdateUserReq)(*models.UpdateUserResp,error)
	DeleteUser(ctx context.Context,req *models.DeleteUserReq) (*models.DeleteUserResp, error)
}
type UserDTOService interface {
	ConvertDBOToDTOUser(ctx context.Context, use *models.User) (*models.UserDTO, error)
}
type FacadeDesignPatternService interface {
	Execute(ctx context.Context, req *models.ReqFacade, reqType string)(*models.RespFacade, error)
}


type RecommendItineraryService interface {
	RecommendItinerary(ctx context.Context, req *models.RecommendItineraryReq)(*models.RecommendItineraryResp,error)
}
type BuildItineraryService interface {
	BuildItinerary(ctx context.Context,req *models.BuildItineraryReq) (*models.BuildItineraryResp, error)
}
type AlgoItineraryDTOService interface {
	AlgoConvertDBOToDTOItinerary(ctx context.Context, iti *models.Itinerary) (*models.ItineraryDTO, error)
}