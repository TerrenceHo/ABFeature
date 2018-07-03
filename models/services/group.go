package services

import (
	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/models/services/stores"
	"github.com/TerrenceHo/ABFeature/utils/id"
)

type IGroupStore interface {
	GetAll() (*models.Group, error)
	GetByID(id string) (*models.Group, error)
	Insert(group *models.Group) error
	Update(group *models.Group) error
	Delete(id string) error
}

type GroupService struct {
	store  IGroupStore
	logger loggers.ILogger
}

func NewGroupService(store IGroupStore, l loggers.ILogger) *GroupService {
	return &GroupService{
		store:  store,
		logger: l,
	}
}

func (g *GroupService) GetAllGroups() (*models.Group, error) {
	// return g.
	return nil, nil
}

func (g *GroupService) GetGroupByID(groupID string) (*models.Group, error) {
	if err := g.idIsValid(groupID); err != nil {
		return nil, err
	}

	group, err := g.store.GetByID(groupID)
	if err != nil {
		if err == stores.ErrNoGroupFound {
			return nil, ErrGroupNotFound
		}
	}

	return group, nil
}

func (g *GroupService) AddGroup(group *models.Group) (*models.Group, error) {
	group.ID = id.New()
	errs := group.Validate()
	if len(errs) != 0 {
		g.logger.Info("GroupService.AddGroup -- validation failed. Errors:", errs)
		return nil, ErrGroupValidation
	}

	if err := g.store.Insert(group); err != nil {
		g.logger.Error("GroupService.AddGroup -- unable to create Group. Error:", err.Error())
		return nil, err
	}
	return group, nil
}

func (g *GroupService) UpdateGroup(group *models.Group) (*models.Group, error) {
	old_group, err := g.GetGroupByID(group.ID)
	if err != nil {
		return nil, err
	}

	if group.Name != "" {
		old_group.Name = group.Name
	}
	if group.Description != "" {
		old_group.Description = group.Description
	}
	errs := old_group.Validate()
	if len(errs) != 0 {
		g.logger.Info("GroupService.UpdateGroup -- validation failed. Errors:", errs)
		return nil, ErrGroupValidation
	}
	if err := g.store.Update(old_group); err != nil {
		g.logger.Error("GroupService.UpdateGroup -- unable to update Group. Error:", err.Error())
		return nil, err
	}

	return old_group, nil
}

func (g *GroupService) DeleteGroup(groupID string) error {
	if err := g.idIsValid(groupID); err != nil {
		return err
	}
	return g.store.Delete(groupID)
}

func (g *GroupService) idIsValid(id string) error {
	if id == "" {
		return ErrIdInvalid
	}

	return nil
}
