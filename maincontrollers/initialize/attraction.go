package initialize
import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/maincontrollers"
	"itineraryplanner/maincontrollers/inf"

)
func InitializeMainAttractionController() (inf.MainAttractionController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	AttractionDal := dal.NewAttractionDal(mainMongoDB)
	AttractionService := service.NewAttractionService(AttractionDal)
	AttractionController := controllers.NewAttractionController(AttractionService)
	mainAttractionController := maincontrollers.NewMainAttractionController(AttractionController)
	return mainAttractionController, nil
}