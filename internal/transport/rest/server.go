package rest

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app  *fiber.App
	addr string
}

func NewServer(addr string, app *fiber.App) *Server {
	return &Server{
		app:  app,
		addr: addr,
	}
}

func (s *Server) Start() {
	if err := s.app.Listen(s.addr); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
}

func (s *Server) StartGracefulShutdown() {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := s.app.Shutdown(); err != nil {
			log.Printf("Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	if err := s.app.Listen(s.addr); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
