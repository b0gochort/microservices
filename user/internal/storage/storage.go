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

func (p *Postgres) UserExists(req model.UserCheckReq) (*model.User, error) {
	var user model.User

	if err := p.db.QueryRow(`
			SELECT 
				id,
				name
			FROM users
			WHERE id = $1
	`, req.UserID).Scan(&user.UserID, &user.UserName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(&customerrors.ErrNotFound, "user not found")
		}
		return nil, errors.Wrap(&customerrors.ErrScan, "scan user")
	}

	return &user, nil
}
