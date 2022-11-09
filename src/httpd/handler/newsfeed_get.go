package handler

import (
	"encoding/json"
	"net/http"

	"github.com/EdwarMontano/back-index-mail/src/platform/enronmail"
)

func MockMailGet(feed enronmail.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		json.NewEncoder(w).Encode(items)
	}
}
