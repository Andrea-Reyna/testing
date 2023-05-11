package products

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServerForUnitaryTestProductHandler(service *MockService) *gin.Engine {
	handler := NewHandler(service)

	gin.SetMode(gin.TestMode)
	server := gin.New()
	server.GET("/api/v1/products", handler.GetProducts)

	return server
}

func TestUnitaryHandler_GetAllBySeller(t *testing.T) {

	t.Run("successful case", func(t *testing.T) {
		// Arrange
		mockService := NewMockService()

		expectedProducts := []Product{{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		}}
		mockService.On("GetAllBySeller", "FEX112AC").Return(expectedProducts, nil)

		server := createServerForUnitaryTestProductHandler(mockService)
		searchId := "FEX112AC"
		expectedStatusCode := http.StatusOK
		expectedHeaders := http.Header{
			"Content-Type": []string{
				"application/json; charset=utf-8",
			}}
		expectedProductsJSON, err := json.Marshal(expectedProducts)
		if err != nil {
			t.Fatalf("Failed to marshal expectedProducts: %v", err)
		}
		expectedResponse := string(expectedProductsJSON)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/products?seller_id="+searchId, nil)
		response := httptest.NewRecorder()

		// Act
		server.ServeHTTP(response, request)

		// Assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
		assert.True(t, len(response.Body.String()) > 0)
		assert.JSONEq(t, expectedResponse, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		// Arrange
		mockService := NewMockService()

		server := createServerForUnitaryTestProductHandler(mockService)
		searchId := ""
		expectedStatusCode := http.StatusBadRequest
		expectedHeaders := http.Header{
			"Content-Type": []string{
				"application/json; charset=utf-8",
			}}
		expectedResponse := `{"error":"seller_id query param is required"}`

		request := httptest.NewRequest(http.MethodGet, "/api/v1/products?seller_id="+searchId, nil)
		response := httptest.NewRecorder()

		// Act
		server.ServeHTTP(response, request)

		// Assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
		assert.True(t, len(response.Body.String()) > 0)
		assert.JSONEq(t, expectedResponse, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		// Arrange
		mockService := NewMockService()
		searchId := "FEX112AC"
		expectedStatusCode := http.StatusInternalServerError
		expectedResponse := `{"error":"something went wrong"}`

		mockService.On("GetAllBySeller", searchId).Return([]Product{}, errors.New("something went wrong"))

		server := createServerForUnitaryTestProductHandler(mockService)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/products?seller_id="+searchId, nil)
		response := httptest.NewRecorder()

		// Act
		server.ServeHTTP(response, request)

		// Assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.True(t, len(response.Body.String()) > 0)
		assert.JSONEq(t, expectedResponse, response.Body.String())

		mockService.AssertExpectations(t)
	})
}
