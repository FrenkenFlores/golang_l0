package main

import (
	"os"

	gol0 "github.com/FrenkenFlores/golang_l0"
	handlers "github.com/FrenkenFlores/golang_l0/pkg/handler"
	"github.com/FrenkenFlores/golang_l0/pkg/repository"
	"github.com/FrenkenFlores/golang_l0/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfigs(); err != nil {
		logrus.Fatalf("Erorr occurred while reading configs: %s", err.Error())
		return
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error occured while loading env: %s", err.Error())
	}
	repository, err := repository.NewRepository(repository.Config{
		Host:     viper.GetStringMapString("db")["host"],
		Port:     viper.GetStringMapString("db")["port"],
		User:     viper.GetStringMapString("db")["user"],
		Password: os.Getenv("DB_PASSWORD"),
		Database: viper.GetStringMapString("db")["database"],
		SslMode:  viper.GetStringMapString("db")["sslmode"],
	})
	if err != nil {
		logrus.Fatalf("Error occurred while creating repository: %s", err.Error())
		return
	}
	service := service.NewService(repository)
	handlers := handlers.NewHandler(service)
	port := viper.GetString("port")
	if err := gol0.NewServer(port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occurred while setting up the server: %s", err.Error())
	}
}
