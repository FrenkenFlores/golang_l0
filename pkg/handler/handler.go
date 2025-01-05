package handler

import (
	"github.com/FrenkenFlores/golang_l0/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/:id", h.getOrder)
	return router // Engine in the Gin framework has the same interface as http.Handler from the standard Go library.
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}
