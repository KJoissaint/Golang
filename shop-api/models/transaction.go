package models

import "time"

type TransactionType string

const (
	TransactionSale       TransactionType = "Sale"
	TransactionExpense    TransactionType = "Expense"
	TransactionWithdrawal TransactionType = "Withdrawal"
)

type Transaction struct {
	ID        int             `json:"id"`
	Type      TransactionType `json:"type"`
	ProductID *int            `json:"product_id,omitempty"` // Optional for expenses/withdrawals
	Quantity  int             `json:"quantity"`
	Amount    float64         `json:"amount"`
	ShopID    int             `json:"shop_id"`
	CreatedAt time.Time       `json:"created_at"`
}
