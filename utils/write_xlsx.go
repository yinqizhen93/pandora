package utils

import (
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"reflect"
)

type XlsxStorage struct {
	file *excelize.File
	//export store the index of "export field" in struct,
	//"export field"指没有export的key 或者export 值为 "-"
	export  []int
	data    []any
	nr      int
	nc      int
	rowType reflect.Type
	rowKind reflect.Kind
}

func NewXlsxStorage(file *excelize.File, data []any) (*XlsxStorage, error) {
	var nc int
	nr := len(data)
	if nr == 0 {
		return nil, errors.New("data can not be empty")
	}
	// 以数据第一行类型为准
	t := reflect.TypeOf(data[0])
	k := t.Kind()
	if k == reflect.Pointer { // 反射类型t为指针，t.Elem()则为pointer指向的Type
		t = t.Elem()
		if t.Kind() != reflect.Struct {
			return nil, errors.New("data can only be []slice or []struct or []*struct")
		}
	}
	nc = t.NumField()
	if nc == 0 {
		return nil, errors.New("struct can not be empty")
	}
	xs := XlsxStorage{
		file:    file,
		data:    data,
		export:  make([]int, 0),
		nr:      nr,
		nc:      nc,
		rowType: t,
		rowKind: k,
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

// WriteHeader 从struct tag 获取table header, tag 标识为"export", 如果值为"-"，则代表不导出
func (f *XlsxStorage) WriteHeader() error {
	// 写入表头
	cell, err := excelize.CoordinatesToCellName(1, 1)
	if err != nil {
		return errors.Wrap(err, "tableHeader excelize.CoordinatesToCellName失败")
	}
	header := make([]string, 0)
	for i := 0; i < f.nc; i++ {
		tag := f.rowType.Field(i).Tag.Get("export")
		// tag 没有 export的key 或者export 值为 "-",则标识不导出
		if tag == "-" || tag == "" {
			continue
		}
		f.export = append(f.export, i)
		header = append(header, tag)
	}
	// 必须至少一个字段是导出的
	if len(header) == 0 {
		return errors.Wrap(err, "no field is export")
	}
	if err = f.file.SetSheetRow("Sheet1", cell, &header); err != nil {
		return errors.Wrap(err, "tableHeader file.SetSheetRow失败")
	}
	return nil
}

func (f *XlsxStorage) WriteStructToXlsx() error {
	for r := 0; r < f.nr; r++ {
		v := reflect.ValueOf(f.data[r])
		if f.rowKind == reflect.Pointer { // 反射类型v为指针，v.Elem()则为pointer指向的Value
			v = v.Elem()
		}
		row := make([]any, len(f.export))
		for i, idx := range f.export {
			row[i] = v.FieldByIndex([]int{idx})
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
