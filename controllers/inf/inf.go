package inf

import (
	"context"

	"itineraryplanner/models"
)

//  mockgen -destination=./controllers/mock/controllers_inf_mock.go -package=mock -source=./controllers/inf GetAttractionController
//go:generate mockgen -source=./inf.go -destination=../mock/controllers_inf_mock.go -package=mock .

type CreateAttractionController interface {
	CreateAttraction(ctx context.Context, req *models.CreateAttractionReq) (*models.CreateAttractionResp, error)
}
type GetAttractionController interface{
	GetAttraction(ctx context.Context, req *models.GetAttractionReq) (*models.GetAttractionResp,error)
}
type GetAttractionByIdController interface{
	GetAttractionById(ctx context.Context, req *models.GetAttractionByIdReq) (*models.GetAttractionByIdResp,error)
}
type UpdateAttractionController interface {
	UpdateAttraction(ctx context.Context, req *models.UpdateAttractionReq) (*models.UpdateAttractionResp, error)
}
type DeleteAttractionController interface{
	DeleteAttraction(ctx context.Context, req *models.DeleteAttractionReq) (*models.DeleteAttractionResp,error)
}
type CreateItineraryController interface {
	CreateItinerary(ctx context.Context, req *models.CreateItineraryReq) (*models.CreateItineraryResp, error)
}
type GetItineraryController interface{
	GetItinerary(ctx context.Context, req *models.GetItineraryReq) (*models.GetItineraryResp,error)
}
type GetItineraryByIdController interface{
	GetItineraryById(ctx context.Context, req *models.GetItineraryByIdReq) (*models.GetItineraryByIdResp,error)
}
type UpdateItineraryController interface {
	UpdateItinerary(ctx context.Context, req *models.UpdateItineraryReq) (*models.UpdateItineraryResp, error)
}
type DeleteItineraryController interface{
	DeleteItinerary(ctx context.Context, req *models.DeleteItineraryReq) (*models.DeleteItineraryResp,error)
}
type CreateEventController interface {
	CreateEvent(ctx context.Context, req *models.CreateEventReq) (*models.CreateEventResp, error)
}
type GetEventController interface{
	GetEvent(ctx context.Context, req *models.GetEventReq) (*models.GetEventResp,error)
}
type GetEventByIdController interface{
	GetEventById(ctx context.Context, req *models.GetEventByIdReq) (*models.GetEventByIdResp,error)
}
type UpdateEventController interface {
	UpdateEvent(ctx context.Context, req *models.UpdateEventReq) (*models.UpdateEventResp, error)
}
type DeleteEventController interface{
	DeleteEvent(ctx context.Context, req *models.DeleteEventReq) (*models.DeleteEventResp,error)
}
type CreateCoordinateController interface {
	CreateCoordinate(ctx context.Context, req *models.CreateCoordinateReq) (*models.CreateCoordinateResp, error)
}
type GetCoordinateController interface{
	GetCoordinate(ctx context.Context, req *models.GetCoordinateReq) (*models.GetCoordinateResp,error)
}
type GetCoordinateByIdController interface{
	GetCoordinateById(ctx context.Context, req *models.GetCoordinateByIdReq) (*models.GetCoordinateByIdResp,error)
}
type UpdateCoordinateController interface {
	UpdateCoordinate(ctx context.Context, req *models.UpdateCoordinateReq) (*models.UpdateCoordinateResp, error)
}
type DeleteCoordinateController interface{
	DeleteCoordinate(ctx context.Context, req *models.DeleteCoordinateReq) (*models.DeleteCoordinateResp,error)
}
type CreateTagController interface {
	CreateTag(ctx context.Context, req *models.CreateTagReq) (*models.CreateTagResp, error)
}
type GetTagController interface{
	GetTag(ctx context.Context, req *models.GetTagReq) (*models.GetTagResp,error)
}
type GetTagByIdController interface{
	GetTagById(ctx context.Context, req *models.GetTagByIdReq) (*models.GetTagByIdResp,error)
}
type UpdateTagController interface {
	UpdateTag(ctx context.Context, req *models.UpdateTagReq) (*models.UpdateTagResp, error)
}
type DeleteTagController interface{
	DeleteTag(ctx context.Context, req *models.DeleteTagReq) (*models.DeleteTagResp,error)
}
type CreateUserController interface {
	CreateUser(ctx context.Context, req *models.CreateUserReq) (*models.CreateUserResp, error)
}
type GetUserController interface{
	GetUser(ctx context.Context, req *models.GetUserReq) (*models.GetUserResp,error)
}
type LoginUserController interface{
	LoginUser(ctx context.Context, req *models.LoginUserReq) (*models.LoginUserResp,error)
}
type GetUserByIdController interface{
	GetUserById(ctx context.Context, req *models.GetUserByIdReq) (*models.GetUserByIdResp,error)
}
type UpdateUserController interface {
	UpdateUser(ctx context.Context, req *models.UpdateUserReq) (*models.UpdateUserResp, error)
}
type DeleteUserController interface{
	DeleteUser(ctx context.Context, req *models.DeleteUserReq) (*models.DeleteUserResp,error)
}
type CreateRatingController interface {
	CreateRating(ctx context.Context, req *models.CreateRatingReq) (*models.CreateRatingResp, error)
}
type GetRatingController interface{
	GetRating(ctx context.Context, req *models.GetRatingReq) (*models.GetRatingResp,error)
}
type GetRatingByIdController interface{
	GetRatingById(ctx context.Context, req *models.GetRatingByIdReq) (*models.GetRatingByIdResp,error)
}
type UpdateRatingController interface {
	UpdateRating(ctx context.Context, req *models.UpdateRatingReq) (*models.UpdateRatingResp, error)
}
type DeleteRatingController interface{
	DeleteRating(ctx context.Context, req *models.DeleteRatingReq) (*models.DeleteRatingResp,error)
}

type RecommendItineraryController interface {
	RecommendItinerary(ctx context.Context, req *models.RecommendItineraryReq)(*models.RecommendItineraryResp,error)
}
type BuildItineraryController interface {
	BuildItinerary(ctx context.Context,req *models.BuildItineraryReq) (*models.BuildItineraryResp, error)
}
