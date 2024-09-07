package initialize

import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/main_controllers"
	"itineraryplanner/main_controllers/inf"

)

func InitializeMainCreateItineraryController() (inf.MainCreateItineraryController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	itineraryDal := dal.NewCreateItineraryDal(mainMongoDB)
	itineraryService := service.NewCreateItineraryService(itineraryDal)
	itineraryController := controllers.NewCreateItineraryController(itineraryService)
	mainItineraryController := main_controllers.NewMainCreateItineraryController(itineraryController)
	return mainItineraryController, nil
}

func InitializeMainGetItineraryController() (inf.MainGetItineraryController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	itineraryDal := dal.NewGetItineraryDal(mainMongoDB)
	itineraryService := service.NewGetItineraryService(itineraryDal)
	itineraryController := controllers.NewGetItineraryController(itineraryService)
	mainItineraryController := main_controllers.NewMainGetItineraryController(itineraryController)
	return mainItineraryController, nil
}
func InitializeMainGetItineraryByIdController() (inf.MainGetItineraryByIdController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	ItineraryDal := dal.NewGetItineraryByIdDal(mainMongoDB)
	ItineraryService := service.NewGetItineraryByIdService(ItineraryDal)
	ItineraryController := controllers.NewGetItineraryByIdController(ItineraryService)
	mainItineraryController := main_controllers.NewMainGetItineraryByIdController(ItineraryController)
	return mainItineraryController, nil
}

func InitializeMainUpdateItineraryController() (inf.MainUpdateItineraryController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	itineraryDal := dal.NewUpdateItineraryDal(mainMongoDB)
	itineraryService := service.NewUpdateItineraryService(itineraryDal)
	itineraryController := controllers.NewUpdateItineraryController(itineraryService)
	mainItineraryController := main_controllers.NewMainUpdateItineraryController(itineraryController)
	return mainItineraryController, nil
}

func InitializeMainDeleteItineraryController() (inf.MainDeleteItineraryController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	itineraryDal := dal.NewDeleteItineraryDal(mainMongoDB)
	itineraryService := service.NewDeleteItineraryService(itineraryDal)
	itineraryController := controllers.NewDeleteItineraryController(itineraryService)
	mainItineraryController := main_controllers.NewMainDeleteItineraryController(itineraryController)
	return mainItineraryController, nil
}

