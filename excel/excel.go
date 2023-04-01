package excel

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

var (
	TemplateErr             = errors.New("template must be struct")
	ColFieldMisalignmentErr = errors.New("field misalignment")
	UnsupportypeErr         = errors.New("unsupported type")
)

func OpenFile(filename string) (file *excelize.File, err error) {
	return excelize.OpenFile(filename)
}

type Scaner interface {
	Scan(value any) error
}

type templateFieldInfo struct {
	Index   int // 结构体字段的index
	ColName string
}

// colNames:key=第几列的字段名,value=该字段对应的第几列的索引
type ColNameIndex func(rows *excelize.Rows) (colNames map[string]int, err error)

// endAt 读取到第几行截至
func SimpleColNameIndex(endAt int) ColNameIndex {
	return func(rows *excelize.Rows) (colNames map[string]int, err error) {
		colNames = map[string]int{}
		for index := 0; index < endAt && rows.Next(); index++ {
			columns, err := rows.Columns()
			if err != nil {
				return nil, err
			}
			for colIndex := 0; colIndex < len(columns); colIndex++ {
				colNames[columns[colIndex]] = colIndex
			}
		}
		return colNames, nil
	}
}

func Read(rows *excelize.Rows, template any, f ColNameIndex) ([]any, error) {
	rt := reflect.TypeOf(template)
	if rt.Kind() != reflect.Struct {
		return nil, TemplateErr
	}
	colNames, err := f(rows)
	if err != nil {
		return nil, err
	}
	columeM := toFieldInfo(template, colNames)
	result := make([]any, 0, 0)
	for index := 0; rows.Next(); index++ {
		columns, err := rows.Columns()
		if err != nil {
			return nil, err
		} else if len(columns) == 0 {
			continue
		}
		if len(columns) < len(columeM) {
			return nil, ColFieldMisalignmentErr
		}
		obj := reflect.New(rt).Interface()
		rv := reflect.ValueOf(obj).Elem()
		for colIndex, col := range columeM {
			colValue := columns[colIndex]
			field := rv.Field(col.Index)
			if !field.CanSet() {
				continue
			}
			var parseErr error
			switch {
			case reflect.PointerTo(field.Type()).Implements(reflect.TypeOf((*Scaner)(nil)).Elem()):
				parseErr = field.Addr().Interface().(Scaner).Scan(colValue)
			case field.CanUint():
				var uv uint64
				uv, parseErr = strconv.ParseUint(colValue, 10, 64)
				field.SetUint(uv)
			case field.CanInt():
				var uv int64
				uv, parseErr = strconv.ParseInt(colValue, 10, 64)
				field.SetInt(uv)
			case field.CanFloat():
				var uv float64
				uv, parseErr = strconv.ParseFloat(colValue, 10)
				field.SetFloat(uv)
			case field.Kind() == reflect.String:
				field.SetString(colValue)
			case field.Kind() == reflect.Bool:
				var uv bool
				strconv.ParseBool(colValue)
				field.SetBool(uv)
			case field.Kind() == reflect.Array:
				parseErr = arrayDecode(field, colValue)
			case field.Kind() == reflect.Slice:
				parseErr = sliceDecode(field, colValue)
			}

			if parseErr != nil {
				return nil, parseErr
			}
		}
		result = append(result, obj)
	}
	return result, nil
}

// 获取结构体字段的名字以及对应的下标
func fieldNameIndex(template any) map[string]int {
	rt := reflect.TypeOf(template)
	kv := make(map[string]int, rt.NumField())
	for fieldIndex := 0; fieldIndex < rt.NumField(); fieldIndex++ {
		tag := rt.Field(fieldIndex).Tag.Get("col")
		if tag == "-" {
			continue
		} else if tag == "" {
			tag = rt.Field(fieldIndex).Name
		}
		kv[tag] = fieldIndex
	}
	return kv
}

// 根据给的结构体，以及表格中字段名以及对应的第几列的下标，生成 columeM k=表格的第几列的下标 v=对应的结构体的相关数据
func toFieldInfo(template any, colNames map[string]int) (columeM map[int]templateFieldInfo) {
	columeM = map[int]templateFieldInfo{}
	kv := fieldNameIndex(template)
	for colName, colIndex := range colNames {
		for name, fieldIndex := range kv {
			if strings.EqualFold(colName, name) {
				columeM[colIndex] = templateFieldInfo{Index: fieldIndex, ColName: name}
				break
			}
		}
	}
	return columeM
}
