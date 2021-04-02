package config

import (
	"encoding/json"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload" // auto import variables from .env
	log "github.com/sirupsen/logrus"
)

// Cfg is a global Config instance
var Cfg Config

// Config represents a configuration object
type Config struct {
	TargetBlockTime uint
	Blocks          uint
	MinerCount      uint
	MinerHashTH     uint
}

// Init instantiates Cfg
func Init() {
	Cfg.TargetBlockTime = uint(getEnvAsInt("TARGET_BLOCK_TIME", 60))
	Cfg.Blocks = uint(getEnvAsInt("BLOCKS", 1000))
	Cfg.MinerCount = uint(getEnvAsInt("MINER_COUNT", 100))
	Cfg.MinerHashTH = uint(getEnvAsInt("MINER_HASH_TH", 100))
}

// Print prints the current configuration in an easy to read format
func Print() {
	s, _ := json.MarshalIndent(Cfg, "", "\t")
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
