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
)

type Scaner interface {
	Scan(s string) error
}

type TemplateFieldInfo struct {
	Index   int // 结构体字段的index
	ColName string
}

// columeM k=第几列的索引
type GetFieldInfo func(rows *excelize.Rows, kv map[string]int) (columeM map[int]TemplateFieldInfo, err error)

// endAt 读取到第几行截至，kv：k=结构体字段名,v=对应的索引下标
func SimpleGetFieldInfo(endAt int) func(rows *excelize.Rows, kv map[string]int) (columeM map[int]TemplateFieldInfo, err error) {
	return func(rows *excelize.Rows, kv map[string]int) (columeM map[int]TemplateFieldInfo, err error) {
		columeM = map[int]TemplateFieldInfo{}
		for index := 0; index < endAt && rows.Next(); index++ {
			columns, err := rows.Columns()
			if err != nil {
				return nil, err
			}
			for colIndex := 0; colIndex < len(columns); colIndex++ {
				for name, fieldIndex := range kv {
					if strings.EqualFold(columns[colIndex], name) {
						columeM[colIndex] = TemplateFieldInfo{Index: fieldIndex, ColName: name}
						break
					}
				}
			}
		}
		return columeM, nil
	}
}

func Read(rows *excelize.Rows, template any, getField GetFieldInfo) ([]any, error) {
	rt := reflect.TypeOf(template)
	if rt.Kind() != reflect.Struct {
		return nil, TemplateErr
	}
	columeM, err := getField(rows, fieldNameIndex(template))
	if err != nil {
		return nil, err
	}
	result := make([]any, 0, 0)
	for index := 0; rows.Next(); index++ {
		columns, err := rows.Columns()
		if err != nil {
			return nil, err
		} else if len(columns) == 0 {
			continue
		}
		if len(columeM) < len(columns) {
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
