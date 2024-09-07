package initialize

import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/main_controllers"
	"itineraryplanner/main_controllers/inf"

)

func InitializeMainCreateTagController() (inf.MainCreateTagController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	tagDal := dal.NewCreateTagDal(mainMongoDB)
	tagService := service.NewCreateTagService(tagDal)
	tagController := controllers.NewCreateTagController(tagService)
	mainTagController := main_controllers.NewMainCreateTagController(tagController)
	return mainTagController, nil
}

func InitializeMainGetTagController() (inf.MainGetTagController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	tagDal := dal.NewGetTagDal(mainMongoDB)
	tagService := service.NewGetTagService(tagDal)
	tagController := controllers.NewGetTagController(tagService)
	mainTagController := main_controllers.NewMainGetTagController(tagController)
	return mainTagController, nil
}
func InitializeMainGetTagByIdController() (inf.MainGetTagByIdController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	TagDal := dal.NewGetTagByIdDal(mainMongoDB)
	TagService := service.NewGetTagByIdService(TagDal)
	TagController := controllers.NewGetTagByIdController(TagService)
	mainTagController := main_controllers.NewMainGetTagByIdController(TagController)
	return mainTagController, nil
}
func InitializeMainUpdateTagController() (inf.MainUpdateTagController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	tagDal := dal.NewUpdateTagDal(mainMongoDB)
	tagService := service.NewUpdateTagService(tagDal)
	tagController := controllers.NewUpdateTagController(tagService)
	mainTagController := main_controllers.NewMainUpdateTagController(tagController)
	return mainTagController, nil
}

func InitializeMainDeleteTagController() (inf.MainDeleteTagController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	tagDal := dal.NewDeleteTagDal(mainMongoDB)
	tagService := service.NewDeleteTagService(tagDal)
	tagController := controllers.NewDeleteTagController(tagService)
	mainTagController := main_controllers.NewMainDeleteTagController(tagController)
	return mainTagController, nil
}

