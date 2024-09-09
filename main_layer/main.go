package main
import (
	"flag"
	"os"
	"github.com/gin-gonic/gin"
	"itineraryplanner/common/config"
	"itineraryplanner/maincontrollers/initialize"
	"itineraryplanner/main_layer/route"
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
	r := gin.Default()
	route.RouteA(r, MA)
	route.RouteI(r, MI)
	route.RouteE(r, ME)
	route.RouteR(r, MR)
	route.RouteU(r, MU)
	route.RouteT(r, MT)
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
	filePath := flag.String("config", "./config/local_config.json", "config path")
	flag.Parse()
	if filePath == nil || *filePath == "" {
		panic("empty config file path")
	}
	return flags{
		configPath: *filePath,
	}
}