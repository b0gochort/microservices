package storage

import (
	"database/sql"

	"github.com/b0gochort/microservices/internal/model"
	"github.com/b0gochort/microservices/pkg/customerrors"
	"github.com/lib/pq"
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

func (p *Postgres) GetEnginesByID(req model.GetEnginesByIDReq) ([]model.Engine, error) {
	rows, err := p.db.Query(`
    		SELECT
        		type,
        		horsepower
   			FROM engines
    		WHERE id = ANY($1)
			`, pq.Array(req.CarsID))
	if err != nil {
		return nil, errors.Wrap(&customerrors.ErrNotFound, "no cars")
	}
	defer rows.Close()

	var responseReq []model.Engine

	for rows.Next() {
		var userEngine model.Engine

		if err = rows.Scan(&userEngine.Type, &userEngine.Horsepower); err != nil {
			return nil, errors.Wrap(&customerrors.ErrNotFound, "scan rows")
		}

		responseReq = append(responseReq, userEngine)
	}

	return responseReq, nil
}
