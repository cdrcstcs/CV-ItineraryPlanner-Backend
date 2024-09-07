package initialize

import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/main_controllers"
	"itineraryplanner/main_controllers/inf"

)

func InitializeMainCoordinateController() (inf.MainCoordinateController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	coordinateDal := dal.NewCoordinateDal(mainMongoDB)
	coordinateService := service.NewCoordinateService(coordinateDal)
	coordinateController := controllers.NewCoordinateController(coordinateService)
	mainCoordinateController := main_controllers.NewMainCoordinateController(coordinateController)
	return mainCoordinateController, nil
}