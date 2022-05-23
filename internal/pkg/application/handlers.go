package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/diwise/api-notify/pkg/models"
	"github.com/go-chi/render"

	"github.com/google/uuid"
)

// /subscriptions - returns list of subscriptions
func (a *notifierApp) listSubscriptions(w http.ResponseWriter, r *http.Request) {
	if s, err := a.db.ListSubscriptions(r.Context(), 0); err == nil {
		if data, err := json.Marshal(s); err == nil {
			Ok(w, data)
			return
		} else {
			InternalServerError(w, err)
			return
		}
	}
	NotFound(w)
}

func (a *notifierApp) getSubscription(w http.ResponseWriter, r *http.Request) {
	subscription := r.Context().Value(subscriptionKey) //.(*models.Subscription)

	if subscription == nil {
		NotFound(w)
		return
	}

	if data, err := json.Marshal(subscription); err == nil {
		Ok(w, data)
	} else {
		InternalServerError(w, err)
	}
}

func (a *notifierApp) createSubscription(w http.ResponseWriter, r *http.Request) {
	subReq := &SubscriptionRequest{}

	if err := render.Bind(r, subReq); err != nil {
		InternalServerError(w, err)
		return
	}

	if subReq.Id == "" {
		u := uuid.New()
		uuid := strings.Replace(u.String(), "-", "", -1)
		subReq.Id = uuid
	}

	if err := a.db.CreateSubscription(r.Context(), subReq.Subscription); err == nil {
		Created(w, fmt.Sprintf("/subscriptions/%s", subReq.Id))
	} else {
		InternalServerError(w, err)
	}
}

func (a *notifierApp) updateSubscription(w http.ResponseWriter, r *http.Request) {
	subscription := r.Context().Value("subscription") //.(*models.Subscription)

	if subscription == nil {
		NotFound(w)
		return
	}

	s := *subscription.(*models.Subscription)

	if err := a.db.UpdateSubscription(r.Context(), &s); err == nil {
		NoContent(w)
	} else {
		InternalServerError(w, err)
	}
}

func (a *notifierApp) deleteSubscription(w http.ResponseWriter, r *http.Request) {
	subscription := r.Context().Value(subscriptionKey) //.(*models.Subscription)

	if subscription == nil {
		NotFound(w)
		return
	}

	s := *subscription.(*models.Subscription)

	if err := a.db.DeleteSubscription(r.Context(), s.Id); err == nil {
		NoContent(w)
	} else {
		InternalServerError(w, err)
	}
}

//TODO: Ändra så att rätt objekt skickas som svar. Använda vanliga ngsi diwise pkg

func Ok(w http.ResponseWriter, data []byte) {
	w.Write(data)
}
func Created(w http.ResponseWriter, location string) {
	w.Header().Add("Location", location)
	w.WriteHeader(http.StatusCreated)
}
func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
func InternalServerError(w http.ResponseWriter, err error) {
	w.Write([]byte(err.Error()))
	w.WriteHeader(http.StatusInternalServerError)
}
