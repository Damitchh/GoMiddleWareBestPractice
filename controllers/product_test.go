package controllers

import (
	"Hacktiv10JWT/models"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define a mock product service that implements the ProductService interface
type mockProductService struct {
	mock.Mock
}

func (m *mockProductService) GetProductByProductID(id string) (*models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *mockProductService) GetProductsByUserID(userID uint) ([]*models.Product, error) {
	args := m.Called(userID)
	return args.Get(0).([]*models.Product), args.Error(1)
}

func (m *mockProductService) GetAllProducts() ([]*models.Product, error) {
	args := m.Called()
	return args.Get(0).([]*models.Product), args.Error(1)
}

func (m *mockProductService) CreateProduct(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *mockProductService) UpdateProduct(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *mockProductService) DeleteProductByID(id string) (int64, error) {
	args := m.Called(id)
	return args.Get(0).(int64), args.Error(1)
}

func TestGetProductByID(t *testing.T) {
	// Initialize mock product service
	mockProduct := &models.Product{
		Title:       "Product 1",
		Description: "Description 1",
		UserID:      1,
	}
	mockProductService := new(mockProductService)
	mockProductService.On("GetProductByProductID", "1").Return(mockProduct, nil)

	// Initialize Gin context
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/products/:ID", GetProductbyID)
	c.Request, _ = http.NewRequest(http.MethodGet, "/products/1", nil)

	// Call the endpoint handler
	r.ServeHTTP(w, c.Request)

	// Check response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, gin.H{
		"ID":          mockProduct.ID,
		"Title":       mockProduct.Title,
		"Description": mockProduct.Description,
		"CreatedAt":   mock.Anything,
		"UpdatedAt":   mock.Anything,
	}, gin.H{})

	// Verify mock
	mockProductService.AssertExpectations(t)
}

func TestGetProductByIDNotFound(t *testing.T) {
	// Initialize mock product service
	mockProductService := new(mockProductService)
	mockProductService.On("GetProductByProductID", "1").Return(nil, errors.New("not found"))

	// Initialize Gin context
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/products/:ID", GetProductbyID)
	c.Request, _ = http.NewRequest(http.MethodGet, "/products/1", nil)

	// Call the endpoint handler
	r.ServeHTTP(w, c.Request)

	// Check response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, gin.H{
		"error message": "Data Not Found",
	}, gin.H{})

	// Verify mock
	mockProductService.AssertExpectations(t)
}

func TestGetAllProducts(t *testing.T) {
	// Initialize mock product service
	mockProduct1 := &models.Product{
		Title:       "Product 1",
		Description: "Description 1",
		UserID:      1,
	}
	mockProduct2 := &models.Product{
		Title:       "Product 2",
		Description: "Description 2",
		UserID:      1,
	}
	mockProductService := new(mockProductService)
	mockProductService.On("GetAllProducts").Return([]*models.Product{mockProduct1, mockProduct2}, nil)

	// Initialize Gin context
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/products", GetAllProducts)
	c.Request, _ = http.NewRequest(http.MethodGet, "/products", nil)

	// Call the endpoint handler
	r.ServeHTTP(w, c.Request)

	// Check response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, gin.H{
		"Product data": []interface{}{
			gin.H{
				"ID":          mockProduct1.ID,
				"Title":       mockProduct1.Title,
				"Description": mockProduct1.Description,
				"CreatedAt":   mock.Anything,
				"UpdatedAt":   mock.Anything,
			},
			gin.H{
				"ID":          mockProduct2.ID,
				"Title":       mockProduct2.Title,
				"Description": mockProduct2.Description,
				"CreatedAt":   mock.Anything,
				"UpdatedAt":   mock.Anything,
			},
		},
	}, gin.H{})

	// Verify mock
	mockProductService.AssertExpectations(t)
}

func TestGetAllProductsWithNoData(t *testing.T) {
	// Initialize mock product service with no data
	mockProductService := new(mockProductService)
	mockProductService.On("GetAllProducts").Return([]*models.Product{}, nil)

	// Initialize Gin context
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/products", GetAllProducts)
	c.Request, _ = http.NewRequest(http.MethodGet, "/products", nil)

	// Call the endpoint handler
	r.ServeHTTP(w, c.Request)

	// Check response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, gin.H{
		"Product data": []interface{}{},
	}, gin.H{})

	// Verify mock
	mockProductService.AssertExpectations(t)
}
