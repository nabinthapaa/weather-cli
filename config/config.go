package config

import (
	"bufio"
	"io/fs"
	"os"
	"strings"
)

type EnvConfig struct {
	WeatherApi string
}

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
