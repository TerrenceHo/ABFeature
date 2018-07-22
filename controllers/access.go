package controllers

import (
	"net/http"

	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/services"
	"github.com/labstack/echo"
)

// Access controller handles the access of experiments, determining whether
// certain users or groups have access to the experiment.
type AccessController struct {
	ps     IProjectService
	es     IExperimentService
	gs     IGroupService
	egs    IExperimentGroupService
	logger loggers.ILogger
}

// Creates and returns a new Access Controller
func NewAccessController(ps IProjectService,
	es IExperimentService,
	gs IGroupService,
	egs IExperimentGroupService,
	l loggers.ILogger) *AccessController {
	return &AccessController{
		ps:     ps,
		es:     es,
		gs:     gs,
		egs:    egs,
		logger: l,
	}
}

func (ac *AccessController) MountRoutes(g *echo.Group) {
	g.GET("/groups", ac.ExperimentGroupAccess)
	// g.GET("/users", ac.ExperimentUserAccess)
}

// Route -- /experiments/groups?experiment=experiment_id&group=group_id
//
// Input -- requires both a GroupID and ExperimentID query parameter.
//
// Output -- Returns a boolean variable whether or not the group has access to
// the experiment.  The group does not have access to the experiment if the
// experiment is not enabled, or if the group is not associated with the
// experiment. Otherwise, it returns true. If the experiment could not be found,
// then the StatusNotFound is returned. If an error occured, then
// StatusInternalServerError is returned.
func (ac *AccessController) ExperimentGroupAccess(c echo.Context) error {
	experiment_id := c.QueryParam("experiment")
	group_id := c.QueryParam("group")

	experiment, err := ac.es.GetExperimentByID(experiment_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if !experiment.Enabled {
		return c.JSON(http.StatusOK, echo.Map{
			"data": false,
		})
	}

	_, err = ac.egs.GetByExperimentAndGroup(experiment_id, group_id)
	if err != nil {
		if err == services.ErrExperimentGroupNotFound {
			return c.JSON(http.StatusOK, echo.Map{
				"data": false,
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": true,
	})
}
