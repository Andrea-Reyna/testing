package products

import "errors"

type repositoryMock struct{}

func NewRepositoryMock() Repository {
	return &repositoryMock{}
}

func (r *repositoryMock) GetAllBySeller(sellerID string) ([]Product, error) {
	var prodList []Product
	prodList = append(prodList, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "Mock product",
		Price:       123.55,
	})
	if sellerID != "FEX112AC" { 
		return nil, errors.New("seller not found") //this error not exist in original repository
	}
	return prodList, nil
}
