package services

import (
	"testing"

	"github.com/TerrenceHo/ABFeature/mocks"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/models/services/stores"
	"github.com/stretchr/testify/assert"
)

func newExperimentService(store *mocks.IExperimentStore) *ExperimentService {
	logger := mocks.Logger{}
	return NewExperimentService(store, logger)
}

func TestExperimentGetAll(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IExperimentStore)
	service := newExperimentService(store)

	testExperiment := models.Experiment{
		ID:          "validID",
		Name:        "Test Experiment",
		Description: "Test",
		ProjectID:   "TestProjectID",
	}
	store.On("GetAll", []stores.QueryModifier{
		stores.QueryMod("project_id", stores.EQ, "TestProjectID"),
	}).Return([]*models.Experiment{&testExperiment}, nil)

	experiments, err := service.GetAllExperimentsByProject("TestProjectID")
	assert.Equal(1, len(experiments), "There should only be one project.")
	assert.Nil(err, "Error should be nil.")

	store.AssertExpectations(t)
}

func TestGetExperimentByID(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IExperimentStore)
	service := newExperimentService(store)

	noExp, err := service.GetExperimentByID("")
	assert.Nil(noExp, "Experiment should be nil on error.")
	assert.EqualError(err, ErrIdInvalid.Error())

	store.On("GetByID", "invalidID").Return(nil, stores.ErrNoExperimentFound)
	noExp, err = service.GetExperimentByID("invalidID")
	assert.Nil(noExp, "Experiment should be nil on error")
	assert.EqualError(err, ErrExperimentNotFound.Error())

	testExperiment := models.Experiment{
		ID:          "validID",
		Name:        "Test Experiment",
		Description: "Test",
	}
	store.On("GetByID", "validID").Return(&testExperiment, nil)
	experiment, err := service.GetExperimentByID("validID")
	assert.Nil(err, "For a valid experiment ID, error should be nil")
	assert.Equal(experiment.ID, testExperiment.ID)
}

func TestAddExperiment(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IExperimentStore)
	service := newExperimentService(store)

	invalidExperiment := models.Experiment{
		Description: "This should be invalid",
		Percentage:  6,
	}
	noExp, err := service.AddExperiment(&invalidExperiment)
	assert.Nil(noExp, "Project should be valid when return with error.")
	assert.EqualError(err, ErrExperimentValidation.Error())

	testExperiment := models.Experiment{
		Name:       "Valid",
		Percentage: 0.5,
		ProjectID:  "ValidProjectID",
	}
	store.On("Insert", &testExperiment).Return(nil)
	experiment, err := service.AddExperiment(&testExperiment)
	assert.Nil(err, "Error should be nil when successful insert.")
	assert.NotEmpty(experiment.ID, "ID should not be empty after being created.")

	store.AssertExpectations(t)
}

func TestUpdateExperiment(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IExperimentStore)
	service := newExperimentService(store)

	invalidExperiment := models.Experiment{
		ID:          "invalidID",
		Description: "This should be invalid",
		Percentage:  6,
	}

	store.On("GetByID", "invalidID").Return(nil, stores.ErrNoExperimentFound)
	noExp, err := service.UpdateExperiment(&invalidExperiment)
	assert.EqualError(err, ErrExperimentNotFound.Error())
	assert.Nil(noExp, "Experiment should be nil if there is an error.")

	invalidExperiment.ID = ""
	noExp, err = service.UpdateExperiment(&invalidExperiment)
	assert.EqualError(err, ErrIdInvalid.Error())
	assert.Nil(noExp, "Should be nil if there is any error.")

	// TODO check actual outputs
	store.AssertExpectations(t)
}

func TestDeleteExperiment(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IExperimentStore)
	service := newExperimentService(store)

	err := service.DeleteExperiment("")
	assert.EqualError(err, ErrIdInvalid.Error())

	store.On("Delete", "validID").Return(nil)
	err = service.DeleteExperiment("validID")
	assert.Nil(err, "There should be no error with a valid ID")

	store.AssertExpectations(t)
}
