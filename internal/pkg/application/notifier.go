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

func (a *notifierApp) notificationMessageHandler() mq.TopicMessageHandler {
	return func(msg amqp.Delivery) {
		var nm NotificationMessage

		log.Printf("message body: %s", string(msg.Body))

		json.Unmarshal(msg.Body, &nm)
		//m := jsonObjectToMap(string(nm.Body))

		//TODO: det ska vara möjligt att hämta baserat på IdPattern också

		entityType := nm.EntityType
		id := nm.EntityID

		if entityType == "" || id == "" {
			log.Printf("bad event without type and id information received")
			return
		}

		subscriptions, err := a.db.GetSubscriptionsByIdOrType(context.Background(), id, entityType)
		if err != nil {
			log.Printf("failed to get matching subscriptions: %s\n", err.Error())
			return
		}

		for _, s := range subscriptions {
			status, err := notify(s, string(nm.Body))
			if err != nil {
				log.Printf("%s \n %s", status, err.Error())
				continue
			}

			log.Printf("notification sent to %s", s.Notification.Endpoint.Uri)
		}
	}
}

/*func jsonObjectToMap(jsonData string) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonData), &result)
	return result
}*/

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

	return resp.Status, nil
}
