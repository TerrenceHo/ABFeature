package services

import (
	"testing"

	"github.com/TerrenceHo/ABFeature/mocks"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/stores"
	"github.com/stretchr/testify/assert"
)

func newUserGroupService(store IUserGroupStore) *UserGroupService {
	logger := mocks.Logger{}

	return NewUserGroupService(store, logger)
}

func TestGetAllUsersByGroup(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IUserGroupStore)
	service := newUserGroupService(store)

	users, err := service.GetAllUsersByGroup("")
	assert.Nil(users, "Users should be nil")
	assert.EqualError(ErrIdInvalid, err.Error(), "Error should be error ID")

	store.On("GetAllUsersByGroup", "groupID").Return([]*models.User{}, nil)
	users, err = service.GetAllUsersByGroup("groupID")
	assert.Nil(err, "Error should be nil")
	assert.Equal(0, len(users), "Length of users should be zero")

	store.AssertExpectations(t)
}

func TestGetAllGroupsByUser(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IUserGroupStore)
	service := newUserGroupService(store)

	groups, err := service.GetAllGroupsByUser("")
	assert.Nil(groups, "Groups should be nil")
	assert.EqualError(ErrIdInvalid, err.Error(), "Error should be error ID")

	store.On("GetAllGroupsByUser", "userID").Return([]*models.Group{}, nil)
	groups, err = service.GetAllGroupsByUser("userID")
	assert.Nil(err, "Error should be nil")
	assert.Equal(0, len(groups), "Length of groups should be zero")

	store.AssertExpectations(t)
}

func TestGetUserGroupByID(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IUserGroupStore)
	service := newUserGroupService(store)

	user_group, err := service.GetUserGroupByID("")
	assert.Nil(user_group, "UserGroup should be nil")
	assert.EqualError(ErrIdInvalid, err.Error(), "Error should be error ID invalid")

	store.On("GetByID", "invalidID").Return(nil, stores.ErrNoUserGroupFound)
	user_group, err = service.GetUserGroupByID("invalidID")
	assert.Nil(user_group, "UserGroup should be nil.")
	assert.EqualError(ErrUserGroupNotFound, err.Error())

	testUserGroup := models.UserGroup{
		ID:      "validID",
		UserID:  "validUserID",
		GroupID: "validGroupID",
	}
	store.On("GetByID", "validID").Return(&testUserGroup, nil)
	user_group, err = service.GetUserGroupByID("validID")
	assert.Nil(err, "Error should be nil.")
	assert.Equal(testUserGroup.ID, user_group.ID)
	assert.Equal(testUserGroup.UserID, user_group.UserID)
	assert.Equal(testUserGroup.GroupID, user_group.GroupID)

	store.AssertExpectations(t)
}

func TestGetByUserAndGroup(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IUserGroupStore)
	service := newUserGroupService(store)

	user_group, err := service.GetByUserAndGroup("", "")
	assert.Nil(user_group, "UserGroup should be nil")
	assert.EqualError(ErrIdInvalid, err.Error(), "Error should be error ID invalid")

	user_group, err = service.GetByUserAndGroup("valid", "")
	assert.Nil(user_group, "UserGroup should be nil")
	assert.EqualError(ErrIdInvalid, err.Error(), "Error should be error ID invalid")

	store.On("GetByUserAndGroup", "invalidUserID", "invalidGroupID").Return(nil, stores.ErrNoUserGroupFound)
	user_group, err = service.GetByUserAndGroup("invalidUserID", "invalidGroupID")
	assert.Nil(user_group, "UserGroup should be nil")
	assert.EqualError(ErrUserGroupNotFound, err.Error())

	testUserGroup := models.UserGroup{
		ID:      "validID",
		UserID:  "validUserID",
		GroupID: "validGroupID",
	}
	store.On("GetByUserAndGroup", "validUserID", "validGroupID").Return(&testUserGroup, nil)
	user_group, err = service.GetByUserAndGroup("validUserID", "validGroupID")
	assert.Nil(err, "Error should be nil")
	assert.Equal(testUserGroup.ID, user_group.ID)
	assert.Equal(testUserGroup.UserID, user_group.UserID)
	assert.Equal(testUserGroup.GroupID, user_group.GroupID)

	store.AssertExpectations(t)
}

func TestAddUserGroup(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IUserGroupStore)
	service := newUserGroupService(store)

	invalidUserGroup := models.UserGroup{}
	user_group, err := service.AddUserGroup(&invalidUserGroup)
	assert.Nil(user_group, "UserGroup should be nil")
	assert.EqualError(ErrUserGroupValidation, err.Error(), "Error should be validation error.")

	testUserGroup := models.UserGroup{
		ID:      "validID",
		UserID:  "validUserID",
		GroupID: "validGroupID",
	}
	store.On("Insert", &testUserGroup).Return(nil)
	user_group, err = service.AddUserGroup(&testUserGroup)
	assert.Nil(err, "Error shoild be nil.")
	assert.Equal(testUserGroup.ID, user_group.ID)
	assert.Equal(testUserGroup.UserID, user_group.UserID)
	assert.Equal(testUserGroup.GroupID, user_group.GroupID)

	store.AssertExpectations(t)
}

func TestDeleteUserGroup(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IUserGroupStore)
	service := newUserGroupService(store)

	err := service.DeleteUserGroup("", "")
	assert.EqualError(ErrIdInvalid, err.Error(), "Error should be error ID invalid")

	err = service.DeleteUserGroup("valid", "")
	assert.EqualError(ErrIdInvalid, err.Error(), "Error should be error ID invalid")

	store.On("Delete", "validUserID", "validGroupID").Return(nil)
	err = service.DeleteUserGroup("validUserID", "validGroupID")
	assert.Nil(err, "Error shoild be nil.")

	store.AssertExpectations(t)
}
