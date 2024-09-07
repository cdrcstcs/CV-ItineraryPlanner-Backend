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
func NewCreateItineraryDal(mainDB *db.MainMongoDB) inf.CreateItineraryDal{
	return &ItineraryDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}
func NewGetItineraryByIdDal(mainDB *db.MainMongoDB) inf.GetItineraryByIdDal{
	return &ItineraryDal{
		MainDB:(*mongo.Database)(mainDB),
	}
}
func NewGetItineraryDal(mainDB *db.MainMongoDB) inf.GetItineraryDal{
	return &ItineraryDal{
		MainDB:(*mongo.Database)(mainDB),
	}
}
func NewUpdateItineraryDal(mainDB *db.MainMongoDB) inf.UpdateItineraryDal {
	return &ItineraryDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}

func NewDeleteItineraryDal(mainDB *db.MainMongoDB) inf.DeleteItineraryDal {
	return &ItineraryDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}
type ItineraryDal struct {
	MainDB *mongo.Database
}

func (i *ItineraryDal) GetItineraryById(ctx context.Context, itineraryId string) (*models.Itinerary, error) {
	if utils.IsEmpty(itineraryId) {
		// TODO logging here
		return nil, custom_errs.DBErrCreateWithID
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
		// Handle cursor error
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
	itinerary.Id = insertedID.String()

	return itinerary, nil
}

func (i *ItineraryDal) UpdateItinerary(ctx context.Context, itinerary *models.Itinerary) (*models.Itinerary, error) {
	if utils.IsEmpty(itinerary.Id) {
		// TODO logging here
		return nil, custom_errs.DBErrUpdateWithID
	}

	collection := i.MainDB.Collection(constant.ItineraryTable)
	ObjectID, err := primitive.ObjectIDFromHex(itinerary.Id)
	if err != nil {
		return nil, custom_errs.DBErrIDConversion
	}

	filter := bson.M{"_id": ObjectID}
	update := bson.M{"$set": itinerary}

	opts := options.Update().SetUpsert(false)

	_, err = collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, custom_errs.DBErrUpdateWithID
	}

	return itinerary, nil
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
