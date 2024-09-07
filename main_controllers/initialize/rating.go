package initialize

import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/main_controllers"
	"itineraryplanner/main_controllers/inf"

)

func InitializeMainRatingController() (inf.MainRatingController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	ratingDal := dal.NewRatingDal(mainMongoDB)
	ratingService := service.NewRatingService(ratingDal)
	ratingController := controllers.NewRatingController(ratingService)
	mainRatingController := main_controllers.NewMainRatingController(ratingController)
	return mainRatingController, nil
}
