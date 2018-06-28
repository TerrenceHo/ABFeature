package services

import (
	"errors"

	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/utils/id"
)

var (
	ErrProjectValidation = errors.New("Project model validation failed.")

	ErrIdInvalid = errors.New("ID cannot be an empty string.")
)

type IProjectStore interface {
	GetAll() ([]*models.Project, error)
	GetByID(id string) (*models.Project, error)
	Insert(project *models.Project) error
	Update(project *models.Project) error
	Delete(id string) error
}

type ProjectService struct {
	store  IProjectStore
	logger loggers.ILogger
}

func NewProjectService(store IProjectStore, l loggers.ILogger) *ProjectService {
	return &ProjectService{
		store:  store,
		logger: l,
	}
}

func (p *ProjectService) GetAllProjects() ([]*models.Project, error) {
	// run validations
	return p.store.GetAll()
}

func (p *ProjectService) GetProjectByID(id string) (*models.Project, error) {
	// run validations
	if err := p.idIsValid(id); err != nil {
		return nil, err
	}

	return p.store.GetByID(id)
}

func (p *ProjectService) AddProject(project *models.Project) (*models.Project, error) {
	project.ID = id.New() // Create a new unique ID for the new project.
	errs := project.Validate()
	if len(errs) != 0 {
		// log errors
		return nil, ErrProjectValidation
	}
	if err := p.store.Insert(project); err != nil {
		return nil, err
	}
	return project, nil
}

func (p *ProjectService) UpdateProject(project *models.Project) error {
	old_project, err := p.GetProjectByID(project.ID)
	if err != nil {
		return err
	}

	// Bad way of doing this
	if project.Name != "" {
		old_project.Name = project.Name
	}
	if project.Description != "" {
		old_project.Description = project.Description
	}

	errs := old_project.Validate()
	if len(errs) != 0 {
		// log errors
		return ErrProjectValidation
	}
	return p.store.Update(old_project)
}

func (p *ProjectService) DeleteProject(id string) error {
	if err := p.idIsValid(id); err != nil {
		return err
	}
	return p.store.Delete(id)
}

func (p *ProjectService) idIsValid(id string) error {
	if id == "" {
		return ErrIdInvalid
	}
	return nil
}
