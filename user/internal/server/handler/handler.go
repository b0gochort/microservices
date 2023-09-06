package handler

import (
	"strconv"

	"github.com/b0gochort/microservices/internal/entity/user"
	"github.com/b0gochort/microservices/internal/model"
	"github.com/b0gochort/microservices/pkg/customerrors"
	"github.com/b0gochort/microservices/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type UserHandler struct {
	user *user.User
}

func New(user *user.User) *UserHandler {
	return &UserHandler{
		user: user,
	}
}

func (u *UserHandler) UserExists(c echo.Context) (err error) {
	var (
		req model.UserCheckReq
		res *model.UserCheckRes
	)

	defer func() {
		c.JSON(utils.HTTPResponse(err, res))
	}()

	req.UserID, err = strconv.Atoi(c.Param("userID"))
	if err != nil {
		return errors.Wrap(&customerrors.ErrValidate, "convert string to integer")
	}

	res, err = u.user.UserExists(req)
	if err != nil {
		return errors.Wrap(err, "user exists")
	}

	return
}
