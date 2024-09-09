package initialize
import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/maincontrollers"
	"itineraryplanner/maincontrollers/inf"

)
func InitializeMainRatingController() (inf.MainRatingController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	RatingDal := dal.NewRatingDal(mainMongoDB)
	RatingService := service.NewRatingService(RatingDal)
	RatingController := controllers.NewRatingController(RatingService)
	mainRatingController := maincontrollers.NewMainRatingController(RatingController)
	return mainRatingController, nil
}