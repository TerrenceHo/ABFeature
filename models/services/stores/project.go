package stores

import (
	"database/sql"
	"errors"

	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/utils/id"
	"github.com/jmoiron/sqlx"
)

var (
	ErrInvalidProjectEntry = errors.New("invalid project entry")

	ErrNoProjectFound = errors.New("no project found")
)

type ProjectStore struct {
	db *sqlx.DB
}

func NewProjectStore(db *sqlx.DB) *ProjectStore {
	return &ProjectStore{
		db: db,
	}
}

func (ps *ProjectStore) GetAll() ([]*models.Project, error) {
	projects := []*models.Project{}

	if err := ps.db.Select(&projects, projectsGetAllSQL); err != nil {
		return nil, err
	}
	return projects, nil
}

func (ps *ProjectStore) GetByID(id string) (*models.Project, error) {
	project, err := ps.getBy(projectsGetByIDSQL, id)
	return project, err
}

func (ps *ProjectStore) Insert(project *models.Project) error {
	project.ID = id.New()

	row := ps.db.QueryRow(
		projectsInsertSQL,
		project.ID,
		project.Name,
		project.Description,
	)
	if err := row.Scan(&project.CreatedAt, &project.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (ps *ProjectStore) Update(project *models.Project) error {
	row := ps.db.QueryRow(projectsUpdateSQL,
		project.ID,
		project.Name,
		project.Description,
	)
	if err := row.Scan(&project.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func (ps *ProjectStore) Delete(id string) error {
	_, err := ps.db.Exec(projectsDeleteSQL, id)

	if err != nil {
		return ErrInvalidProjectEntry
	}

	return nil
}

func (ps *ProjectStore) getBy(query string, args interface{}) (*models.Project, error) {
	var project models.Project

	if err := ps.db.Get(&project, query, args); err != nil {
		if err == sql.ErrNoRows {
			err = ErrNoProjectFound
		}
		return nil, err
	}
	return &project, nil
}

func (ps *ProjectStore) migrate() {
	ps.db.MustExec(projectsCreateTableSQL)
}
