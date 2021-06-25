// // Code generated by MockGen. DO NOT EDIT.
// // Source: initiator.go

// // Package user_mock is a generated GoMock package.
// package user_mock

// import (
// 	context "context"
// 	gomock "github.com/golang/mock/gomock"
// 	model "github.com/iDevoid/stygis/internal/constant/model"
// 	reflect "reflect"
// )

// // MockUsecase is a mock of Usecase interface
// type MockUsecase struct {
// 	ctrl     *gomock.Controller
// 	recorder *MockUsecaseMockRecorder
// }

// // MockUsecaseMockRecorder is the mock recorder for MockUsecase
// type MockUsecaseMockRecorder struct {
// 	mock *MockUsecase
// }

// // NewMockUsecase creates a new mock instance
// func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
// 	mock := &MockUsecase{ctrl: ctrl}
// 	mock.recorder = &MockUsecaseMockRecorder{mock}
// 	return mock
// }

// // EXPECT returns an object that allows the caller to indicate expected use
// func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
// 	return m.recorder
// }

// // Registration mocks base method
// func (m *MockUsecase) Registration(ctx context.Context, user *model.User) error {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "Registration", ctx, user)
// 	ret0, _ := ret[0].(error)
// 	return ret0
// }

// // Registration indicates an expected call of Registration
// func (mr *MockUsecaseMockRecorder) Registration(ctx, user interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Registration", reflect.TypeOf((*MockUsecase)(nil).Registration), ctx, user)
// }