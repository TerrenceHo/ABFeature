package services

import (
	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/stores"
	"github.com/TerrenceHo/ABFeature/utils/id"
)

type IUserStore interface {
	GetAll() ([]*models.User, error)
	GetByID(id string) (*models.User, error)
	Insert(user *models.User) error
	Update(user *models.User) error
	Delete(id string) error
}

type UserService struct {
	store  IUserStore
	logger loggers.ILogger
}

func NewUserService(store IUserStore, l loggers.ILogger) *UserService {
	return &UserService{
		store:  store,
		logger: l,
	}
}

func (u *UserService) GetAllUsers() ([]*models.User, error) {
	return u.store.GetAll()
}

func (u *UserService) GetUserByID(userID string) (*models.User, error) {
	if err := u.idIsValid(userID); err != nil {
		return nil, err
	}

	user, err := u.store.GetByID(userID)
	if err != nil {
		if err == stores.ErrNoUserFound {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (u *UserService) AddUser(user *models.User) (*models.User, error) {
	user.ID = id.New()
	errs := user.Validate()
	if len(errs) != 0 {
		u.logger.Info("UserService.AddUser -- validation failed. Errors:", errs)
		return nil, ErrUserValidation
	}

	if err := u.store.Insert(user); err != nil {
		u.logger.Error("UserService.AddUser -- unable to create User. Error:", err.Error())
		return nil, err
	}
	return user, nil
}

func (u *UserService) UpdateUser(user *models.User) (*models.User, error) {
	old_user, err := u.GetUserByID(user.ID)
	if err != nil {
		return nil, err
	}

	errs := old_user.UpdateFields(user)
	if len(errs) != 0 {
		u.logger.Info("UserService.UpdateUser -- validation failed. Errors:", errs)
		return nil, ErrUserValidation
	}
	if err := u.store.Update(old_user); err != nil {
		u.logger.Error("UserService.UpdateUser -- unable to update User. Error:", err.Error())
		return nil, err
	}

	return old_user, nil
}

func (u *UserService) DeleteUser(userID string) error {
	if err := u.idIsValid(userID); err != nil {
		return err
	}
	return u.store.Delete(userID)
}

func (u *UserService) idIsValid(id string) error {
	if id == "" {
		return ErrIdInvalid
	}
	return nil
}
