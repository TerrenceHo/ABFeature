package stores

import (
	"database/sql"
	"errors"

	"github.com/TerrenceHo/ABFeature/models"
	"github.com/jmoiron/sqlx"
)

var (
	ErrNoUserExperimentFound = errors.New("no user_experiment found")
)

type UserExperimentStore struct {
	db *sqlx.DB
}

func NewUserExperimentStore(db *sqlx.DB) *UserExperimentStore {
	return &UserExperimentStore{
		db: db,
	}
}

func (ues *UserExperimentStore) GetAllUsersByExperiment(experiment_id string) ([]*models.User, error) {
	users := []*models.User{}

	err := ues.db.Select(&users, usersExperimentsGetAllByExperimentSQL, experiment_id)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ues *UserExperimentStore) GetAllExperimentsByUser(user_id string) ([]*models.Experiment, error) {
	experiments := []*models.Experiment{}

	err := ues.db.Select(&experiments, usersExperimentsGetAllByUserSQL, user_id)
	if err != nil {
		return nil, err
	}
	return experiments, nil
}

func (ues *UserExperimentStore) GetById(id string) (*models.UserExperiment, error) {
	user_exp, err := ues.getBy(usersExperimentsGetByIDSQL, id)
	return user_exp, err
}

func (ues *UserExperimentStore) GetByUserAndExperiment(user_id, exp_id string) (*models.UserExperiment, error) {
	user_exp, err := ues.getBy(usersExperimentsGetByUserAndExperimentSQL,
		user_id,
		exp_id,
	)
	return user_exp, err
}

func (ues *UserExperimentStore) Insert(user_exp *models.UserExperiment) error {
	row := ues.db.QueryRow(
		usersExperimentsInsertSQL,
		user_exp.ID,
		user_exp.UserID,
		user_exp.ExperimentID,
	)
	if err := row.Scan(&user_exp.CreatedAt, &user_exp.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (ues *UserExperimentStore) Delete(user_id, exp_id string) error {
	_, err := ues.db.Exec(usersExperimentsDeleteSQL, user_id, exp_id)
	if err != nil {
		return ErrNoUserExperimentFound
	}
	return nil
}

func (ues *UserExperimentStore) getBy(query string, args ...interface{}) (*models.UserExperiment, error) {
	var user_exp models.UserExperiment
	if err := ues.db.Get(&user_exp, query, args...); err != nil {
		if err == sql.ErrNoRows {
			err = ErrNoUserExperimentFound
		}
		return nil, err
	}
	return &user_exp, nil
}

func (ues *UserExperimentStore) migrate() {
	ues.db.MustExec(usersExperimentsCreateTableSQL)
}
