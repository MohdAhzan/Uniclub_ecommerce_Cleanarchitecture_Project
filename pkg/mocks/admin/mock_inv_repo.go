// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repository/interface/inventory.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "project/pkg/utils/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInventoryRepository is a mock of InventoryRepository interface.
type MockInventoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryRepositoryMockRecorder
}

// MockInventoryRepositoryMockRecorder is the mock recorder for MockInventoryRepository.
type MockInventoryRepositoryMockRecorder struct {
	mock *MockInventoryRepository
}

// NewMockInventoryRepository creates a new mock instance.
func NewMockInventoryRepository(ctrl *gomock.Controller) *MockInventoryRepository {
	mock := &MockInventoryRepository{ctrl: ctrl}
	mock.recorder = &MockInventoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInventoryRepository) EXPECT() *MockInventoryRepositoryMockRecorder {
	return m.recorder
}

// AddInventory mocks base method.
func (m *MockInventoryRepository) AddInventory(Inventory models.AddInventory, URL string) (models.InventoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddInventory", Inventory, URL)
	ret0, _ := ret[0].(models.InventoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddInventory indicates an expected call of AddInventory.
func (mr *MockInventoryRepositoryMockRecorder) AddInventory(Inventory, URL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInventory", reflect.TypeOf((*MockInventoryRepository)(nil).AddInventory), Inventory, URL)
}

// CheckCategoryID mocks base method.
func (m *MockInventoryRepository) CheckCategoryID(CategoryID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCategoryID", CategoryID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckCategoryID indicates an expected call of CheckCategoryID.
func (mr *MockInventoryRepositoryMockRecorder) CheckCategoryID(CategoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCategoryID", reflect.TypeOf((*MockInventoryRepository)(nil).CheckCategoryID), CategoryID)
}

// CheckProduct mocks base method.
func (m *MockInventoryRepository) CheckProduct(productName, size string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckProduct", productName, size)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckProduct indicates an expected call of CheckProduct.
func (mr *MockInventoryRepositoryMockRecorder) CheckProduct(productName, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckProduct", reflect.TypeOf((*MockInventoryRepository)(nil).CheckProduct), productName, size)
}

// CheckStock mocks base method.
func (m *MockInventoryRepository) CheckStock(pid int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckStock", pid)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckStock indicates an expected call of CheckStock.
func (mr *MockInventoryRepositoryMockRecorder) CheckStock(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckStock", reflect.TypeOf((*MockInventoryRepository)(nil).CheckStock), pid)
}

// DeleteInventory mocks base method.
func (m *MockInventoryRepository) DeleteInventory(pid int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInventory", pid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInventory indicates an expected call of DeleteInventory.
func (mr *MockInventoryRepositoryMockRecorder) DeleteInventory(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInventory", reflect.TypeOf((*MockInventoryRepository)(nil).DeleteInventory), pid)
}

// EditInventory mocks base method.
func (m *MockInventoryRepository) EditInventory(pid int, model models.EditInventory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditInventory", pid, model)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditInventory indicates an expected call of EditInventory.
func (mr *MockInventoryRepositoryMockRecorder) EditInventory(pid, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditInventory", reflect.TypeOf((*MockInventoryRepository)(nil).EditInventory), pid, model)
}

// FindPrice mocks base method.
func (m *MockInventoryRepository) FindPrice(pid int) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPrice", pid)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPrice indicates an expected call of FindPrice.
func (mr *MockInventoryRepositoryMockRecorder) FindPrice(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPrice", reflect.TypeOf((*MockInventoryRepository)(nil).FindPrice), pid)
}

// FindStock mocks base method.
func (m *MockInventoryRepository) FindStock(pid int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindStock", pid)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindStock indicates an expected call of FindStock.
func (mr *MockInventoryRepositoryMockRecorder) FindStock(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindStock", reflect.TypeOf((*MockInventoryRepository)(nil).FindStock), pid)
}

// GetCategoryID mocks base method.
func (m *MockInventoryRepository) GetCategoryID(pid int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryID", pid)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryID indicates an expected call of GetCategoryID.
func (mr *MockInventoryRepositoryMockRecorder) GetCategoryID(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryID", reflect.TypeOf((*MockInventoryRepository)(nil).GetCategoryID), pid)
}

// GetProductImages mocks base method.
func (m *MockInventoryRepository) GetProductImages(pid int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductImages", pid)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductImages indicates an expected call of GetProductImages.
func (mr *MockInventoryRepositoryMockRecorder) GetProductImages(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductImages", reflect.TypeOf((*MockInventoryRepository)(nil).GetProductImages), pid)
}

// ListProducts mocks base method.
func (m *MockInventoryRepository) ListProducts() ([]models.Inventories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts")
	ret0, _ := ret[0].([]models.Inventories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockInventoryRepositoryMockRecorder) ListProducts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockInventoryRepository)(nil).ListProducts))
}

// SearchProducts mocks base method.
func (m *MockInventoryRepository) SearchProducts(pdtName string) ([]models.Inventories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchProducts", pdtName)
	ret0, _ := ret[0].([]models.Inventories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchProducts indicates an expected call of SearchProducts.
func (mr *MockInventoryRepositoryMockRecorder) SearchProducts(pdtName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchProducts", reflect.TypeOf((*MockInventoryRepository)(nil).SearchProducts), pdtName)
}
