package initialize
import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/maincontrollers"
	"itineraryplanner/maincontrollers/inf"

)
func InitializeMainTagController() (inf.MainTagController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	TagDal := dal.NewTagDal(mainMongoDB)
	TagService := service.NewTagService(TagDal)
	TagController := controllers.NewTagController(TagService)
	mainTagController := maincontrollers.NewMainTagController(TagController)
	return mainTagController, nil
}