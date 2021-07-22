package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/gmazay/xls2json"
)

func main() {
	var filetype string
	path := flag.String("file", "testa.xlsx", "xlsx file")
	mode := flag.String("mode", "map", "map or array")
	delimiter := flag.String("delimiter", ",", "csv delimiter")
	flag.Parse()

	if strings.HasSuffix(*path, ".xlsx") {
		filetype = "xlsx"
	} else if strings.HasSuffix(*path, ".xls") {
		filetype = "xls"
	} else if strings.HasSuffix(*path, ".ods") {
		filetype = "ods"
	} else if strings.HasSuffix(*path, ".csv") {
		filetype = "csv"
	} else {
		log.Fatal(`Unknown file suffix, expected xlsx or csv`)
		return
	}

	file, err := os.Open(*path)
	if err != nil {
		log.Fatalf(`unable to open file, error: %s`, err)
	}
	defer file.Close()
	//r := bufio.NewReader(file)

	resReader, err := xls2json.DataToJSON(file, filetype, *mode, []rune(*delimiter)[0])
	if err != nil {
		log.Fatalf(`Unable to parse file, error: %s`, err)
	}

	if b, err := io.ReadAll(resReader); err == nil {
		fmt.Println(string(b))
	}
}
