package inf
import(
	"github.com/gin-gonic/gin"
)
type MainCreateAttractionController interface {
	CreateAttraction(c *gin.Context)
}
type MainGetAttractionController interface {
	GetAttraction(c *gin.Context)
}
type MainGetAttractionByIdController interface {
	GetAttractionById(c *gin.Context)
}
type MainUpdateAttractionController interface {
	UpdateAttraction(c *gin.Context)
}
type MainDeleteAttractionController interface {
	DeleteAttraction(c *gin.Context)
}
type MainCreateEventController interface {
	CreateEvent(c *gin.Context)
}
type MainGetEventController interface {
	GetEvent(c *gin.Context)
}
type MainGetEventByIdController interface {
	GetEventById(c *gin.Context)
}
type MainUpdateEventController interface {
	UpdateEvent(c *gin.Context)
}
type MainDeleteEventController interface {
	DeleteEvent(c *gin.Context)
}
type MainCreateItineraryController interface {
	CreateItinerary(c *gin.Context)
}
type MainGetItineraryController interface {
	GetItinerary(c *gin.Context)
}
type MainGetItineraryByIdController interface {
	GetItineraryById(c *gin.Context)
}
type MainUpdateItineraryController interface {
	UpdateItinerary(c *gin.Context)
}
type MainDeleteItineraryController interface {
	DeleteItinerary(c *gin.Context)
}
type MainCreateCoordinateController interface {
	CreateCoordinate(c *gin.Context)
}
type MainGetCoordinateController interface {
	GetCoordinate(c *gin.Context)
}
type MainGetCoordinateByIdController interface {
	GetCoordinateById(c *gin.Context)
}

type MainUpdateCoordinateController interface {
	UpdateCoordinate(c *gin.Context)
}
type MainDeleteCoordinateController interface {
	DeleteCoordinate(c *gin.Context)
}
type MainCreateUserController interface {
	CreateUser(c *gin.Context)
}
type MainGetUserController interface {
	GetUser(c *gin.Context)
}
type MainLoginUserController interface {
	LoginUser(c *gin.Context)
}
type MainGetUserByIdController interface {
	GetUserById(c *gin.Context)
}

type MainUpdateUserController interface {
	UpdateUser(c *gin.Context)
}
type MainDeleteUserController interface {
	DeleteUser(c *gin.Context)
}
type MainCreateRatingController interface {
	CreateRating(c *gin.Context)
}
type MainGetRatingController interface {
	GetRating(c *gin.Context)
}
type MainGetRatingByIdController interface {
	GetRatingById(c *gin.Context)
}

type MainUpdateRatingController interface {
	UpdateRating(c *gin.Context)
}
type MainDeleteRatingController interface {
	DeleteRating(c *gin.Context)
}
type MainCreateTagController interface {
	CreateTag(c *gin.Context)
}
type MainGetTagController interface {
	GetTag(c *gin.Context)
}
type MainGetTagByIdController interface {
	GetTagById(c *gin.Context)
}

type MainUpdateTagController interface {
	UpdateTag(c *gin.Context)
}
type MainDeleteTagController interface {
	DeleteTag(c *gin.Context)
}

type MainRecommendItineraryController interface {
	RecommendItinerary(c *gin.Context)
}
type MainBuildItineraryController interface {
	BuildItinerary(c *gin.Context)
}