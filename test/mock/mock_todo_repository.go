package mock

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/jobpay/todo/internal/domain/entity/todo"
	"github.com/jobpay/todo/internal/domain/entity/todo/valueobject"
)

type MockTodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryMockRecorder
}

type MockTodoRepositoryMockRecorder struct {
	mock *MockTodoRepository
}

func NewMockTodoRepository(ctrl *gomock.Controller) *MockTodoRepository {
	mock := &MockTodoRepository{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryMockRecorder{mock}
	return mock
}

func (m *MockTodoRepository) EXPECT() *MockTodoRepositoryMockRecorder {
	return m.recorder
}

func (m *MockTodoRepository) Save(t *todo.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", t)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockTodoRepositoryMockRecorder) Save(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockTodoRepository)(nil).Save), t)
}

func (m *MockTodoRepository) FindByID(id valueobject.ID) (*todo.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*todo.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTodoRepositoryMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockTodoRepository)(nil).FindByID), id)
}

func (m *MockTodoRepository) FindAll() ([]*todo.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*todo.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTodoRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockTodoRepository)(nil).FindAll))
}

func (m *MockTodoRepository) Update(t *todo.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", t)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockTodoRepositoryMockRecorder) Update(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodoRepository)(nil).Update), t)
}

func (m *MockTodoRepository) Delete(id valueobject.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockTodoRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodoRepository)(nil).Delete), id)
}
