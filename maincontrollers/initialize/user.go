package initialize
import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/maincontrollers"
	"itineraryplanner/maincontrollers/inf"

)
func InitializeMainUserController() (inf.MainUserController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	UserDal := dal.NewUserDal(mainMongoDB)
	UserService := service.NewUserService(UserDal)
	UserController := controllers.NewUserController(UserService)
	mainUserController := maincontrollers.NewMainUserController(UserController)
	return mainUserController, nil
}