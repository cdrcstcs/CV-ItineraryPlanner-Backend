package main_controller

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
	MCA, err := initialize.InitializeMainCreateAttractionController()
	if err != nil {
		os.Exit(1)
		return
	}
	MGA, err := initialize.InitializeMainGetAttractionController()
	if err != nil {
		os.Exit(1)
		return
	}
	MGAB, err := initialize.InitializeMainGetAttractionByIdController()
	if err != nil {
		os.Exit(1)
		return
	}
	MUA, err := initialize.InitializeMainUpdateAttractionController()
	if err != nil {
		os.Exit(1)
		return
	}
	MDA, err := initialize.InitializeMainDeleteAttractionController()
	if err != nil {
		os.Exit(1)
		return
	}
	MCI, err := initialize.InitializeMainCreateItineraryController()
	if err != nil {
		os.Exit(1)
		return
	}
	MGI, err := initialize.InitializeMainGetItineraryController()
	if err != nil {
		os.Exit(1)
		return
	}
	MGIB, err := initialize.InitializeMainGetItineraryByIdController()
	if err != nil {
		os.Exit(1)
		return
	}
	MUI, err := initialize.InitializeMainUpdateItineraryController()
	if err != nil {
		os.Exit(1)
		return
	}
	MDI, err := initialize.InitializeMainDeleteItineraryController()
	if err != nil {
		os.Exit(1)
		return
	}
	MCE, err := initialize.InitializeMainCreateEventController()
	if err != nil {
		os.Exit(1)
		return
	}
	MGE, err := initialize.InitializeMainGetEventController()
	if err != nil {
		os.Exit(1)
		return
	}
	MGEB, err := initialize.InitializeMainGetEventByIdController()
	if err != nil {
		os.Exit(1)
		return
	}
	MUE, err := initialize.InitializeMainUpdateEventController()
	if err != nil {
		os.Exit(1)
		return
	}
	MDE, err := initialize.InitializeMainDeleteEventController()
	if err != nil {
		os.Exit(1)
		return
	}
	MCR, err := initialize.InitializeMainCreateRatingController()
	if err != nil {
		os.Exit(1)
		return
	}
	MGR, err := initialize.InitializeMainGetRatingController()
	if err != nil {
		os.Exit(1)
		return
	}
	MGRB, err := initialize.InitializeMainGetRatingByIdController()
	if err != nil {
		os.Exit(1)
		return
	}
	MUR, err := initialize.InitializeMainUpdateRatingController()
	if err != nil {
		os.Exit(1)
		return
	}
	MDR, err := initialize.InitializeMainDeleteRatingController()
	if err != nil {
		os.Exit(1)
		return
	}
	MCC, err := initialize.InitializeMainCreateCoordinateController()
	if err != nil {
		os.Exit(1)
		return
	}
	MGC, err := initialize.InitializeMainGetCoordinateController()
	if err != nil {
		os.Exit(1)
		return
	}
	MGCB, err := initialize.InitializeMainGetCoordinateByIdController()
	if err != nil {
		os.Exit(1)
		return
	}
	MUC, err := initialize.InitializeMainUpdateCoordinateController()
	if err != nil {
		os.Exit(1)
		return
	}
	MDC, err := initialize.InitializeMainDeleteCoordinateController()
	if err != nil {
		os.Exit(1)
		return
	}
	MCU, err := initialize.InitializeMainCreateUserController()
	if err != nil {
		os.Exit(1)
		return
	}
	MGU, err := initialize.InitializeMainGetUserController()
	if err != nil {
		os.Exit(1)
		return
	}
	MGUB, err := initialize.InitializeMainGetUserByIdController()
	if err != nil {
		os.Exit(1)
		return
	}
	MUU, err := initialize.InitializeMainUpdateUserController()
	if err != nil {
		os.Exit(1)
		return
	}
	MDU, err := initialize.InitializeMainDeleteUserController()
	if err != nil {
		os.Exit(1)
		return
	}
	MCT, err := initialize.InitializeMainCreateTagController()
	if err != nil {
		os.Exit(1)
		return
	}
	MGT, err := initialize.InitializeMainGetTagController()
	if err != nil {
		os.Exit(1)
		return
	}
	MGTB, err := initialize.InitializeMainGetTagByIdController()
	if err != nil {
		os.Exit(1)
		return
	}
	MUT, err := initialize.InitializeMainUpdateTagController()
	if err != nil {
		os.Exit(1)
		return
	}
	MDT, err := initialize.InitializeMainDeleteTagController()
	if err != nil {
		os.Exit(1)
		return
	}
	MRI, err := initialize.InitializeMainRecommendItineraryController()
	if err != nil {
		os.Exit(1)
		return
	}
	MBI, err := initialize.InitializeMainBuildItineraryController()
	if err != nil {
		os.Exit(1)
		return
	}
	MLU, err := initialize.InitializeMainLoginUserController()
	if err != nil {
		os.Exit(1)
		return
	}
	r := gin.Default()
	RouteAlgo(r, MRI, MBI)
	RouteA(r, MCA, MGA, MGAB, MUA, MDA)
	RouteI(r, MCI, MGI, MGIB, MUI, MDI)
	RouteE(r, MCE, MGE, MGEB, MUE, MDE)
	RouteR(r, MCR, MGR, MGRB, MUR, MDR)
	RouteC(r, MCC, MGC, MGCB, MUC, MDC)
	RouteU(r, MCU, MGU, MGUB, MUU, MDU, MLU)
	RouteT(r, MCT, MGT, MGTB, MUT, MDT)


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
func RouteAlgo(r *gin.Engine, ri inf.MainRecommendItineraryController, b inf.MainBuildItineraryController) {
	r.POST("/algo/recommend", ri.RecommendItinerary)
	r.POST("/algo/build", b.BuildItinerary)
}

func RouteA(r *gin.Engine, c inf.MainCreateAttractionController, g inf.MainGetAttractionController, gb inf.MainGetAttractionByIdController , u inf.MainUpdateAttractionController, d inf.MainDeleteAttractionController) {
	r.POST("/attraction", c.CreateAttraction)
	r.GET("/attraction", g.GetAttraction)
	r.GET("/attraction/:id", gb.GetAttractionById)
	r.PUT("/attraction/:id", u.UpdateAttraction)
	r.DELETE("/attraction/:id", d.DeleteAttraction)
}

func RouteI(r *gin.Engine, c inf.MainCreateItineraryController, g inf.MainGetItineraryController, gb inf.MainGetItineraryByIdController, u inf.MainUpdateItineraryController, d inf.MainDeleteItineraryController) {
	r.POST("/itinerary", c.CreateItinerary)
	r.GET("/itinerary", g.GetItinerary)
	r.GET("/itinerary/:id", gb.GetItineraryById)
	r.PUT("/itinerary/:id", u.UpdateItinerary)
	r.DELETE("/itinerary/:id", d.DeleteItinerary)
}
func RouteE(r *gin.Engine, c inf.MainCreateEventController, g inf.MainGetEventController, gb inf.MainGetEventByIdController, u inf.MainUpdateEventController, d inf.MainDeleteEventController) {
	r.POST("/event", c.CreateEvent)
	r.GET("/event", g.GetEvent)
	r.GET("/event/:id", gb.GetEventById)
	r.PUT("/event/:id", u.UpdateEvent)
	r.DELETE("/event/:id", d.DeleteEvent)
}
func RouteU(r *gin.Engine, c inf.MainCreateUserController, g inf.MainGetUserController, gb inf.MainGetUserByIdController, u inf.MainUpdateUserController, d inf.MainDeleteUserController, l inf.MainLoginUserController) {
	r.POST("/user", c.CreateUser)
	r.POST("/user/login",l.LoginUser)
	r.GET("/user", g.GetUser)
	r.GET("/user/:id", gb.GetUserById)
	r.PUT("/user/:id", u.UpdateUser)
	r.DELETE("/user/:id", d.DeleteUser)
}
func RouteC(r *gin.Engine, c inf.MainCreateCoordinateController, g inf.MainGetCoordinateController, gb inf.MainGetCoordinateByIdController, u inf.MainUpdateCoordinateController, d inf.MainDeleteCoordinateController) {
	r.POST("/coordinate", c.CreateCoordinate)
	r.GET("/coordinate", g.GetCoordinate)
	r.GET("/coordinate/:id", gb.GetCoordinateById)
	r.PUT("/coordinate/:id", u.UpdateCoordinate)
	r.DELETE("/coordinate/:id", d.DeleteCoordinate)
}
func RouteR(r *gin.Engine, c inf.MainCreateRatingController, g inf.MainGetRatingController, gb inf.MainGetRatingByIdController, u inf.MainUpdateRatingController, d inf.MainDeleteRatingController) {
	r.POST("/rating", c.CreateRating)
	r.GET("/rating", g.GetRating)
	r.GET("/rating/:id", gb.GetRatingById)
	r.PUT("/rating/:id", u.UpdateRating)
	r.DELETE("/rating/:id", d.DeleteRating)
}
func RouteT(r *gin.Engine, c inf.MainCreateTagController, g inf.MainGetTagController, gb inf.MainGetTagByIdController, u inf.MainUpdateTagController, d inf.MainDeleteTagController) {
	r.POST("/tag", c.CreateTag)
	r.GET("/tag", g.GetTag)
	r.GET("/tag/:id", gb.GetTagById)
	r.PUT("/tag/:id", u.UpdateTag)
	r.DELETE("/tag/:id", d.DeleteTag)
}
