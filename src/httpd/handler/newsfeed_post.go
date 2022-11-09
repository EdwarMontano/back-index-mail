package handler

import (
	"encoding/json"
	"net/http"

	"github.com/EdwarMontano/back-index-mail/src/platform/enronmail"
)

func MockMailPost(feed enronmail.Adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		feed.Add(enronmail.Item{
			IdMsg:      request["IdMsg"],
			DateMsg:    request["DateMsg"],
			FromMsg:    request["FromMsg"],
			ToMsg:      request["ToMsg"],
			SubjectMsg: request["SubjectMsg"],
			CcMsg:      request["CcMsg"],
			BccMsg:     request["BccMsg"],
			XFromMsg:   request["XFromMsg"],
			XToMsg:     request["XToMsg"],
			XccMsg:     request["XccMsg"],
			XbccMsg:    request["XbccMsg"],
		})

		w.Write([]byte("Proceso exitoso!"))
	}
}
