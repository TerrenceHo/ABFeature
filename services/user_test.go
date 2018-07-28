package services

import (
	"testing"

	"github.com/TerrenceHo/ABFeature/mocks"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/stores"
	"github.com/stretchr/testify/assert"
)

func newUserService(store IUserStore) *UserService {
	logger := mocks.Logger{}

	return NewUserService(store, logger)
}

func TestGetAllUsers(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IUserStore)
	service := newUserService(store)

	store.On("GetAll").Return([]*models.User{}, nil)
	users, err := service.GetAllUsers()
	assert.Nil(err, "Error should be nil.")
	assert.NotNil(users, "Users should not be nil")

	store.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IUserStore)
	service := newUserService(store)

	user, err := service.GetUserByID("")
	assert.Nil(user, "User should be nil")
	assert.EqualError(ErrIdInvalid, err.Error())

	store.On("GetByID", "invalidUserID").Return(nil, stores.ErrNoUserFound)
	user, err = service.GetUserByID("invalidUserID")
	assert.Nil(user, "User should be nil")
	assert.EqualError(ErrUserNotFound, err.Error())

	test_user := models.User{
		ID:   "userID",
		Name: "Name",
	}
	store.On("GetByID", "userID").Return(&test_user, nil)
	user, err = service.GetUserByID("userID")
	assert.Nil(err, "Error should be nil")
	assert.Equal(user.ID, test_user.ID)
	assert.Equal(user.Name, test_user.Name)

	store.AssertExpectations(t)
}

func TestAddUser(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IUserStore)
	service := newUserService(store)

	badUser := models.User{}
	user, err := service.AddUser(&badUser)
	assert.Nil(user, "User should be nil")
	assert.EqualError(ErrUserValidation, err.Error())

	validUser := models.User{
		ID:          "userID",
		Name:        "User",
		Description: "Description",
	}
	store.On("Insert", &validUser).Return(nil)
	user, err = service.AddUser(&validUser)
	assert.Nil(err, "Err should be nil")
	assert.Equal(user.ID, validUser.ID)
	assert.Equal(user.Name, validUser.Name)
	assert.Equal(user.Description, validUser.Description)

	store.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IUserStore)
	service := newUserService(store)

	invalidUser := models.User{
		ID: "invalidUser",
	}
	store.On("GetByID", "invalidUser").Return(nil, stores.ErrNoUserFound)
	user, err := service.UpdateUser(&invalidUser)
	assert.Nil(user, "User should be nil.")
	assert.EqualError(ErrUserNotFound, err.Error())

	inputUser := models.User{
		ID:   "UserID",
		Name: "New User Name",
	}
	testUser := models.User{
		ID:          "UserID",
		Name:        "User",
		Description: "Description",
	}
	updatedUser := models.User{
		ID:          "UserID",
		Name:        "New User Name",
		Description: "Description",
	}
	store.On("GetByID", "UserID").Return(&testUser, nil)
	store.On("Update", &updatedUser).Return(nil)
	user, err = service.UpdateUser(&inputUser)
	assert.Nil(err, "Error should be nil")

	store.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IUserStore)
	service := newUserService(store)

	err := service.DeleteUser("")
	assert.EqualError(ErrIdInvalid, err.Error())

	store.On("Delete", "validID").Return(nil)
	err = service.DeleteUser("validID")
	assert.Nil(err, "Error should be nil.")

	store.AssertExpectations(t)
}
