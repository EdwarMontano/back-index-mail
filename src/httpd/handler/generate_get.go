package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/EdwarMontano/back-index-mail/src/etldb"
)

func GenerateMailPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HI ALIENs")
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)
		fmt.Println(request["option"])
		option := request["option"]
		fmt.Println(reflect.TypeOf(option))
		// defer track(runningtime("execute"))
		etldb.ExtractData(option)
		w.Write([]byte("Proceso se ejecuto!"))
	}
}
