package application

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	db "github.com/diwise/api-notify/internal/pkg/database"

	mq "github.com/diwise/messaging-golang/pkg/messaging"
)

type Application interface {
	Start(port string) error
}

type notifierApp struct {
	router chi.Router
	db     db.Db
	mq     mq.Context
}

func NewApplication(r chi.Router, db db.Db, mq mq.Context) Application {
	return newNotifierApp(r, db, mq)
}

func newNotifierApp(r chi.Router, db db.Db, mq mq.Context) *notifierApp {
	a := &notifierApp{
		router: r,
		db:     db,
		mq:     mq,
	}

	r.Use(middleware.Logger)

	r.Post("/notify/{entityType}", a.notify)

	r.Route("/subscriptions", func(r chi.Router) {
		r.Get("/", a.listSubscriptions)
		r.Post("/", a.createSubscription)
		r.Route("/{subscriptionId}", func(r chi.Router) {
			r.Use(a.subscriptionCtx)
			r.Get("/", a.getSubscription)
			r.Put("/", a.updateSubscription)
			r.Delete("/", a.deleteSubscription)
		})
	})

	a.mq.RegisterTopicMessageHandler("ngsi-entity-created", a.notificationMessageHandler())
	a.mq.RegisterTopicMessageHandler("ngsi-entity-updated", a.notificationMessageHandler())

	return a
}

func (a *notifierApp) Start(port string) error {
	return http.ListenAndServe(fmt.Sprintf(":%s", port), a.router)
}

func (a *notifierApp) notify(w http.ResponseWriter, r *http.Request) {
	// En metod för att kunna trigga en händelse som läggs på MQ.

	entityType := chi.URLParam(r, "entityType")
	if entityType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	str := string(body)
	err := a.mq.PublishOnTopic(NewNotificationMessage(str))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type contextKey string

const (
	subscriptionKey contextKey = "subscription"
)

func (a *notifierApp) subscriptionCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if subscriptionId := chi.URLParam(r, "subscriptionId"); subscriptionId != "" {
			if s, err := a.db.GetSubscriptionById(subscriptionId); err == nil {
				ctx := context.WithValue(r.Context(), subscriptionKey, s)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		ctx := context.WithValue(r.Context(), subscriptionKey, nil)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
