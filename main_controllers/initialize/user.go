package initialize

import (
	"itineraryplanner/controllers"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/service"
	"itineraryplanner/main_controllers"
	"itineraryplanner/main_controllers/inf"

)

func InitializeMainCreateUserController() (inf.MainCreateUserController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	userDal := dal.NewCreateUserDal(mainMongoDB)
	userService := service.NewCreateUserService(userDal)
	userController := controllers.NewCreateUserController(userService)
	mainUserController := main_controllers.NewMainCreateUserController(userController)
	return mainUserController, nil
}

func InitializeMainGetUserController() (inf.MainGetUserController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	userDal := dal.NewGetUserDal(mainMongoDB)
	userService := service.NewGetUserService(userDal)
	userController := controllers.NewGetUserController(userService)
	mainUserController := main_controllers.NewMainGetUserController(userController)
	return mainUserController, nil
}

func InitializeMainGetUserByIdController() (inf.MainGetUserByIdController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	UserDal := dal.NewGetUserByIdDal(mainMongoDB)
	UserService := service.NewGetUserByIdService(UserDal)
	UserController := controllers.NewGetUserByIdController(UserService)
	mainUserController := main_controllers.NewMainGetUserByIdController(UserController)
	return mainUserController, nil
}

func InitializeMainUpdateUserController() (inf.MainUpdateUserController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	userDal := dal.NewUpdateUserDal(mainMongoDB)
	userService := service.NewUpdateUserService(userDal)
	userController := controllers.NewUpdateUserController(userService)
	mainUserController := main_controllers.NewMainUpdateUserController(userController)
	return mainUserController, nil
}

func InitializeMainDeleteUserController() (inf.MainDeleteUserController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	userDal := dal.NewDeleteUserDal(mainMongoDB)
	userService := service.NewDeleteUserService(userDal)
	userController := controllers.NewDeleteUserController(userService)
	mainUserController := main_controllers.NewMainDeleteUserController(userController)
	return mainUserController, nil
}

func InitializeMainLoginUserController() (inf.MainLoginUserController, error) {
	mainMongoDB := db.GetMainMongoDatabase()
	userDal := dal.NewLoginUserDal(mainMongoDB)
	userService := service.NewLoginUserService(userDal)
	userController := controllers.NewLoginUserController(userService)
	mainUserController := main_controllers.NewMainLoginUserController(userController)
	return mainUserController, nil
}

