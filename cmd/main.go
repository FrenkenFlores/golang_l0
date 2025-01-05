package main

import (
	"log"

	gol0 "github.com/FrenkenFlores/golang_l0"
	"github.com/spf13/viper"
)

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalf("Erorr occurred while reading configs: %s", err.Error())
		return
	}
	port := viper.GetString("port")
	if err := gol0.NewServer(port); err != nil {
		log.Fatalf("Error occurred while setting up the server: %s", err.Error())
	}
}
