package application

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	db "github.com/diwise/api-notify/internal/pkg/database"

	mq "github.com/diwise/messaging-golang/pkg/messaging"
)

type App struct {
	router chi.Router
	db     db.Db
	mq     mq.Context
}

func NewApplication(r chi.Router, db db.Db, mq mq.Context) *App {
	a := &App{
		router: r,
		db:     db,
		mq:     mq,
	}

	r.Use(middleware.Logger)

	r.Post("/notify/{entityType}", a.notify)

	r.Route("/subscriptions", func(r chi.Router) {
		r.Get("/", a.ListSubscriptions)
		r.Post("/", a.CreateSubscription)
		r.Route("/{subscriptionId}", func(r chi.Router) {
			r.Use(a.subscriptionCtx)
			r.Get("/", a.GetSubscription)
			r.Put("/", a.UpdateSubscription)
			r.Delete("/", a.DeleteSubscription)
		})
	})

	a.mq.RegisterTopicMessageHandler("ngsi-entity-created", a.notificationMessageHandler())
	a.mq.RegisterTopicMessageHandler("ngsi-entity-updated", a.notificationMessageHandler())

	return a
}

func (a *App) Start(port string) error {
	return http.ListenAndServe(fmt.Sprintf(":%s", port), a.router)
}

func (a *App) notify(w http.ResponseWriter, r *http.Request) {
	// En metod för att kunna trigga en händelse som läggs på MQ.
	// Att entityType ska vara i URL är egentligen inte nödvändigt.

	if entityType := chi.URLParam(r, "entityType"); entityType != "" {
		body, _ := ioutil.ReadAll(r.Body)
		str := string(body)
		a.mq.PublishOnTopic(NewNotificationMessage(str))
	}
}

func (a *App) subscriptionCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if subscriptionId := chi.URLParam(r, "subscriptionId"); subscriptionId != "" {
			if s, err := a.db.GetSubscriptionById(subscriptionId); err == nil {
				ctx := context.WithValue(r.Context(), "subscription", s)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		ctx := context.WithValue(r.Context(), "subscription", nil)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func JsonObjectToMap(jsonData string) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonData), &result)
	return result
}

func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}

	return cp
}
