package dal
import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"itineraryplanner/common/custom_errs"
	"itineraryplanner/constant"
	"itineraryplanner/dal/db"
	"itineraryplanner/dal/inf"
	"itineraryplanner/models"
	"go.mongodb.org/mongo-driver/mongo/options"

)
func NewRecommendItineraryDal(mainDB *db.MainMongoDB) inf.RecommendItineraryDal {
	return &AlgoDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}

func NewBuildItineraryDal(mainDB *db.MainMongoDB) inf.BuildItineraryDal {
	return &AlgoDal{
		MainDB: (*mongo.Database)(mainDB),
	}
}

type AlgoDal struct {
	MainDB *mongo.Database
}

func (a *AlgoDal) RecommendItinerary(ctx context.Context, cor *models.Coordinate) ([]*models.Itinerary, error){
	itiIds, err := findCoordinatesWithConditions(a.MainDB, constant.CoordinateTable, cor.X, cor.Y, true)
	if err != nil {
		return nil, custom_errs.DbErrors
	}
	var itis []*models.Itinerary
	for _,v := range itiIds{
		iti, err:= NewGetItineraryByIdDal((*db.MainMongoDB)(a.MainDB)).GetItineraryById(ctx, v)
		if err != nil {
			return nil, custom_errs.DBErrGetWithID
		}
		itis = append(itis, iti)
	}
	return itis, nil
}

func (a *AlgoDal) BuildItinerary(ctx context.Context, cor *models.Coordinate) (*models.Itinerary, error) {
    eveIds, err := findCoordinatesWithConditions(a.MainDB, constant.CoordinateTable, cor.X, cor.Y, false)
    if err != nil {
        return nil, custom_errs.DbErrors
    }

    var minStartTime, maxEndTime time.Time

    for _, v := range eveIds {
        eve, err := NewGetEventByIdDal((*db.MainMongoDB)(a.MainDB)).GetEventById(ctx, v)
        if err != nil {
            return nil, custom_errs.DBErrGetWithID
        }

        minStartTime = minTime(minStartTime, eve.StartTime)
        maxEndTime = maxTime(maxEndTime, eve.EndTime)
    }
	iti := &models.Itinerary{
		StartTime: minStartTime,
		EndTime: maxEndTime,
		EventIds: eveIds,
	}
	i, err := NewCreateItineraryDal((*db.MainMongoDB)(a.MainDB)).CreateItinerary(ctx, iti)
	if err !=nil{
		return nil, custom_errs.DBErrCreateWithID
	}

    return i, nil
}

func minTime(a, b time.Time) time.Time {
	if a.IsZero() {
        return b
    }
    if a.Before(b) {
        return a
    }
    return b
}

func maxTime(a, b time.Time) time.Time {
	if a.IsZero() {
        return b
    }
    if a.After(b) {
        return a
    }
    return b
}


func createIndex(db *mongo.Database, collectionName string) error {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{"x", 1},
			{"y", 1},
		},
	}

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second) 

	indexView := db.Collection(collectionName).Indexes()

	_, err := indexView.CreateOne(context.Background(), indexModel, opts)
	return err
}


func sortDocuments(db *mongo.Database, collectionName string) error {
	findOptions := options.Find().SetSort(bson.D{
		{"x", 1},
		{"y", 1},
	})

	_, err := db.Collection(collectionName).Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		return err
	}
	return nil
}



func findCoordinatesWithConditions(db *mongo.Database, collectionName string, specificX, specificY int, t bool) ([]string, error) {

	err := createIndex(db, collectionName)
	if err != nil {
		return nil, err
	}
	err = sortDocuments(db, collectionName)
	if err != nil {
		return nil, err
	}

	filter := bson.D{
		{"x", bson.D{{"$gt", specificX}}},
		{"y", bson.D{{"$gt", specificY}}},
	}

	findOptions := options.Find().SetLimit(20) 

	cur, err := db.Collection(collectionName).Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var result []models.Coordinate
	if err := cur.All(context.Background(), &result); err != nil {
		return nil, err
	}

	var ret []string
	for _,v := range result{
		if t {
			ret = append(ret, v.ItineraryId)
		} else {
			ret = append(ret, v.EventId)
		}
	}

	return ret, nil
}
