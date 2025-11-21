package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type CollectRequest struct {
	PhoneNumber string `json:"from"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
}

type PaymentResponse struct {
	TransactionID string `json:"reference"`
	Status        string `json:"status"`
	Message       string `json:"message"`
}

func isCameroonNumber(num string) bool {
    return strings.HasPrefix(num, "237") && len(num) == 12
}


func MakeMTNMobileTransferFromCampayAccount(req CollectRequest) (PaymentResponse, error) {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			return PaymentResponse{}, fmt.Errorf("failed to load env file: %w", err)
		}
	}

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		return PaymentResponse{}, fmt.Errorf("API_KEY is not set")
	}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return PaymentResponse{}, fmt.Errorf("failed to marshal json: %w", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}

	// Create the POST request
	httpReq, err := http.NewRequest(
		"POST",
		"https://demo.campay.net/api/collect/",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return PaymentResponse{}, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", apiKey)

	// Send request
	resp, err := client.Do(httpReq)
	if err != nil {
		return PaymentResponse{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	fmt.Println("Please complete the transaction on your mobile device...")
	time.Sleep(10 * time.Second) //give the user time to complete the transaction

	// Parse response
	var paymentResp PaymentResponse
	err = json.NewDecoder(resp.Body).Decode(&paymentResp)
	if err != nil {
		return PaymentResponse{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return paymentResp, nil

}