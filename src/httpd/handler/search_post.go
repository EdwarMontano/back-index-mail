package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/EdwarMontano/back-index-mail/src/models"
	// "github.com/EdwarMontano/back-index-mail/src/httpd/model"
)

func SearchMailPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request := map[string]string{}
		// json.NewDecoder(r.Body).Decode(&request)
		// find := request["word"]
		// fmt.Println(find)
		busquedad, _ := search_post()
		/////////

		// rta := search_post()
		// decoder := json.NewDecoder(resp.Body)
		// var result ResultZincSearch
		// err = decoder.Decode(&ResultZincSearch)
		// if err != nil {
		// 	fmt.Fprintf(w, "error: %v", err)
		// 	return
		// }

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
		json.NewEncoder(w).Encode(busquedad)
		// w.Write([]byte())
	}
}

func search_post() (results []models.Mail, err error) {
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
	// fmt.Println(string(body))
	fmt.Println(string(body))
	// fmt.Println(reflect.TypeOf(body))
	fmt.Println(reflect.TypeOf(resp.Body))
	fmt.Println(resp.StatusCode)
	byt := []byte(body)
	var eRes map[string]interface{}
	if err := json.Unmarshal(byt, &eRes); err != nil {
		panic(err)
	}
	fmt.Println(eRes)
	num := eRes["took"].(float64)
	fmt.Println(num)

	var mails []models.Mail
	for _, hit := range eRes["hits"].(map[string]interface{})["hits"].([]interface{}) {
		fmt.Println("HI ALIENS")
		mail := models.Mail{}
		source := hit.(map[string]interface{})["_source"]
		marshal, err := json.Marshal(source)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(marshal, &mail); err == nil {
			// fmt.Println(mail.To)
			mails = append(mails, mail)
		}
		fmt.Println(mail.To)
	}
	return mails, nil

	// var eRes map[string]interface{}
	// if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
	// 	return nil, err
	// }

	// var resultados []model.ResultZincSearch
	// json.Unmarshal(body, &model.ResultZincSearch)

	// for _, resultado := range resultados {
	// 	fmt.Println(resultado.Took)
	// }

	// json.NewDecoder(resp.Body).Decode(&ResultZincSearch)
	// fmt.Println(string(ResultZincSearch["hits"]))
	// fmt.Println(string(body))
	// for i := range body["hits"]["hits"] {
	// 	fmt.Println(body["hits"]["hits"][i]["_source"]["To"])
	// }
}
