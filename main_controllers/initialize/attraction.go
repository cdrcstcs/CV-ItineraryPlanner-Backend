package initialize

import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/main_controllers/inf"
	"itineraryplanner/main_controllers"
)

func InitializeMainAttractionController() (inf.MainAttractionController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	attractionDal := dal.NewAttractionDal(mainMongoDB)
	attractionService := service.NewAttractionService(attractionDal)
	attractionController := controllers.NewAttractionController(attractionService)
	mainAttractionController := main_controllers.NewMainAttractionController(attractionController)
	return mainAttractionController, nil
}