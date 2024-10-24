// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repository/interface/category.go

// Package mocks is a generated GoMock package.
package mocks

import (
	domain "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCategoryRepository is a mock of CategoryRepository interface.
type MockCategoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryRepositoryMockRecorder
}

// MockCategoryRepositoryMockRecorder is the mock recorder for MockCategoryRepository.
type MockCategoryRepositoryMockRecorder struct {
	mock *MockCategoryRepository
}

// NewMockCategoryRepository creates a new mock instance.
func NewMockCategoryRepository(ctrl *gomock.Controller) *MockCategoryRepository {
	mock := &MockCategoryRepository{ctrl: ctrl}
	mock.recorder = &MockCategoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryRepository) EXPECT() *MockCategoryRepositoryMockRecorder {
	return m.recorder
}

// AddCategory mocks base method.
func (m *MockCategoryRepository) AddCategory(category string) (domain.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCategory", category)
	ret0, _ := ret[0].(domain.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCategory indicates an expected call of AddCategory.
func (mr *MockCategoryRepositoryMockRecorder) AddCategory(category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCategory", reflect.TypeOf((*MockCategoryRepository)(nil).AddCategory), category)
}

// CheckCategory mocks base method.
func (m *MockCategoryRepository) CheckCategory(current string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCategory", current)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckCategory indicates an expected call of CheckCategory.
func (mr *MockCategoryRepositoryMockRecorder) CheckCategory(current interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCategory", reflect.TypeOf((*MockCategoryRepository)(nil).CheckCategory), current)
}

// CheckCategoryByID mocks base method.
func (m *MockCategoryRepository) CheckCategoryByID(categoryID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCategoryByID", categoryID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckCategoryByID indicates an expected call of CheckCategoryByID.
func (mr *MockCategoryRepositoryMockRecorder) CheckCategoryByID(categoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCategoryByID", reflect.TypeOf((*MockCategoryRepository)(nil).CheckCategoryByID), categoryID)
}

// DeleteCategory mocks base method.
func (m *MockCategoryRepository) DeleteCategory(CategoryID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", CategoryID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategory indicates an expected call of DeleteCategory.
func (mr *MockCategoryRepositoryMockRecorder) DeleteCategory(CategoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockCategoryRepository)(nil).DeleteCategory), CategoryID)
}

// GetCategories mocks base method.
func (m *MockCategoryRepository) GetCategories() ([]domain.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories")
	ret0, _ := ret[0].([]domain.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockCategoryRepositoryMockRecorder) GetCategories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockCategoryRepository)(nil).GetCategories))
}

// UpdateCategory mocks base method.
func (m *MockCategoryRepository) UpdateCategory(current, new string) (domain.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategory", current, new)
	ret0, _ := ret[0].(domain.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCategory indicates an expected call of UpdateCategory.
func (mr *MockCategoryRepositoryMockRecorder) UpdateCategory(current, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockCategoryRepository)(nil).UpdateCategory), current, new)
}
