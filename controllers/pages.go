package controllers

import (
	"net/http"

	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/labstack/echo"
)

type PagesController struct {
	logger loggers.ILogger
}

// Creates controller with logger for extra pages that are not specific to any
// other controller
func NewPagesController(l loggers.ILogger) *PagesController {
	return &PagesController{
		logger: l,
	}
}

// Mounts routes for extra pages
func (p *PagesController) MountRoutes(c *echo.Group) {
	c.GET("/", p.Home)
	c.GET("/health", p.Health)
}

// Home page, only return "App Running"
func (p *PagesController) Home(c echo.Context) error {
	return c.String(http.StatusOK, "App Running")
}

// Route -- /health
//
// Input -- nothing.
//
// Output -- JSON object indicating the service is alive
func (p *PagesController) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"data": "Alive",
	})
}
