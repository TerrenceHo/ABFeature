package services

import (
	"testing"

	"github.com/TerrenceHo/ABFeature/mocks"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/models/services/stores"
	"github.com/stretchr/testify/assert"
)

var (
	testProject = models.Project{
		ID:          "testProject1",
		Name:        "Testing Project",
		Description: "Description about project here",
	}
)

func newProjectService(store *mocks.IProjectStore) *ProjectService {
	logger := new(mocks.ILogger)
	return NewProjectService(store, logger)
}

func TestGetAll(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IProjectStore)
	service := newProjectService(store)

	// store.On("GetAll").Return([]*models.Project{}, nil)
	// noProjects, err := service.GetAll()
	// assert.Equal(0, len(noProjects), "No Projects should be returned when there are none.")
	// assert.Nil(err, "There should be no error returning zero projects.")

	store.On("GetAll").Return([]*models.Project{&testProject}, nil)
	oneProject, err := service.GetAllProjects()
	assert.Equal(1, len(oneProject), "There should only be one project.")
	assert.Nil(err, "There should be no error returning one project.")

	store.AssertExpectations(t)
}

func TestGetProjectByID(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IProjectStore)
	service := newProjectService(store)

	store.On("GetByID", "invalidID").Return(nil, stores.ErrNoProjectFound)
	noProj, err := service.GetProjectByID("invalidID")
	assert.Nil(noProj, "Querying with invalid ID should return no Projects.")
	assert.EqualError(err, "no project found")

	store.On("GetByID", testProject.ID).Return(&testProject, nil)
	project, err := service.GetProjectByID(testProject.ID)
	assert.Nil(err, "For a valid car ID, error should be nil")
	assert.Equal(project.ID, testProject.ID)

	store.AssertExpectations(t)
}

func TestAddProject(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IProjectStore)
	service := newProjectService(store)

	invalidProject := models.Project{
		ID:          "",
		Name:        "",
		Description: "This Project should be invalid",
	}
	noProject, err := service.AddProject(&invalidProject)
	assert.Nil(noProject, "Project should not be valid when returned with err.")
	assert.EqualError(err, "Project model validation failed.")

	store.On("Insert", &testProject).Return(nil)
	retProject, err := service.AddProject(&testProject)
	assert.Nil(err, "Error should not exist when inserting a Project")
	assert.NotEmpty(retProject.ID, "ID should not be empty after being created.")
	assert.Equal(retProject.Name, testProject.Name, "Returned name should be equal")
	assert.Equal(retProject.Description, testProject.Description, "Returned Description should be equal")

	store.AssertExpectations(t)
}

func TestUpdateProject(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IProjectStore)
	service := newProjectService(store)

	invalidProject := models.Project{
		ID:          "invalidID",
		Name:        "",
		Description: "This Project should be invalid",
	}

	store.On("GetByID", "invalidID").Return(nil, stores.ErrNoProjectFound)
	err := service.UpdateProject(&invalidProject)
	assert.EqualError(err, "no project found")

	invalidProject.ID = ""
	err = service.UpdateProject(&invalidProject)
	assert.EqualError(err, "ID cannot be an empty string.")

	validInput := models.Project{
		ID:          testProject.ID,
		Name:        "",
		Description: "New Description",
	}

	updatedProject := models.Project{
		ID:          testProject.ID,
		Name:        testProject.Name,
		Description: validInput.Description,
	}
	store.On("GetByID", testProject.ID).Return(&testProject, nil)
	store.On("Update", &updatedProject).Return(nil)
	err = service.UpdateProject(&validInput)
	assert.Nil(err, "Project updating correctly should return nil")

	store.AssertExpectations(t)
}

func TestDeleteProject(t *testing.T) {
	assert := assert.New(t)
	store := new(mocks.IProjectStore)
	service := newProjectService(store)

	err := service.DeleteProject("")
	assert.EqualError(err, "ID cannot be an empty string.")

	store.On("Delete", "validID").Return(nil)
	err = service.DeleteProject("validID")
	assert.Nil(err, "There should be no error with valid ID.")
	store.AssertExpectations(t)
}
