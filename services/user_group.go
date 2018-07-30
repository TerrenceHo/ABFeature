package services

import (
	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/stores"
	"github.com/TerrenceHo/ABFeature/utils/id"
)

type IUserGroupStore interface {
	GetAllUsersByGroup(group_id string) ([]*models.User, error)
	GetAllGroupsByUser(user_id string) ([]*models.Group, error)
	GetByID(id string) (*models.UserGroup, error)
	GetByUserAndGroup(user_id, group_id string) (*models.UserGroup, error)
	Insert(user_group *models.UserGroup) error
	Delete(user_id, group_id string) error
}

type UserGroupService struct {
	store  IUserGroupStore
	logger loggers.ILogger
}

func NewUserGroupService(store IUserGroupStore, l loggers.ILogger) *UserGroupService {
	return &UserGroupService{
		store:  store,
		logger: l,
	}
}

func (ug *UserGroupService) GetAllUsersByGroup(group_id string) ([]*models.User, error) {
	if err := ug.idIsValid(group_id); err != nil {
		return nil, err
	}
	return ug.store.GetAllUsersByGroup(group_id)
}

func (ug *UserGroupService) GetAllGroupsByUser(user_id string) ([]*models.Group, error) {
	if err := ug.idIsValid(user_id); err != nil {
		return nil, err
	}
	return ug.store.GetAllGroupsByUser(user_id)
}

func (ug *UserGroupService) GetUserGroupByID(id string) (*models.UserGroup, error) {
	if err := ug.idIsValid(id); err != nil {
		return nil, err
	}

	user_group, err := ug.store.GetByID(id)
	if err != nil {
		if err == stores.ErrNoUserGroupFound {
			return nil, ErrUserGroupNotFound
		}
		return nil, err
	}
	return user_group, nil
}

func (ug *UserGroupService) GetByUserAndGroup(user_id, group_id string) (*models.UserGroup, error) {
	if err := ug.idIsValid(user_id); err != nil {
		return nil, err
	}
	if err := ug.idIsValid(group_id); err != nil {
		return nil, err
	}

	user_group, err := ug.store.GetByUserAndGroup(user_id, group_id)
	if err != nil {
		if err == stores.ErrNoUserGroupFound {
			return nil, ErrUserGroupNotFound
		}
		return nil, err
	}
	return user_group, nil
}

func (ug *UserGroupService) AddUserGroup(user_group *models.UserGroup) (*models.UserGroup, error) {
	user_group.ID = id.New()
	errs := user_group.Validate()
	if len(errs) != 0 {
		ug.logger.Info("UserGroupService.AddUserGroup -- validate failed. Errors:", errs)
		return nil, ErrUserGroupValidation
	}
	if err := ug.store.Insert(user_group); err != nil {
		return nil, err
	}
	return user_group, nil
}

func (ug *UserGroupService) DeleteUserGroup(user_id, group_id string) error {
	if err := ug.idIsValid(user_id); err != nil {
		return err
	}
	if err := ug.idIsValid(group_id); err != nil {
		return err
	}
	return ug.store.Delete(user_id, group_id)
}

func (ug *UserGroupService) idIsValid(id string) error {
	if id == "" {
		return ErrIdInvalid
	}
	return nil
}
