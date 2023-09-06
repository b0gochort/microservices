package handler

import (
	"github.com/b0gochort/microservices/internal/entity/engine"
	"github.com/b0gochort/microservices/internal/model"
	"github.com/b0gochort/microservices/pkg/customerrors"
	"github.com/b0gochort/microservices/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type EngineHandler struct {
	engine *engine.Engine
}

func New(engine *engine.Engine) *EngineHandler {
	return &EngineHandler{
		engine: engine,
	}
}

// /api/engine-service/engines
func (e *EngineHandler) GetEnginesByID(c echo.Context) (err error) {
	var (
		req model.GetEnginesByIDReq
		res *model.GetEnginesByIDRes
	)

	defer func() {
		c.JSON(utils.HTTPResponse(err, res))
	}()

	if err = c.Bind(&req); err != nil {
		return errors.Wrap(&customerrors.ErrBindReq, "bind req")
	}

	res, err = e.engine.GetEnginesByID(req)
	if err != nil {
		return errors.Wrap(err, "user exists")
	}

	return
}
