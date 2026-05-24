package services

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func POST(url string, body map[string]any) (error, map[string]any) {
	postBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("Failed to marshal request body: %v", err)
		return err, nil
	}

	log.Printf("POST %s | Request Body: %s", url, string(postBody))

	requestBody := bytes.NewBuffer(postBody)
	request, err := http.NewRequest(http.MethodPost, url, requestBody)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return err, nil
	}

	apiKey := os.Getenv("WHATSAPP_API_KEY")
	request.Header.Add("Authorization", "Bearer "+apiKey)
	request.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Request failed: %v", err)
		return err, nil
	}
	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return err, nil
	}

	log.Printf("POST %s | Status: %d | Response Body: %s", url, response.StatusCode, string(responseBytes))

	responseBody := map[string]any{}
	err = json.Unmarshal(responseBytes, &responseBody)
	if err != nil {
		log.Printf("Failed to unmarshal response body: %v", err)
		return err, nil
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		log.Printf("Request returned non-2xx status: %d", response.StatusCode)
		return err, nil
	}

	return nil, responseBody
}
