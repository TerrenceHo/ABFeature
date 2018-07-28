// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/TerrenceHo/ABFeature/models"

// IExperimentService is an autogenerated mock type for the IExperimentService type
type IExperimentService struct {
	mock.Mock
}

// AddExperiment provides a mock function with given fields: experiment
func (_m *IExperimentService) AddExperiment(experiment *models.Experiment) (*models.Experiment, error) {
	ret := _m.Called(experiment)

	var r0 *models.Experiment
	if rf, ok := ret.Get(0).(func(*models.Experiment) *models.Experiment); ok {
		r0 = rf(experiment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Experiment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Experiment) error); ok {
		r1 = rf(experiment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteExperiment provides a mock function with given fields: experimentID
func (_m *IExperimentService) DeleteExperiment(experimentID string) error {
	ret := _m.Called(experimentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(experimentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllExperimentsByProject provides a mock function with given fields: projectID
func (_m *IExperimentService) GetAllExperimentsByProject(projectID string) ([]*models.Experiment, error) {
	ret := _m.Called(projectID)

	var r0 []*models.Experiment
	if rf, ok := ret.Get(0).(func(string) []*models.Experiment); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Experiment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExperimentByID provides a mock function with given fields: experimentID
func (_m *IExperimentService) GetExperimentByID(experimentID string) (*models.Experiment, error) {
	ret := _m.Called(experimentID)

	var r0 *models.Experiment
	if rf, ok := ret.Get(0).(func(string) *models.Experiment); ok {
		r0 = rf(experimentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Experiment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(experimentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateExperiment provides a mock function with given fields: experiment
func (_m *IExperimentService) UpdateExperiment(experiment *models.Experiment) (*models.Experiment, error) {
	ret := _m.Called(experiment)

	var r0 *models.Experiment
	if rf, ok := ret.Get(0).(func(*models.Experiment) *models.Experiment); ok {
		r0 = rf(experiment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Experiment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Experiment) error); ok {
		r1 = rf(experiment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
