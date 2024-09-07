package config
import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)
func TestInitGlobalConfig(t *testing.T) {
	filePath := "../../conf/local_config.json"
	InitGlobalConfig(filePath)
	wantConfig := Config{
		ServerURL: "localhost:8100",
		MongoURL:  os.Getenv("MONGO_URL"),
	}
	assert.Equal(t, wantConfig, GlobalConfig)
}