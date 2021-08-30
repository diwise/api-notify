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

		log.Printf("message body: %s", string(msg.Body))

		json.Unmarshal(msg.Body, &nm)
		//m := JsonObjectToMap(string(nm.Body))

		//TODO: det ska vara möjligt att hämta baserat på IdPattern också

		entityType := nm.EntityType
		id := nm.EntityID

		if entityType == "" || id == "" {
			log.Printf("bad event without type and id information received")
			return
		}

		subscriptions := a.db.GetSubscriptionsByIdOrType(context.Background(), id, entityType)

		for _, s := range subscriptions {
			if status, err := notify(s, string(nm.Body)); err != nil {
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

	if err != nil {
		return "", err
	}

	log.Printf("POST returned code %d", resp.StatusCode)

	return resp.Status, nil
}
