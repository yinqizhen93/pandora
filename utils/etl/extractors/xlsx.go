package extractors

import (
	"strconv"
)

type Xlsx struct {
	file string
	sht  string
}

type Type string

const (
	String  Type = "string"
	Int     Type = "int"
	Float64 Type = "float64"
	//Unknown Type = "unknown"
)

//func (xl Xlsx) extract() error {
//	f, err := excelize.OpenFile(xl.file)
//	if err != nil {
//		return err
//	}
//	rows, err := f.GetRows("Sheet1")
//	header := rows[0]
//}

func detectType(s string) (interface{}, Type) {
	if v, err := strconv.Atoi(s); err != nil {
		return v, String
	}
	if v, err := strconv.ParseFloat(s, 64); err != nil {
		return v, Float64
	}
	return s, String
}
