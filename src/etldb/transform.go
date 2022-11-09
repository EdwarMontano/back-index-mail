package etldb

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Correo struct {
	IdMsg      string
	DateMsg    string
	FromMsg    string
	ToMsg      string
	SubjectMsg string
	CcMsg      string
	BccMsg     string
	XFromMsg   string
	XToMsg     string
	XccMsg     string
	XbccMsg    string
	// XfolderMsg   string
	// XoriginMsg   string
	// XfileNameMsg string
	// Idiomas      []string
}

// Info da info del archivo
func InfoContent(file string) Correo {
	var id, date, from, to, subject, cc, bcc, xfrom, xto, xcc, xbcc, fieldCurrent string
	info := strings.Split(file, "\n")
	memory := []string{}
	for i, row := range info {
		// fmt.Println(i, "^^^", row)
		if len(row) < 3 {
			continue
		}
		if !(strings.Contains(row, ":")) {
			if len(memory) == 0 {
				rowPreview := strings.SplitAfterN(info[i-1], ": ", 2)
				fieldCurrent = rowPreview[0]
			}
			// fmt.Println("#######", fieldCurrent, "#######", len(fieldCurrent))
			memory = append(memory, row)
			//yyyyyyyyyy
			switch fieldCurrent {
			case "To: ":
				to = to + strings.TrimSpace(row)
			case "Subject: ":
				subject = subject + strings.TrimSpace(row)
			}
			//yyyyyyyyy
		}

		if strings.Contains(row, "Message-ID:") {
			sliceSplit := strings.SplitAfterN(row, ": ", 2)
			id = strings.TrimSpace(sliceSplit[1])
		}
		if strings.Contains(row, "Date:") {
			sliceSplit := strings.SplitAfterN(row, ": ", 2)
			date = strings.TrimSpace(sliceSplit[1])
		}
		if strings.Contains(row, "Subject:") {
			sliceSplit := strings.SplitAfterN(row, ": ", 2)
			// fmt.Println(len(sliceSplit))
			subject = strings.TrimSpace(sliceSplit[1])
		}
		if len(row) > 3 {
			if strings.Contains(row[0:3], "To:") {
				sliceSplit := strings.SplitAfterN(row, ": ", 2)
				to = strings.TrimSpace(sliceSplit[1])
			}
			if strings.Contains(row[0:3], "Cc:") {
				sliceSplit := strings.SplitAfterN(row, ": ", 2)
				cc = strings.TrimSpace(sliceSplit[1])
			}
		}
		if len(row) > 4 {
			if strings.Contains(row[0:4], "Bcc:") {
				sliceSplit := strings.SplitAfterN(row, ": ", 2)
				bcc = strings.TrimSpace(sliceSplit[1])
			}
		}
		if len(row) > 5 {
			if strings.Contains(row[0:5], "From:") {
				sliceSplit := strings.SplitAfterN(row, ": ", 2)
				from = strings.TrimSpace(sliceSplit[1])
			}
			if strings.Contains(row[0:5], "X-To:") {
				sliceSplit := strings.SplitAfterN(row, ": ", 2)

				xto = strings.TrimSpace(sliceSplit[1])
			}
			if strings.Contains(row[0:5], "X-cc:") {
				sliceSplit := strings.SplitAfterN(row, ": ", 2)
				xcc = strings.TrimSpace(sliceSplit[1])
			}
		}
		if len(row) > 6 {
			if strings.Contains(row[0:6], "X-bcc:") {
				sliceSplit := strings.SplitAfterN(row, ": ", 2)
				xbcc = strings.TrimSpace(sliceSplit[1])
			}
		}
		if len(row) > 7 {
			if strings.Contains(row[0:7], "X-From:") {
				sliceSplit := strings.SplitAfterN(row, ": ", 2)
				xfrom = strings.TrimSpace(sliceSplit[1])
			}
		}
		// if strings.Contains(row, "X-From:") {
		// 	sliceSplit := strings.SplitAfter(row, ": ")
		// 	from = strings.TrimSpace(sliceSplit[1])
		// }
		// fmt.Println(i, "***", row) // este recorrido lo puede hacer de manera aleatoria
	}
	data := Correo{
		IdMsg:      id,
		DateMsg:    date,
		FromMsg:    from,
		ToMsg:      to,
		SubjectMsg: subject,
		CcMsg:      cc,
		BccMsg:     bcc,
		XFromMsg:   xfrom,
		XToMsg:     xto,
		XccMsg:     xcc,
		XbccMsg:    xbcc,
	}
	return data
	// fmt.Println(len(info), "^^^^^", info[0])
}

// Info_part split infomation part out

func TransformData(corpusMail, ruta string, i uint32) {
	// fmt.Println("--->Empezó a transformarse la DATA")
	var porcentaje float32
	content := strings.SplitAfterN(corpusMail, "\n\r", 2)
	data := InfoContent(content[0])
	cutRuta := strings.SplitAfter(ruta, "db/")
	ruta = cutRuta[1]
	contenido := strings.TrimSpace(content[1])
	contenido = strings.ReplaceAll(contenido, "\n", " ")
	contenido = strings.ReplaceAll(contenido, "\r", " ")

	//  \"Path\":\"%s\"
	dataString := fmt.Sprintf("{\"IdMSG\":\"%s\",\"Path\":\"%s\",\"DateMSG\":\"%s\",\"From\":\"%s\",\"To\":\"%s\",\"Subject\":\"%s\",\"Cc\":\"%s\",\"Bcc\":\"%s\",\"X-From\":\"%s\",\"X-To\":\"%s\",\"X-cc\":\"%s\",\"X-bcc\":\"%s\",\"Content\":\"%s\"}\n", data.IdMsg, ruta, data.DateMsg, data.FromMsg, data.ToMsg, data.SubjectMsg, data.CcMsg, data.BccMsg, data.XFromMsg, data.XToMsg, data.XccMsg, data.XbccMsg, contenido)
	LoadData(dataString)
	porcentaje = (float32(100) / float32(517426)) * float32(i)
	porcentajeString := fmt.Sprintf("%.3f", porcentaje)
	fmt.Println("Transformado..", porcentajeString, "% ", ruta)
	f, err := os.OpenFile("/home/chocoplot/Documents/codeLAB/IndexMail_GoVue/back-index-mail/output/EnronMail.ndjson", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("{ \"index\" : { \"_index\" : \"enron_mail\" } }\n")); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(dataString)); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	// content := strings.SplitAfterN(corpusMail, "\r\n\n", 2)
	// fmt.Println(">>>>>", len(content))
	// InfoContent(content[0])
	// fmt.Println("-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-")
	// fmt.Println(content[1])
	// for i, row := range content {
	// 	// switch i {
	// 	// case 0:

	// 	// 	fmt.Println(row)

	// 	// }
	// 	fmt.Println(i, "***", row) // este recorrido lo puede hacer de manera aleatoria
	// }
	// fmt.Println("Termino de transformarse la DATA<---")

}
