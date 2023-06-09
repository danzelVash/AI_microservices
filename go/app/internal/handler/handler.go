package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"microservices/app/internal/service"
	"microservices/app/libraries/logging"
)

type Handler struct {
	services *service.Service
	logger   *logging.Logger
}

func NewHandler(services *service.Service, logger *logging.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) InitRouts() *echo.Echo {
	router := echo.New()
	router.Use(middleware.Recover())
	//router.Use(middleware.Logger())

	router.Static("/css", "app/templates/css")
	router.Static("/js", "app/templates/js")
	router.Static("/img", "app/templates/img")
	router.Static("/txt", "app/templates/txt")

	router.GET("/", h.indexPage)
	router.POST("/", h.getAdvices)

	return router
}
