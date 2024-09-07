package initialize

import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal/searchEngine"
	"itineraryplanner/service"
	"itineraryplanner/main_controllers"
	"itineraryplanner/main_controllers/inf"

)

func InitializeMainSearchEngineController() (inf.MainSearchEngineController, error) {
	Dal := searchEngine.NewSearchEngineDal()
	Service := service.NewSearchEngineService(Dal)
	Controller := controllers.NewSearchEngineController(Service)
	mainSController := main_controllers.NewMainSearchEngineController(Controller)
	return mainSController, nil
}