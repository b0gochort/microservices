package app

import (
	"github.com/b0gochort/microservices/internal/entity/car"
	"github.com/b0gochort/microservices/internal/server"
	"github.com/b0gochort/microservices/internal/server/handler"
	"github.com/b0gochort/microservices/internal/server/route"
	"github.com/b0gochort/microservices/internal/storage"
	"github.com/b0gochort/microservices/pkg/postgres"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func Start() error {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "config")
	}

	conn, err := postgres.NewPostgres()
	if err != nil {
		return errors.Wrap(err, "connect")
	}

	db := storage.New(conn)

	car := car.New(db)
	carHandler := handler.New(car)
	carHandlerRoute := route.New(carHandler)

	server := server.New()
	server.Start(carHandlerRoute)

	return nil
}
