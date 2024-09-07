package dal

import (
	"context"
	"itineraryplanner/dal/db"

	"itineraryplanner/dal/inf"

	"go.mongodb.org/mongo-driver/mongo"
	"itineraryplanner/models"
	"itineraryplanner/common/utils"
	"itineraryplanner/common/custom_errs"
	"itineraryplanner/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

)
func NewCreateUserDal(mainDB *db.MainMongoDB) inf.CreateUserDal{
	return &UserDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}
func NewGetUserByIdDal(mainDB *db.MainMongoDB) inf.GetUserByIdDal{
	return &UserDal{
		MainDB:(*mongo.Database)(mainDB),
	}
}
func NewGetUserDal(mainDB *db.MainMongoDB) inf.GetUserDal{
	return &UserDal{
		MainDB:(*mongo.Database)(mainDB),
	}
}
func NewUpdateUserDal(mainDB *db.MainMongoDB) inf.UpdateUserDal {
	return &UserDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}

func NewDeleteUserDal(mainDB *db.MainMongoDB) inf.DeleteUserDal {
	return &UserDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}

func NewLoginUserDal(mainDB *db.MainMongoDB) inf.LoginUserDal {
	return &UserDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}
type UserDal struct {
	MainDB *mongo.Database
}
func (u *UserDal) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	if utils.IsEmpty(userId) {
		// TODO logging here
		return nil, custom_errs.DBErrCreateWithID
	}
	collection := u.MainDB.Collection(constant.UserTable)
	ObjectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, custom_errs.DBErrIDConversion
	}

	result := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	if result.Err() != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	var user *models.User
	if err := result.Decode(&user); err != nil {
		return nil, custom_errs.DecodeErr
	}

	return user, nil
}
func (u *UserDal) GetUser(ctx context.Context) ([]*models.User, error) {
	collection := u.MainDB.Collection(constant.UserTable)

	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	defer result.Close(ctx)

	var users []*models.User
	for result.Next(ctx) {
		var user *models.User
		if err := result.Decode(&user); err != nil {
			return nil, custom_errs.DecodeErr
		}
		users = append(users, user)
	}

	if err := result.Err(); err != nil {
		// Handle cursor error
		return nil, custom_errs.DBErrGetWithID
	}

	return users, nil
}

func (u *UserDal) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if !utils.IsEmpty(user.Id){
		return nil, custom_errs.DBErrCreateWithID
	}
	collection := u.MainDB.Collection(constant.UserTable)
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, custom_errs.DbErrors
	}
	insertedID, ok:=result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, custom_errs.DBErrIDConversion
	}
	user.Id = insertedID.String()
	return user, nil
}

func (u *UserDal) LoginUser(ctx context.Context, user *models.User) (bool, error) {
	if !utils.IsEmpty(user.Id){
		return false, custom_errs.DBErrCreateWithID
	}
	collection := u.MainDB.Collection(constant.UserTable)
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return false, custom_errs.DbErrors
	}
	defer result.Close(ctx)
	if err := result.Err(); err != nil {
		// Handle cursor error
		return false, custom_errs.DBErrGetWithID
	}
	for result.Next(ctx) {
		var user1 *models.User
		if err := result.Decode(&user1); err != nil {
			return false, custom_errs.DecodeErr
		}
		if user1.Name == user.Name && user1.Password == user.Password && user1.Email == user.Email && user1.EmailPassword == user.EmailPassword {
			return true, nil
		}
	}
	return false, custom_errs.DbErrors
}
func (u *UserDal) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if utils.IsEmpty(user.Id) {
		// TODO logging here
		return nil, custom_errs.DBErrUpdateWithID
	}

	collection := u.MainDB.Collection(constant.UserTable)
	ObjectID, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return nil, custom_errs.DBErrIDConversion
	}

	filter := bson.M{"_id": ObjectID}
	update := bson.M{"$set": user}

	opts := options.Update().SetUpsert(false)

	_, err = collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, custom_errs.DBErrUpdateWithID
	}

	return user, nil
}


func (u *UserDal) DeleteUser(ctx context.Context, userId string) (*models.User, error) {
	if utils.IsEmpty(userId) {
		return nil, custom_errs.DBErrDeleteWithID
	}
	collection := u.MainDB.Collection(constant.UserTable)

	ObjectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, custom_errs.DBErrIDConversion
	}

	result := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	if result.Err() != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id": ObjectID})
	if err != nil {
		return nil, custom_errs.DBErrDeleteWithID
	}

	var user1 *models.User
	if err := result.Decode(&user1); err != nil {
		return nil, custom_errs.DecodeErr
	}

	return user1, nil
}
