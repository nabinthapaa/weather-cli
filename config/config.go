package config

import (
	"bufio"
	"io/fs"
	"os"
	"strings"
)

// EnvConfig holds the configuration values loaded from the environment file.
type EnvConfig struct {
	WeatherApi string // WeatherApi is the API key for accessing the weather service.
}

// loadEnv reads the environment file provided at the given envFilePath
// and returns a map containing key-value pairs from the env file.
//
// It skips lines that are comments (starting with "#") or empty.
//
// Parameters:
// - envFilePath: The file path to the .env file to be read.
//
// Returns:
// - A map of environment variables, where the keys are the variable names and the values are the corresponding values.
// - An error if the file cannot be opened or read.
func loadEnv(envFilePath string) (map[string]string, error) {
	envMap := make(map[string]string)
	file, err := os.OpenFile(envFilePath, os.O_RDONLY, fs.FileMode(os.O_RDONLY))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}
		value := strings.Split(line, "=")
		envMap[value[0]] = value[1]
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return envMap, nil
}

// GetConfig loads the configuration from the specified environment file
// and returns an EnvConfig struct populated with the loaded values.
//
// Parameters:
// - envFilePath: The file path to the .env file to be read.
//
// Returns:
// - An EnvConfig struct containing the loaded configuration values.
// - An error if the environment file cannot be read or parsed.
func GetConfig(envFilePath string) (EnvConfig, error) {
	env, err := loadEnv(envFilePath)
	if err != nil {
		return EnvConfig{}, err
	}

	envConfig := EnvConfig{
		WeatherApi: env["WEATHER_API"],
	}

	return envConfig, nil
}
