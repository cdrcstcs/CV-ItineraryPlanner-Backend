package dal
import (
	"context"
	"itineraryplanner/dal/inf"
	"go.mongodb.org/mongo-driver/mongo"
	"itineraryplanner/models"
	"itineraryplanner/common/utils"
	"itineraryplanner/common/custom_errs"
	"itineraryplanner/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"itineraryplanner/dal/db"
)
func NewItineraryDal(mainDB *db.MainMongoDB) inf.ItineraryDal{
	return &ItineraryDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}
type ItineraryDal struct {
	MainDB *mongo.Database
}
func (i *ItineraryDal) GetDB() *mongo.Database{
	return i.MainDB
}
func (i *ItineraryDal) GetItineraryById(ctx context.Context, itineraryId string) (*models.Itinerary, error) {
	if utils.IsEmpty(itineraryId) {
		return nil, custom_errs.DBErrGetWithID
	}
	collection := i.MainDB.Collection(constant.ItineraryTable)
	ObjectID, err := primitive.ObjectIDFromHex(itineraryId)
	if err != nil {
		return nil, custom_errs.DBErrIDConversion
	}
	result := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	if result.Err() != nil {
		return nil, result.Err()
	}
	var itinerary *models.Itinerary
	if err := result.Decode(&itinerary); err != nil {
		return nil, custom_errs.DecodeErr
	}
	return itinerary, nil
}
func (i *ItineraryDal) GetItinerary(ctx context.Context) ([]*models.Itinerary, error) {
	collection := i.MainDB.Collection(constant.ItineraryTable)
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	defer result.Close(ctx)
	var itineraries []*models.Itinerary
	for result.Next(ctx) {
		var itinerary *models.Itinerary
		if err := result.Decode(&itinerary); err != nil {
			return nil, custom_errs.DecodeErr
		}
		itineraries = append(itineraries, itinerary)
	}
	if err := result.Err(); err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	return itineraries, nil
}
func (i *ItineraryDal) CreateItinerary(ctx context.Context, itinerary *models.Itinerary) (*models.Itinerary, error) {
	if !utils.IsEmpty(itinerary.Id){
		return nil, custom_errs.DBErrCreateWithID
	}
	collection := i.MainDB.Collection(constant.ItineraryTable)
	result, err := collection.InsertOne(ctx, itinerary)
	if err != nil {
		return nil, custom_errs.DbErrors
	}
	insertedID, ok:=result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, custom_errs.DBErrIDConversion
	}
	itinerary.Id = insertedID.Hex()
	return itinerary, nil
}
func (i *ItineraryDal) UpdateItinerary(ctx context.Context, itinerary *models.Itinerary) (*models.Itinerary, error) {
    if utils.IsEmpty(itinerary.Id) {
        return nil, custom_errs.DBErrUpdateWithID
    }
    collection := i.MainDB.Collection(constant.ItineraryTable)
    ObjectID, err := primitive.ObjectIDFromHex(itinerary.Id)
    if err != nil {
        return nil, custom_errs.DBErrIDConversion
    }
    filter := bson.M{"_id": ObjectID}
    itineraryBson, err := bson.Marshal(itinerary)
    if err != nil {
        return nil, custom_errs.DecodeErr
    }
    var updateDoc bson.M
    if err := bson.Unmarshal(itineraryBson, &updateDoc); err != nil {
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
    var itineraryUpdated models.Itinerary
    if err := result.Decode(&itineraryUpdated); err != nil {
        return nil, custom_errs.DecodeErr
    }
    return &itineraryUpdated, nil
}
func (i *ItineraryDal) DeleteItinerary(ctx context.Context, itineraryId string) (*models.Itinerary, error) {
	if utils.IsEmpty(itineraryId) {
		return nil, custom_errs.DBErrDeleteWithID
	}
	collection := i.MainDB.Collection(constant.ItineraryTable)
	ObjectID, err := primitive.ObjectIDFromHex(itineraryId)
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
	var itinerary1 *models.Itinerary
	if err := result.Decode(&itinerary1); err != nil {
		return nil, custom_errs.DecodeErr
	}
	return itinerary1, nil
}