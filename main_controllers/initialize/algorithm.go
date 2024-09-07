package initialize

import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/main_controllers/inf"
	"itineraryplanner/main_controllers"
)

func InitializeMainRecommendItineraryController() (inf.MainRecommendItineraryController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	itineraryDal := dal.NewRecommendItineraryDal(mainMongoDB)
	itineraryService := service.NewRecommendItineraryService(itineraryDal)
	itineraryController := controllers.NewRecommendItineraryController(itineraryService)
	mainitineraryController := main_controllers.NewMainRecommendItineraryController(itineraryController)
	return mainitineraryController, nil
}

func InitializeMainBuildItineraryController() (inf.MainBuildItineraryController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	itineraryDal := dal.NewBuildItineraryDal(mainMongoDB)
	itineraryService := service.NewBuildItineraryService(itineraryDal)
	itineraryController := controllers.NewBuildItineraryController(itineraryService)
	mainitineraryController := main_controllers.NewMainBuildItineraryController(itineraryController)
	return mainitineraryController, nil
}