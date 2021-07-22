package xls2json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/knieriem/odf/ods"
)

// ParseODS reads all data from 1st sheet of ods file and returns it in JSON map or array
func ParseODS(r io.Reader, mode string) ([]byte, error) {
	var (
		headers []string
		resMap  []*map[string]string
		result  []byte
		wb      = new(ods.File)
		doc     ods.Doc
		err     error
	)

	buffer, err := io.ReadAll(r)
	//buffer, err := ioutil.ReadAll(r) // for Go<1.16
	if err != nil {
		return nil, fmt.Errorf("Read error")
	}

	fileBytes := bytes.NewReader(buffer)

	if wb, err = ods.NewReader(fileBytes, int64(len(buffer))); err != nil {
		return nil, err
	}

	if err := wb.ParseContent(&doc); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	if len(doc.Table) < 1 {
		return nil, fmt.Errorf("No data")
	}

	rows := doc.Table[0].Strings()
	if len(rows) < 1 {
		return nil, fmt.Errorf("No data in first sheet")
	}

	if mode == "map" {
		headers = rows[0]
		for _, row := range rows[1:] {
			var tmpMap = make(map[string]string)
			for j, v := range row {
				tmpMap[strings.Join(strings.Split(headers[j], " "), "")] = v
			}
			resMap = append(resMap, &tmpMap)
		}
		result, err = json.Marshal(resMap)
	} else {
		result, err = json.Marshal(rows)
	}
	return result, nil
}
