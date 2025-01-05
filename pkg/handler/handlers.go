package handlers

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/:id", func(c *gin.Context) { c.Writer.Write([]byte("Hello world!")) })
	return router // Engine in the Gin framework has the same interface as http.Handler from the standard Go library.
}

func NewHandler() *Handler {
	return &Handler{}
}
