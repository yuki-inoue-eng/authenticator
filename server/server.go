package server

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yuki-inoue-eng/authenticator/configs"
	"github.com/yuki-inoue-eng/authenticator/infrastructure/db"
	"github.com/yuki-inoue-eng/authenticator/internal/handlers"
	"github.com/yuki-inoue-eng/authenticator/internal/route"
)

type server struct {
	Server  *echo.Echo
	Context context.Context
	Port    string
}

type Server interface {
	Start(errC chan error)
	Shutdown(ctx context.Context) error
}

func New(ctx context.Context, cfg configs.Configs) Server {
	// Echo instance
	e := echo.New()

	// Establish db connection.
	cfgs := cfg.Get()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// handler
	repositories, err := db.NewRepositories(ctx,cfgs)
	if err != nil {
		panic(fmt.Errorf("failed to generate repositories: %v", err))
	}
	handler := handlers.NewHandlers(*repositories)

	// Routes
	route.Routing(e, handler)

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())

	// use echo default logger
	e.Use(middleware.Logger())

	// error handling
	return &server{
		Server:  e,
		Context: ctx,
		Port:    fmt.Sprintf(":%v", cfgs.Server.Port),
	}
}

func (s *server) Start(errC chan error) {
	err := s.Server.Start(s.Port)
	errC <- err
}

func (s *server) Shutdown(ctx context.Context) error {
	if err := s.Server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
