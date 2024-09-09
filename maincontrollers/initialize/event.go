package initialize
import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/maincontrollers"
	"itineraryplanner/maincontrollers/inf"

)
func InitializeMainEventController() (inf.MainEventController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	EventDal := dal.NewEventDal(mainMongoDB)
	EventService := service.NewEventService(EventDal)
	EventController := controllers.NewEventController(EventService)
	mainEventController := maincontrollers.NewMainEventController(EventController)
	return mainEventController, nil
}