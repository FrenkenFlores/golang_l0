package main

import (
	"os"
	"os/signal"

	gol0 "github.com/FrenkenFlores/golang_l0"
	handlerkafka "github.com/FrenkenFlores/golang_l0/internal/handler-kafka"
	"github.com/FrenkenFlores/golang_l0/internal/kafka"
	"github.com/FrenkenFlores/golang_l0/pkg/handler"
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

const (
	consumerGroup = "orders"
	topic         = "orders"
)

var address = []string{"localhost:9092"}

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
	handlers := handler.NewHandler(service)
	port := viper.GetString("port")

	kafkaHandler := handlerkafka.NewHandler()
	consumer, err := kafka.NewConsumer(kafkaHandler, address, topic, consumerGroup)
	if err != nil {
		logrus.Fatalf("Error occurred while creating consumer: %s", err.Error())
	}
	go func() {
		consumer.Start()
	}()
	if err := gol0.NewServer(port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occurred while setting up the server: %s", err.Error())
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
	consumer.Stop()
}
