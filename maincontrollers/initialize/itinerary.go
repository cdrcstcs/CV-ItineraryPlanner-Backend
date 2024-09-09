package initialize
import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/maincontrollers"
	"itineraryplanner/maincontrollers/inf"

)
func InitializeMainItineraryController() (inf.MainItineraryController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	ItineraryDal := dal.NewItineraryDal(mainMongoDB)
	ItineraryService := service.NewItineraryService(ItineraryDal)
	ItineraryController := controllers.NewItineraryController(ItineraryService)
	mainItineraryController := maincontrollers.NewMainItineraryController(ItineraryController)
	return mainItineraryController, nil
}