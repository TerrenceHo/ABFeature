package services

import (
	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/models/services/stores"
	"github.com/TerrenceHo/ABFeature/utils/id"
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
	return p.store.GetAll()
}

func (p *ProjectService) GetProjectByID(projectID string) (*models.Project, error) {
	if err := p.idIsValid(projectID); err != nil {
		return nil, err
	}

	project, err := p.store.GetByID(projectID)
	if err != nil {
		if err == stores.ErrNoProjectFound {
			return nil, ErrProjectNotFound
		}
		return nil, err
	}

	return project, nil
}

func (p *ProjectService) AddProject(project *models.Project) (*models.Project, error) {
	project.ID = id.New() // Create a new unique ID for the new project.
	errs := project.Validate()
	if len(errs) != 0 {
		p.logger.Info("ProjectService.AddProject -- validation failed. Errors: ", errs)
		return nil, ErrProjectValidation
	}
	if err := p.store.Insert(project); err != nil {
		p.logger.Error("ProjectService.AddProject -- unable to create Project. Error:", err.Error())
		return nil, err
	}
	return project, nil
}

func (p *ProjectService) UpdateProject(project *models.Project) (*models.Project, error) {
	old_project, err := p.GetProjectByID(project.ID)
	if err != nil {
		return nil, err
	}

	errs := old_project.UpdateFields(project)
	if len(errs) != 0 {
		p.logger.Info("ProjectService.UpdateProject -- validation failed. Errors:", errs)
		return nil, ErrProjectValidation
	}
	if err := p.store.Update(old_project); err != nil {
		p.logger.Error("ProjectService.UpdateProject -- unable to update Project. Error:", err.Error())
		return nil, err
	}
	return old_project, nil
}

func (p *ProjectService) DeleteProject(projectID string) error {
	if err := p.idIsValid(projectID); err != nil {
		return err
	}
	return p.store.Delete(projectID)
}

func (p *ProjectService) idIsValid(id string) error {
	if id == "" {
		return ErrIdInvalid
	}
	return nil
}
