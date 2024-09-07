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

func NewCreateAttractionDal(mainDB *db.MainMongoDB) inf.CreateAttractionDal {
	return &AttractionDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}

func NewGetAttractionByIdDal(mainDB *db.MainMongoDB) inf.GetAttractionByIdDal {
	return &AttractionDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}
func NewGetAttractionDal(mainDB *db.MainMongoDB) inf.GetAttractionDal {
	return &AttractionDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}

func NewUpdateAttractionDal(mainDB *db.MainMongoDB) inf.UpdateAttractionDal {
	return &AttractionDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}

func NewDeleteAttractionDal(mainDB *db.MainMongoDB) inf.DeleteAttractionDal {
	return &AttractionDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}

type AttractionDal struct {
	MainDB *mongo.Database
}


func (a *AttractionDal) GetAttractionById(ctx context.Context, attractionId string) (*models.Attraction, error) {
	if utils.IsEmpty(attractionId) {
		// TODO logging here
		return nil, custom_errs.DBErrCreateWithID
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
		// Handle cursor error
		return nil, custom_errs.DBErrGetWithID
	}

	return attractions, nil
}

func (a *AttractionDal) CreateAttraction(ctx context.Context, attraction *models.Attraction) (*models.Attraction, error) {
	if !utils.IsEmpty(attraction.Id) {
		// TODO logging here
		return nil, custom_errs.DBErrCreateWithID
	}
	attractionCollection := a.MainDB.Collection(constant.AttractionTable)
		// Insert newAttraction into MongoDB
	result, err := attractionCollection.InsertOne(ctx, attraction)

	// Access the generated IDs if needed

	if err != nil {
		// TODO logging here
		return nil, custom_errs.DbErrors
	}

	// Extract inserted ID
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		// TODO logging here
		return nil, custom_errs.DBErrIDConversion
	}

	attraction.Id = insertedID.String()
	return attraction, nil
}

func (a *AttractionDal) UpdateAttraction(ctx context.Context, attraction *models.Attraction) (*models.Attraction, error) {
	if utils.IsEmpty(attraction.Id) {
		// TODO logging here
		return nil, custom_errs.DBErrUpdateWithID
	}

	collection := a.MainDB.Collection(constant.AttractionTable)
	ObjectID, err := primitive.ObjectIDFromHex(attraction.Id)
	if err != nil {
		return nil, custom_errs.DBErrIDConversion
	}

	filter := bson.M{"_id": ObjectID}
	update := bson.M{"$set": attraction}

	opts := options.Update().SetUpsert(false) 

	_, err = collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, custom_errs.DBErrUpdateWithID
	}

	result:= collection.FindOne(ctx, bson.M{"_id":ObjectID})  
	if result.Err() != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	var Attraction1 models.Attraction
	if err := result.Decode(Attraction1); err != nil {
		return nil, custom_errs.DecodeErr
	}
	return &Attraction1, nil
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
