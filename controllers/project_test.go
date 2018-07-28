package controllers

import (
	"testing"

	"github.com/TerrenceHo/ABFeature/mocks"
)

func newProjectController(ps IProjectService) *ProjectController {
	logger := mocks.Logger{}
	return NewProjectController(ps, logger)
}

func TestGetAllProjects(t *testing.T) {
	// assert := assert.New(t)
	// service := new(mocks.IProjectService)
	// controller := newProjectController(service)
	// recorder := httptest.NewRecorder()
	// req := httptest.NewRequest(echo.GET, "/", nil)
	// e := echo.New()
	// c := e.NewContext(req, recorder)

	// bad_err := errors.New("Database failed.")
	// service.On("GetAllProjects").Return(nil, bad_err)
	// if err := controller.GetAllProjects(c); err != nil {
	// 	fmt.Println(recorder)
	// 	fmt.Println(recorder.Code)
	// 	fmt.Println(err.Error())
	// 	// assert.Equal(http.StatusInternalServerError, recorder.Code)
	// 	// assert.Equal(
	// }
}
