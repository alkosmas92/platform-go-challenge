// Code generated by MockGen. DO NOT EDIT.
// Source: repository/favorite_repository.go

// Package mocksn is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/alkosmas92/platform-go-challenge/internal/app/models"
	gomock "github.com/golang/mock/gomock"
)

// MockFavoriteRepository is a mock of FavoriteRepository interface.
type MockFavoriteRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFavoriteRepositoryMockRecorder
}

// MockFavoriteRepositoryMockRecorder is the mock recorder for MockFavoriteRepository.
type MockFavoriteRepositoryMockRecorder struct {
	mock *MockFavoriteRepository
}

// NewMockFavoriteRepository creates a new mock instance.
func NewMockFavoriteRepository(ctrl *gomock.Controller) *MockFavoriteRepository {
	mock := &MockFavoriteRepository{ctrl: ctrl}
	mock.recorder = &MockFavoriteRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFavoriteRepository) EXPECT() *MockFavoriteRepositoryMockRecorder {
	return m.recorder
}

// CreateFavorite mocks base method.
func (m *MockFavoriteRepository) CreateFavorite(ctx context.Context, favorite *models.Favorite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFavorite", ctx, favorite)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFavorite indicates an expected call of CreateFavorite.
func (mr *MockFavoriteRepositoryMockRecorder) CreateFavorite(ctx, favorite interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFavorite", reflect.TypeOf((*MockFavoriteRepository)(nil).CreateFavorite), ctx, favorite)
}

// DeleteFavorite mocks base method.
func (m *MockFavoriteRepository) DeleteFavorite(ctx context.Context, userID, assetID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFavorite", ctx, userID, assetID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFavorite indicates an expected call of DeleteFavorite.
func (mr *MockFavoriteRepositoryMockRecorder) DeleteFavorite(ctx, userID, assetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFavorite", reflect.TypeOf((*MockFavoriteRepository)(nil).DeleteFavorite), ctx, userID, assetID)
}

// GetFavoritesByUserID mocks base method.
func (m *MockFavoriteRepository) GetFavoritesByUserID(ctx context.Context, userID string, limit, offset int) ([]*models.Favorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoritesByUserID", ctx, userID, limit, offset)
	ret0, _ := ret[0].([]*models.Favorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoritesByUserID indicates an expected call of GetFavoritesByUserID.
func (mr *MockFavoriteRepositoryMockRecorder) GetFavoritesByUserID(ctx, userID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoritesByUserID", reflect.TypeOf((*MockFavoriteRepository)(nil).GetFavoritesByUserID), ctx, userID, limit, offset)
}

// UpdateFavorite mocks base method.
func (m *MockFavoriteRepository) UpdateFavorite(ctx context.Context, userID, assetID string, favorite *models.Favorite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFavorite", ctx, userID, assetID, favorite)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFavorite indicates an expected call of UpdateFavorite.
func (mr *MockFavoriteRepositoryMockRecorder) UpdateFavorite(ctx, userID, assetID, favorite interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFavorite", reflect.TypeOf((*MockFavoriteRepository)(nil).UpdateFavorite), ctx, userID, assetID, favorite)
}
