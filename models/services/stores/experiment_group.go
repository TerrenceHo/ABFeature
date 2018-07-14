package stores

import (
	"database/sql"
	"errors"

	"github.com/TerrenceHo/ABFeature/models"
	"github.com/jmoiron/sqlx"
)

var (
	ErrInvalidExperimentGroupEntry = errors.New("invalid experiment_group entry")

	ErrNoExperimentGroupFound = errors.New("no experiment_group found")
)

type ExperimentGroupStore struct {
	db *sqlx.DB
}

func NewExperimentGroupStore(db *sqlx.DB) *ExperimentGroupStore {
	return &ExperimentGroupStore{
		db: db,
	}
}

func (egs *ExperimentGroupStore) GetAll(queryModifiers []QueryModifier) ([]*models.ExperimentGroup, error) {
	experiment_groups := []*models.ExperimentGroup{}
	var queryString string
	query, vals := generateWhereStatement(&queryModifiers)
	if len(queryModifiers) > 0 {
		queryString = experimentsGroupsGetAllSQL + query
	} else {
		queryString = ""
	}
	err := egs.db.Select(&experiment_groups, queryString, vals)
	if err != nil {
		return nil, err
	}

	return experiment_groups, nil
}

func (egs *ExperimentGroupStore) GetByID(id string) (*models.Experiment, error) {
	exp_group, err := egs.getBy(experimentsGroupsGetByIDSQL, id)
	return exp_group, err
}

func (egs *ExperimentGroupStore) Insert(exp_group *models.ExperimentGroup) error {
	row := egs.db.QueryRow(
		experimentsGroupsInsertSQL,
		exp_group.ID,
		exp_group.ExperimentID,
		exp_group.GroupID,
	)
	if err := row.Scan(&exp_group.CreatedAt, &exp_group.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func (egs *ExperimentGroupStore) Delete(id string) error {
	_, err := es.db.Exec(experimentsGroupsDeleteSQL, id)
	if err != nil {
		return ErrInvalidExperimentGroupEntry
	}
	return nil
}

func (egs *ExperimentGroupStore) getBy(query string, args interface{}) (*models.ExperimentGroup, error) {
	var exp_group models.ExperimentGroup

	if err := egs.db.Get(&exp, query, args); err != nil {
		if err == sql.ErrNoRows {
			err = ErrNoExperimentGroupFound
		}
		return nil, err
	}
	return &exp_group, nil
}
