// TODO Write Documentation for controllers
//
// TODO standard errors and return codes
package controllers

import (
	"net/http"

	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/models/services"
	"github.com/TerrenceHo/ABFeature/models/services/stores"
	"github.com/labstack/echo"
)

// Interface for services, to interact with database related functions
type IProjectService interface {
	GetAllProjects() ([]*models.Project, error)
	GetProjectByID(projectID string) (*models.Project, error)
	AddProject(project *models.Project) (*models.Project, error)
	UpdateProject(project *models.Project) (*models.Project, error)
	DeleteProject(projectID string) error
}

type ProjectController struct {
	service           IProjectService
	experimentService IExperimentService
	logger            loggers.ILogger
}

func NewProjectController(ps IProjectService, es IExperimentService,
	l loggers.ILogger) *ProjectController {
	return &ProjectController{
		service: ps,
		logger:  l,
	}
}

// Mounts Routes to the Echo Router
func (pc *ProjectController) MountRoutes(g *echo.Group) {
	g.GET("", pc.GetProjects)
	g.POST("", pc.CreateProject)
	g.PUT("", pc.UpdateProject)
	g.DELETE("", pc.DeleteProject)
	g.GET("/experiments", pc.GetAllExperiments)
}

// Wrapper depending on existance of project query param, holding "ID"
func (pc *ProjectController) GetProjects(c echo.Context) error {
	project_id := c.QueryParam("project")
	if project_id != "" {
		return pc.GetProject(c)
	}
	return pc.GetAllProjects(c)
}

// Route -- /projects GET
//
// Input -- Requires no input parameters.
//
// Output -- Returns all projects, as an array of JSON objects, with all objects
// under the key "data".  If an error
// occured, then StatusInternalServerError is returned, with the error
// description. Otherwise, returns with a 200 status request.
func (pc *ProjectController) GetAllProjects(c echo.Context) error {
	projects, err := pc.service.GetAllProjects()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": projects,
	})
}

// Route -- /projects?project=project_id GET
//
// Input -- requires no input, only a query parameter of a valid Project ID.
//
// Output -- Return JSON object of Project with specified Project ID.  If
// successful, it returns a 200 status response. If the ID provided was invalid,
// it returns with a StatusBadRequest(400).  If the project was not found, then
// StatusNotFound(404) is returned.  Otherwise, a StatusInternalServerError is
// returned, indicating a server failure. The appropriate error message is
// provided under the key "message".
func (pc *ProjectController) GetProject(c echo.Context) error {
	project_id := c.QueryParam("project")

	project, err := pc.service.GetProjectByID(project_id)

	if err != nil {
		if err == services.ErrIdInvalid {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else if err == stores.ErrNoProjectFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": project,
	})
}

// Route -- /projects POST
//
// Input -- Requires a valid Project JSON Object, and creates and inserts the
// Project into the database.
//
// Output -- Returns the created object with ID, and timestamp fields
// instantiated, with a StatusOK(200) response under the key "data".  If the
// provided JSON object was invalid or malformed, it returns a
// StatusBadRequest(400) with the appropriate message under the key "message".
func (pc *ProjectController) CreateProject(c echo.Context) error {
	var body models.Project

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	project, err := pc.service.AddProject(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": project,
	})
}

// Route -- /projects PUT
//
// Input -- Requires a valid Project JSON object, and the ID of the project must
// correspond to an already existing project.  Takes the Project and updates its
// fields.
//
// Output -- Returns the updated object with updated fields, including the
// UpdatedAt timestamp field, with a StatusOK(200) response under the key
// "data". If the provided JSON object was invalid or malformed, it returns a
// StatusBadRequest(400) under the key "message".
func (pc *ProjectController) UpdateProject(c echo.Context) error {
	var body models.Project
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	project, err := pc.service.UpdateProject(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": project,
	})
}

// Route -- /projects DELETE
//
// Input -- requires a valid ProjectID query parameter.  Deletes the project
// with the associated Project ID.
//
// Output -- Returns StatusNoContent(204) if successful, otherwise returns
// StatusBadRequest(400) with the appropriate message under "message"
func (pc *ProjectController) DeleteProject(c echo.Context) error {
	project_id := c.QueryParam("project")
	err := pc.service.DeleteProject(project_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

// TODO Add get all experiments, utilizing experiments service
func (pc *ProjectController) GetAllExperiments(c echo.Context) error {
	project_id := c.QueryParam("project")
	if project_id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No Project ID")
	}

	experiments, err := pc.experimentService.GetAllExperimentsByProject(project_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": experiments,
	})

	return nil
}
