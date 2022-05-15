package main

import (
	"context"
	"errors"
	"net/http"
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
	Port       int
	Production bool
	MongoUri   string
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Error("error loading env variables from .env: %v", err.Error())
	}

	config, err := loadConfig()
	if err != nil {
		logrus.Fatalf("failed to load config: %s", err.Error())
	}

	if config.Production {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := mongo.NewMongoDB(config.MongoUri)
	if err != nil {
		logrus.Fatalf("failed to init db: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	server := new(domain.Server)

	go func() {
		if err := server.Run(strconv.Itoa(config.Port), handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
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
	config := new(Config)

	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 32)
	if err != nil {
		return nil, err
	}

	config.Port = int(port)

	if config.Port > 65535 || config.Port < 0 {
		return nil, errors.New("port is too big")
	}

	production, err := strconv.ParseBool(os.Getenv("PRODUCTION"))
	if err != nil {
		return nil, err
	}

	config.Production = production
	config.MongoUri = os.Getenv("MONGO_URI")

	return config, nil
}
