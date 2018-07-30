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

func (egs *ExperimentGroupStore) GetAllGroupsByExperiment(experiment_id string) ([]*models.Group, error) {
	groups := []*models.Group{}

	err := egs.db.Select(&groups, experimentsGroupsGetAllByExperimentSQL, experiment_id)
	if err != nil {
		if err == sql.ErrNoRows {
			// TODO: replace with proper error
			return nil, err
		}
		return nil, err
	}
	return groups, nil
}

func (egs *ExperimentGroupStore) GetAllExperimentsByGroup(group_id string) ([]*models.Experiment, error) {
	experiments := []*models.Experiment{}

	err := egs.db.Select(&experiments, experimentsGroupsGetAllByGroupSQL, group_id)
	if err != nil {
		return nil, err
	}
	return experiments, nil
}

func (egs *ExperimentGroupStore) GetByID(id string) (*models.ExperimentGroup, error) {
	exp_group, err := egs.getBy(experimentsGroupsGetByIDSQL, id)
	return exp_group, err
}

func (egs *ExperimentGroupStore) GetByExperimentAndGroup(experimentID, groupID string) (*models.ExperimentGroup, error) {
	exp_group, err := egs.getBy(
		experimentsGroupGetByExperimentAndGroupSQL,
		experimentID,
		groupID,
	)
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

func (egs *ExperimentGroupStore) Delete(experimentID, groupID string) error {
	_, err := egs.db.Exec(experimentsGroupsDeleteSQL, experimentID, groupID)
	if err != nil {
		return ErrInvalidExperimentGroupEntry
	}
	return nil
}

func (egs *ExperimentGroupStore) getBy(query string, args ...interface{}) (*models.ExperimentGroup, error) {
	var exp_group models.ExperimentGroup

	if err := egs.db.Get(&exp_group, query, args...); err != nil {
		if err == sql.ErrNoRows {
			err = ErrNoExperimentGroupFound
		}
		return nil, err
	}
	return &exp_group, nil
}

func (egs *ExperimentGroupStore) migrate() {
	egs.db.MustExec(experimentsGroupsCreateTableSQL)
}
