package controllers

import (
	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/models"
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

}

// Wrapper function to determine GetGroup or GetAllGroups
// func (uc *UserController) GetUsers(c echo.Context) error {
// 	if c.QueryParam("group") != "" {
// 		return uc.GetUser(c)
// 	}
// 	return uc.GetAllGroups(c)
// }

// func (uc *UserController) GetUser(c echo.Context) error {}

// func (uc *UserController) GetAllUsers(c echo.Context) error {}
