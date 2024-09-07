package inf
import (
	"context"
	"itineraryplanner/models"
	"go.mongodb.org/mongo-driver/mongo"
)
//go:generate mockgen -source=./inf.go -destination=../mock/dal_inf_mock.go -package=mock .
type AttractionDal interface {
	CreateAttraction(ctx context.Context, attraction *models.Attraction) (*models.Attraction, error)
	GetAttractionById(ctx context.Context, attractionId string)(*models.Attraction, error)
	GetAttraction(ctx context.Context)([]*models.Attraction, error)
	UpdateAttraction(ctx context.Context, attraction *models.Attraction) (*models.Attraction, error)
	DeleteAttraction(ctx context.Context, attractionId string)(*models.Attraction, error)
	GetDB() *mongo.Database
}
type TagDal interface {
	GetTagById(ctx context.Context, tagId string)(*models.Tag, error)
	GetTag(ctx context.Context)([]*models.Tag, error)
	UpdateTag(ctx context.Context, tag *models.Tag) (*models.Tag, error)
	DeleteTag(ctx context.Context, tagId string)(*models.Tag, error)
	CreateTag(ctx context.Context, tag *models.Tag) (*models.Tag, error)
	GetDB() *mongo.Database
}
type EventDal interface {
	CreateEvent(ctx context.Context, event *models.Event)(*models.Event, error)
	GetEventById(ctx context.Context, eventId string)(*models.Event, error)
	GetEvent(ctx context.Context)([]*models.Event, error)
	UpdateEvent(ctx context.Context, event *models.Event) (*models.Event, error)
	DeleteEvent(ctx context.Context, eventId string)(*models.Event, error)
	GetDB() *mongo.Database
}
type RatingDal interface {
	CreateRating(ctx context.Context, rating *models.Rating) (*models.Rating, error)
	GetRatingById(ctx context.Context, attractionId string)(*models.Rating, error)
	GetRating(ctx context.Context)([]*models.Rating, error)
	UpdateRating(ctx context.Context, rating *models.Rating) (*models.Rating, error)
	DeleteRating(ctx context.Context, ratingId string)(*models.Rating, error)
	GetDB() *mongo.Database
}
type ItineraryDal interface {
	CreateItinerary(ctx context.Context, itinerary *models.Itinerary)(*models.Itinerary, error)
	GetItineraryById(ctx context.Context, itineraryId string)(*models.Itinerary, error)
	GetItinerary(ctx context.Context)([]*models.Itinerary, error)
	UpdateItinerary(ctx context.Context, itinerary *models.Itinerary) (*models.Itinerary, error)
	DeleteItinerary(ctx context.Context, itineraryId string)(*models.Itinerary, error)
	GetDB() *mongo.Database
}
type UserDal interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserById(ctx context.Context, userId string)(*models.User, error)
	GetUser(ctx context.Context)([]*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, userId string)(*models.User, error)
	GetDB() *mongo.Database
}