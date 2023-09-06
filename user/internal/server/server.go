package server

import (
	"github.com/b0gochort/microservices/internal/server/route"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Server struct {
	echo *echo.Echo
}

func New() *Server {
	return &Server{}
}

func (s *Server) Start(r *route.Route) {
	s.echo = echo.New()

	r.Register(s.echo)

	s.echo.Logger.Fatal(s.echo.Start(viper.GetString("port")))
}
