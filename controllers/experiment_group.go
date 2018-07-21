package controllers

import (
	"net/http"

	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/services"
	"github.com/labstack/echo"
)

// type IExperimentGroupService interface {
// 	GetExperimentGroupByID(id string) (*models.ExperimentGroup, error)
// 	AddExperimentGroup(exp_group *models.ExperimentGroup) (*models.ExperimentGroup, error)
// 	DeleteExperimentGroup(id string) error
// }

type ExperimentGroupController struct {
	service IExperimentGroupService
	logger  loggers.ILogger
}

func NewExperimentGroupController(egs IExperimentGroupService, l loggers.ILogger) *ExperimentGroupController {
	return &ExperimentGroupController{
		service: egs,
		logger:  l,
	}
}

func (egc *ExperimentGroupController) MountRoutes(g *echo.Group) {
	g.GET("", egc.GetExperimentGroups)
	// g.POST("", egc.CreateExperimentGroup)
	// g.DELETE("", egc.DeleteExperimentGroup)
}

// Route -- /groups?group=
func (egc *ExperimentGroupController) GetExperimentGroups(c echo.Context) error {
	exp_group_id := c.QueryParam("experiment_group")
	exp_group, err := egc.service.GetExperimentGroupByID(exp_group_id)
	if err != nil {
		if err == services.ErrIdInvalid {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else if err == services.ErrExperimentGroupNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": exp_group,
	})
}
