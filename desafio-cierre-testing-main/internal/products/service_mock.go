package products

import "github.com/stretchr/testify/mock"

type MockService struct {
	mock.Mock
}

func NewMockService() *MockService {
	return new(MockService)
}

func (m *MockService) GetAllBySeller(sellerID string) ([]Product, error) {
	args := m.Called(sellerID)
	return args.Get(0).([]Product), args.Error(1)
}
