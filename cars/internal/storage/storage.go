package storage

import (
	"database/sql"

	"github.com/b0gochort/microservices/internal/model"
	"github.com/b0gochort/microservices/pkg/customerrors"
	"github.com/pkg/errors"
)

type Postgres struct {
	db *sql.DB
}

func New(db *sql.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}

func (p *Postgres) GetCar(req model.GetUserCarsReq) ([]model.Car, error) {
	rows, err := p.db.Query(`
			SELECT 
				id,
				brand,
				model,
				year
			FROM cars
			WHERE
				user_id = $1
	`, req.UserID)
	if err != nil {
		return nil, errors.Wrap(&customerrors.ErrNotFound, "no cars")
	}
	defer rows.Close()

	var cars []model.Car

	for rows.Next() {
		var userCar model.Car

		if err = rows.Scan(&userCar.Id, &userCar.Brand, &userCar.Model, &userCar.Year); err != nil {
			return nil, errors.Wrap(&customerrors.ErrNotFound, "scan rows")
		}

		cars = append(cars, userCar)
	}

	return cars, nil
}

func (p *Postgres) GetIdCarsByBrand(req model.GetBrandCarsIDReq) ([]int, error) {
	rows, err := p.db.Query(`
			SELECT 
				id
			FROM cars
			WHERE
				brand = $1
	`, req.Brand)
	if err != nil {
		return nil, errors.Wrap(&customerrors.ErrNotFound, "no cars")
	}
	defer rows.Close()

	var carsID []int

	for rows.Next() {
		var carID int

		if err = rows.Scan(&carID); err != nil {
			return nil, errors.Wrap(&customerrors.ErrNotFound, "scan rows")
		}

		carsID = append(carsID, carID)
	}

	return carsID, nil
}

func (p *Postgres) GetIdCarsByModel(req model.GetModelCarsIDReq) ([]int, error) {
	rows, err := p.db.Query(`
			SELECT 
				id
			FROM cars
			WHERE 
				brand = $1 AND model = $2
	`, req.Brand, req.Model)
	if err != nil {
		return nil, errors.Wrap(&customerrors.ErrNotFound, "no cars")
	}
	defer rows.Close()

	var carsID []int

	for rows.Next() {
		var carID int

		if err = rows.Scan(&carID); err != nil {
			return nil, errors.Wrap(&customerrors.ErrNotFound, "scan rows")
		}

		carsID = append(carsID, carID)
	}

	return carsID, nil
}

func (p *Postgres) GetIdsCarsByUser(req model.GetIDsUserCarsReq) ([]int, error) {
	rows, err := p.db.Query(`
			SELECT 
				id
			FROM cars
			WHERE 
				user_id = $1
	`, req.UserID)
	if err != nil {
		return nil, errors.Wrap(&customerrors.ErrNotFound, "no cars")
	}
	defer rows.Close()

	var carsID []int

	for rows.Next() {
		var carID int

		if err = rows.Scan(&carID); err != nil {
			return nil, errors.Wrap(&customerrors.ErrNotFound, "scan rows")
		}

		carsID = append(carsID, carID)
	}

	return carsID, nil
}
