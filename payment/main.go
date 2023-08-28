package main

import (
	"distributed-systems-ghc/payment/config"
	"distributed-systems-ghc/payment/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"path/filepath"
)

func main() {

	config, err := getConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		db, err := setupDB(config)
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}
		c.Set("db", db) // Store the DB connection in the context
		c.Next()
	})
	routes.SetupRoutes(router)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupDB(config config.Config) (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host, config.Database.Port, config.Database.User,
		config.Database.Password, config.Database.DBName, config.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Printf("Db connection established")
	return db, nil
}

func getConfig() (config.Config, error) {
	configFile := "config/config.json"
	absConfigPath, err := filepath.Abs(configFile)
	if err != nil {
		log.Fatal("Error getting absolute path:", err)
	}
	log.Printf(absConfigPath)

	return config.Load(absConfigPath)
}
