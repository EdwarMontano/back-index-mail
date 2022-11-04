package etldb

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Info da info del archivo
func InfoLoad(file string) {
	fmt.Println(file)
}
func LoadData(data string) {
	fmt.Println("--->Empezó a cargarse la DATA")

	req, err := http.NewRequest("POST", "http://localhost:4080/api/enron_mail/_doc", strings.NewReader(data))
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
	fmt.Println("Termino de cargarse la DATA<---")
}
