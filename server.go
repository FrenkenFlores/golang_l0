package gol0

import (
	"net/http"
	"time"
)

type Server struct {
	Server *http.Server
}

func NewServer(port string) error {
	server := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}
	return server.ListenAndServe()
}
