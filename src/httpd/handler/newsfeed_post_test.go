package handler

import (
	"net/http"
	"testing"

	"github.com/EdwarMontano/back-index-mail/src/platform/enronmail"
	"github.com/EdwarMontano/back-index-mail/src/platform/mock_http"
)

func TestMockMailPost(t *testing.T) {
	feed := enronmail.New()

	headers := http.Header{}
	headers.Add("content-type", "application/json")

	w := &mock_http.ResponseWriter{}
	r := &http.Request{
		Header: headers,
	}

	r.Body = mock_http.RequestBody(map[string]string{
		"IdMsg":      "hello",
		"DateMsg":    "world",
		"FromMsg":    "C1",
		"ToMsg":      "C2",
		"SubjectMsg": "C3",
		"CcMsg":      "C4",
		"BccMsg":     "C5",
		"XFromMsg":   "C6",
		"XToMsg":     "C7",
		"XccMsg":     "C8",
		"XbccMsg":    "C9",
	})

	handler := MockMailPost(feed)
	handler(w, r)

	result := w.GetBodyString()

	if result != "Proceso exitoso!" {
		t.Errorf("Handler did not complete")
	}

	if len(feed.GetAll()) != 1 {
		t.Errorf("Item did not add")
	}

	if feed.GetAll()[0].IdMsg != "hello" {
		t.Errorf("Item bad")
	}
}
