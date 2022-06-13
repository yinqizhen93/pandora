package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"os"
)

func ReadXlsx(file string, sheet string) ([][]string, error) {
	f, err := excelize.OpenFile(file)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func ReadCsv(file string) ([][]string, error) {
	csvFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var data [][]string
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
			return nil, err
		}
		data = append(data, line)
	}
	return data, nil
}

//func main() {
//	file := "/Users/tison.yin/Desktop/tison/600000.csv"
//	c, err := ReadCsv(file)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(c)
//}
