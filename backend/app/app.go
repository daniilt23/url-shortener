package app

import (
	"log"
	"url-shortener/internal/database/postgres"
	"url-shortener/internal/database/postgres/url"
	"url-shortener/internal/handler"
	"url-shortener/internal/service"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (app *App) Start() {
	log.Println("START APP")

	db := postgres.InitDB()

	urlRepo := url.NewRepoSQL(db)

	service := service.NewService(urlRepo)

	handler := handler.NewHandler(service)

	router := handler.InitRoutes()

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
