package handler

import (
	"url-shortener/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{"GET", "POST", "OPTIONS"},
			AllowHeaders: []string{"content-type"},
		},
	))

	router.POST("/url", h.CreateShortUrl)
	router.GET("/:short_url", h.GetFullUrl)

	return router
}
