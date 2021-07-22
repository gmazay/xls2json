package xls2json

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// ParseXLSX reads all data from 1st sheet of xlsx file and returns it in JSON map or array
func ParseXLSX(r io.Reader, mode string) ([]byte, error) {
	var (
		headers   []string
		resMap    []*map[string]string
		wb        = new(excelize.File)
		err       error
		sheetName string
		result    []byte
	)

	if wb, err = excelize.OpenReader(r); err != nil {
		return nil, err
	}
	// Get all the rows in the Sheet.
	sheetName = wb.GetSheetMap()[1]
	//fmt.Printf("SheetName: %v", sheetName)
	rows := wb.GetRows(sheetName)
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
