package engine

import (
	"github.com/b0gochort/microservices/internal/model"
	"github.com/b0gochort/microservices/internal/storage"
	"github.com/pkg/errors"
)

type Engine struct {
	postgres *storage.Postgres
}

func New(p *storage.Postgres) *Engine {
	return &Engine{
		postgres: p,
	}
}

func (e Engine) GetEnginesByID(req model.GetEnginesByIDReq) (*model.GetEnginesByIDRes, error) {
	var res model.GetEnginesByIDRes

	engines, err := e.postgres.GetEnginesByID(req)
	if err != nil {
		return nil, errors.Wrap(err, "get user engines")
	}

	res.Engines = engines

	return &res, nil
}
