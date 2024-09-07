package initialize

import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/main_controllers"
	"itineraryplanner/main_controllers/inf"

)

func InitializeMainTagController() (inf.MainTagController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	tagDal := dal.NewTagDal(mainMongoDB)
	tagService := service.NewTagService(tagDal)
	tagController := controllers.NewTagController(tagService)
	mainTagController := main_controllers.NewMainTagController(tagController)
	return mainTagController, nil
}
