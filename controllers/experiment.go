package controllers

import (
	"net/http"

	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/services"
	"github.com/labstack/echo"
)

// Interface defined for a Experiment Service
type IExperimentService interface {
	GetAllExperimentsByProject(projectID string) ([]*models.Experiment, error)
	GetExperimentByID(experimentID string) (*models.Experiment, error)
	AddExperiment(experiment *models.Experiment) (*models.Experiment, error)
	UpdateExperiment(experiment *models.Experiment) (*models.Experiment, error)
	DeleteExperiment(experimentID string) error
}

// Interface defined for a ExperimentGroupService
type IExperimentGroupService interface {
	GetAllGroupsByExperiment(experimentID string) ([]*models.Group, error)
	GetAllExperimentsByGroup(groupID string) ([]*models.Experiment, error)
	GetExperimentGroupByID(id string) (*models.ExperimentGroup, error)
	AddExperimentGroup(exp_group *models.ExperimentGroup) (*models.ExperimentGroup, error)
	DeleteExperimentGroup(experimentID, groupID string) error
}

type ExperimentController struct {
	service      IExperimentService
	expGrService IExperimentGroupService
	logger       loggers.ILogger
}

// Creates and returns an ExperimentController struct with a ExperimentService
// and logger.
func NewExperimentController(es IExperimentService, egs IExperimentGroupService, l loggers.ILogger) *ExperimentController {
	return &ExperimentController{
		service:      es,
		expGrService: egs,
		logger:       l,
	}
}

// Mounts ExperimentController routes to Echo
func (ec *ExperimentController) MountRoutes(g *echo.Group) {
	g.GET("", ec.GetExperiments)
	g.POST("", ec.CreateExperiment)
	g.PUT("", ec.UpdateExperiment)
	g.DELETE("", ec.DeleteExperiment)
	g.GET("/groups", ec.GetAllExperimentsByGroup)
	g.POST("/groups", ec.CreateExperimentGroup)
	g.DELETE("/groups", ec.DeleteExperimentGroup)
}

// Wrapper function to determine for GetExperiment or GetAllExperiments
func (ec *ExperimentController) GetExperiments(c echo.Context) error {
	experiment_id := c.QueryParam("experiment")
	if experiment_id != "" {
		return ec.GetExperiment(c)
	}
	return ec.GetAllExperiments(c)
}

// Route -- /experiments?experiment=experiment_id
//
// Input -- Requires query parameter of a valid Experiment ID.
//
// Output -- Returns a JSON Object of Experiment with specified Experiment ID.
// If successful it returns a StatusOK(200). If the ID provided was invalid,
// it returns with a StatusBadRequest(400).  If the project was not found, then
// StatusNotFound(404) is returned.  Otherwise, a StatusInternalServerError is
// returned, indicating a server failure. The appropriate error message is
// provided under the key "message".
func (ec *ExperimentController) GetExperiment(c echo.Context) error {
	experiment_id := c.QueryParam("experiment")
	experiment, err := ec.service.GetExperimentByID(experiment_id)
	if err != nil {
		if err == services.ErrIdInvalid {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else if err == services.ErrExperimentNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": experiment,
	})
}

// Route -- /experiments?project=project_id GET
//
// Input -- Requires project_id, to see which project's experiments you want o
// retrieve.
//
// Output -- Returns all experiments associated with the project_id, as an array
// of JSON objects, with all objects under the key "data".  If an error
// occured, then StatusInternalServerError is returned, with the error
// description. Otherwise, returns with a 200 status request.
func (ec *ExperimentController) GetAllExperiments(c echo.Context) error {
	project_id := c.QueryParam("project")
	if project_id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No Project ID")
	}

	experiments, err := ec.service.GetAllExperimentsByProject(project_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": experiments,
	})
}

// Route --/experiments POST
//
// Input -- Requires a valid Experiment JSON Object, and creates and inserts the
// Experiment into the database.
//
// Output -- Returns the created object with ID, and timestamp fields
// instantiated, with a StatusOK(200) response under the key "data".  If the
// provided JSON object was invalid or malformed, it returns a
// StatusBadRequest(400) with the appropriate message under the key "message".
func (ec *ExperimentController) CreateExperiment(c echo.Context) error {
	var body models.Experiment

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	experiment, err := ec.service.AddExperiment(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": experiment,
	})
}

// Route -- /projects PUT
//
// Input -- Requires a valid Experiment JSON object, and the ID of the
// experiment must correspond to an already existing experiment.  Takes
// the Experiment and updates its fields.
//
// Output -- Returns the updated object with updated fields, including the
// UpdatedAt timestamp field, with a StatusOK(200) response under the key
// "data". If the provided JSON object was invalid or malformed, it returns a
// StatusBadRequest(400) under the key "message".
func (ec *ExperimentController) UpdateExperiment(c echo.Context) error {
	var body models.Experiment

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	experiment, err := ec.service.UpdateExperiment(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": experiment,
	})
}

// Route -- /experiments DELETE
//
// Input -- requires a valid Experiment ID query parameter.  Deletes the
// experiment with the associated Experiment ID.
//
// Output -- Returns StatusNoContent(204) if successful, otherwise returns
// StatusBadRequest(400) with the appropriate message under "message"
func (ec *ExperimentController) DeleteExperiment(c echo.Context) error {
	experiment_id := c.QueryParam("experiment")
	if experiment_id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No Experiment ID")
	}

	err := ec.service.DeleteExperiment(experiment_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

// Route -- /experiments/groups?group=group_id GET
//
// Input -- requires a valid Group ID query parameter.
//
// Output -- Returns a list of all experiments the group is associated with all
// objects under the kep "data". If an error occured, then
// StatusInternalServerError is returned, with the error description. Otherwise,
// returns with a 200 status request.
func (ec *ExperimentController) GetAllExperimentsByGroup(c echo.Context) error {
	group_id := c.QueryParam("group")
	experiments, err := ec.expGrService.GetAllExperimentsByGroup(group_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": experiments,
	})
}

// Route /experiments/groups POST
//
// Input -- requires a valid ExperimentGroupObject. Serves as a joining entry
// between a group and an experiment.
//
// Output -- Returns the created object with ID, and timestamp fields
// instantiated, with a StatusOK(200) response under the key "data".  If the
// provided JSON object was invalid or malformed, it returns a
// StatusBadRequest(400) with the appropriate message under the key "message".
func (ec *ExperimentController) CreateExperimentGroup(c echo.Context) error {
	var body models.ExperimentGroup

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	exp_group, err := ec.expGrService.AddExperimentGroup(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": exp_group,
	})
}

// Route -- /experiments/groups?experiment_group DELETE
//
// Input -- requires a valid experiment_group id. Deletes the association
// between group and project.
//
// Output -- Returns StatusNoContent(204) if successful, otherwise returns
// StatusBadRequest(400) with the appropriate message under "message"
func (ec *ExperimentController) DeleteExperimentGroup(c echo.Context) error {
	experiment_id := c.QueryParam("experiment")
	group_id := c.QueryParam("group")
	err := ec.expGrService.DeleteExperimentGroup(experiment_id, group_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
