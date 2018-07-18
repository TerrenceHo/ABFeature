package stores

import (
	"database/sql"
	"errors"

	"github.com/TerrenceHo/ABFeature/models"
	"github.com/jmoiron/sqlx"
)

var (
	ErrInvalidExperimentEntry = errors.New("invalid experiment entry")

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

func (es *ExperimentStore) GetAll(queryModifiers []QueryModifier) ([]*models.Experiment, error) {
	experiments := []*models.Experiment{}

	query, vals := generateWhereStatement(&queryModifiers)
	queryString := experimentsGetAllSQL + query
	err := es.db.Select(&experiments, queryString, vals...)
	if err != nil {
		return nil, err
	}

	return experiments, nil
}

func (es *ExperimentStore) GetByID(id string) (*models.Experiment, error) {
	experiment, err := es.getBy(experimentsGetByIDSQL, id)
	return experiment, err
}

func (es *ExperimentStore) GetCount(queryModifiers []QueryModifier) (int, error) {
	var count int

	query, vals := generateWhereStatement(&queryModifiers)
	queryString := experimentsGetCountSQL + query

	err := es.db.Get(&count, queryString, vals...)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func (es *ExperimentStore) Insert(exp *models.Experiment) error {
	row := es.db.QueryRow(
		experimentsInsertSQL,
		exp.ID,
		exp.Name,
		exp.Description,
		exp.Percentage,
		exp.Enabled,
		exp.ProjectID,
	)
	if err := row.Scan(&exp.CreatedAt, &exp.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func (es *ExperimentStore) Update(exp *models.Experiment) error {
	row := es.db.QueryRow(
		experimentsUpdateSQL,
		exp.ID,
		exp.Name,
		exp.Description,
		exp.Percentage,
		exp.Enabled,
		exp.ProjectID,
	)
	if err := row.Scan(&exp.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func (es *ExperimentStore) Delete(id string) error {
	_, err := es.db.Exec(experimentsDeleteSQL, id)

	if err != nil {
		return ErrInvalidExperimentEntry
	}

	return nil
}

func (es *ExperimentStore) getBy(query string, args interface{}) (*models.Experiment, error) {
	var exp models.Experiment

	if err := es.db.Get(&exp, query, args); err != nil {
		if err == sql.ErrNoRows {
			err = ErrNoExperimentFound
		}
		return nil, err
	}
	return &exp, nil
}

func (es *ExperimentStore) migrate() {
	es.db.MustExec(experimentsCreateTableSQL)
}
