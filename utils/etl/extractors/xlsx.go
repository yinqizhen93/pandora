package extractors

import (
	"fmt"
	"github.com/xuri/excelize/v2"
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

func (xl Xlsx) extract() error {
	f, err := excelize.OpenFile(xl.file)
	if err != nil {
		return err
	}
	rows, err := f.GetRows(xl.sht)
	header := rows[0]
	fmt.Println(header)
	return nil
}
