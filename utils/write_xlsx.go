package utils

import (
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"reflect"
)

type XlsxStorage struct {
	file    *excelize.File
	header  []TableHeader
	data    []any
	nr      int
	nc      int
	rowType reflect.Type
	rowKind reflect.Kind
}

type TableHeader struct {
	StructFieldName string
	XlsxColName     string
}

func NewXlsxStorage(file *excelize.File, header []TableHeader, data []any) (*XlsxStorage, error) {
	nr := len(data)
	if nr == 0 {
		return nil, errors.New("data can not be empty")
	}
	nc := len(header)
	if nc == 0 {
		return nil, errors.New("header can not be empty")
	}
	// 以数据第一行类型为准
	t := reflect.TypeOf(data[0])
	xs := XlsxStorage{
		file:    file,
		header:  header,
		data:    data,
		nr:      nr,
		nc:      nc,
		rowType: t,
		rowKind: t.Kind(),
	}
	return &xs, nil
}

// WriteXlsx write data to Excel
func (f *XlsxStorage) WriteXlsx() error {
	if err := f.WriteHeader(); err != nil {
		return errors.Wrap(err, "写入header失败")
	}
	switch f.rowKind {
	// any为Struct
	case reflect.Struct:
		if err := f.WriteStructToXlsx(); err != nil {
			return err
		}
	case reflect.Ptr:
		if f.rowType.Elem().Kind() != reflect.Struct {
			return errors.New("data can only be []slice or []struct or []*struct")
		}
		if err := f.WriteStructToXlsx(); err != nil {
			return err
		}
	case reflect.Slice:
		if err := f.WriteSliceToXlsx(); err != nil {
			return err
		}
	default:
		return errors.New("data can only be []slice or []struct or []*struct")
	}
	return nil
}

func (f *XlsxStorage) WriteHeader() error {
	// 写入表头
	cell, err := excelize.CoordinatesToCellName(1, 1)
	if err != nil {
		return errors.Wrap(err, "tableHeader excelize.CoordinatesToCellName失败")
	}
	header := make([]string, len(f.header))
	for i, h := range f.header {
		header[i] = h.XlsxColName
	}
	if err = f.file.SetSheetRow("Sheet1", cell, &header); err != nil {
		return errors.Wrap(err, "tableHeader file.SetSheetRow失败")
	}
	return nil
}

func (f *XlsxStorage) WriteStructToXlsx() error {
	t := f.rowType
	if f.rowKind == reflect.Pointer { // 反射类型t为指向Struct指针，t.Elem()则为pointer指向的Type
		t = t.Elem()
	}
	n := t.NumField()
	// m 存储struct里filed name-> index 映射
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		m[t.Field(i).Name] = i
	}
	for r := 0; r < f.nr; r++ {
		v := reflect.ValueOf(f.data[r])
		if f.rowKind == reflect.Pointer { // 反射类型v为指针，v.Elem()则为pointer指向的Value
			v = v.Elem()
		}
		row := make([]any, f.nc)
		// 匹配header, fieldName都不匹配则为""
		for i, head := range f.header {
			if idx, ok := m[head.StructFieldName]; ok {
				row[i] = v.FieldByIndex([]int{idx})
			} else {
				row[i] = "null"
			}
		}
		cell, err := excelize.CoordinatesToCellName(1, r+2) //从第二行开始写入
		if err != nil {
			return errors.Wrap(err, "reflect.Struct excelize.CoordinatesToCellName失败")
		}
		// 写入行数据
		if err = f.file.SetSheetRow("Sheet1", cell, &row); err != nil {
			return errors.Wrap(err, "reflect.Struct file.SetSheetRow失败")
		}
	}
	return nil
}

func (f *XlsxStorage) WriteSliceToXlsx() error {
	for r := 0; r < f.nr; r++ {
		v := reflect.ValueOf(f.data[r])
		row := make([]any, f.nc)
		for c := 0; c < f.nc; c++ {
			row[c] = v.Elem()
		}
		cell, err := excelize.CoordinatesToCellName(1, r+2)
		if err != nil {
			return errors.Wrap(err, "reflect.Slice excelize.CoordinatesToCellName失败")
		}
		if err = f.file.SetSheetRow("Sheet1", cell, &row); err != nil {
			return errors.Wrap(err, "reflect.Slice file.SetSheetRow失败")
		}
	}
	return nil
}
