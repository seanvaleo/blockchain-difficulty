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
	TargetBlockTimeSeconds         uint64
	SimulationDays                 uint32
	InitialNetworkHashPower        uint64 // Hashes per second
	LimitNetworkHashPowerPctChange int    // How much the network power can vary from the initial power
}

// InitConfig instantiates Cfg with defaults or environment variables
func InitConfig() {
	Config.TargetBlockTimeSeconds = getEnvAsUint64("TARGET_BLOCK_TIME_SECONDS", 600)
	Config.SimulationDays = getEnvAsUint32("SIMULATION_DAYS", 365)
	Config.InitialNetworkHashPower = getEnvAsUint64("INITIAL_NETWORK_HASH_POWER", 1000000)
	Config.LimitNetworkHashPowerPctChange = getEnvAsInt("LIMIT_NETWORK_HASH_POWER_PCT_CHANGE", 10)
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

// getEnvAsUint64 is a simple helper function to read an environment variable into integer or return a default value
func getEnvAsUint64(name string, defaultVal uint64) uint64 {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseUint(valueStr, 10, 64); err == nil {
		return value
	}

	return defaultVal
}

// getEnvAsUint32 is a simple helper function to read an environment variable into integer or return a default value
func getEnvAsUint32(name string, defaultVal uint32) uint32 {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseUint(valueStr, 10, 32); err == nil {
		return uint32(value)
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
