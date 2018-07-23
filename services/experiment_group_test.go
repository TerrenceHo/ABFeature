package services

import (
	"testing"

	"github.com/TerrenceHo/ABFeature/mocks"
)

func newExperimentGroupService(store *mocks.IExperimentGroupStore) *ExperimentGroupService {
	logger := mocks.Logger{}
	return NewExperimentGroupService(store, logger)
}

func TestGetAllGroupsByExperiment(t *testing.T) {
	// assert := assert.New(t)
	// store := new(mocks.IExperimentGroupStore)
	// service := newExperimentGroupService(store)
}
