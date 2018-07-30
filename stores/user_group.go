package stores

import (
	"database/sql"
	"errors"

	"github.com/TerrenceHo/ABFeature/models"
	"github.com/jmoiron/sqlx"
)

var (
	ErrNoUserGroupFound = errors.New("no user_group found")
)

type UserGroupStore struct {
	db *sqlx.DB
}

func NewUserGroupStore(db *sqlx.DB) *UserGroupStore {
	return &UserGroupStore{
		db: db,
	}
}

func (ugs *UserGroupStore) GetAllUsersByGroup(group_id string) ([]*models.User, error) {
	users := []*models.User{}

	err := ugs.db.Select(&users, usersGroupsGetAllByGroupSQL, group_id)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ugs *UserGroupStore) GetAllGroupsByUser(user_id string) ([]*models.Group, error) {
	groups := []*models.Group{}

	err := ugs.db.Select(&groups, usersGroupsGetAllByUserSQL, user_id)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (ugs *UserGroupStore) GetById(id string) (*models.UserGroup, error) {
	user_group, err := ugs.getBy(usersGroupsGetByIDSQL, id)
	return user_group, err
}

func (ugs *UserGroupStore) GetByUserAndGroup(userID, groupID string) (*models.UserGroup, error) {
	user_group, err := ugs.getBy(usersGroupsGetByUserAndGroupSQL,
		userID,
		groupID,
	)
	return user_group, err
}

func (ugs *UserGroupStore) Insert(user_group *models.UserGroup) error {
	row := ugs.db.QueryRow(
		usersGroupsInsertSQL,
		user_group.ID,
		user_group.UserID,
		user_group.GroupID,
	)
	if err := row.Scan(&user_group.CreatedAt, &user_group.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func (ugs *UserGroupStore) Delete(userID, groupID string) error {
	_, err := ugs.db.Exec(usersGroupsDeleteSQL, userID, groupID)
	if err != nil {
		return ErrNoUserGroupFound
	}
	return nil
}

func (ugs *UserGroupStore) getBy(query string, args ...interface{}) (*models.UserGroup, error) {
	var user_group models.UserGroup

	if err := ugs.db.Get(&user_group, query, args...); err != nil {
		if err == sql.ErrNoRows {
			err = ErrNoUserGroupFound
		}
		return nil, err
	}

	return &user_group, nil
}

func (ugs *UserGroupStore) migrate() {
	ugs.db.MustExec(usersGroupsCreateTableSQL)
}
