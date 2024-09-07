package inf

import (
	"context"

	"itineraryplanner/models"
)

//go:generate mockgen -source=./inf.go -destination=../mock/service_inf_mock.go -package=mock .
type AttractionService interface {
	CreateAttraction(ctx context.Context, req *models.CreateAttractionReq) (*models.CreateAttractionResp, error)
	GetAttractionById(ctx context.Context, req *models.GetAttractionByIdReq) (*models.GetAttractionByIdResp, error)
	GetAttraction(ctx context.Context, req *models.GetAttractionReq) (*models.GetAttractionResp, error)
	UpdateAttraction(ctx context.Context, req *models.UpdateAttractionReq) (*models.UpdateAttractionResp, error)
	DeleteAttraction(ctx context.Context, req *models.DeleteAttractionReq) (*models.DeleteAttractionResp, error)

}
type EventService interface {
	CreateEvent(ctx context.Context, req *models.CreateEventReq) (*models.CreateEventResp, error)
	GetEventById(ctx context.Context,req *models.GetEventByIdReq) (*models.GetEventByIdResp, error)
	GetEvent(ctx context.Context, req *models.GetEventReq) (*models.GetEventResp, error)
	UpdateEvent(ctx context.Context, req *models.UpdateEventReq)(*models.UpdateEventResp,error)
	DeleteEvent(ctx context.Context,req *models.DeleteEventReq) (*models.DeleteEventResp, error)
}
type RatingService interface {
	CreateRating(ctx context.Context, req *models.CreateRatingReq) (*models.CreateRatingResp, error)
	GetRatingById(ctx context.Context, req *models.GetRatingByIdReq) (*models.GetRatingByIdResp, error)
	GetRating(ctx context.Context, req *models.GetRatingReq) (*models.GetRatingResp, error)
	UpdateRating(ctx context.Context, req *models.UpdateRatingReq)(*models.UpdateRatingResp,error)
	DeleteRating(ctx context.Context,req *models.DeleteRatingReq) (*models.DeleteRatingResp, error)
}
type ItineraryService interface {
	CreateItinerary(ctx context.Context, req *models.CreateItineraryReq) (*models.CreateItineraryResp, error)
	GetItineraryById(ctx context.Context, req *models.GetItineraryByIdReq) (*models.GetItineraryByIdResp, error)
	GetItinerary(ctx context.Context, req *models.GetItineraryReq) (*models.GetItineraryResp, error)
	UpdateItinerary(ctx context.Context, req *models.UpdateItineraryReq)(*models.UpdateItineraryResp,error)
	DeleteItinerary(ctx context.Context,req *models.DeleteItineraryReq) (*models.DeleteItineraryResp, error)
}
type TagService interface {
	CreateTag(ctx context.Context, req *models.CreateTagReq) (*models.CreateTagResp, error)
	GetTagById(ctx context.Context, req *models.GetTagByIdReq) (*models.GetTagByIdResp, error)
	GetTag(ctx context.Context, req *models.GetTagReq) (*models.GetTagResp, error)
	UpdateTag(ctx context.Context, req *models.UpdateTagReq)(*models.UpdateTagResp,error)
	DeleteTag(ctx context.Context,req *models.DeleteTagReq) (*models.DeleteTagResp, error)
}
type UserService interface {
	CreateUser(ctx context.Context, req *models.CreateUserReq) (*models.CreateUserResp, error)
	GetUserById(ctx context.Context, req *models.GetUserByIdReq) (*models.GetUserByIdResp, error)
	GetUser(ctx context.Context, req *models.GetUserReq) (*models.GetUserResp, error)
	UpdateUser(ctx context.Context, req *models.UpdateUserReq)(*models.UpdateUserResp,error)
	DeleteUser(ctx context.Context,req *models.DeleteUserReq) (*models.DeleteUserResp, error)
}
type SearchEngineService interface{
	SearchEngine(ctx context.Context,req *models.SearchEngineReq) (*models.SearchEngineResp, error)
}


