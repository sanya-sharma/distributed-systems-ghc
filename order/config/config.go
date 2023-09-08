package config

import (
	"encoding/json"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"fmt"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
		SSLMode  string `json:"sslmode"`
	} `json:"database"`
}

func Load(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

func ReadServiceConfig(service string) (string, error) {
	configFile := "config/config.yaml"
	absConfigPath, err := filepath.Abs(configFile)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
	}
    viper.SetConfigFile(absConfigPath)

    if err := viper.ReadInConfig(); err != nil {
        return "", err
    }

	serviceConfig := "service_url."+service
    serviceURL := viper.GetString(serviceConfig)

    return serviceURL, nil
}


func ReadAPIConfig(api string) (string, error) {
	configFile := "config/config.yaml"
	absConfigPath, err := filepath.Abs(configFile)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
	}
    viper.SetConfigFile(absConfigPath)

    if err := viper.ReadInConfig(); err != nil {
        return "", err
    }

	apiConfig := "api_routes."+api
    apiURL := viper.GetString(apiConfig)

    return apiURL, nil
}
