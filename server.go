package gol0

import (
	"net/http"
	"time"
)

type Server struct {
	Server *http.Server
}

func NewServer(port string, handler http.Handler) error {
	server := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
		Handler:        handler,
	}
	return server.ListenAndServe()
}
