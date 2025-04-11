package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) SetWebhook(webhookURL string) error {
	payload := map[string]string{"url": webhookURL}
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/setWebhook", c.apiURL),
		"application/json",
		bytes.NewBuffer(data))

	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("setWebhook failed: %d", resp.StatusCode)
	}

	log.Println("Telegram Bot webhook set successfully")
	return nil
}
