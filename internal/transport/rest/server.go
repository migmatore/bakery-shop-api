package rest

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app  *fiber.App
	addr string
	pool *pgxpool.Pool
}

func NewServer(addr string, app *fiber.App, pool *pgxpool.Pool) *Server {
	return &Server{
		app:  app,
		addr: addr,
		pool: pool,
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

		log.Printf("Close all database connections...")
		s.pool.Close()
		log.Printf("All database connections have been closed!")

		if err := s.app.Shutdown(); err != nil {
			log.Printf("Server is not shutting down! Reason: %v", err)
		}

		log.Printf("Server has successfully shut down!")

		close(idleConnsClosed)
	}()

	if err := s.app.Listen(s.addr); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
