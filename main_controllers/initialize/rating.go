package initialize

import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/main_controllers"
	"itineraryplanner/main_controllers/inf"

)

func InitializeMainCreateRatingController() (inf.MainCreateRatingController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	ratingDal := dal.NewCreateRatingDal(mainMongoDB)
	ratingService := service.NewCreateRatingService(ratingDal)
	ratingController := controllers.NewCreateRatingController(ratingService)
	mainRatingController := main_controllers.NewMainCreateRatingController(ratingController)
	return mainRatingController, nil
}

func InitializeMainGetRatingController() (inf.MainGetRatingController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	ratingDal := dal.NewGetRatingDal(mainMongoDB)
	ratingService := service.NewGetRatingService(ratingDal)
	ratingController := controllers.NewGetRatingController(ratingService)
	mainRatingController := main_controllers.NewMainGetRatingController(ratingController)
	return mainRatingController, nil
}

func InitializeMainGetRatingByIdController() (inf.MainGetRatingByIdController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	RatingDal := dal.NewGetRatingByIdDal(mainMongoDB)
	RatingService := service.NewGetRatingByIdService(RatingDal)
	RatingController := controllers.NewGetRatingByIdController(RatingService)
	mainRatingController := main_controllers.NewMainGetRatingByIdController(RatingController)
	return mainRatingController, nil
}

func InitializeMainUpdateRatingController() (inf.MainUpdateRatingController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	ratingDal := dal.NewUpdateRatingDal(mainMongoDB)
	ratingService := service.NewUpdateRatingService(ratingDal)
	ratingController := controllers.NewUpdateRatingController(ratingService)
	mainRatingController := main_controllers.NewMainUpdateRatingController(ratingController)
	return mainRatingController, nil
}

func InitializeMainDeleteRatingController() (inf.MainDeleteRatingController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	ratingDal := dal.NewDeleteRatingDal(mainMongoDB)
	ratingService := service.NewDeleteRatingService(ratingDal)
	ratingController := controllers.NewDeleteRatingController(ratingService)
	mainRatingController := main_controllers.NewMainDeleteRatingController(ratingController)
	return mainRatingController, nil
}

