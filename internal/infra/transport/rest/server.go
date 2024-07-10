package rest

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewServer(router *gin.Engine, port string) *HTTPServer {
	addr := ":" + port

	server := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &HTTPServer{
		Router: router,
		Server: server,
	}
}

type HTTPServer struct {
	*http.Server
	Router *gin.Engine
}

func (s *HTTPServer) Start() error {
	slog.Info("HTTP Server is running on port " + s.Addr + "...")

	if err := s.Server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
