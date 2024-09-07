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
func NewCreateCoordinateDal(mainDB *db.MainMongoDB) inf.CreateCoordinateDal{
	return &CoordinateDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}
func NewGetCoordinateByIdDal(mainDB *db.MainMongoDB) inf.GetCoordinateByIdDal{
	return &CoordinateDal{
		MainDB:(*mongo.Database)(mainDB),
	}
}
func NewGetCoordinateDal(mainDB *db.MainMongoDB) inf.GetCoordinateDal{
	return &CoordinateDal{
		MainDB:(*mongo.Database)(mainDB),
	}
}
func NewUpdateCoordinateDal(mainDB *db.MainMongoDB) inf.UpdateCoordinateDal {
	return &CoordinateDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}

func NewDeleteCoordinateDal(mainDB *db.MainMongoDB) inf.DeleteCoordinateDal {
	return &CoordinateDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}

type CoordinateDal struct {
	MainDB *mongo.Database
}

func (c *CoordinateDal) GetCoordinateById(ctx context.Context, coordinateId string) (*models.Coordinate, error) {
	if utils.IsEmpty(coordinateId) {
		// TODO logging here
		return nil, custom_errs.DBErrCreateWithID
	}
	collection := c.MainDB.Collection(constant.CoordinateTable)
	ObjectID, err := primitive.ObjectIDFromHex(coordinateId)
	if err != nil {
		return nil, custom_errs.DBErrIDConversion
	}

	result := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	if result.Err() != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	var coordinate *models.Coordinate
	if err := result.Decode(&coordinate); err != nil {
		return nil, custom_errs.DecodeErr
	}

	return coordinate, nil
}
func (c *CoordinateDal) GetCoordinate(ctx context.Context) ([]*models.Coordinate, error) {
	collection := c.MainDB.Collection(constant.CoordinateTable)

	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	defer result.Close(ctx)

	var coordinates []*models.Coordinate
	for result.Next(ctx) {
		var coordinate *models.Coordinate
		if err := result.Decode(&coordinate); err != nil {
			return nil, custom_errs.DecodeErr
		}
		coordinates = append(coordinates, coordinate)
	}

	if err := result.Err(); err != nil {
		// Handle cursor error
		return nil, custom_errs.DBErrGetWithID
	}

	return coordinates, nil
}

func (c *CoordinateDal) CreateCoordinate(ctx context.Context, coordinate *models.Coordinate) (*models.Coordinate, error) {
	if !utils.IsEmpty(coordinate.Id){
		return nil, custom_errs.DBErrCreateWithID
	}
	collection := c.MainDB.Collection(constant.CoordinateTable)
	result, err := collection.InsertOne(ctx, coordinate)
	if err != nil {
		return nil, custom_errs.DbErrors
	}
	insertedID, ok:=result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, custom_errs.DBErrIDConversion
	}
	coordinate.Id = insertedID.String()
	return coordinate, nil

}

func (c *CoordinateDal) UpdateCoordinate(ctx context.Context, coordinate *models.Coordinate) (*models.Coordinate, error) {
	if utils.IsEmpty(coordinate.Id) {
		// TODO logging here
		return nil, custom_errs.DBErrUpdateWithID
	}

	collection := c.MainDB.Collection(constant.CoordinateTable)
	ObjectID, err := primitive.ObjectIDFromHex(coordinate.Id)
	if err != nil {
		return nil, custom_errs.DBErrIDConversion
	}

	filter := bson.M{"_id": ObjectID}
	update := bson.M{"$set": coordinate}

	opts := options.Update().SetUpsert(false)

	_, err = collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, custom_errs.DBErrUpdateWithID
	}

	return coordinate, nil
}


func (c *CoordinateDal) DeleteCoordinate(ctx context.Context, coordinateId string) (*models.Coordinate, error) {
	if utils.IsEmpty(coordinateId) {
		return nil, custom_errs.DBErrDeleteWithID
	}
	collection := c.MainDB.Collection(constant.CoordinateTable)

	ObjectID, err := primitive.ObjectIDFromHex(coordinateId)
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

	var coordinate1 *models.Coordinate
	if err := result.Decode(&coordinate1); err != nil {
		return nil, custom_errs.DecodeErr
	}

	return coordinate1, nil
}
