package initialize

import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/main_controllers/inf"
	"itineraryplanner/main_controllers"
)

func InitializeMainCreateAttractionController() (inf.MainCreateAttractionController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	attractionDal := dal.NewCreateAttractionDal(mainMongoDB)
	attractionService := service.NewCreateAttractionService(attractionDal)
	attractionController := controllers.NewCreateAttractionController(attractionService)
	mainAttractionController := main_controllers.NewMainCreateAttractionController(attractionController)
	return mainAttractionController, nil
}

func InitializeMainGetAttractionController() (inf.MainGetAttractionController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	attractionDal := dal.NewGetAttractionDal(mainMongoDB)
	attractionService := service.NewGetAttractionService(attractionDal)
	attractionController := controllers.NewGetAttractionController(attractionService)
	mainAttractionController := main_controllers.NewMainGetAttractionController(attractionController)
	return mainAttractionController, nil
}

func InitializeMainGetAttractionByIdController() (inf.MainGetAttractionByIdController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	attractionDal := dal.NewGetAttractionByIdDal(mainMongoDB)
	attractionService := service.NewGetAttractionByIdService(attractionDal)
	attractionController := controllers.NewGetAttractionByIdController(attractionService)
	mainAttractionController := main_controllers.NewMainGetAttractionByIdController(attractionController)
	return mainAttractionController, nil
}

func InitializeMainUpdateAttractionController() (inf.MainUpdateAttractionController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	attractionDal := dal.NewUpdateAttractionDal(mainMongoDB)
	attractionService := service.NewUpdateAttractionService(attractionDal)
	attractionController := controllers.NewUpdateAttractionController(attractionService)
	mainAttractionController := main_controllers.NewMainUpdateAttractionController(attractionController)
	return mainAttractionController, nil
}

func InitializeMainDeleteAttractionController() (inf.MainDeleteAttractionController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	attractionDal := dal.NewDeleteAttractionDal(mainMongoDB)
	attractionService := service.NewDeleteAttractionService(attractionDal)
	attractionController := controllers.NewDeleteAttractionController(attractionService)
	mainAttractionController := main_controllers.NewMainDeleteAttractionController(attractionController)
	return mainAttractionController, nil
}

