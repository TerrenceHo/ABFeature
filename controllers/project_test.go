package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/TerrenceHo/ABFeature/mocks"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/services"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var bad_err error = errors.New("Database failed.")

func newProjectController(ps IProjectService) *ProjectController {
	logger := mocks.Logger{}
	return NewProjectController(ps, logger)
}

func TestGetAllProjects200(t *testing.T) {
	assert := assert.New(t)
	service := new(mocks.IProjectService)
	controller := newProjectController(service)
	e := echo.New()
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(echo.GET, "/", nil)
	c := e.NewContext(req, recorder)

	service.On("GetAllProjects").Return([]*models.Project{}, nil)
	if assert.NoError(controller.GetAllProjects(c)) {
		assert.Equal(http.StatusOK, recorder.Code)
		assert.Equal(`{"data":[]}`, recorder.Body.String())
	}
}

func TestGetAllProjects500(t *testing.T) {
	assert := assert.New(t)
	service := new(mocks.IProjectService)
	controller := newProjectController(service)
	e := echo.New()
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(echo.GET, "/", nil)
	c := e.NewContext(req, recorder)

	service.On("GetAllProjects").Return(nil, bad_err)
	err := controller.GetAllProjects(c)
	herr, ok := err.(*echo.HTTPError)
	if ok {
		assert.Equal(http.StatusInternalServerError, herr.Code)
		assert.Equal(herr.Message, bad_err.Error())
	}

	service.AssertExpectations(t)
}

func TestGetProject200(t *testing.T) {
	assert := assert.New(t)
	service := new(mocks.IProjectService)
	controller := newProjectController(service)
	e := echo.New()
	recorder := httptest.NewRecorder()
	q := make(url.Values)
	q.Set("project", "projectID")
	req := httptest.NewRequest(echo.GET, "/?"+q.Encode(), nil)
	c := e.NewContext(req, recorder)

	testProject := models.Project{
		ID:   "projectID",
		Name: "Project",
	}
	service.On("GetProjectByID", "projectID").Return(&testProject, nil)
	if assert.NoError(controller.GetProject(c)) {
		assert.Equal(http.StatusOK, recorder.Code)

		var project models.Project
		body := map[string]*models.Project{
			"data": &project,
		}
		if err := json.NewDecoder(recorder.Body).Decode(&body); err != nil {
			t.Fatal("JSON failed to decode")
		}
		fmt.Println(body)

		assert.Equal(testProject.ID, body["data"].ID)
		assert.Equal(testProject.Name, body["data"].Name)
		assert.Equal("", body["data"].Description)
	}

	service.AssertExpectations(t)
}

func TestGetProject(t *testing.T) {
	assert := assert.New(t)
	service := new(mocks.IProjectService)
	controller := newProjectController(service)
	e := echo.New()
	recorder := httptest.NewRecorder()
	q := make(url.Values)
	req := httptest.NewRequest(echo.GET, "/?"+q.Encode(), nil)
	c := e.NewContext(req, recorder)

	service.On("GetProjectByID", "").Return(nil, services.ErrIdInvalid)
	err := controller.GetProject(c)
	herr, ok := err.(*echo.HTTPError)
	if ok {
		assert.Equal(http.StatusBadRequest, herr.Code)
		assert.Equal(services.ErrIdInvalid.Error(), herr.Message)
	}
}

func TestGetProject404(t *testing.T) {
	assert := assert.New(t)
	service := new(mocks.IProjectService)
	controller := newProjectController(service)
	e := echo.New()
	recorder := httptest.NewRecorder()
	q := make(url.Values)
	q.Set("project", "invalidProjectID")
	req := httptest.NewRequest(echo.GET, "/?"+q.Encode(), nil)
	c := e.NewContext(req, recorder)

	service.On("GetProjectByID", "invalidProjectID").Return(nil, services.ErrProjectNotFound)
	err := controller.GetProject(c)
	herr, ok := err.(*echo.HTTPError)
	if ok {
		assert.Equal(http.StatusNotFound, herr.Code)
		assert.Equal(services.ErrProjectNotFound.Error(), herr.Message)
	}

	service.AssertExpectations(t)
}

func TestGetProject500(t *testing.T) {
	assert := assert.New(t)
	service := new(mocks.IProjectService)
	controller := newProjectController(service)
	e := echo.New()
	recorder := httptest.NewRecorder()
	q := make(url.Values)
	q.Set("project", "validID")
	req := httptest.NewRequest(echo.GET, "/?"+q.Encode(), nil)
	c := e.NewContext(req, recorder)

	service.On("GetProjectByID", "validID").Return(nil, bad_err)
	err := controller.GetProject(c)
	herr, ok := err.(*echo.HTTPError)
	if ok {
		assert.Equal(http.StatusInternalServerError, herr.Code)
		assert.Equal(bad_err.Error(), herr.Message)
	}
}

// TODO: Test Create Project
func TestCreateProject200(t *testing.T) {}
func TestCreateProject400(t *testing.T) {}

// TODO: Test Update Project
func TestUpdateProject200(t *testing.T) {}
func TestUpdateProject400(t *testing.T) {}

func TestDeleteProject204(t *testing.T) {
	assert := assert.New(t)
	service := new(mocks.IProjectService)
	controller := newProjectController(service)
	e := echo.New()
	recorder := httptest.NewRecorder()
	q := make(url.Values)
	q.Set("project", "validID")
	req := httptest.NewRequest(echo.DELETE, "/?"+q.Encode(), nil)
	c := e.NewContext(req, recorder)

	service.On("DeleteProject", "validID").Return(nil)
	if assert.NoError(controller.DeleteProject(c)) {
		assert.Equal(http.StatusNoContent, recorder.Code)
	}
}

func TestDeleteProject400(t *testing.T) {
	assert := assert.New(t)
	service := new(mocks.IProjectService)
	controller := newProjectController(service)
	e := echo.New()
	recorder := httptest.NewRecorder()
	q := make(url.Values)
	req := httptest.NewRequest(echo.DELETE, "/?"+q.Encode(), nil)
	c := e.NewContext(req, recorder)

	service.On("DeleteProject", "").Return(services.ErrIdInvalid)
	err := controller.DeleteProject(c)
	if herr, ok := err.(*echo.HTTPError); ok {
		assert.Equal(http.StatusBadRequest, herr.Code)
	}
}
