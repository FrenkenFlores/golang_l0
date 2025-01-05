package main

import (
	"log"

	gol0 "github.com/FrenkenFlores/golang_l0"
	handlers "github.com/FrenkenFlores/golang_l0/pkg/handler"
	"github.com/spf13/viper"
)

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	handlers := handlers.NewHandler()
	if err := initConfigs(); err != nil {
		log.Fatalf("Erorr occurred while reading configs: %s", err.Error())
		return
	}
	port := viper.GetString("port")
	if err := gol0.NewServer(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occurred while setting up the server: %s", err.Error())
	}
}
