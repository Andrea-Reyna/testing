package storage

import "github.com/stretchr/testify/mock"

type StorageMock struct {
	GetValueFunc func(key string) interface{}
}

func (s *StorageMock) GetValue(key string) interface{} {
	if condition := s.GetValueFunc != nil; condition {
		return s.GetValueFunc(key)
	}
	return nil
}


//MockStorageTestify - Example of testify mock
type MockStorageTestify struct {
	Mock *mock.Mock
}

func NewMockStorageTestify() *MockStorageTestify {
	return &MockStorageTestify{Mock: &mock.Mock{}}
}

func (mock *MockStorageTestify) GetValue(key string) interface{} {
	args := mock.Mock.Called(key)
	return args.Get(0)
}
