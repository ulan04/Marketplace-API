package services

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type NotificationService struct {
	client *resty.Client
}

func NewNotificationService() *NotificationService {
	client := resty.New()
	return &NotificationService{client: client}
}

func (s *NotificationService) SendItemCreatedNotification(itemID uint, title string) error {

	webhookURL := "https://webhook.site/655efcc7-a4bb-4504-94b9-5a5ea79f40f7"

	payload := map[string]interface{}{
		"event":   "item_created",
		"item_id": itemID,
		"title":   title,
	}

	log.Println("Sending notification with resty...")

	resp, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(webhookURL)

	if err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}

	log.Println("Webhook response status:", resp.Status())

	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("notification failed with status: %d", resp.StatusCode())
	}

	return nil
}