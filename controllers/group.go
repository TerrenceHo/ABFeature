package controllers

import (
	"net/http"

	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/models/services"
	"github.com/labstack/echo"
)

type IGroupService interface {
	GetAllGroups() ([]*models.Group, error)
	GetGroupByID(groupID string) (*models.Group, error)
	AddGroup(group *models.Group) (*models.Group, error)
	UpdateGroup(group *models.Group) (*models.Group, error)
	DeleteGroup(groupID string) error
}

type GroupController struct {
	service IGroupService
	logger  loggers.ILogger
}

func NewGroupController(gs IGroupService, l loggers.ILogger) *GroupController {
	return &GroupController{
		service: gs,
		logger:  l,
	}
}

// Mounts ExperimentController routes to Echo
func (gc *GroupController) MountRoutes(g *echo.Group) {
	g.GET("", gc.GetGroups)
	g.POST("", gc.CreateGroup)
	g.PUT("", gc.UpdateGroup)
	g.DELETE("", gc.DeleteGroup)
}

// Wrapper function to determine GetGroup or GetAllGroups
func (gc *GroupController) GetGroups(c echo.Context) error {
	if c.QueryParam("group") != "" {
		return gc.GetGroup(c)
	}
	return gc.GetAllGroups(c)
}

// Route -- /groups?group=group_id GET
//
// Input -- Requires query parameter of valid Group ID.
//
// Output -- Returns a JSON object of specified Group.
// If successful it returns a StatusOK(200). If the ID provided was invalid,
// it returns with a StatusBadRequest(400). Otherwise, a
// StatusInternalServerError is returned, indicating a server failure. The
// appropriate error message is provided under the key "message".
func (gc *GroupController) GetGroup(c echo.Context) error {
	group_id := c.QueryParam("group")
	group, err := gc.service.GetGroupByID(group_id)
	if err != nil {
		if err == services.ErrIdInvalid {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else if err == services.ErrGroupNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": group,
	})
}

// Route -- /groups GET
//
// Input -- No input required.
//
// Output -- Returns all groups, as an array of JSON objects under the key
// "data".  If an error occured, then StatusInternalServerError is returned,
// with the error description. Otherwise, returns with a StatusOK(200).
func (gc *GroupController) GetAllGroups(c echo.Context) error {
	groups, err := gc.service.GetAllGroups()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": groups,
	})
}

// Route -- /groups POST
//
// Input -- Requires a valid Group JSON object.  Creates and inserts the Group
// into the database.
//
// Output -- Returns the created object with the ID, and timestamp fields
// instantiated, with a StatusOK(200) response under the key "data". If he
// provided JSON object was invalid or malformed, it returns a
// StatusBadRequest(400) with the appropriate message under the key "message".
func (gc *GroupController) CreateGroup(c echo.Context) error {
	var body models.Group

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	group, err := gc.service.AddGroup(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": group,
	})
}

// Route -- /groups PUT
//
// Input -- Requires a valid Group JSON object, and the ID of the group must
// correspond to an already existing group. Takes the Group and updates its
// fields.
//
// Output -- Returns the updated object with updated fields, including the
// UpdatedAt timestamp field, with a StatusOK(200) response under the key
// "data". If the provided JSON object was invalid or malformed, it returns a
// StatusBadRequest(400) under the key "message".
func (gc *GroupController) UpdateGroup(c echo.Context) error {
	var body models.Group

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	group, err := gc.service.UpdateGroup(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": group,
	})
}

// Route -- /groups DELETE
//
// Input -- requires a valid Group ID query parameters. Deletes the group with
// the associated Group ID.
//
// Output -- Returns StatusNoContent(204) if successful, otherwise returns
// StatusBadRequest(400) with the appropriate message under "message"
func (gc *GroupController) DeleteGroup(c echo.Context) error {
	group_id := c.QueryParam("group")
	if group_id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No Group ID")
	}

	err := gc.service.DeleteGroup(group_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
