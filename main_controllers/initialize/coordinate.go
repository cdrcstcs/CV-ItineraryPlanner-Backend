package initialize

import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/main_controllers"
	"itineraryplanner/main_controllers/inf"

)

func InitializeMainCreateCoordinateController() (inf.MainCreateCoordinateController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	coordinateDal := dal.NewCreateCoordinateDal(mainMongoDB)
	coordinateService := service.NewCreateCoordinateService(coordinateDal)
	coordinateController := controllers.NewCreateCoordinateController(coordinateService)
	mainCoordinateController := main_controllers.NewMainCreateCoordinateController(coordinateController)
	return mainCoordinateController, nil
}

func InitializeMainGetCoordinateController() (inf.MainGetCoordinateController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	coordinateDal := dal.NewGetCoordinateDal(mainMongoDB)
	coordinateService := service.NewGetCoordinateService(coordinateDal)
	coordinateController := controllers.NewGetCoordinateController(coordinateService)
	mainCoordinateController := main_controllers.NewMainGetCoordinateController(coordinateController)
	return mainCoordinateController, nil
}

func InitializeMainGetCoordinateByIdController() (inf.MainGetCoordinateByIdController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	CoordinateDal := dal.NewGetCoordinateByIdDal(mainMongoDB)
	CoordinateService := service.NewGetCoordinateByIdService(CoordinateDal)
	CoordinateController := controllers.NewGetCoordinateByIdController(CoordinateService)
	mainCoordinateController := main_controllers.NewMainGetCoordinateByIdController(CoordinateController)
	return mainCoordinateController, nil
}

func InitializeMainUpdateCoordinateController() (inf.MainUpdateCoordinateController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	coordinateDal := dal.NewUpdateCoordinateDal(mainMongoDB)
	coordinateService := service.NewUpdateCoordinateService(coordinateDal)
	coordinateController := controllers.NewUpdateCoordinateController(coordinateService)
	mainCoordinateController := main_controllers.NewMainUpdateCoordinateController(coordinateController)
	return mainCoordinateController, nil
}

func InitializeMainDeleteCoordinateController() (inf.MainDeleteCoordinateController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	coordinateDal := dal.NewDeleteCoordinateDal(mainMongoDB)
	coordinateService := service.NewDeleteCoordinateService(coordinateDal)
	coordinateController := controllers.NewDeleteCoordinateController(coordinateService)
	mainCoordinateController := main_controllers.NewMainDeleteCoordinateController(coordinateController)
	return mainCoordinateController, nil
}

