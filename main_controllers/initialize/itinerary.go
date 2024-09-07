package initialize

import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/main_controllers"
	"itineraryplanner/main_controllers/inf"

)

func InitializeMainItineraryController() (inf.MainItineraryController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	itineraryDal := dal.NewItineraryDal(mainMongoDB)
	itineraryService := service.NewItineraryService(itineraryDal)
	itineraryController := controllers.NewItineraryController(itineraryService)
	mainItineraryController := main_controllers.NewMainItineraryController(itineraryController)
	return mainItineraryController, nil
}
