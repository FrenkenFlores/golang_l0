package gol0

import (
	"net/http"
	"time"

	handlers "github.com/FrenkenFlores/golang_l0/pkg/handler"
)

type Server struct {
	Server *http.Server
}

func NewServer(port string, handler *handlers.Handler) error {
	server := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
		Handler:        handler,
	}
	return server.ListenAndServe()
}
