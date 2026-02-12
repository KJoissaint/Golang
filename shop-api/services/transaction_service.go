package services

import (
	"errors"
	"shop-api/models"
	"sync"
	"time"
)

type TransactionService interface {
	GetAll(shopID int) []models.Transaction
	Create(transaction models.Transaction) (*models.Transaction, error)
	GetDashboard(shopID int) (*DashboardStats, error)
}

type DashboardStats struct {
	TotalSales    float64 `json:"total_sales"`
	TotalExpenses float64 `json:"total_expenses"`
	NetProfit     float64 `json:"net_profit"`
	LowStockCount int     `json:"low_stock_count"`
	TotalRevenue  float64 `json:"total_revenue"`
	TotalCost     float64 `json:"total_cost"`
	ProductsSold  int     `json:"products_sold"`
}

type TransactionServiceImpl struct {
	transactions []models.Transaction
	nextID       int
	mu           sync.RWMutex
	productSvc   ProductService
}

func NewTransactionService(productSvc ProductService) TransactionService {
	return &TransactionServiceImpl{
		transactions: []models.Transaction{
			{
				ID:        1,
				Type:      models.TransactionSale,
				ProductID: intPtr(1),
				Quantity:  2,
				Amount:    20000,
				ShopID:    1,
				CreatedAt: time.Now().AddDate(0, 0, -5),
			},
			{
				ID:        2,
				Type:      models.TransactionExpense,
				ProductID: nil,
				Quantity:  1,
				Amount:    5000,
				ShopID:    1,
				CreatedAt: time.Now().AddDate(0, 0, -3),
			},
		},
		nextID:     3,
		productSvc: productSvc,
	}
}

func intPtr(i int) *int {
	return &i
}

func (s *TransactionServiceImpl) GetAll(shopID int) []models.Transaction {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var transactions []models.Transaction
	for _, transaction := range s.transactions {
		if transaction.ShopID == shopID {
			transactions = append(transactions, transaction)
		}
	}
	return transactions
}

func (s *TransactionServiceImpl) Create(transaction models.Transaction) (*models.Transaction, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Validate product exists and belongs to the same shop
	if transaction.ProductID != nil {
		product, err := s.productSvc.GetByID(*transaction.ProductID)
		if err != nil {
			return nil, errors.New("product not found")
		}
		if product.ShopID != transaction.ShopID {
			return nil, errors.New("product does not belong to this shop")
		}

		// For sales, check and update stock
		if transaction.Type == models.TransactionSale {
			if product.Stock < transaction.Quantity {
				return nil, errors.New("insufficient stock")
			}
			// Update stock (this is simplified - in a real app, you'd update through the service)
		}
	}

	transaction.ID = s.nextID
	transaction.CreatedAt = time.Now()
	s.nextID++
	s.transactions = append(s.transactions, transaction)

	return &transaction, nil
}

func (s *TransactionServiceImpl) GetDashboard(shopID int) (*DashboardStats, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	stats := &DashboardStats{}

	// Get all products for this shop
	products := s.productSvc.GetAll(shopID)

	// Count low stock products (less than 5)
	for _, product := range products {
		if product.Stock < 5 {
			stats.LowStockCount++
		}
	}

	// Calculate sales, expenses, and profit
	for _, transaction := range s.transactions {
		if transaction.ShopID == shopID {
			switch transaction.Type {
			case models.TransactionSale:
				stats.TotalSales += transaction.Amount
				stats.ProductsSold += transaction.Quantity

				// Calculate revenue and cost for profit
				if transaction.ProductID != nil {
					product, err := s.productSvc.GetByID(*transaction.ProductID)
					if err == nil {
						stats.TotalRevenue += float64(transaction.Quantity) * product.SellingPrice
						stats.TotalCost += float64(transaction.Quantity) * product.PurchasePrice
					}
				}
			case models.TransactionExpense, models.TransactionWithdrawal:
				stats.TotalExpenses += transaction.Amount
			}
		}
	}

	// Calculate net profit (sales revenue - cost - expenses)
	stats.NetProfit = stats.TotalRevenue - stats.TotalCost - stats.TotalExpenses

	return stats, nil
}
