package main

import (
	"flag"
	"os"

	"github.com/gin-gonic/gin"

	"itineraryplanner/common/config"
	"itineraryplanner/main_controllers/initialize"
	"itineraryplanner/main_controllers/inf"
)

func main() {
	flags := getFlags()
	config.InitGlobalConfig(flags.configPath)
	
	MA, err := initialize.InitializeMainAttractionController()
	if err != nil {
		os.Exit(1)
		return
	}
	MI, err := initialize.InitializeMainItineraryController()
	if err != nil {
		os.Exit(1)
		return
	}
	ME, err := initialize.InitializeMainEventController()
	if err != nil {
		os.Exit(1)
		return
	}
	MR, err := initialize.InitializeMainRatingController()
	if err != nil {
		os.Exit(1)
		return
	}
	MC, err := initialize.InitializeMainCoordinateController()
	if err != nil {
		os.Exit(1)
		return
	}
	MU, err := initialize.InitializeMainUserController()
	if err != nil {
		os.Exit(1)
		return
	}
	MT, err := initialize.InitializeMainTagController()
	if err != nil {
		os.Exit(1)
		return
	}
	MS, err := initialize.InitializeMainSearchEngineController()
	if err != nil {
		os.Exit(1)
		return
	}
	
	r := gin.Default()
	RouteS(r, MS)
	RouteA(r, MA)
	RouteI(r, MI)
	RouteE(r, ME)
	RouteR(r, MR)
	RouteC(r, MC)
	RouteU(r, MU)
	RouteT(r, MT)


	// Start the Gin server
	err = r.Run("localhost:8100")
	if err != nil {
		os.Exit(1)
		return
	}
}

type flags struct {
	configPath string
}

func getFlags() flags {
	filePath := flag.String("config", "./config/local_config", "config path")
	flag.Parse()

	if filePath == nil || *filePath == "" {
		panic("empty config file path")
	}

	return flags{
		configPath: *filePath,
	}
}

func RouteS(r *gin.Engine, s inf.MainSearchEngineController){
	r.POST("/search", s.SearchEngine)
}
func RouteA(r *gin.Engine, a inf.MainAttractionController) {
	r.POST("/attraction", a.CreateAttraction)
	r.GET("/attraction", a.GetAttraction)
	r.GET("/attraction/:id", a.GetAttractionById)
	r.PUT("/attraction/:id", a.UpdateAttraction)
	r.DELETE("/attraction/:id", a.DeleteAttraction)
}

func RouteI(r *gin.Engine, i inf.MainItineraryController) {
	r.POST("/itinerary", i.CreateItinerary)
	r.GET("/itinerary", i.GetItinerary)
	r.GET("/itinerary/:id", i.GetItineraryById)
	r.PUT("/itinerary/:id", i.UpdateItinerary)
	r.DELETE("/itinerary/:id", i.DeleteItinerary)
}
func RouteE(r *gin.Engine, e inf.MainEventController) {
	r.POST("/event", e.CreateEvent)
	r.GET("/event", e.GetEvent)
	r.GET("/event/:id", e.GetEventById)
	r.PUT("/event/:id", e.UpdateEvent)
	r.DELETE("/event/:id", e.DeleteEvent)
}

func RouteU(r *gin.Engine, u inf.MainUserController) {
	r.POST("/user", u.CreateUser)
	r.GET("/user", u.GetUser)
	r.GET("/user/:id", u.GetUserById)
	r.PUT("/user/:id", u.UpdateUser)
	r.DELETE("/user/:id", u.DeleteUser)
}
func RouteC(r *gin.Engine, c inf.MainCoordinateController) {
	r.POST("/coordinate", c.CreateCoordinate)
	r.GET("/coordinate", c.GetCoordinate)
	r.GET("/coordinate/:id", c.GetCoordinateById)
	r.PUT("/coordinate/:id", c.UpdateCoordinate)
	r.DELETE("/coordinate/:id", c.DeleteCoordinate)
}
func RouteR(r *gin.Engine, R inf.MainRatingController) {
	r.POST("/rating", R.CreateRating)
	r.GET("/rating", R.GetRating)
	r.GET("/rating/:id", R.GetRatingById)
	r.PUT("/rating/:id", R.UpdateRating)
	r.DELETE("/rating/:id", R.DeleteRating)
}
func RouteT(r *gin.Engine, t inf.MainTagController) {
	r.POST("/tag", t.CreateTag)
	r.GET("/tag", t.GetTag)
	r.GET("/tag/:id", t.GetTagById)
	r.PUT("/tag/:id", t.UpdateTag)
	r.DELETE("/tag/:id", t.DeleteTag)
}
