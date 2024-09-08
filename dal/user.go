package dal
import (
	"context"
	"itineraryplanner/common/custom_errs"
	"itineraryplanner/common/utils"
	"itineraryplanner/constant"
	"itineraryplanner/dal/db"
	"itineraryplanner/dal/inf"
	"itineraryplanner/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func NewUserDal(mainDB *db.MainMongoDB) inf.UserDal {
	return &UserDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}
type UserDal struct {
	MainDB *mongo.Database
}
func (u *UserDal) GetDB() *mongo.Database {
	return u.MainDB
}
func (u *UserDal) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	if utils.IsEmpty(userId) {
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
		return nil, custom_errs.DBErrGetWithID
	}
	return users, nil
}
func (t *UserDal) CreateUser(ctx context.Context, User *models.User) (*models.User, error) {
	if !utils.IsEmpty(User.Id) {
		return nil, custom_errs.DBErrCreateWithID
	}
	collection := t.MainDB.Collection(constant.UserTable)
	result, err := collection.InsertOne(ctx, User)
	if err != nil {
		return nil, custom_errs.DbErrors
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, custom_errs.DBErrIDConversion
	}
	User.Id = insertedID.Hex()
	return User, nil
}
func (u *UserDal) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
    if utils.IsEmpty(user.Id) {
        return nil, custom_errs.DBErrUpdateWithID
    }
    collection := u.MainDB.Collection(constant.UserTable)
    ObjectID, err := primitive.ObjectIDFromHex(user.Id)
    if err != nil {
        return nil, custom_errs.DBErrIDConversion
    }
    filter := bson.M{"_id": ObjectID}
    userBson, err := bson.Marshal(user)
    if err != nil {
        return nil, custom_errs.DecodeErr
    }
    var updateDoc bson.M
    if err := bson.Unmarshal(userBson, &updateDoc); err != nil {
        return nil, custom_errs.DecodeErr
    }
    delete(updateDoc, "_id")
    update := bson.M{"$set": updateDoc}
    opts := options.Update().SetUpsert(false)
    ret, err := collection.UpdateOne(ctx, filter, update, opts)
    if err != nil {
        return nil, custom_errs.DBErrUpdateWithID
    }
    if ret.ModifiedCount == 0 {
        return nil, custom_errs.DBErrUpdateWithID
    }
    result := collection.FindOne(ctx, filter)
    if result.Err() != nil {
        return nil, custom_errs.DBErrGetWithID
    }
    var userUpdated models.User
    if err := result.Decode(&userUpdated); err != nil {
        return nil, custom_errs.DecodeErr
    }
    return &userUpdated, nil
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