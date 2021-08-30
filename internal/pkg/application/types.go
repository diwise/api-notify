package application

import (
	"errors"
	"net/http"

	"github.com/diwise/api-notify/pkg/models"
)

type SubscriptionRequest struct {
	*models.Subscription
}

func (s *SubscriptionRequest) Bind(r *http.Request) error {
	if s.Subscription == nil {
		return errors.New("missing required Subscription fields")
	}

	return nil
}

type NotificationMessage struct {
	EntityType string `json:"type"`
	EntityID   string `json:"id"`
	Body       string `json:"body"`
}

func (nm NotificationMessage) TopicName() string {
	return "notify"
}

func (nm NotificationMessage) ContentType() string {
	return "application/json"
}

func NewNotificationMessage(body string) *NotificationMessage {
	return &NotificationMessage{
		EntityType: "Device",
		EntityID:   "someid",
		Body:       body,
	}
}
