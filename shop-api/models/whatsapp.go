package models

import (
	"net/url"
	"strings"
)

// GenerateWhatsAppLink creates a WhatsApp link for product inquiries
func GenerateWhatsAppLink(whatsappNumber, productName string) string {
	// Clean the WhatsApp number (remove spaces, dashes, etc.)
	cleanNumber := strings.ReplaceAll(whatsappNumber, " ", "")
	cleanNumber = strings.ReplaceAll(cleanNumber, "-", "")
	cleanNumber = strings.ReplaceAll(cleanNumber, "(", "")
	cleanNumber = strings.ReplaceAll(cleanNumber, ")", "")

	// Create the message text
	message := "Bonjour je veux plus d'information sur " + productName

	// URL encode the message
	encodedMessage := url.QueryEscape(message)

	// Return the formatted WhatsApp link
	return "https://wa.me/" + cleanNumber + "?text=" + encodedMessage
}
