package car

import (
	"github.com/b0gochort/microservices/internal/model"
	"github.com/b0gochort/microservices/internal/storage"
	"github.com/pkg/errors"
)

type Car struct {
	postgres *storage.Postgres
}

func New(p *storage.Postgres) *Car {
	return &Car{
		postgres: p,
	}
}

func (c *Car) GetUserCars(req model.GetUserCarsReq) (*model.GetUserCarsRes, error) {
	var res model.GetUserCarsRes

	cars, err := c.postgres.GetCar(req)
	if err != nil {
		return nil, errors.Wrap(err, "car exists")
	}

	res.Cars = cars

	return &res, nil
}

func (c *Car) GetIDsUserCars(req model.GetIDsUserCarsReq) (*model.GetIDsUserCarsRes, error) {
	var res model.GetIDsUserCarsRes

	carsID, err := c.postgres.GetIdsCarsByUser(req)
	if err != nil {
		return nil, errors.Wrap(err, "car exists")
	}

	res.CarsID = carsID

	return &res, nil
}

func (c *Car) GetCarsIDByBrand(req model.GetBrandCarsIDReq) (*model.GetBrandCarsIDRes, error) {
	var res model.GetBrandCarsIDRes

	carsID, err := c.postgres.GetIdCarsByBrand(req)
	if err != nil {
		return nil, errors.Wrap(err, "car exists")
	}

	res.CarsID = carsID

	return &res, nil
}

func (c *Car) GetCarsIDByModel(req model.GetModelCarsIDReq) (*model.GetModelCarsIDRes, error) {
	var res model.GetModelCarsIDRes

	carsID, err := c.postgres.GetIdCarsByModel(req)
	if err != nil {
		return nil, errors.Wrap(err, "car exists")
	}

	res.CarsID = carsID

	return &res, nil
}
