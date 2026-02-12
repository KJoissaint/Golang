package models

import "time"

type Product struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Category      string    `json:"category"`
	PurchasePrice float64   `json:"purchase_price,omitempty"` // Only for SuperAdmin
	SellingPrice  float64   `json:"selling_price"`
	Stock         int       `json:"stock"`
	ImageURL      string    `json:"image_url"`
	ShopID        int       `json:"shop_id"`
	CreatedAt     time.Time `json:"created_at"`
}

// PublicProductResponse is used for public API (guests)
type PublicProductResponse struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Category     string  `json:"category"`
	SellingPrice float64 `json:"selling_price"`
	Stock        int     `json:"stock"`
	ImageURL     string  `json:"image_url"`
	WhatsAppLink string  `json:"whatsapp_link"`
}

// AdminProductResponse is for Admin users (no purchase price)
type AdminProductResponse struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Category     string    `json:"category"`
	SellingPrice float64   `json:"selling_price"`
	Stock        int       `json:"stock"`
	ImageURL     string    `json:"image_url"`
	ShopID       int       `json:"shop_id"`
	CreatedAt    time.Time `json:"created_at"`
}

func (p *Product) ToPublicResponse(whatsappNumber string) PublicProductResponse {
	return PublicProductResponse{
		ID:           p.ID,
		Name:         p.Name,
		Description:  p.Description,
		Category:     p.Category,
		SellingPrice: p.SellingPrice,
		Stock:        p.Stock,
		ImageURL:     p.ImageURL,
		WhatsAppLink: GenerateWhatsAppLink(whatsappNumber, p.Name),
	}
}

func (p *Product) ToAdminResponse() AdminProductResponse {
	return AdminProductResponse{
		ID:           p.ID,
		Name:         p.Name,
		Description:  p.Description,
		Category:     p.Category,
		SellingPrice: p.SellingPrice,
		Stock:        p.Stock,
		ImageURL:     p.ImageURL,
		ShopID:       p.ShopID,
		CreatedAt:    p.CreatedAt,
	}
}
