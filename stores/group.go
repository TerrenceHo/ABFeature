package stores

import (
	"database/sql"
	"errors"

	"github.com/TerrenceHo/ABFeature/models"
	"github.com/jmoiron/sqlx"
)

type GroupStore struct {
	db *sqlx.DB
}

var (
	ErrInvalidGroupEntry = errors.New("invalid group entry")

	ErrNoGroupFound = errors.New("no group found")
)

func NewGroupStore(db *sqlx.DB) *GroupStore {
	return &GroupStore{
		db: db,
	}
}

// func (gs *GroupStore) GetAll(queryModifiers []QueryModifier) ([]*models.Group, error) {
func (gs *GroupStore) GetAll() ([]*models.Group, error) {
	groups := []*models.Group{}

	// query, vals := generateWhereStatement(&queryModifiers)
	// queryString := groupsGetAllSQL + query
	err := gs.db.Select(&groups, groupsGetAllSQL)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (gs *GroupStore) GetByID(id string) (*models.Group, error) {
	group, err := gs.getBy(groupsGetByIDSQL, id)
	return group, err
}

func (gs *GroupStore) Insert(group *models.Group) error {
	row := gs.db.QueryRow(
		groupsInsertSQL,
		group.ID,
		group.Name,
		group.Description,
	)

	if err := row.Scan(&group.CreatedAt, &group.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func (gs *GroupStore) Update(group *models.Group) error {
	row := gs.db.QueryRow(
		groupsUpdateSQL,
		group.ID,
		group.Name,
		group.Description,
	)
	if err := row.Scan(&group.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func (gs *GroupStore) Delete(id string) error {
	_, err := gs.db.Exec(groupsDeleteSQL, id)

	if err != nil {
		return ErrInvalidGroupEntry
	}

	return nil
}

func (gs *GroupStore) getBy(query string, args interface{}) (*models.Group, error) {
	var group models.Group

	if err := gs.db.Get(&group, query, args); err != nil {
		if err == sql.ErrNoRows {
			err = ErrNoGroupFound
		}
		return nil, err
	}
	return &group, nil
}

func (gs *GroupStore) migrate() {
	gs.db.MustExec(groupsCreateTableSQL)
}
