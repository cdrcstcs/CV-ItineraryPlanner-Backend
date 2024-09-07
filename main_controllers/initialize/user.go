package initialize

import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/main_controllers"
	"itineraryplanner/main_controllers/inf"

)

func InitializeMainUserController() (inf.MainUserController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	userDal := dal.NewUserDal(mainMongoDB)
	userService := service.NewUserService(userDal)
	userController := controllers.NewUserController(userService)
	mainUserController := main_controllers.NewMainUserController(userController)
	return mainUserController, nil
}