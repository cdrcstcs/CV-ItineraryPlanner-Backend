package inf

import (
	"context"

	"itineraryplanner/models"
)

//go:generate mockgen -source=./inf.go -destination=../mock/controllers_inf_mock.go -package=mock .

type AttractionController interface {
	CreateAttraction(ctx context.Context, req *models.CreateAttractionReq) (*models.CreateAttractionResp, error)
	GetAttraction(ctx context.Context, req *models.GetAttractionReq) (*models.GetAttractionResp,error)
	GetAttractionById(ctx context.Context, req *models.GetAttractionByIdReq) (*models.GetAttractionByIdResp,error)
	UpdateAttraction(ctx context.Context, req *models.UpdateAttractionReq) (*models.UpdateAttractionResp, error)
	DeleteAttraction(ctx context.Context, req *models.DeleteAttractionReq) (*models.DeleteAttractionResp,error)
}
type ItineraryController interface {
	CreateItinerary(ctx context.Context, req *models.CreateItineraryReq) (*models.CreateItineraryResp, error)
	GetItinerary(ctx context.Context, req *models.GetItineraryReq) (*models.GetItineraryResp,error)
	GetItineraryById(ctx context.Context, req *models.GetItineraryByIdReq) (*models.GetItineraryByIdResp,error)
	UpdateItinerary(ctx context.Context, req *models.UpdateItineraryReq) (*models.UpdateItineraryResp, error)
	DeleteItinerary(ctx context.Context, req *models.DeleteItineraryReq) (*models.DeleteItineraryResp,error)
}
type EventController interface {
	CreateEvent(ctx context.Context, req *models.CreateEventReq) (*models.CreateEventResp, error)
	GetEvent(ctx context.Context, req *models.GetEventReq) (*models.GetEventResp,error)
	GetEventById(ctx context.Context, req *models.GetEventByIdReq) (*models.GetEventByIdResp,error)
	UpdateEvent(ctx context.Context, req *models.UpdateEventReq) (*models.UpdateEventResp, error)
	DeleteEvent(ctx context.Context, req *models.DeleteEventReq) (*models.DeleteEventResp,error)
}
type CoordinateController interface {
	CreateCoordinate(ctx context.Context, req *models.CreateCoordinateReq) (*models.CreateCoordinateResp, error)
	GetCoordinate(ctx context.Context, req *models.GetCoordinateReq) (*models.GetCoordinateResp,error)
	GetCoordinateById(ctx context.Context, req *models.GetCoordinateByIdReq) (*models.GetCoordinateByIdResp,error)
	UpdateCoordinate(ctx context.Context, req *models.UpdateCoordinateReq) (*models.UpdateCoordinateResp, error)
	DeleteCoordinate(ctx context.Context, req *models.DeleteCoordinateReq) (*models.DeleteCoordinateResp,error)
}
type TagController interface {
	CreateTag(ctx context.Context, req *models.CreateTagReq) (*models.CreateTagResp, error)
	GetTag(ctx context.Context, req *models.GetTagReq) (*models.GetTagResp,error)
	GetTagById(ctx context.Context, req *models.GetTagByIdReq) (*models.GetTagByIdResp,error)
	UpdateTag(ctx context.Context, req *models.UpdateTagReq) (*models.UpdateTagResp, error)
	DeleteTag(ctx context.Context, req *models.DeleteTagReq) (*models.DeleteTagResp,error)
}
type UserController interface {
	CreateUser(ctx context.Context, req *models.CreateUserReq) (*models.CreateUserResp, error)
	GetUser(ctx context.Context, req *models.GetUserReq) (*models.GetUserResp,error)
	GetUserById(ctx context.Context, req *models.GetUserByIdReq) (*models.GetUserByIdResp,error)
	UpdateUser(ctx context.Context, req *models.UpdateUserReq) (*models.UpdateUserResp, error)
	DeleteUser(ctx context.Context, req *models.DeleteUserReq) (*models.DeleteUserResp,error)
}
type RatingController interface {
	CreateRating(ctx context.Context, req *models.CreateRatingReq) (*models.CreateRatingResp, error)
	GetRating(ctx context.Context, req *models.GetRatingReq) (*models.GetRatingResp,error)
	GetRatingById(ctx context.Context, req *models.GetRatingByIdReq) (*models.GetRatingByIdResp,error)
	UpdateRating(ctx context.Context, req *models.UpdateRatingReq) (*models.UpdateRatingResp, error)
	DeleteRating(ctx context.Context, req *models.DeleteRatingReq) (*models.DeleteRatingResp,error)
}
type SearchEngineController interface {
	SearchEngine(ctx context.Context, req *models.SearchEngineReq) (*models.SearchEngineResp, error)
}


