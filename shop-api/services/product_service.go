package services

import (
	"errors"
	"shop-api/models"
	"sync"
	"time"
)

type ProductService interface {
	GetAll(shopID int) []models.Product
	GetByID(id int) (*models.Product, error)
	GetPublicProducts(shopID int) []models.Product
	Create(product models.Product) (*models.Product, error)
	Update(id int, product models.Product) (*models.Product, error)
	Delete(id int) error
}

type ProductServiceImpl struct {
	products []models.Product
	nextID   int
	mu       sync.RWMutex
}

func NewProductService() ProductService {
	return &ProductServiceImpl{
		products: []models.Product{
			{
				ID:            1,
				Name:          "iPhone 14 Pro",
				Description:   "Latest iPhone with advanced camera system",
				Category:      "Smartphones",
				PurchasePrice: 8000,
				SellingPrice:  10000,
				Stock:         15,
				ImageURL:      "https://example.com/iphone14.jpg",
				ShopID:        1,
				CreatedAt:     time.Now(),
			},
			{
				ID:            2,
				Name:          "MacBook Pro M2",
				Description:   "Powerful laptop for professionals",
				Category:      "Laptops",
				PurchasePrice: 15000,
				SellingPrice:  18000,
				Stock:         8,
				ImageURL:      "https://example.com/macbook.jpg",
				ShopID:        1,
				CreatedAt:     time.Now(),
			},
			{
				ID:            3,
				Name:          "Samsung Galaxy S23",
				Description:   "Premium Android smartphone",
				Category:      "Smartphones",
				PurchasePrice: 6000,
				SellingPrice:  7500,
				Stock:         20,
				ImageURL:      "https://example.com/samsung.jpg",
				ShopID:        2,
				CreatedAt:     time.Now(),
			},
			{
				ID:            4,
				Name:          "AirPods Pro",
				Description:   "Wireless earbuds with noise cancellation",
				Category:      "Accessories",
				PurchasePrice: 1500,
				SellingPrice:  2000,
				Stock:         3,
				ImageURL:      "https://example.com/airpods.jpg",
				ShopID:        1,
				CreatedAt:     time.Now(),
			},
		},
		nextID: 5,
	}
}

func (s *ProductServiceImpl) GetAll(shopID int) []models.Product {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var products []models.Product
	for _, product := range s.products {
		if product.ShopID == shopID {
			products = append(products, product)
		}
	}
	return products
}

func (s *ProductServiceImpl) GetByID(id int) (*models.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, product := range s.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}

func (s *ProductServiceImpl) GetPublicProducts(shopID int) []models.Product {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var products []models.Product
	for _, product := range s.products {
		if product.ShopID == shopID {
			products = append(products, product)
		}
	}
	return products
}

func (s *ProductServiceImpl) Create(product models.Product) (*models.Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	product.ID = s.nextID
	product.CreatedAt = time.Now()
	s.nextID++
	s.products = append(s.products, product)

	return &product, nil
}

func (s *ProductServiceImpl) Update(id int, updated models.Product) (*models.Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.products {
		if s.products[i].ID == id {
			// Keep the original ID, ShopID, and CreatedAt
			updated.ID = s.products[i].ID
			updated.ShopID = s.products[i].ShopID
			updated.CreatedAt = s.products[i].CreatedAt
			s.products[i] = updated
			return &s.products[i], nil
		}
	}
	return nil, errors.New("product not found")
}

func (s *ProductServiceImpl) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, product := range s.products {
		if product.ID == id {
			s.products = append(s.products[:i], s.products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}
