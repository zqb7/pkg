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

func Read(rows *excelize.Rows, template any, beginAt int) ([]any, error) {
	rt := reflect.TypeOf(template)
	if rt.Kind() != reflect.Struct {
		return nil, TemplateErr
	}
	type columnFieldInfo struct {
		Index   int // 结构体字段的index
		ColName string
	}
	columeM := map[int]columnFieldInfo{} // key=colIndex
	result := make([]any, 0, 0)
	for index := 0; rows.Next(); index++ {
		if index < beginAt {
			continue
		}
		columns, err := rows.Columns()
		if err != nil {
			return nil, err
		} else if len(columns) == 0 {
			continue
		}
		if index == beginAt {
			rv := reflect.ValueOf(reflect.New(rt).Interface()).Elem()
			for fieldIndex := 0; fieldIndex < rv.NumField(); fieldIndex++ {
				field := rv.Field(fieldIndex)
				if !field.CanSet() {
					continue
				}
				for colIndex := 0; colIndex < len(columns); colIndex++ {
					tag := rt.Field(fieldIndex).Tag.Get("col")
					if tag == "-" {
						continue
					} else if tag == "" {
						tag = rt.Field(fieldIndex).Name
					}
					if strings.EqualFold(columns[colIndex], tag) {
						columeM[colIndex] = columnFieldInfo{Index: fieldIndex, ColName: tag}
					}
				}
			}
		}

		if index <= beginAt {
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
			var parseErr error
			switch {
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
