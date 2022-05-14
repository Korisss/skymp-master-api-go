package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Korisss/skymp-master-api-go/internal/domain"
	"github.com/Korisss/skymp-master-api-go/internal/handler"
	"github.com/Korisss/skymp-master-api-go/internal/repository"
	"github.com/Korisss/skymp-master-api-go/internal/repository/mongo"
	"github.com/Korisss/skymp-master-api-go/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Port       int  `json:"port"`
	Production bool `json:"production"`
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	config, err := loadConfig()
	if err != nil {
		logrus.Fatalf("failed to load config: %s", err.Error())
	}

	if config.Production {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Error("error loading env variables: %v", err.Error())
	}

	mongoUri := os.Getenv("MONGO_URI")

	db, err := mongo.NewMongoDB(mongoUri)
	if err != nil {
		logrus.Fatalf("failed to init db: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	server := new(domain.Server)

	go func() {
		if err := server.Run(strconv.Itoa(config.Port), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("SkyMP Master Server started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("SkyMP Master Server shutting down...")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Error("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Disconnect(context.Background()); err != nil {
		logrus.Error("error occured on db connection close: %s", err.Error())
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