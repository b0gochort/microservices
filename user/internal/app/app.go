package app

import (
	"github.com/b0gochort/microservices/internal/entity/user"
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

	user := user.New(db)
	userHandler := handler.New(user)
	userHandlerRoute := route.New(userHandler)

	server.New().Start(userHandlerRoute)

	return nil
}
