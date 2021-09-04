package application

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/diwise/api-notify/internal/pkg/database"
	"github.com/diwise/messaging-golang/pkg/messaging"
	"github.com/go-chi/chi/v5"
	"github.com/matryer/is"
)

func TestThatHealthEndpointReturns204(t *testing.T) {
	is := is.New(t)

	_, _, app := newAppForTesting()
	ts := httptest.NewServer(app.router)
	defer ts.Close()

	resp, _ := testRequest(is, ts, "GET", "/health", nil)

	is.Equal(resp.StatusCode, http.StatusNoContent) // health endpoint status code not ok
}

func TestThatNotifyInvocationPostsMessageToMQ(t *testing.T) {
	is := is.New(t)

	_, mq, app := newAppForTesting()
	ts := httptest.NewServer(app.router)
	defer ts.Close()

	resp, _ := testRequest(is, ts, "POST", "/notify/WaterQualityObserved", nil)

	is.Equal(resp.StatusCode, http.StatusNoContent) // returned status code
	is.Equal(len(mq.PublishOnTopicCalls()), 1)      // number of calls to publish on topic
}

func TestThatNotifyInvocationReturns500OnMqFailure(t *testing.T) {
	is := is.New(t)

	_, mq, app := newAppForTesting()
	ts := httptest.NewServer(app.router)
	defer ts.Close()

	mq.PublishOnTopicFunc = func(messaging.TopicMessage) error { return errors.New("failure") }

	resp, _ := testRequest(is, ts, "POST", "/notify/WaterQualityObserved", nil)

	is.Equal(resp.StatusCode, http.StatusInternalServerError) // returned status code
}

func newAppForTesting() (*database.DbMock, *messaging.ContextMock, *notifierApp) {
	r := chi.NewRouter()
	db := &database.DbMock{}
	mq := &messaging.ContextMock{}
	return db, mq, newNotifierApp(r, db, mq)
}

func testRequest(is *is.I, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) {
	req, _ := http.NewRequest(method, ts.URL+path, body)
	resp, _ := http.DefaultClient.Do(req)
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return resp, string(respBody)
}
