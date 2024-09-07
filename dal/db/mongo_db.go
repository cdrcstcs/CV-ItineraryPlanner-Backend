package db
import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "itineraryplanner/constant"
    "github.com/joho/godotenv"
    "log"
    "os"
)
type MainMongoDB mongo.Database
func GetMainMongoDatabase() *MainMongoDB {
    return (*MainMongoDB)(GetMongoClient().Database(constant.MainMongoDB))
}
func GetMongoClient() *mongo.Client {
    err := godotenv.Load("D:/CV-Projects/MainCV/CV-ItineraryPlanner/CV-ItineraryPlanner-Backend/.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
    clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
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