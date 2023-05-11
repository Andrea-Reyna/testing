package products

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServerForTestProductHandler() *gin.Engine {
	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.TestMode)
	server := gin.New()
	server.GET("/api/v1/products", handler.GetProducts)

	return server
}

func TestHandler_GetAllBySeller(t *testing.T) {

	t.Run("successful case", func(t *testing.T) {
		//arrange
		server := createServerForTestProductHandler()
		searchId := "FEX112AC"
		expectedStatusCode := http.StatusOK
		expectedHeaders := http.Header{
			"Content-Type": []string{
				"application/json; charset=utf-8",
			}}
		expectedResponse := `[
				{
				  "ID": "mock",
				  "SellerID": "FEX112AC",
				  "Description": "generic product",
				  "Price": 123.55
				}
			  ]`
		request := httptest.NewRequest(http.MethodGet, "/api/v1/products?seller_id="+searchId, nil)
		response := httptest.NewRecorder()

		//act
		server.ServeHTTP(response, request)

		//assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
		assert.True(t, len(response.Body.String()) > 0)
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})

	t.Run("bad request", func(t *testing.T) {
		//arrange
		server := createServerForTestProductHandler()
		searchId := ""
		expectedStatusCode := http.StatusBadRequest
		expectedHeaders := http.Header{
			"Content-Type": []string{
				"application/json; charset=utf-8",
			}}
		expectedResponse := `{"error":"seller_id query param is required"}`

		request := httptest.NewRequest(http.MethodGet, "/api/v1/products?seller_id="+searchId, nil)
		response := httptest.NewRecorder()

		//act
		server.ServeHTTP(response, request)

		//assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
		assert.True(t, len(response.Body.String()) > 0)
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})

}
