package route 
import(
	"itineraryplanner/maincontrollers/inf"
	"github.com/gin-gonic/gin"
)
func RouteA(r *gin.Engine, a inf.MainAttractionController) {
	r.POST("/attraction", a.CreateAttraction)
	r.GET("/attraction", a.GetAttraction)
	r.GET("/attraction/:id", a.GetAttractionById)
	r.PUT("/attraction", a.UpdateAttraction)
	r.DELETE("/attraction/:id", a.DeleteAttraction)
}
func RouteI(r *gin.Engine, i inf.MainItineraryController) {
	r.POST("/itinerary", i.CreateItinerary)
	r.GET("/itinerary", i.GetItinerary)
	r.GET("/itinerary/:id", i.GetItineraryById)
	r.PUT("/itinerary", i.UpdateItinerary)
	r.DELETE("/itinerary/:id", i.DeleteItinerary)
}
func RouteE(r *gin.Engine, e inf.MainEventController) {
	r.POST("/event", e.CreateEvent)
	r.GET("/event", e.GetEvent)
	r.GET("/event/:id", e.GetEventById)
	r.PUT("/event", e.UpdateEvent)
	r.DELETE("/event/:id", e.DeleteEvent)
}
func RouteU(r *gin.Engine, u inf.MainUserController) {
	r.POST("/user", u.CreateUser)
	r.GET("/user", u.GetUser)
	r.GET("/user/:id", u.GetUserById)
	r.PUT("/user", u.UpdateUser)
	r.DELETE("/user/:id", u.DeleteUser)
}
func RouteR(r *gin.Engine, R inf.MainRatingController) {
	r.POST("/rating", R.CreateRating)
	r.GET("/rating", R.GetRating)
	r.GET("/rating/:id", R.GetRatingById)
	r.PUT("/rating", R.UpdateRating)
	r.DELETE("/rating/:id", R.DeleteRating)
}
func RouteT(r *gin.Engine, t inf.MainTagController) {
	r.POST("/tag", t.CreateTag)
	r.GET("/tag", t.GetTag)
	r.GET("/tag/:id", t.GetTagById)
	r.PUT("/tag", t.UpdateTag)
	r.DELETE("/tag/:id", t.DeleteTag)
}