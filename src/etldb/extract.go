package etldb

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	output_path = filepath.Join("./output")
	bash_script = filepath.Join("_script.sh")
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
func exe_cmd(cmds []string) {
	os.RemoveAll(output_path)
	err := os.MkdirAll(output_path, os.ModePerm|os.ModeDir)
	checkError(err)
	file, err := os.Create(filepath.Join(output_path, bash_script))
	checkError(err)
	defer file.Close()
	file.WriteString("#!/bin/sh\n")
	file.WriteString(strings.Join(cmds, "\n"))
	err = os.Chdir(output_path)
	checkError(err)
	out, err := exec.Command("sh", bash_script).Output()
	checkError(err)
	fmt.Println(string(out))
}

func filePathDB() {
	commands := []string{
		// "echo newline >> foo.o",
		"tree /home/chocoplot/Documents/codeLAB/IndexMail_GoVue/db/enron_mail_20110402/maildir >> file.txt",
		"du -a /home/chocoplot/Documents/codeLAB/IndexMail_GoVue/db/enron_mail_20110402/maildir >> filepath.txt",
	}
	exe_cmd(commands)
}

func readfileAll(ruta string, i uint32) {
	// fmt.Println("******** Leyendo archivo ********")

	archivo, err := os.ReadFile(ruta)
	if err != nil {
		log.Fatal(err)
	}

	text := string(archivo)
	// fmt.Println(text)
	TransformData(text, ruta, i)
	// InfoExtract(text)
	// fmt.Println("******** Leyendo archivo ********")
}
func clearPath(linea string) string {

	// fmt.Println(linea)
	indexChar := strings.Index(linea, "/")
	lastChar := linea[len(linea)-1:]
	// fmt.Println(lastChar)
	if lastChar == "." {
		linea = linea[indexChar:]
		return linea
	} else {
		return "IS-FOLDER"
	}
}
func readFileLine() {
	fmt.Println("******** Empezó lectura de filePath.TXT ********")
	archivo, error := os.Open("/home/chocoplot/Documents/codeLAB/IndexMail_GoVue/back-index-mail/output/filepath.txt")
	defer archivo.Close()
	if error != nil {
		fmt.Println("Hubo Error")
	}
	scanner := bufio.NewScanner(archivo)
	var i uint32

	for scanner.Scan() {
		i++
		linea := scanner.Text()
		// fmt.Println("-", linea)
		pathfile := clearPath(linea)
		if pathfile == "IS-FOLDER" {
			// fmt.Println("Hola soy folder")
			continue
		}
		// fmt.Print(i)
		if i == 1 {
			// fmt.Println("Hola se creo NDJSON")

			_, err := os.Create("/home/chocoplot/Documents/codeLAB/IndexMail_GoVue/back-index-mail/output/EnronMail.ndjson")

			if err != nil {
				fmt.Println(err)
				return
			}
		}

		readfileAll(pathfile, i)

		// if i == 2 {
		// 	fmt.Println(len(linea))
		// 	s := strings.Split(linea, " ")
		// 	fmt.Println(s)
		// }
	}
	fmt.Println("******** Terminó lectura de filePath.TXT ********")
}

// Info da info del archivo
func InfoExtract(file string) {
	fmt.Println(file)
}

func ExtractData(numSelect string) {
	// var numSelect int
	fmt.Println("--->Empezó a extraerse la DATA")
	fmt.Printf("1. Desea generar el .TXT con las rutas de los archivos.\n2. Ya tiene el .TXT con las rutas.\n")
	fmt.Print("Escriba un número con la opción: ")
	fmt.Print(numSelect)
	// fmt.Scanln(&numSelect)
	switch numSelect {
	case "1":
		filePathDB()
	case "2":
		fmt.Println("continuo...")
		readFileLine()
	}
	//
	fmt.Println("Termino de extraerse la DATA<---")
}
