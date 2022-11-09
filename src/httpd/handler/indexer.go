package handler

import (
	"fmt"
	"log"
	"time"

	"github.com/EdwarMontano/back-index-mail/src/etldb"
)

func runningtime(s string) (string, time.Time) {
	log.Println("Start:	", s)
	return s, time.Now()
}

func track(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("End:	", s, "took", endTime.Sub(startTime))
}

func IndexerEnron() {
	fmt.Println("******** Empieza 3. Ejecutando ETL para enron_mail_2011040 ******** ")
	defer track(runningtime("execute"))
	etldb.ExtractData()
	// etldb.LoadData()
	fmt.Println(" ------------- Se ejecuto Correctamente 3.LEER FILE ----------")
}
