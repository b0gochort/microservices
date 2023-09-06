package model

import (
	"github.com/b0gochort/microservices/pkg/customerrors"
	"github.com/pkg/errors"
)

type GetUserCarsReq struct {
	UserId int `json:"user_id"`
}

func (r *GetUserCarsReq) Validate() error {
	if r.UserId <= 0 {
		return errors.Wrap(&customerrors.ErrValidate, "user id")
	}
	return nil
}

type GetUserEnginesReq struct {
	UserId int `json:"user_id"`
}

func (r *GetUserEnginesReq) Validate() error {
	if r.UserId <= 0 {
		return errors.Wrap(&customerrors.ErrValidate, "user id")
	}
	return nil
}

type GetEnginesByidReq struct {
	Carsid []int `json:"cars_id"`
}

type CheckUserExists struct {
	User User `json:"user"`
}

type GetUserCarsRes struct {
	Cars []Car `json:"cars"`
}

type GetBrandCarsidReq struct {
	Brand string `json:"brand"`
}

func (r *GetBrandCarsidReq) Validate() error {
	if len(r.Brand) == 0 {
		return errors.Wrap(&customerrors.ErrValidate, "brand")
	}
	return nil
}

type GetBrandCarsidRes struct {
	Carsid []int `json:"cars_id"`
}

type GetModelCarsidReq struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
}

func (r *GetModelCarsidReq) Validate() error {
	if len(r.Brand) == 0 || len(r.Model) == 0 {
		return errors.Wrap(&customerrors.ErrValidate, "brand")
	}
	return nil
}

type GetModelCarsidRes struct {
	Carsid []int `json:"cars_id"`
}

type User struct {
	Userid   int    `json:"id"`
	UserName string `json:"name"`
}

type Car struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	id    int    `json:"id"`
}

type GetidsUserCarsRes struct {
	Data struct {
		Carsid []int `json:"cars_id"`
	} `json:"data"`
}

type HTTPResponse struct {
	Data any    `json:"data"`
	Err  string `json:"err"`
}
