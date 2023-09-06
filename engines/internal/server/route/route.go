package route

import (
	"net/http"

	"github.com/b0gochort/microservices/internal/server/handler"
	"github.com/labstack/echo/v4"
)

type Route struct {
	userHandler *handler.EngineHandler
}

func New(h *handler.EngineHandler) *Route {
	return &Route{
		userHandler: h,
	}
}

func (r *Route) Register(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})

	apiGroup := e.Group("/api/engine-service")

	apiGroup.POST("/engines", r.userHandler.GetEnginesByID) // /api/engine-service/engines
}
