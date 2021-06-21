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
	topicName   string
	contentType string
	Payload     string
}

func (nm NotificationMessage) TopicName() string {
	return nm.topicName
}

func (nm NotificationMessage) ContentType() string {
	return nm.contentType
}

func NewNotificationMessage(payload string) *NotificationMessage {
	return &NotificationMessage{
		contentType: "application/json",
		topicName:   "notify",
		Payload:     payload,
	}
}
