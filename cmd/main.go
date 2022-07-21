package main

import (
	E_shop_go "E-shop-go"
	"E-shop-go/pkg/handler"
	"E-shop-go/pkg/repository"
	"E-shop-go/pkg/service"
	"fmt"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "4eJ7sKwB",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	fmt.Println("error_____")
	fmt.Println(err)
	fmt.Println("_____error")
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(E_shop_go.Server)
	if err = server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
