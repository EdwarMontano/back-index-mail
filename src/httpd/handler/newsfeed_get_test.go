package handler

import (
	"net/http"
	"testing"

	"github.com/EdwarMontano/back-index-mail/src/platform/enronmail"
	"github.com/EdwarMontano/back-index-mail/src/platform/mock_http"
)

func TestMockMailGet(t *testing.T) {
	feed := enronmail.New()
	feed.Add(enronmail.Item{"helloo", "world", "C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9"})

	handler := MockMailGet(feed)

	w := &mock_http.ResponseWriter{}
	r := &http.Request{}

	handler(w, r)

	result := w.GetBodyJSONArray()
	// fmt.Println(result)
	if len(result) != 1 {
		t.Errorf("Item was not added to the datastore")
	}

	if result[0]["IdMsg"] != "helloo" {
		t.Errorf("Item was not properly set")
	}
}
