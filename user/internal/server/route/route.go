package route

import (
	"net/http"

	"github.com/b0gochort/microservices/internal/server/handler"
	"github.com/labstack/echo/v4"
)

type Route struct {
	userHandler *handler.UserHandler
}

func New(h *handler.UserHandler) *Route {
	return &Route{
		userHandler: h,
	}
}

func (r *Route) Register(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})

	apiGroup := e.Group("/api/user-service")

	userGroup := apiGroup.Group("/users")

	userGroup.GET("/:userID/exists", r.userHandler.UserExists) // api/user-service/users/id/exists
}
