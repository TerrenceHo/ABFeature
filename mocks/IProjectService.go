// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/TerrenceHo/ABFeature/models"

// IProjectService is an autogenerated mock type for the IProjectService type
type IProjectService struct {
	mock.Mock
}

// AddProject provides a mock function with given fields: project
func (_m *IProjectService) AddProject(project *models.Project) (*models.Project, error) {
	ret := _m.Called(project)

	var r0 *models.Project
	if rf, ok := ret.Get(0).(func(*models.Project) *models.Project); ok {
		r0 = rf(project)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Project) error); ok {
		r1 = rf(project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProject provides a mock function with given fields: projectID
func (_m *IProjectService) DeleteProject(projectID string) error {
	ret := _m.Called(projectID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(projectID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllProjects provides a mock function with given fields:
func (_m *IProjectService) GetAllProjects() ([]*models.Project, error) {
	ret := _m.Called()

	var r0 []*models.Project
	if rf, ok := ret.Get(0).(func() []*models.Project); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectByID provides a mock function with given fields: projectID
func (_m *IProjectService) GetProjectByID(projectID string) (*models.Project, error) {
	ret := _m.Called(projectID)

	var r0 *models.Project
	if rf, ok := ret.Get(0).(func(string) *models.Project); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
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

// UpdateProject provides a mock function with given fields: project
func (_m *IProjectService) UpdateProject(project *models.Project) (*models.Project, error) {
	ret := _m.Called(project)

	var r0 *models.Project
	if rf, ok := ret.Get(0).(func(*models.Project) *models.Project); ok {
		r0 = rf(project)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Project) error); ok {
		r1 = rf(project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
