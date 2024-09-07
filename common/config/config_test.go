package config
import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
	"log"
)
func TestInitGlobalConfig(t *testing.T) {
	err := godotenv.Load("D:/CV-Projects/MainCV/CV-ItineraryPlanner/CV-ItineraryPlanner-Backend/.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
	filePath := "../../conf/local_config.json"
	InitGlobalConfig(filePath)
	wantConfig := Config{
		ServerURL: "localhost:8100",
		MongoURL:  os.Getenv("MONGO_URL"),
	}
	assert.Equal(t, wantConfig, GlobalConfig)
}