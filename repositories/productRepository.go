package repositories

import (
	"Hacktiv10JWT/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (repo *ProductRepository) CreateProduct(product *models.Product) error {
	return repo.db.Debug().Create(product).Error
}

func (repo *ProductRepository) GetProductsByUserID(userID uint) ([]models.Product, error) {
	var results []models.Product
	err := repo.db.Debug().Preload("User").Find(&results, "user_id = ?", userID).Error
	return results, err
}

func (repo *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var results []models.Product
	err := repo.db.Debug().Preload("User").Find(&results).Error
	return results, err
}

func (repo *ProductRepository) GetProductByProductID(ProductID uint) (*models.Product, error) {
	var result models.Product
	err := repo.db.Debug().First(&result, "id = ?", ProductID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo *ProductRepository) DeleteProductByID(productID int) (int64, error) {
	result := repo.db.Debug().Delete(&models.Product{}, productID)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *ProductRepository) UpdateProduct(product *models.Product) error {
	result := repo.db.Debug().Preload("User").Model(product).Updates(models.Product{Title: product.Title, Description: product.Description})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
