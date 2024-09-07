//go:build wireinject
// +build wireinject
package main_controllers
import (
	"github.com/google/wire"

	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
)
//go:generate

func InitializeMainCreateAttractionController() (MainAttractionController, error) {
	wire.Build(
		db.GetMainMongoDatabase,

		dal.NewCreateAttractionDal,

		service.NewCreateAttractionService,

		controllers.NewCreateAttractionController,

		NewMainCreateAttractionController,
	)
	return MainAttractionController{}, nil
}
