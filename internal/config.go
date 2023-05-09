package internal

import (
	"encoding/json"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload" // auto import variables from .env
	log "github.com/sirupsen/logrus"
)

// Cfg is a global Config instance
var Config Configuration

// Configuration represents a configuration object
type Configuration struct {
	TargetBlockTime uint64
	Blocks          uint64
	StartMinerCount uint64
	MinerHashTH     uint64
}

// InitConfig instantiates Cfg
func InitConfig() {
	Config.TargetBlockTime = uint64(getEnvAsInt("TARGET_BLOCK_TIME", 60))
	Config.Blocks = uint64(getEnvAsInt("BLOCKS", 1000))
	Config.StartMinerCount = uint64(getEnvAsInt("START_MINER_COUNT", 100))
	Config.MinerHashTH = uint64(getEnvAsInt("MINER_HASH_TH", 100))
}

// PrintConfig prints the current configuration in an easy to read format
func PrintConfig() {
	s, _ := json.MarshalIndent(Config, "", "\t")
	log.Info("Configuration: \n", string(s))
}

// getEnv is a simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// getEnvAsInt is a simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// getEnvAsBool is a simple helper function to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
