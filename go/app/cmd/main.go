package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"microservices"
	"microservices/app/internal/handler"
	"microservices/app/internal/repository"
	"microservices/app/internal/service"
	"microservices/app/libraries/logging"
)

func main() {
	logger := logging.GetLogger()
	if err := initConfig(); err != nil {
		logger.Fatalf("error reading config file %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logger.Fatal("can`t read env files")
	}

	repos := repository.NewRepository()
	services := service.NewService(repos, logger)
	handlers := handler.NewHandler(services, logger)

	srv := &microservices.Server{}
	logger.Info("trying to start server")
	if err := srv.Run(viper.GetString("port"), handlers.InitRouts()); err != nil {
		log.Fatalf("error while trying run server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("app/internal/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
