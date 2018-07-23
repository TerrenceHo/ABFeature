package services

import (
	"testing"

	"github.com/TerrenceHo/ABFeature/mocks"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/stretchr/testify/assert"
)

func newExperimentGroupService(store *mocks.IExperimentGroupStore) *ExperimentGroupService {
	logger := mocks.Logger{}
	return NewExperimentGroupService(store, logger)
}

func TestGetAllGroupsByExperiment(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IExperimentGroupStore)
	service := newExperimentGroupService(store)

	groups, err := service.GetAllGroupsByExperiment("")
	assert.Nil(groups, "Groups should be nil.")
	assert.EqualError(ErrIdInvalid, err.Error(), "ID should be invalid.")

	store.On("GetAllGroupsByExperiment", "validID").Return([]*models.Group{}, nil)
	groups, err = service.GetAllGroupsByExperiment("validID")
	assert.Nil(err, "Error should be nil.")
	assert.Equal(len(groups), 0, "Length of groups should be 0")
	store.AssertExpectations(t)
}

func TestGetAllExperimentsByGroup(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IExperimentGroupStore)
	service := newExperimentGroupService(store)

	experiments, err := service.GetAllExperimentsByGroup("")
	assert.Nil(experiments, "Experiments should be nil.")
	assert.EqualError(ErrIdInvalid, err.Error(), "ID should be invalid.")

	store.On("GetAllExperimentsByGroup", "validID").Return([]*models.Experiment{}, nil)
	experiments, err = service.GetAllExperimentsByGroup("validID")
	assert.Nil(err, "Error should be nil.")
	assert.Equal(len(experiments), 0, "Length of experiments should be 0")

	store.AssertExpectations(t)
}

func TestGetExperimentGetByID(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IExperimentGroupStore)
	service := newExperimentGroupService(store)

	exp_group, err := service.GetExperimentGroupByID("")
	assert.Nil(exp_group, "ExperimentGroup should be nil.")
	assert.EqualError(ErrIdInvalid, err.Error(), "ID should be invalid.")

	eg := models.ExperimentGroup{
		ID:           "validID",
		ExperimentID: "ID",
		GroupID:      "ID",
	}
	store.On("GetByID", "validID").Return(&eg, nil)
	exp_group, err = service.GetExperimentGroupByID("validID")
	assert.Nil(err, "Error should be nil.")
	assert.Equal(eg.ID, exp_group.ID)
	assert.Equal(eg.ExperimentID, exp_group.ExperimentID)
	assert.Equal(eg.GroupID, exp_group.GroupID)

	store.AssertExpectations(t)
}

func TestAddExperimentGroup(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IExperimentGroupStore)
	service := newExperimentGroupService(store)

	invalid_exp_group := models.ExperimentGroup{}
	exp_group, err := service.AddExperimentGroup(&invalid_exp_group)
	assert.Nil(exp_group)
	assert.EqualError(ErrExperimentGroupValidation, err.Error())

	valid_exp_group := models.ExperimentGroup{
		ID:           "validID",
		ExperimentID: "validExperimentID",
		GroupID:      "validGroupID",
	}
	store.On("Insert", &valid_exp_group).Return(nil)
	eg, err := service.AddExperimentGroup(&valid_exp_group)
	assert.Nil(err)
	assert.Equal(eg.ID, valid_exp_group.ID)
	assert.Equal(eg.ExperimentID, valid_exp_group.ExperimentID)
	assert.Equal(eg.GroupID, valid_exp_group.GroupID)

	store.AssertExpectations(t)
}

func TestDeleteExperimentGroup(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IExperimentGroupStore)
	service := newExperimentGroupService(store)

	err := service.DeleteExperimentGroup("", "")
	assert.EqualError(ErrIdInvalid, err.Error())
	err = service.DeleteExperimentGroup("validExperimentID", "")
	assert.EqualError(ErrIdInvalid, err.Error())

	store.On("Delete", "validExperimentID", "validGroupID").Return(nil)
	err = service.DeleteExperimentGroup("validExperimentID", "validGroupID")
	assert.Nil(err)

	store.AssertExpectations(t)
}
