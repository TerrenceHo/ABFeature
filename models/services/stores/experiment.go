package stores

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	ErrInalidExperimentEntry = errors.New("invalid experiment entry")

	ErrNoExperimentFound = errors.New("no experiment found")
)

type ExperimentStore struct {
	db *sqlx.DB
}

func NewExperimentStore(db *sqlx.DB) *ExperimentStore {
	return &ExperimentStore{
		db: db,
	}
}

// func (es *ExperimentStore) GetAll() ([]*models.Experiment, error) {

// }
