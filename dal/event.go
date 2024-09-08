package dal
import(
	"itineraryplanner/common/custom_errs"
	"itineraryplanner/common/utils"
	"itineraryplanner/constant"
	"itineraryplanner/models"
	"itineraryplanner/dal/db"
	"itineraryplanner/dal/inf"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func NewEventDal(mainDB *db.MainMongoDB) inf.EventDal{
	return &EventDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}
type EventDal struct {
	MainDB *mongo.Database
}
func (e *EventDal) GetDB() *mongo.Database{
	return e.MainDB
}
func (e *EventDal) GetEventById(ctx context.Context, eventId string) (*models.Event, error) {
	if utils.IsEmpty(eventId) {
		return nil, custom_errs.DBErrCreateWithID
	}
	collection := e.MainDB.Collection(constant.EventTable)
	ObjectID, err := primitive.ObjectIDFromHex(eventId)
	if err != nil {
		return nil, custom_errs.DBErrIDConversion
	}
	result := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	if result.Err() != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	var event *models.Event
	if err := result.Decode(&event); err != nil {
		return nil, custom_errs.DecodeErr
	}
	return event, nil
}
func (e *EventDal) GetEvent(ctx context.Context) ([]*models.Event, error) {
	collection := e.MainDB.Collection(constant.EventTable)
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	defer result.Close(ctx)
	var events []*models.Event
	for result.Next(ctx) {
		var event *models.Event
		if err := result.Decode(&event); err != nil {
			return nil, custom_errs.DecodeErr
		}
		events = append(events, event)
	}
	if err := result.Err(); err != nil {
		return nil, custom_errs.DBErrGetWithID
	}
	return events, nil
}
func (e *EventDal) CreateEvent(ctx context.Context, event *models.Event) (*models.Event, error){
	if !utils.IsEmpty(event.Id){
		return nil, custom_errs.DBErrCreateWithID
	}
	eventCollection := e.MainDB.Collection(constant.EventTable)
	result, err := eventCollection.InsertOne(ctx, event)
	if err != nil {
		return nil, custom_errs.DbErrors
	}
	insertedID, ok:=result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, custom_errs.DBErrIDConversion
	}
	event.Id = insertedID.Hex()
	return event, nil
}
func (e *EventDal) UpdateEvent(ctx context.Context, event *models.Event) (*models.Event, error) {
    if utils.IsEmpty(event.Id) {
        return nil, custom_errs.DBErrUpdateWithID
    }
    collection := e.MainDB.Collection(constant.EventTable)
    ObjectID, err := primitive.ObjectIDFromHex(event.Id)
    if err != nil {
        return nil, custom_errs.DBErrIDConversion
    }
    filter := bson.M{"_id": ObjectID}
    eventBson, err := bson.Marshal(event)
    if err != nil {
        return nil, custom_errs.DecodeErr
    }
    var updateDoc bson.M
    if err := bson.Unmarshal(eventBson, &updateDoc); err != nil {
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
    var eventUpdated models.Event
    if err := result.Decode(&eventUpdated); err != nil {
        return nil, custom_errs.DecodeErr
    }
    return &eventUpdated, nil
}
func (e *EventDal) DeleteEvent(ctx context.Context, eventId string) (*models.Event, error) {
	if utils.IsEmpty(eventId) {
		return nil, custom_errs.DBErrDeleteWithID
	}
	collection := e.MainDB.Collection(constant.EventTable)
	ObjectID, err := primitive.ObjectIDFromHex(eventId)
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
	var event1 *models.Event
	if err := result.Decode(&event1); err != nil {
		return nil, custom_errs.DecodeErr
	}
	return event1, nil
}