package xls2json

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"strings"
)

//ParseCSV reads all data from csv file and returns it in JSON map or array
func ParseCSV(r io.Reader, mode string, delimiter rune) ([]byte, error) {
	var (
		resMap  []*map[string]string
		err     error
		headers []string
		data    [][]string
		result  []byte
	)

	//if delimiter == nil {
	//	delimiter = ','
	//}
	// initialize reader
	reader := csv.NewReader(r)
	reader.Comma = rune(delimiter)
	reader.Comment = '#'
	if data, err = reader.ReadAll(); err != nil {
		return nil, err
	}

	if mode == "map" {
		// get headers
		headers = data[0]
		for _, row := range data[1:] {
			var tmpMap = make(map[string]string)
			for j, v := range row {
				tmpMap[strings.Join(strings.Split(headers[j], " "), "")] = v
			}
			resMap = append(resMap, &tmpMap)
		}

		result, err = json.Marshal(resMap)
	} else {
		result, err = json.Marshal(data)
	}

	return result, err
}
