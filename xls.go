package xls2json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/extrame/xls"
)

// ParseXLS reads all data from 1st sheet of xls file and returns it in JSON map or array
func ParseXLS(r io.Reader, mode string) ([]byte, error) {
	var (
		headers []string
		resMap  []*map[string]string
		result  []byte
		wb      = new(xls.WorkBook)
		err     error
	)

	buffer, err := io.ReadAll(r)
	//buffer, err := ioutil.ReadAll(r) // for Go<1.16
	if err != nil {
		return nil, fmt.Errorf("Read error")
	}

	fileBytes := bytes.NewReader(buffer) // converted to io.ReadSeeker type

	if wb, err = xls.OpenReader(fileBytes, "utf-8"); err != nil {
		return nil, err
	}

	sheet := wb.GetSheet(0)

	if sheet.MaxRow < 1 {
		return nil, fmt.Errorf("No data in first sheet")
	}

	if mode == "map" {
		for j := 0; j < sheet.Row(0).LastCol(); j++ {
			headers = append(headers, sheet.Row(0).Col(j))
		}
		for i := 1; i <= int(sheet.MaxRow); i++ {
			var tmpMap = make(map[string]string)
			for j := 0; j < sheet.Row(i).LastCol(); j++ {
				tmpMap[strings.Join(strings.Split(headers[j], " "), "")] = sheet.Row(i).Col(j)
			}
			resMap = append(resMap, &tmpMap)
		}
		result, err = json.Marshal(resMap)
	} else {
		result, err = json.Marshal(wb.ReadAllCells(524288))
	}
	return result, nil
}
