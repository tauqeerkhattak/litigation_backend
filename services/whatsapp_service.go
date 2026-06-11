package services

import (
	"log"
	"os"
)

func SendWhatsappReminder(phoneNumber string) error {
	phoneNumberId := os.Getenv("WHATSAPP_PHONE_NUMBER_ID")
	url := "https://graph.facebook.com/v25.0/" + phoneNumberId + "/messages"
	body := map[string]any{
		"messaging_product": "whatsapp",
		"to":                phoneNumber,
		"type":              "template",
		"template": map[string]any{
			"name":     "hello_world",
			"language": map[string]any{"code": "en_US"},
		},
	}
	err, response := POST(url, body)
	if err != nil {
		return err
	}
	log.Println("RESPONSE: ", response)
	return nil
}
