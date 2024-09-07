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
func NewTagDal(mainDB *db.MainMongoDB) inf.TagDal{
	return &TagDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}
type TagDal struct {
	MainDB *mongo.Database
}
func (t *TagDal) GetDB() *mongo.Database{
	return t.MainDB
}
func (t *TagDal) GetTagById(ctx context.Context, tagId string) (*models.Tag, error) {
	if utils.IsEmpty(tagId) {
		return nil, custom_errs.DBErrCreateWithID
	}
	collection := t.MainDB.Collection(constant.TagTable)
	ObjectID, err := primitive.ObjectIDFromHex(tagId)
	if err != nil {
		return nil, custom_errs.DBErrIDConversion
	}
	result := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	if result.Err() != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	var tag *models.Tag
	if err := result.Decode(&tag); err != nil {
		return nil, custom_errs.DecodeErr
	}
	return tag, nil
}
func (t *TagDal) GetTag(ctx context.Context) ([]*models.Tag, error) {
	collection := t.MainDB.Collection(constant.TagTable)
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	defer result.Close(ctx)
	var tags []*models.Tag
	for result.Next(ctx) {
		var tag *models.Tag
		if err := result.Decode(&tag); err != nil {
			return nil, custom_errs.DecodeErr
		}
		tags = append(tags, tag)
	}
	if err := result.Err(); err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	return tags, nil
}
func (t *TagDal) CreateTag(ctx context.Context, tag *models.Tag) (*models.Tag, error) {
	if !utils.IsEmpty(tag.Id){
		return nil, custom_errs.DBErrCreateWithID
	}
	collection := t.MainDB.Collection(constant.TagTable)
	result, err := collection.InsertOne(ctx, tag)
	if err != nil {
		return nil, custom_errs.DbErrors
	}
	insertedID, ok:=result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, custom_errs.DBErrIDConversion
	}
	tag.Id = insertedID.String()
	return tag, nil
}
func (t *TagDal) UpdateTag(ctx context.Context, tag *models.Tag) (*models.Tag, error) {
    if utils.IsEmpty(tag.Id) {
        return nil, custom_errs.DBErrUpdateWithID
    }
    collection := t.MainDB.Collection(constant.TagTable)
    ObjectID, err := primitive.ObjectIDFromHex(tag.Id)
    if err != nil {
        return nil, custom_errs.DBErrIDConversion
    }
    filter := bson.M{"_id": ObjectID}
    tagBson, err := bson.Marshal(tag)
    if err != nil {
        return nil, custom_errs.DecodeErr
    }
    var updateDoc bson.M
    if err := bson.Unmarshal(tagBson, &updateDoc); err != nil {
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
    var tagUpdated models.Tag
    if err := result.Decode(&tagUpdated); err != nil {
        return nil, custom_errs.DecodeErr
    }
    return &tagUpdated, nil
}
func (t *TagDal) DeleteTag(ctx context.Context, tagId string) (*models.Tag, error) {
	if utils.IsEmpty(tagId) {
		return nil, custom_errs.DBErrDeleteWithID
	}
	collection := t.MainDB.Collection(constant.TagTable)
	ObjectID, err := primitive.ObjectIDFromHex(tagId)
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
	var tag1 *models.Tag
	if err := result.Decode(&tag1); err != nil {
		return nil, custom_errs.DecodeErr
	}
	return tag1, nil
}