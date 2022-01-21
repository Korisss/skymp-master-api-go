package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"

	master_api "github.com/Korisss/skymp-master-api-go"
	"github.com/Korisss/skymp-master-api-go/internal/handler"
	"github.com/Korisss/skymp-master-api-go/internal/repository"
	"github.com/Korisss/skymp-master-api-go/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Port       int  `json:"port"`
	Production bool `json:"production"`
	DBConfig   struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Username string `json:"username"`
		DBName   string `json:"db_name"`
		SSLMode  string `json:"ssl_mode"`
	} `json:"db_config"`
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	config, err := loadConfig()
	if err != nil {
		logrus.Fatalf("failed to load config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %v", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     config.DBConfig.Host,
		Port:     config.DBConfig.Port,
		Username: config.DBConfig.Username,
		DBName:   config.DBConfig.DBName,
		SSLMode:  config.DBConfig.SSLMode,
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to init db: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	server := new(master_api.Server)
	if err := server.Run(strconv.Itoa(config.Port), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while running http server: %s", err.Error())
	}
}

func loadConfig() (*Config, error) {
	buffer, err := ioutil.ReadFile("configs/config.json")

	if err != nil {
		return nil, err
	}

	config := new(Config)
	json.Unmarshal(buffer, config)

	if config.Port > 65535 || config.Port < 0 {
		return nil, errors.New("port is too big")
	}

	return config, nil
}
