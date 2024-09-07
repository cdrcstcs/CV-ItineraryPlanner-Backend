package initialize

import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/main_controllers"
	"itineraryplanner/main_controllers/inf"

)

func InitializeMainEventController() (inf.MainEventController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	eventDal := dal.NewEventDal(mainMongoDB)
	eventService := service.NewEventService(eventDal)
	eventController := controllers.NewEventController(eventService)
	mainEventController := main_controllers.NewMainEventController(eventController)
	return mainEventController, nil
}
