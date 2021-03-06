package services

import (
	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/stores"
	"github.com/TerrenceHo/ABFeature/utils/id"
)

type IExperimentGroupStore interface {
	GetAllGroupsByExperiment(experiment_id string) ([]*models.Group, error)
	GetAllExperimentsByGroup(group_id string) ([]*models.Experiment, error)
	GetByID(id string) (*models.ExperimentGroup, error)
	GetByExperimentAndGroup(experimentID, groupID string) (*models.ExperimentGroup, error)
	Insert(exp_group *models.ExperimentGroup) error
	Delete(experimentID, groupID string) error
}

type ExperimentGroupService struct {
	store  IExperimentGroupStore
	logger loggers.ILogger
}

func NewExperimentGroupService(store IExperimentGroupStore, l loggers.ILogger) *ExperimentGroupService {
	return &ExperimentGroupService{
		store:  store,
		logger: l,
	}
}

func (eg *ExperimentGroupService) GetAllGroupsByExperiment(experimentID string) ([]*models.Group, error) {
	if err := eg.idIsValid(experimentID); err != nil {
		return nil, err
	}
	return eg.store.GetAllGroupsByExperiment(experimentID)
}

func (eg *ExperimentGroupService) GetAllExperimentsByGroup(groupID string) ([]*models.Experiment, error) {
	if err := eg.idIsValid(groupID); err != nil {
		return nil, err
	}
	return eg.store.GetAllExperimentsByGroup(groupID)
}

func (eg *ExperimentGroupService) GetExperimentGroupByID(id string) (*models.ExperimentGroup, error) {
	if err := eg.idIsValid(id); err != nil {
		return nil, err
	}

	exp_group, err := eg.store.GetByID(id)
	if err != nil {
		if err == stores.ErrNoExperimentGroupFound {
			return nil, ErrExperimentGroupNotFound
		}
		return nil, err
	}
	return exp_group, nil
}

func (eg *ExperimentGroupService) GetByExperimentAndGroup(experimentID, groupID string) (*models.ExperimentGroup, error) {
	if err := eg.idIsValid(experimentID); err != nil {
		return nil, err
	}
	if err := eg.idIsValid(groupID); err != nil {
		return nil, err
	}

	exp_group, err := eg.store.GetByExperimentAndGroup(experimentID, groupID)
	if err != nil {
		if err == stores.ErrNoExperimentGroupFound {
			return nil, ErrExperimentGroupNotFound
		}
		return nil, err
	}
	return exp_group, nil
}

func (eg *ExperimentGroupService) AddExperimentGroup(exp_group *models.ExperimentGroup) (*models.ExperimentGroup, error) {
	exp_group.ID = id.New()
	errs := exp_group.Validate()
	if len(errs) != 0 {
		eg.logger.Info("ExperimentGroupService.AddExperimentGroup -- validate failed. Errors:", errs)
		return nil, ErrExperimentGroupValidation
	}
	if err := eg.store.Insert(exp_group); err != nil {
		return nil, err
	}
	return exp_group, nil
}

func (eg *ExperimentGroupService) DeleteExperimentGroup(experimentID, groupID string) error {
	if err := eg.idIsValid(experimentID); err != nil {
		return err
	}
	if err := eg.idIsValid(groupID); err != nil {
		return err
	}
	return eg.store.Delete(experimentID, groupID)
}

func (eg *ExperimentGroupService) idIsValid(id string) error {
	if id == "" {
		return ErrIdInvalid
	}
	return nil
}
