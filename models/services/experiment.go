package services

import (
	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/models/services/stores"
	"github.com/TerrenceHo/ABFeature/utils/id"
)

type IExperimentStore interface {
	GetAll(queryModifiers []stores.QueryModifier) ([]*models.Experiment, error)
	GetByID(id string) (*models.Experiment, error)
	Insert(exp *models.Experiment) error
	Update(exp *models.Experiment) error
	Delete(id string) error
}

type ExperimentService struct {
	store  IExperimentStore
	logger loggers.ILogger
}

func NewExperimentService(store IExperimentStore, l loggers.ILogger) *ExperimentService {
	return &ExperimentService{
		store:  store,
		logger: l,
	}
}

func (e *ExperimentService) GetAllExperimentsByProject(projectID string) ([]*models.Experiment, error) {
	queryModifiers := []stores.QueryModifier{
		stores.QueryMod("project_id", stores.EQ, projectID),
	}
	return e.store.GetAll(queryModifiers)
}

func (e *ExperimentService) GetExperimentByID(experimentID string) (*models.Experiment, error) {
	if err := e.idIsValid(experimentID); err != nil {
		return nil, err
	}
	experiment, err := e.store.GetByID(experimentID)
	if err != nil {
		if err == stores.ErrNoExperimentFound {
			return nil, ErrExperimentNotFound
		}
		return nil, err
	}
	return experiment, nil
}

func (e *ExperimentService) AddExperiment(experiment *models.Experiment) (*models.Experiment, error) {
	experiment.ID = id.New()
	errs := experiment.Validate()
	if len(errs) != 0 {
		e.logger.Info("ExperimentService.AddExperiment -- validation failed.  Errors: ", errs)
		return nil, ErrExperimentValidation
	}
	if err := e.store.Insert(experiment); err != nil {
		return nil, err
	}
	return experiment, nil
}

func (e *ExperimentService) UpdateExperiment(experiment *models.Experiment) (*models.Experiment, error) {
	old_experiment, err := e.GetExperimentByID(experiment.ID)
	if err != nil {
		return nil, err
	}

	if experiment.Name != "" {
		old_experiment.Name = experiment.Name
	}
	if experiment.Description != "" {
		old_experiment.Description = experiment.Description
	}
	if experiment.Percentage != -1 {
		old_experiment.Percentage = experiment.Percentage
	}
	old_experiment.Enabled = experiment.Enabled

	errs := old_experiment.Validate()
	if len(errs) != 0 {
		// log errors
		return nil, ErrExperimentValidation
	}
	if err := e.store.Update(old_experiment); err != nil {
		return nil, err
	}
	return old_experiment, nil
}

func (e *ExperimentService) DeleteExperiment(experimentID string) error {
	if err := e.idIsValid(experimentID); err != nil {
		return err
	}
	return e.store.Delete(experimentID)
}

func (e *ExperimentService) idIsValid(id string) error {
	if id == "" {
		return ErrIdInvalid
	}
	return nil
}
