package main

import (
	"log"

	GenPass "github.com/mikromolekula2002/key_generate_api"
	"github.com/mikromolekula2002/key_generate_api/internal/config"
	"github.com/mikromolekula2002/key_generate_api/internal/generator"
	"github.com/mikromolekula2002/key_generate_api/internal/handler"
	"github.com/mikromolekula2002/key_generate_api/internal/logger"
	"github.com/mikromolekula2002/key_generate_api/internal/repo"
	"github.com/mikromolekula2002/key_generate_api/internal/service"
)

func main() {
	srv := new(GenPass.Server)

	//Загрузка конфига
	cfg := config.LoadConfig("/config.yaml")
	// Инициализация базы данных
	logger := logger.Init(cfg.Logger.Level, cfg.Logger.FilePath, cfg.Logger.Output)

	gen := new(generator.GenValue)

	//инициализация репозитория
	repos, err := repo.InitDB(cfg)
	if err != nil {
		log.Fatalf("Ошибка инициализации Базы Данных: \n%v", err)
	}

	service := service.InitService(logger.Logrus, repos, gen) // написать сервисный слой

	handl := handler.Init(service)

	err = srv.Run(cfg.Server.Port, handl.Gin) // порт берем из конфига!!!
	if err != nil {
		log.Fatal("Error while starting server", err)
	}
}
