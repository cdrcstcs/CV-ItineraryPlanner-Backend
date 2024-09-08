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
func NewRatingDal(mainDB *db.MainMongoDB) inf.RatingDal{
	return &RatingDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}
type RatingDal struct {
	MainDB *mongo.Database
}
func (r *RatingDal) GetDB() *mongo.Database{
	return r.MainDB
}
func (r *RatingDal) GetRatingById(ctx context.Context, ratingId string) (*models.Rating, error) {
	if utils.IsEmpty(ratingId) {
		return nil, custom_errs.DBErrCreateWithID
	}
	collection := r.MainDB.Collection(constant.RatingTable)
	ObjectID, err := primitive.ObjectIDFromHex(ratingId)
	if err != nil {
		return nil, custom_errs.DBErrIDConversion
	}
	result := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	if result.Err() != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	var rating *models.Rating
	if err := result.Decode(&rating); err != nil {
		return nil, custom_errs.DecodeErr
	}
	return rating, nil
}
func (r *RatingDal) GetRating(ctx context.Context) ([]*models.Rating, error) {
	collection := r.MainDB.Collection(constant.RatingTable)
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	defer result.Close(ctx)
	var ratings []*models.Rating
	for result.Next(ctx) {
		var rating *models.Rating
		if err := result.Decode(&rating); err != nil {
			return nil, custom_errs.DecodeErr
		}
		ratings = append(ratings, rating)
	}
	if err := result.Err(); err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	return ratings, nil
}
func (r *RatingDal) CreateRating(ctx context.Context, rating *models.Rating) (*models.Rating, error) {
	if !utils.IsEmpty(rating.Id){
		return nil, custom_errs.DBErrCreateWithID
	}
	collection := r.MainDB.Collection(constant.RatingTable)
	result, err := collection.InsertOne(ctx, rating)
	if err != nil {
		return nil, custom_errs.DbErrors
	}
	insertedID, ok:=result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, custom_errs.DBErrIDConversion
	}
	rating.Id = insertedID.Hex()
	return rating, nil
}
func (r *RatingDal) UpdateRating(ctx context.Context, rating *models.Rating) (*models.Rating, error) {
    if utils.IsEmpty(rating.Id) {
        return nil, custom_errs.DBErrUpdateWithID
    }
    collection := r.MainDB.Collection(constant.RatingTable)
    ObjectID, err := primitive.ObjectIDFromHex(rating.Id)
    if err != nil {
        return nil, custom_errs.DBErrIDConversion
    }
    filter := bson.M{"_id": ObjectID}
    ratingBson, err := bson.Marshal(rating)
    if err != nil {
        return nil, custom_errs.DecodeErr
    }
    var updateDoc bson.M
    if err := bson.Unmarshal(ratingBson, &updateDoc); err != nil {
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
    var ratingUpdated models.Rating
    if err := result.Decode(&ratingUpdated); err != nil {
        return nil, custom_errs.DecodeErr
    }
    return &ratingUpdated, nil
}
func (r *RatingDal) DeleteRating(ctx context.Context, ratingId string) (*models.Rating, error) {
	if utils.IsEmpty(ratingId) {
		return nil, custom_errs.DBErrDeleteWithID
	}
	collection := r.MainDB.Collection(constant.RatingTable)
	ObjectID, err := primitive.ObjectIDFromHex(ratingId)
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
	var rating1 *models.Rating
	if err := result.Decode(&rating1); err != nil {
		return nil, custom_errs.DecodeErr
	}
	return rating1, nil
}