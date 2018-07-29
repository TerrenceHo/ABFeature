package controllers

import (
	"net/http"

	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/models"
	"github.com/TerrenceHo/ABFeature/services"
	"github.com/labstack/echo"
)

type IUserService interface {
	GetAllUsers() ([]*models.User, error)
	GetUserByID(userID string) (*models.User, error)
	AddUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(userID string) error
}

type UserController struct {
	service IUserService
	logger  loggers.ILogger
}

func NewUserController(us IUserService, l loggers.ILogger) *UserController {
	return &UserController{
		service: us,
		logger:  l,
	}
}

func (uc *UserController) MountRoutes(g *echo.Group) {
	g.GET("", uc.GetUsers)
	g.POST("", uc.CreateUser)
	g.PUT("", uc.UpdateUser)
	g.DELETE("", uc.DeleteUser)
}

// Wrapper function to determine GetGroup or GetAllGroups
func (uc *UserController) GetUsers(c echo.Context) error {
	if c.QueryParam("user") != "" {
		return uc.GetUser(c)
	}
	return uc.GetAllUsers(c)
}

// Route -- /users?user=user_id GET
//
// Input -- requires no input, only a query parameter of a valid User ID.
//
// Output == Return JSON objectof User withj specified User ID. If successful,
// it returns a 200 status response. If the ID provided was invalid,
// it returns with a StatusBadRequest(400).  If the project was not found, then
// StatusNotFound(404) is returned.  Otherwise, a StatusInternalServerError is
// returned, indicating a server failure. The appropriate error message is
// provided under the key "message".
func (uc *UserController) GetUser(c echo.Context) error {
	user_id := c.QueryParam("user")
	user, err := uc.service.GetUserByID(user_id)

	if err != nil {
		if err == services.ErrIdInvalid {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else if err == services.ErrUserNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": user,
	})
}

// Route -- /users GET
//
// Input -- Requires no input parameters.
//
// Output -- Returns all useres, as an array of JSON objects, under the key
// "data". If an error occured, then StatusInternalServerError is returned, with
// the error description. Otherwise, returns with a 200 satus request.
func (uc *UserController) GetAllUsers(c echo.Context) error {
	users, err := uc.service.GetAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": users,
	})
}

// Route -- /users POST
//
// Input -- Requires valid User JSON Object, creates and inserts the User into
// the database.
//
// Output -- Returns the created object with ID, and timestamp fields
// instantiated, with a StatusOK(200) response under the key "data".  If the
// provided JSON object was invalid or malformed, it returns a
// StatusBadRequest(400) with the appropriate message under the key "message".
func (uc *UserController) CreateUser(c echo.Context) error {
	var body models.User

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := uc.service.AddUser(&body)
	if err != nil {
		if err == services.ErrUserValidation {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": user,
	})
}

// Route /users PUT
//
// Input -- Requires a valid User JSON object, and the ID of the user must
// correspond to an already existing user.  Takes the User and updates its
// fields.
//
// Output -- Returns the updated object with updated fields, including the
// UpdatedAt timestamp field, with a StatusOK(200) response under the key
// "data". If the provided JSON object was invalid or malformed, it returns a
// StatusBadRequest(400) under the key "message".
func (uc *UserController) UpdateUser(c echo.Context) error {
	var body models.User
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := uc.service.UpdateUser(&body)
	if err != nil {
		if err == services.ErrUserValidation {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": user,
	})
}

// Route /users?user=user_id DELETE
//
// Input -- requires a valid UserID query parameter. Deletes the user associated
// with the user ID.
//
// Output -- Returns StatusNoContent(204) if successful, otherwise returns
// StatusBadRequest(400) with the appropriate message under "message"
func (uc *UserController) DeleteUser(c echo.Context) error {
	user_id := c.QueryParam("user")
	err := uc.service.DeleteUser(user_id)
	if err != nil {
		if err == services.ErrIdInvalid {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else if err == services.ErrUserNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
