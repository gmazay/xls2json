package xls2json

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	path := flag.String("file", "test.xlsx", "xlsx file")
	flag.Parse()

	file, err := os.Open(*path)
	if err != nil {
		log.Fatalf(`unable to open file, error: %s`, err)
	}
	defer file.Close()
	r := bufio.NewReader(file)

	//for _, val := range result {
	resReader, err := DataToJSON(r, "csv", "map", ',')
	if err != nil {
		log.Fatalf(`unable to parse file, error: %s`, err)
	}
	if b, err := io.ReadAll(resReader); err == nil {
		fmt.Println(string(b))
	}
}

// DataToJSON is method to parse xlsx, xls, ods, csv to json
// filetype: xlsx, xls, ods, csv
// mode: map, array
// delimiter: for csv parsing
func DataToJSON(r io.Reader, filetype, mode string, delimiter rune) (io.Reader, error) {
	var (
		resBytes []byte
		err      error
	)

	switch filetype {
	case "xlsx":
		resBytes, err = ParseXLSX(r, mode)
		if err != nil {
			log.Fatalf(`unable to parse xlsx file, error: %s`, err)
			return nil, err
		}
	case "xls":
		resBytes, err = ParseXLS(r, mode)
		if err != nil {
			log.Fatalf(`unable to parse xls file, error: %s`, err)
			return nil, err
		}
	case "ods":
		resBytes, err = ParseODS(r, mode)
		if err != nil {
			log.Fatalf(`unable to parse ods file, error: %s`, err)
			return nil, err
		}
	case "csv":
		resBytes, err = ParseCSV(r, mode, delimiter)
		if err != nil {
			log.Fatalf(`unable to parse csv file, error: %s`, err)
			return nil, err
		}
	default:
		log.Fatalf(`wrong filetype, expected: ods, xlsx, xls, csv`)
		return nil, err
	}

	return bytes.NewReader(resBytes), err
}
