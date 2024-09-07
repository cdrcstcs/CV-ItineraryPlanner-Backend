package dal
import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"itineraryplanner/common/custom_errs"
	"itineraryplanner/common/utils"
	"itineraryplanner/constant"
	"itineraryplanner/dal/db"
	"itineraryplanner/dal/inf"
	"itineraryplanner/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func NewAttractionDal(mainDB *db.MainMongoDB) inf.AttractionDal {
	return &AttractionDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}
type AttractionDal struct {
	MainDB *mongo.Database
}
func (a *AttractionDal) GetAttractionById(ctx context.Context, attractionId string) (*models.Attraction, error) {
	if utils.IsEmpty(attractionId) {
		return nil, custom_errs.DBErrGetWithID
	}
	collection := a.MainDB.Collection(constant.AttractionTable)
	ObjectID, err := primitive.ObjectIDFromHex(attractionId)
	if err != nil {
		return nil, custom_errs.DBErrIDConversion
	}
	result := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	if result.Err() != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	var attraction *models.Attraction
	if err := result.Decode(&attraction); err != nil {
		return nil, custom_errs.DecodeErr
	}
	return attraction, nil
}
func (a *AttractionDal) GetDB() *mongo.Database{
	return a.MainDB
}
func (a *AttractionDal) GetAttraction(ctx context.Context) ([]*models.Attraction, error) {
	collection := a.MainDB.Collection(constant.AttractionTable)
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	defer result.Close(ctx)
	var attractions []*models.Attraction
	for result.Next(ctx) {
		var attraction *models.Attraction
		if err := result.Decode(&attraction); err != nil {
			return nil, custom_errs.DecodeErr
		}
		attractions = append(attractions, attraction)
	}
	if err := result.Err(); err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	return attractions, nil
}
func (a *AttractionDal) CreateAttraction(ctx context.Context, attraction *models.Attraction) (*models.Attraction, error) {
	if !utils.IsEmpty(attraction.Id) {
		return nil, custom_errs.DBErrCreateWithID
	}
	attractionCollection := a.MainDB.Collection(constant.AttractionTable)
	result, err := attractionCollection.InsertOne(ctx, attraction)
	if err != nil {
		return nil, custom_errs.DbErrors
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, custom_errs.DBErrIDConversion
	}
	attraction.Id = insertedID.Hex()
	return attraction, nil
}
func (a *AttractionDal) UpdateAttraction(ctx context.Context, attraction *models.Attraction) (*models.Attraction, error) {
    if utils.IsEmpty(attraction.Id) {
        return nil, custom_errs.DBErrUpdateWithID
    }
    collection := a.MainDB.Collection(constant.AttractionTable)
    ObjectID, err := primitive.ObjectIDFromHex(attraction.Id)
    if err != nil {
        return nil, custom_errs.DBErrIDConversion
    }
    filter := bson.M{"_id": ObjectID}
    attractionBson, err := bson.Marshal(attraction)
    if err != nil {
        return nil, custom_errs.DecodeErr
    }
    var updateDoc bson.M
    if err := bson.Unmarshal(attractionBson, &updateDoc); err != nil {
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
    var attractionUpdated models.Attraction
    if err := result.Decode(&attractionUpdated); err != nil {
        return nil, custom_errs.DecodeErr
    }
    return &attractionUpdated, nil
}
func (a *AttractionDal) DeleteAttraction(ctx context.Context, attractionId string) (*models.Attraction, error) {
	if utils.IsEmpty(attractionId) {
		return nil, custom_errs.DBErrDeleteWithID
	}
	collection := a.MainDB.Collection(constant.AttractionTable)
	ObjectID, err := primitive.ObjectIDFromHex(attractionId)
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
	var attraction1 *models.Attraction
	if err := result.Decode(&attraction1); err != nil {
		return nil, custom_errs.DecodeErr
	}
	return attraction1, nil
}