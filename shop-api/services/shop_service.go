package services

import (
	"errors"
	"shop-api/models"
	"sync"
	"time"
)

type ShopService interface {
	GetByID(id int) (*models.Shop, error)
	GetAll() []models.Shop
	Create(shop models.Shop) (*models.Shop, error)
	UpdateWhatsApp(shopID int, whatsappNumber string) error
}

type ShopServiceImpl struct {
	shops  []models.Shop
	nextID int
	mu     sync.RWMutex
}

func NewShopService() ShopService {
	return &ShopServiceImpl{
		shops: []models.Shop{
			{
				ID:             1,
				Name:           "TechStore Casablanca",
				Active:         true,
				WhatsAppNumber: "212600000001",
				CreatedAt:      time.Now(),
			},
			{
				ID:             2,
				Name:           "ElectroShop Rabat",
				Active:         true,
				WhatsAppNumber: "212600000002",
				CreatedAt:      time.Now(),
			},
		},
		nextID: 3,
	}
}

func (s *ShopServiceImpl) GetByID(id int) (*models.Shop, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, shop := range s.shops {
		if shop.ID == id {
			return &shop, nil
		}
	}
	return nil, errors.New("shop not found")
}

func (s *ShopServiceImpl) GetAll() []models.Shop {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.shops
}

func (s *ShopServiceImpl) Create(shop models.Shop) (*models.Shop, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	shop.ID = s.nextID
	shop.CreatedAt = time.Now()
	s.nextID++
	s.shops = append(s.shops, shop)

	return &shop, nil
}

func (s *ShopServiceImpl) UpdateWhatsApp(shopID int, whatsappNumber string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.shops {
		if s.shops[i].ID == shopID {
			s.shops[i].WhatsAppNumber = whatsappNumber
			return nil
		}
	}
	return errors.New("shop not found")
}
