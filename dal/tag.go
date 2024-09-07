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
func NewCreateTagDal(mainDB *db.MainMongoDB) inf.CreateTagDal{
	return &TagDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}
func NewGetTagByIdDal(mainDB *db.MainMongoDB) inf.GetTagByIdDal {
	return &TagDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}
func NewGetTagDal(mainDB *db.MainMongoDB) inf.GetTagDal {
	return &TagDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}
func NewUpdateTagDal(mainDB *db.MainMongoDB) inf.UpdateTagDal {
	return &TagDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}

func NewDeleteTagDal(mainDB *db.MainMongoDB) inf.DeleteTagDal {
	return &TagDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}
type TagDal struct {
	MainDB *mongo.Database
}

func (t *TagDal) GetTagById(ctx context.Context, tagId string) (*models.Tag, error) {
	if utils.IsEmpty(tagId) {
		// TODO logging here
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
		// Handle cursor error
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
		// TODO logging here
		return nil, custom_errs.DBErrUpdateWithID
	}

	collection := t.MainDB.Collection(constant.TagTable)
	ObjectID, err := primitive.ObjectIDFromHex(tag.Id)
	if err != nil {
		return nil, custom_errs.DBErrIDConversion
	}

	filter := bson.M{"_id": ObjectID}
	update := bson.M{"$set": tag}

	opts := options.Update().SetUpsert(false)

	_, err = collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, custom_errs.DBErrUpdateWithID
	}

	return tag, nil
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
