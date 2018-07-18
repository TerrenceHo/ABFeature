package services

import (
	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/stores"
	"github.com/TerrenceHo/ABFeature/utils/id"
)

type IExperimentGroupStore interface {
	GetAll(queryModifiers []stores.QueryModifier) ([]*models.ExperimentGroup, error)
	GetByID(id string) (*models.ExperimentGroup, error)
	Insert(exp_group *models.ExperimentGroup) error
	Delete(id string) error
}

type ExperimentGroupService struct {
	store  IExperimentGroupStore
	logger loggers.ILogger
}

func NewExperimentGroupStore(store IExperimentGroupStore, l loggers.ILogger) *ExperimentGroupService {
	return &ExperimentGroupService{
		store:  store,
		logger: l,
	}
}

// func (eg *ExperimentGroupService) GetAll
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

func (eg *ExperimentGroupService) DeleteExperimentGroup(id string) error {
	if err := eg.idIsValid(id); err != nil {
		return err
	}
	return eg.store.Delete(id)
}

func (eg *ExperimentGroupService) idIsValid(id string) error {
	if id == "" {
		return ErrIdInvalid
	}
	return nil
}
