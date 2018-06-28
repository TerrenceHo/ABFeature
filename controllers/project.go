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
	GetProjectByID(id string) (*models.Project, error)
	AddProject(project *models.Project) (*models.Project, error)
	UpdateProject(project *models.Project) error
	DeleteProject(id string) error
}

type ProjectController struct {
	service IProjectService
	logger  loggers.ILogger
}

func NewProjectController(ps IProjectService, l loggers.ILogger) *ProjectController {
	return &ProjectController{
		service: ps,
		logger:  l,
	}
}

// Mounts Routes to the Echo Router
func (pc *ProjectController) MountRoutes(g *echo.Group) {
	g.GET("", pc.getProjects)
	g.POST("", pc.createProject)
	g.PUT("", pc.updateProject)
	g.DELETE("", pc.deleteProject)
}

// Wrapper depending on existance of project query param, holding "ID"
func (pc *ProjectController) getProjects(c echo.Context) error {
	project_id := c.QueryParam("project")
	if project_id != "" {
		return pc.getProject(c)
	}
	return pc.getAllProjects(c)
}

func (pc *ProjectController) getAllProjects(c echo.Context) error {
	projects, err := pc.service.GetAllProjects()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": projects,
	})
}

func (pc *ProjectController) getProject(c echo.Context) error {
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

func (pc *ProjectController) createProject(c echo.Context) error {
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

func (pc *ProjectController) updateProject(c echo.Context) error {
	var body models.Project
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := pc.service.UpdateProject(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (pc *ProjectController) deleteProject(c echo.Context) error {
	project_id := c.QueryParam("project")
	err := pc.service.DeleteProject(project_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

// TODO Add get all experiments, utilizing experiments service
func (pc *ProjectController) getAllExperiments(c echo.Context) error {
	// get all experiments
	return nil
}
