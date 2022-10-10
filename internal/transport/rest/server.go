package rest

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	logger *logging.Logger
	app    *fiber.App
	addr   string
	pool   *pgxpool.Pool
}

func NewServer(addr string, app *fiber.App, pool *pgxpool.Pool, logger *logging.Logger) *Server {
	return &Server{
		app:    app,
		addr:   addr,
		pool:   pool,
		logger: logger,
	}
}

func (s *Server) Start() {
	if err := s.app.Listen(s.addr); err != nil {
		s.logger.Fatalf("Server is not running! Reason: %v", err)
	}
}

func (s *Server) StartWithGracefulShutdown() {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		s.logger.Info("Close all database connections...")
		s.pool.Close()
		s.logger.Info("All database connections have been closed!")

		if err := s.app.Shutdown(); err != nil {
			s.logger.Errorf("Server is not shutting down! Reason: %v", err)
		}

		s.logger.Info("Server has successfully shut down!")

		close(idleConnsClosed)
	}()

	if err := s.app.Listen(s.addr); err != nil {
		s.logger.Fatalf("Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
