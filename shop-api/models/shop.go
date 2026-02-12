package models

import "time"

type Shop struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Active         bool      `json:"active"`
	WhatsAppNumber string    `json:"whatsapp_number"`
	CreatedAt      time.Time `json:"created_at"`
}
