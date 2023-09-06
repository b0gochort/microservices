package main

import (
	"log"
	"net/http"

	"github.com/b0gochort/microservices/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("config")
	}

	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})

	apiGroup := e.Group("/api")

	userGroup := apiGroup.Group("/users")
	enginesGroup := apiGroup.Group("/engines")

	userGroup.GET("/:userid/cars", handler.GetUserCars)       // api/users/id/cars
	userGroup.GET("/:userid/engines", handler.GetUserEngines) // api/users/id/engines

	enginesGroup.GET("/:brand", handler.GetEnginesByBrand)        // /api/engines/brand
	enginesGroup.GET("/:brand/:model", handler.GetEnginesByModel) // /api/engines/brand/model

	if err := e.Start(viper.GetString("port")); err != nil {
		log.Fatal(err)
	}
}
