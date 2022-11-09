package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/EdwarMontano/back-index-mail/src/platform/model"
)

func SearchMailPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request := map[string]string{}
		// json.NewDecoder(r.Body).Decode(&request)
		// find := request["word"]
		// fmt.Println(find)
		search_post()
		/////////

		// rta := search_post()
		// decoder := json.NewDecoder(resp.Body)
		// var result ResultZincSearch
		// err = decoder.Decode(&ResultZincSearch)
		// if err != nil {
		// 	fmt.Fprintf(w, "error: %v", err)
		// 	return
		// }
		// fmt.Println(result.Took)
		// fmt.Fprintf(w, "payload: %v \n", result)

		/////////
		// decoder := json.NewDecoder(r.Body)
		// var metadata MetaData
		// err := decoder.Decode(&metadata)
		// if err != nil {
		// 	fmt.Fprintf(w, "error: %v", err)
		// 	return
		// }
		// fmt.Fprintf(w, "payload: %v \n", metadata)
		w.Write([]byte("Busqueda exitosa!"))
	}
}

func search_post() http.HandlerFunc {
	query := `{
        "search_type": "match",
        "query":
        {
            "term": "please",
            "start_time": "2021-12-25T15:08:48.777Z",
            "end_time": "2022-12-07T16:08:48.777Z"
        },
        "from": 0,
        "max_results": 2,
        "_source": []
    }`
	req, err := http.NewRequest("POST", "http://localhost:4080/api/enron_mail/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#4321")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	fmt.Println(reflect.TypeOf(body))
	fmt.Println(reflect.TypeOf(resp.Body))

	var resultados []model.ResultZincSearch
	json.Unmarshal(body, &model.ResultZincSearch)

	for _, resultado := range resultados {
		fmt.Println(resultado.Took)
	}

	// json.NewDecoder(resp.Body).Decode(&ResultZincSearch)
	// fmt.Println(string(ResultZincSearch["hits"]))
	// fmt.Println(string(body))
	// for i := range body["hits"]["hits"] {
	// 	fmt.Println(body["hits"]["hits"][i]["_source"]["To"])
	// }

}
