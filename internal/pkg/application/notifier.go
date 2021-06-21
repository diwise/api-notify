package application

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/diwise/api-notify/pkg/models"
	mq "github.com/diwise/messaging-golang/pkg/messaging"
	"github.com/streadway/amqp"
)

func (a *App) notificationMessageHandler() mq.TopicMessageHandler {
	return func(msg amqp.Delivery) {
		var nm NotificationMessage
		json.Unmarshal(msg.Body, &nm)
		m := JsonObjectToMap(nm.Payload)

		//TODO: det sak vara möjligt att hämta baserat på IdPattern också

		entityType := m["type"].(string)
		id := m["id"].(string)

		subscriptions := a.db.GetSubscriptionsByIdOrType(context.Background(), id, entityType)

		for _, s := range subscriptions {
			if status, err := notify(s, nm.Payload); err != nil {
				log.Printf("%s \n %s", status, err.Error())
			} else {
				log.Printf("Notification sent to %s", s.Notification.Endpoint.Uri)
			}
		}
	}
}

func notify(subscription models.Subscription, jsonData string) (string, error) {

	s := fmt.Sprintf(`
	{ 
		"subscriptionId": "%s", 
		"data": [%s] 
	}`, subscription.Id, jsonData)

	log.Printf("POST to %s\n%s", subscription.Notification.Endpoint.Uri, s)

	resp, err := http.Post(subscription.Notification.Endpoint.Uri, "application/json", strings.NewReader(s))

	return resp.Status, err
}
