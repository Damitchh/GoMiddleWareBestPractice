package services

import (
	"Hacktiv10JWT/models"
	"Hacktiv10JWT/repositories"
	"strconv"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService(productRepo *repositories.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.productRepo.CreateProduct(product)
}

func (s *ProductService) GetProductsByUserID(userID uint) ([]models.Product, error) {
	return s.productRepo.GetProductsByUserID(userID)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.productRepo.GetAllProducts()
}

func (s *ProductService) GetProductByProductID(productID string) (*models.Product, error) {
	id, err := strconv.Atoi(productID)
	if err != nil {
		return nil, err
	}
	return s.productRepo.GetProductByProductID(uint(id))
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.productRepo.UpdateProduct(product)
}

func (s *ProductService) DeleteProductByID(productID string) (int64, error) {
	id, err := strconv.Atoi(productID)
	if err != nil {
		return 0, err
	}
	return s.productRepo.DeleteProductByID(id)
}
