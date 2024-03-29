package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saravanastar/go-hello-world/src/database"
	"github.com/saravanastar/go-hello-world/src/models"
)

type Server interface {
	Start() error
	Readiness(contex echo.Context) error
	Liveness(contex echo.Context) error
}

type EchoServer struct {
	echo     *echo.Echo
	DbClient database.DbClient
}

func NewEchoServer(db database.DbClient) Server {
	server := &EchoServer{
		echo:     echo.New(),
		DbClient: db,
	}
	server.registerRoutes()
	return server
}

func (server *EchoServer) registerRoutes() {
	server.echo.GET("/readiness", server.Readiness)
	server.echo.GET("/liveness", server.Liveness)
}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occurred: %s", err)
		return err
	}
	return nil
}

func (server *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}

func (server *EchoServer) Readiness(ctx echo.Context) error {
	ready := server.DbClient.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
	}
	return ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failed"})
}
