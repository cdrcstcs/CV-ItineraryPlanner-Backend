package initialize

import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/main_controllers"
	"itineraryplanner/main_controllers/inf"

)

func InitializeMainCreateEventController() (inf.MainCreateEventController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	eventDal := dal.NewCreateEventDal(mainMongoDB)
	eventService := service.NewCreateEventService(eventDal)
	eventController := controllers.NewCreateEventController(eventService)
	mainEventController := main_controllers.NewMainCreateEventController(eventController)
	return mainEventController, nil
}

func InitializeMainGetEventController() (inf.MainGetEventController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	eventDal := dal.NewGetEventDal(mainMongoDB)
	eventService := service.NewGetEventService(eventDal)
	eventController := controllers.NewGetEventController(eventService)
	mainEventController := main_controllers.NewMainGetEventController(eventController)
	return mainEventController, nil
}
func InitializeMainGetEventByIdController() (inf.MainGetEventByIdController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	EventDal := dal.NewGetEventByIdDal(mainMongoDB)
	EventService := service.NewGetEventByIdService(EventDal)
	EventController := controllers.NewGetEventByIdController(EventService)
	mainEventController := main_controllers.NewMainGetEventByIdController(EventController)
	return mainEventController, nil
}

func InitializeMainUpdateEventController() (inf.MainUpdateEventController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	eventDal := dal.NewUpdateEventDal(mainMongoDB)
	eventService := service.NewUpdateEventService(eventDal)
	eventController := controllers.NewUpdateEventController(eventService)
	mainEventController := main_controllers.NewMainUpdateEventController(eventController)
	return mainEventController, nil
}

func InitializeMainDeleteEventController() (inf.MainDeleteEventController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	eventDal := dal.NewDeleteEventDal(mainMongoDB)
	eventService := service.NewDeleteEventService(eventDal)
	eventController := controllers.NewDeleteEventController(eventService)
	mainEventController := main_controllers.NewMainDeleteEventController(eventController)
	return mainEventController, nil
}

