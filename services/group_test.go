package services

import (
	"testing"

	"github.com/TerrenceHo/ABFeature/mocks"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/stores"
	"github.com/stretchr/testify/assert"
)

func newGroupService(store *mocks.IGroupStore) *GroupService {
	logger := mocks.Logger{}
	return NewGroupService(store, logger)
}

func TestGetAllGroups(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IGroupStore)
	service := newGroupService(store)

	store.On("GetAll").Return([]*models.Group{}, nil)
	groups, err := service.GetAllGroups()
	assert.Nil(err, "There should be no error returning all groups.")
	assert.NotNil(groups, "All Groups should not return nil.")

	store.AssertExpectations(t)
}

func TestGetGroupByID(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IGroupStore)
	service := newGroupService(store)

	test_group := models.Group{
		ID:   "groupID",
		Name: "Group Name",
	}
	store.On("GetByID", "groupID").Return(&test_group, nil)
	group, err := service.GetGroupByID("groupID")
	assert.Nil(err, "There should be no error returning group by ID")
	assert.Equal(group.ID, test_group.ID, "Returned IDs should match.")

	store.On("GetByID", "invalidID").Return(nil, stores.ErrNoGroupFound)
	group, err = service.GetGroupByID("invalidID")
	assert.Nil(group, "If invalid, group should be nil.")
	assert.EqualError(ErrGroupNotFound, err.Error(), "Errors should be equal if not found.")

	store.AssertExpectations(t)
}

func TestAddGroup(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IGroupStore)
	service := newGroupService(store)

	badGroup := models.Group{}
	group, err := service.AddGroup(&badGroup)
	assert.EqualError(ErrGroupValidation, err.Error(), "Group should not be validated.")
	assert.Nil(group, "No group should be returned.")

	validGroup := models.Group{
		Name:        "Group Name",
		Description: "Descript me",
	}
	store.On("Insert", &validGroup).Return(nil)
	group, err = service.AddGroup(&validGroup)
	assert.Nil(err, "No error should be returned.")
	assert.NotEmpty(group.ID, "ID should have been created.")
	assert.Equal(group.Name, validGroup.Name)
	assert.Equal(group.Description, validGroup.Description)

	store.AssertExpectations(t)
}

func TestUpdateGroup(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IGroupStore)
	service := newGroupService(store)

	test_group := models.Group{
		ID:          "invalidID",
		Name:        "Group Name",
		Description: "Descript me",
	}
	store.On("GetByID", "invalidID").Return(nil, stores.ErrNoGroupFound)
	group, err := service.UpdateGroup(&test_group)
	assert.EqualError(ErrGroupNotFound, err.Error())
	assert.Nil(group, "Group should be nil.")

	test_group.ID = "validID"
	new_group := models.Group{
		ID:   "validID",
		Name: "Updated Group Name",
	}
	updated_group := models.Group{
		ID:          "validID",
		Name:        "Updated Group Name",
		Description: "Descript me",
	}
	store.On("GetByID", "validID").Return(&test_group, nil)
	store.On("Update", &updated_group).Return(nil)
	_, err = service.UpdateGroup(&new_group)
	assert.Nil(err, "There should be no error.")

	store.AssertExpectations(t)
}

func TestDeleteGroup(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IGroupStore)
	service := newGroupService(store)

	err := service.DeleteGroup("")
	assert.EqualError(ErrIdInvalid, err.Error())

	store.On("Delete", "validID").Return(nil)
	err = service.DeleteGroup("validID")
	assert.Nil(err, "Error should be nil")

	store.AssertExpectations(t)
}
