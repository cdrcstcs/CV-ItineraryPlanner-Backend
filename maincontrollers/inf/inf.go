package inf
import(
	"github.com/gin-gonic/gin"
)
//go:generate mockgen -source=./inf.go -destination=../mock/main_controllers_inf_mock.go -package=mock .
type MainAttractionController interface {
	CreateAttraction(c *gin.Context)
	GetAttraction(c *gin.Context)
	GetAttractionById(c *gin.Context)
	UpdateAttraction(c *gin.Context)
	DeleteAttraction(c *gin.Context)
}
type MainEventController interface {
	CreateEvent(c *gin.Context)
	GetEvent(c *gin.Context)
	GetEventById(c *gin.Context)
	UpdateEvent(c *gin.Context)
	DeleteEvent(c *gin.Context)
}
type MainItineraryController interface {
	CreateItinerary(c *gin.Context)
	GetItinerary(c *gin.Context)
	GetItineraryById(c *gin.Context)
	UpdateItinerary(c *gin.Context)
	DeleteItinerary(c *gin.Context)
}
type MainCoordinateController interface {
	CreateCoordinate(c *gin.Context)
	GetCoordinate(c *gin.Context)
	GetCoordinateById(c *gin.Context)
	UpdateCoordinate(c *gin.Context)
	DeleteCoordinate(c *gin.Context)
}
type MainUserController interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
type MainRatingController interface {
	CreateRating(c *gin.Context)
	GetRating(c *gin.Context)
	GetRatingById(c *gin.Context)
	UpdateRating(c *gin.Context)
	DeleteRating(c *gin.Context)
}
type MainTagController interface {
	CreateTag(c *gin.Context)
	GetTag(c *gin.Context)
	GetTagById(c *gin.Context)
	UpdateTag(c *gin.Context)
	DeleteTag(c *gin.Context)
}
type MainSearchEngineController interface{
	SearchEngine(c *gin.Context)
}