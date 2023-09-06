package handler

import (
	"strconv"

	"github.com/b0gochort/microservices/internal/entity/car"
	"github.com/b0gochort/microservices/internal/model"
	"github.com/b0gochort/microservices/pkg/customerrors"
	"github.com/b0gochort/microservices/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type CarHandler struct {
	user *car.Car
}

func New(user *car.Car) *CarHandler {
	return &CarHandler{
		user: user,
	}
}

// /api/cars-service/users/id/cars
func (u *CarHandler) GetUserCars(c echo.Context) (err error) {
	var (
		req model.GetUserCarsReq
		res *model.GetUserCarsRes
	)

	defer func() {
		c.JSON(utils.HTTPResponse(err, res))
	}()

	req.UserID, err = strconv.Atoi(c.Param("userID"))
	if err != nil {
		return errors.Wrap(&customerrors.ErrValidate, "convert string to integer")
	}

	res, err = u.user.GetUserCars(req)
	if err != nil {
		return errors.Wrap(err, "get user cars")
	}

	return
}

// /api/cars-service/users/id/carsID
func (u *CarHandler) GetIDsUserCars(c echo.Context) (err error) {
	var (
		req model.GetIDsUserCarsReq
		res *model.GetIDsUserCarsRes
	)

	defer func() {
		c.JSON(utils.HTTPResponse(err, res))
	}()

	req.UserID, err = strconv.Atoi(c.Param("userID"))
	if err != nil {
		return errors.Wrap(&customerrors.ErrValidate, "convert string to integer")
	}

	res, err = u.user.GetIDsUserCars(req)
	if err != nil {
		return errors.Wrap(err, "get user cars")
	}

	return
}

// /api/cars-service/cars/brand
func (u *CarHandler) GetCarsIDByBrand(c echo.Context) (err error) {
	var (
		req model.GetBrandCarsIDReq
		res *model.GetBrandCarsIDRes
	)

	defer func() {
		c.JSON(utils.HTTPResponse(err, res))
	}()

	req.Brand = c.Param("brand")

	res, err = u.user.GetCarsIDByBrand(req)
	if err != nil {
		return errors.Wrap(err, "get user cars")
	}

	return
}

// /api/cars-service/cars/brand/model
func (u *CarHandler) GetCarsIDByModel(c echo.Context) (err error) {
	var (
		req model.GetModelCarsIDReq
		res *model.GetModelCarsIDRes
	)
	defer func() {
		c.JSON(utils.HTTPResponse(err, res))
	}()

	req.Brand = c.Param("brand")
	req.Model = c.Param("model")

	res, err = u.user.GetCarsIDByModel(req)
	if err != nil {
		return errors.Wrap(err, "get user cars")
	}

	return
}
