package controllers

import (
	"testing"

	"github.com/TerrenceHo/ABFeature/mocks"
)

func newProjectController(ps IProjectService, es IExperimentService) *ProjectController {
	logger := mocks.Logger{}
	return NewProjectController(ps, es, logger)
}

func TestGetAllProjects(t *testing.T) {
	// assert := assert.New(t)

}
