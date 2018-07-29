package stores

import (
	"database/sql"
	"errors"

	"github.com/TerrenceHo/ABFeature/models"
	"github.com/jmoiron/sqlx"
)

var (
	ErrInvalidUserEntry = errors.New("invalid user entry")

	ErrNoUserFound = errors.New("no user found")
)

type UserStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) GetAll() ([]*models.User, error) {
	users := []*models.User{}
	if err := us.db.Select(&users, usersGetAllSQL); err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserStore) GetByID(id string) (*models.User, error) {
	user, err := us.getBy(usersGetByIDSQL, id)
	return user, err
}

func (us *UserStore) Insert(user *models.User) error {
	row := us.db.QueryRow(
		usersInsertSQL,
		user.ID,
		user.Name,
		user.Description,
	)
	if err := row.Scan(&user.CreatedAt, &user.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (us *UserStore) Update(user *models.User) error {
	row := us.db.QueryRow(
		usersUpdateSQL,
		user.ID,
		user.Name,
		user.Description,
	)
	if err := row.Scan(&user.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func (us *UserStore) Delete(id string) error {
	_, err := us.db.Exec(usersDeleteSQL, id)

	if err != nil {
		return ErrInvalidUserEntry
	}
	return nil
}

func (us *UserStore) getBy(query string, args interface{}) (*models.User, error) {
	var user models.User

	if err := us.db.Get(&user, query, args); err != nil {
		if err == sql.ErrNoRows {
			err = ErrNoUserFound
		}
		return nil, err
	}
	return &user, nil
}

func (us *UserStore) migrate() {
	us.db.MustExec(usersCreateTableSQL)
}
