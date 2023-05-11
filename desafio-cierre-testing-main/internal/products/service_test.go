package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_GetAllBySeller(t *testing.T) {
	t.Run("successful case", func(t *testing.T) {
		//arrange
		service := NewService(NewRepositoryMock())
		searchId := "FEX112AC"
		expectedResult := []Product{{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "Mock product",
			Price:       123.55,
		}}

		//act
		result, err := service.GetAllBySeller(searchId)

		//assert
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("error case", func(t *testing.T) {
		//arrange
		service := NewService(NewRepositoryMock())
		searchId := "ASD"
		expectedResult := "seller not found"

		//act
		result, err := service.GetAllBySeller(searchId)

		//assert
		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedResult, err.Error())
	})

}
