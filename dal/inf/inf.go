package inf

import (
	"context"

	"itineraryplanner/models"
)

//go:generate mockgen -source=./inf.go -destination=../mock/dal_inf_mock.go -package=mock .
type CreateAttractionDal interface {
	CreateAttraction(ctx context.Context, attraction *models.Attraction) (*models.Attraction, error)
}
type GetAttractionByIdDal interface {
	GetAttractionById(ctx context.Context, attractionId string)(*models.Attraction, error)
}
type GetAttractionDal interface {
	GetAttraction(ctx context.Context)([]*models.Attraction, error)
}
type UpdateAttractionDal interface {
	UpdateAttraction(ctx context.Context, attraction *models.Attraction) (*models.Attraction, error)
}
type DeleteAttractionDal interface {
	DeleteAttraction(ctx context.Context, attractionId string)(*models.Attraction, error)
}
type GetTagByIdDal interface {
	GetTagById(ctx context.Context, tagId string)(*models.Tag, error)
}
type GetTagDal interface {
	GetTag(ctx context.Context)([]*models.Tag, error)
}
type UpdateTagDal interface {
	UpdateTag(ctx context.Context, tag *models.Tag) (*models.Tag, error)
}
type DeleteTagDal interface {
	DeleteTag(ctx context.Context, tagId string)(*models.Tag, error)
}
type CreateEventDal interface {
	CreateEvent(ctx context.Context, event *models.Event)(*models.Event, error)
}
type GetEventByIdDal interface {
	GetEventById(ctx context.Context, eventId string)(*models.Event, error)
}
type GetEventDal interface {
	GetEvent(ctx context.Context)([]*models.Event, error)
}
type UpdateEventDal interface {
	UpdateEvent(ctx context.Context, event *models.Event) (*models.Event, error)
}
type DeleteEventDal interface {
	DeleteEvent(ctx context.Context, eventId string)(*models.Event, error)
}
type CreateRatingDal interface {
	CreateRating(ctx context.Context, rating *models.Rating) (*models.Rating, error)
}
type GetRatingByIdDal interface {
	GetRatingById(ctx context.Context, attractionId string)(*models.Rating, error)
}
type GetRatingDal interface {
	GetRating(ctx context.Context)([]*models.Rating, error)
}
type UpdateRatingDal interface {
	UpdateRating(ctx context.Context, rating *models.Rating) (*models.Rating, error)
}
type DeleteRatingDal interface {
	DeleteRating(ctx context.Context, ratingId string)(*models.Rating, error)
}
type CreateTagDal interface {
	CreateTag(ctx context.Context, tag *models.Tag) (*models.Tag, error)
}
type CreateItineraryDal interface {
	CreateItinerary(ctx context.Context, itinerary *models.Itinerary)(*models.Itinerary, error)
}
type GetItineraryByIdDal interface {
	GetItineraryById(ctx context.Context, itineraryId string)(*models.Itinerary, error)
}
type GetItineraryDal interface {
	GetItinerary(ctx context.Context)([]*models.Itinerary, error)
}
type UpdateItineraryDal interface {
	UpdateItinerary(ctx context.Context, itinerary *models.Itinerary) (*models.Itinerary, error)
}
type DeleteItineraryDal interface {
	DeleteItinerary(ctx context.Context, itineraryId string)(*models.Itinerary, error)
}
type CreateUserDal interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
}
type GetUserByIdDal interface {
	GetUserById(ctx context.Context, userId string)(*models.User, error)
}
type GetUserDal interface {
	GetUser(ctx context.Context)([]*models.User, error)
}
type UpdateUserDal interface {
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
}
type LoginUserDal interface {
	LoginUser(ctx context.Context, user *models.User) (bool, error)
}

type DeleteUserDal interface {
	DeleteUser(ctx context.Context, userId string)(*models.User, error)
}
type CreateCoordinateDal interface {
	CreateCoordinate(ctx context.Context, coordinate *models.Coordinate) (*models.Coordinate, error)
}
type GetCoordinateByIdDal interface {
	GetCoordinateById(ctx context.Context, coordinateId string)(*models.Coordinate, error)
}
type GetCoordinateDal interface {
	GetCoordinate(ctx context.Context)([]*models.Coordinate, error)
}
type UpdateCoordinateDal interface {
	UpdateCoordinate(ctx context.Context, coordinate *models.Coordinate) (*models.Coordinate, error)
}
type DeleteCoordinateDal interface {
	DeleteCoordinate(ctx context.Context, coordinateId string)(*models.Coordinate, error)
}
type RecommendItineraryDal interface {
	RecommendItinerary(ctx context.Context, cor *models.Coordinate) ([]*models.Itinerary, error)
}
type BuildItineraryDal interface {
	BuildItinerary(ctx context.Context, cor *models.Coordinate) (*models.Itinerary, error)
}
