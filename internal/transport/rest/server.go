package rest

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
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

func (s *Server) Start(ctx context.Context) {
	if err := s.app.Listen(s.addr); err != nil {
		logging.GetLogger(ctx).Fatalf("Server is not running! Reason: %v", err)
	}
}

func (s *Server) StartWithGracefulShutdown(ctx context.Context) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		logging.GetLogger(ctx).Info("Close all database connections...")
		s.pool.Close()
		logging.GetLogger(ctx).Info("All database connections have been closed!")

		if err := s.app.Shutdown(); err != nil {
			logging.GetLogger(ctx).Errorf("Server is not shutting down! Reason: %v", err)
		}

		logging.GetLogger(ctx).Info("Server has successfully shut down!")

		close(idleConnsClosed)
	}()

	if err := s.app.Listen(s.addr); err != nil {
		logging.GetLogger(ctx).Fatalf("Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
