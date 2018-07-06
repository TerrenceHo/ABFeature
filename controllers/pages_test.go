package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TerrenceHo/ABFeature/mocks"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	assert := assert.New(t)

	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(echo.GET, "/", nil)
	assert.Nil(err, "Making request error should be nil.")

	e := echo.New()
	c := e.NewContext(req, recorder)

	pagesController := NewPagesController(mocks.Logger{})
	if assert.NoError(pagesController.Home(c)) {
		assert.Equal(http.StatusOK, recorder.Code)
		assert.Equal("App Running", recorder.Body.String())
	}
}

func TestHealth(t *testing.T) {
	assert := assert.New(t)

	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(echo.GET, "/", nil)
	assert.Nil(err, "Making request error should be nil.")

	e := echo.New()
	c := e.NewContext(req, recorder)

	pagesController := NewPagesController(mocks.Logger{})
	if assert.NoError(pagesController.Health(c)) {
		assert.Equal(http.StatusOK, recorder.Code)
		assert.Equal(`{"data":"Alive"}`, recorder.Body.String())
	}
}
