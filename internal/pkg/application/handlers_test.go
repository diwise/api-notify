package application

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestListSubscriptions(t *testing.T) {
	is := is.New(t)

	_, _, app := newAppForTesting()
	ts := httptest.NewServer(app.router)
	defer ts.Close()

	resp, _ := testRequest(is, ts, "GET", "/subscriptions", nil)

	is.Equal(resp.StatusCode, http.StatusOK) // Check status code
}
