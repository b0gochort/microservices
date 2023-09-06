package user

import (
	"github.com/b0gochort/microservices/internal/model"
	"github.com/b0gochort/microservices/internal/storage"
	"github.com/pkg/errors"
)

type User struct {
	postgres *storage.Postgres
}

func New(p *storage.Postgres) *User {
	return &User{
		postgres: p,
	}
}

func (u *User) UserExists(req model.UserCheckReq) (*model.UserCheckRes, error) {
	var res model.UserCheckRes

	user, err := u.postgres.UserExists(req)
	if err != nil {
		return nil, errors.Wrap(err, "user exists")
	}

	res.User = *user

	return &res, nil
}
