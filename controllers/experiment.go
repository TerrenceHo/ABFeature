package controllers

import (
	"net/http"

	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/models/services"
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

type ExperimentController struct {
	service IExperimentService
	logger  loggers.ILogger
}

// Creates and returns an ExperimentController struct with a ExperimentService
// and logger.
func NewExperimentController(es IExperimentService, l loggers.ILogger) *ExperimentController {
	return &ExperimentController{
		service: es,
		logger:  l,
	}
}

// Mounts ExperimentController routes to Echo
func (ec *ExperimentController) MountRoutes(g *echo.Group) {
	g.GET("", ec.GetExperiments)
	g.POST("", ec.CreateExperiment)
	g.PUT("", ec.UpdateExperiment)
	g.DELETE("", ec.DeleteExperiment)
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
