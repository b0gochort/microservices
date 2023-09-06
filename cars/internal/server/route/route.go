package route

import (
	"net/http"

	"github.com/b0gochort/microservices/internal/server/handler"
	"github.com/labstack/echo/v4"
)

type Route struct {
	userHandler *handler.CarHandler
}

func New(h *handler.CarHandler) *Route {
	return &Route{
		userHandler: h,
	}
}

func (r *Route) Register(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})

	apiGroup := e.Group("/api/cars-service")

	apiGroup.GET("/users/:userID/carsID", r.userHandler.GetIDsUserCars) // /api/cars-service/users/id/carsID
	apiGroup.GET("/users/:userID/cars", r.userHandler.GetUserCars)      // /api/cars-service/users/id/cars
	apiGroup.GET("/cars/:brand", r.userHandler.GetCarsIDByBrand)        // /api/cars-service/cars/brand
	apiGroup.GET("/cars/:brand/:model", r.userHandler.GetCarsIDByModel) // /api/cars-service/cars/brand/model
}
