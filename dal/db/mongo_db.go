package db
import (
    "context"
    "log"
    "os"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "itineraryplanner/common/config"
    "itineraryplanner/constant"
)
type MainMongoDB mongo.Database
func init() {
    err := godotenv.Load("D:/CV-Projects/MainCV/CV-ItineraryPlanner/API-Golang/.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
    config.GlobalConfig.MongoURL = os.Getenv("MONGO_URL")
}
func GetMainMongoDatabase() *MainMongoDB {
    return (*MainMongoDB)(GetMongoClient().Database(constant.MainMongoDB))
}
func GetMongoClient() *mongo.Client {
    clientOptions := options.Client().ApplyURI(config.GlobalConfig.MongoURL)
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        panic(err)
    }
    err = client.Ping(context.Background(), nil)
    if err != nil {
        panic(err)
    }
    return client
}
var testDB *mongo.Database
func initTestDB() {
    testDB = GetMongoClient().Database(constant.TestMongoDB)
}
func GetMemoMongo() *mongo.Database {
    if testDB == nil {
        initTestDB()
    }
    return testDB
}