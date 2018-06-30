package controllers

import (
	"net/http"

	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/labstack/echo"
)

type PagesController struct {
	logger loggers.ILogger
}

func NewPagesController(l loggers.ILogger) *PagesController {
	return &PagesController{
		logger: l,
	}
}

func (p *PagesController) MountRoutes(c *echo.Group) {
	c.GET("/", p.home)
	c.GET("/health", p.health)
}

func (p *PagesController) home(c echo.Context) error {
	return c.String(http.StatusOK, "App Running")
}

func (p *PagesController) health(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"status": "Alive",
	})
}
