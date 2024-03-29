package excel

import (
	"errors"
	"io"
	"reflect"
	"sort"
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

func OpenReader(r io.Reader) (file *excelize.File, err error) {
	return excelize.OpenReader(r)
}

type Scaner interface {
	Scan(value any) error
}

type templateFieldInfo struct {
	Index   int // 结构体字段的index
	ColName string
}

// endAt:截止读取到(含)第几行(1为第一行开始) colNames:key=第几列的字段名,value=该字段对应的第几列的索引
type ColNameIndex func() (endAt int, callback func(rows [][]string) (colNames map[string]int, err error))

// breakOff:当为true时，表明对该字段的解析结束了(无论成功与否)
type Decoder func(rv reflect.Value, colValue string) (breakOff bool, err error)

var decoders = make([]Decoder, 0, 32)

func SimpleColNameIndex() ColNameIndex {
	return func() (endAt int, callback func(rows [][]string) (colNames map[string]int, err error)) {
		return 2, func(rows [][]string) (colNames map[string]int, err error) {
			colNames = map[string]int{}
			for _, row := range rows {
				for colIndex, colValue := range row {
					colNames[colValue] = colIndex
				}
			}
			return colNames, nil
		}
	}
}

func Read(rows *excelize.Rows, template any, f ColNameIndex) ([]any, error) {
	rt := reflect.TypeOf(template)
	if rt.Kind() != reflect.Struct {
		return nil, TemplateErr
	}
	colNames, err := getColNames(rows, f)
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
		obj := reflect.New(rt).Interface()
		rv := reflect.ValueOf(obj).Elem()
	walk:
		for colIndex, col := range columeM {
			var colValue string
			if colIndex < len(columns) {
				colValue = columns[colIndex]
			}

			field := rv.Field(col.Index)
			if !field.CanSet() {
				continue
			}
			var parseErr error
			var breakOff bool
			for _, f := range decoders {
				breakOff, parseErr = f(field, colValue)
				if breakOff {
					continue walk
				}
				if parseErr != nil {
					return nil, parseErr
				}
			}

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
			case field.Kind() == reflect.Map:
				parseErr = mapDecode(field, colValue)
			}

			if parseErr != nil {
				return nil, parseErr
			}
		}
		result = append(result, obj)
	}
	return result, nil
}

func GetRows(rows *excelize.Rows, f ColNameIndex) ([][]string, error) {
	var result [][]string
	colNames, err := getColNames(rows, f)
	if err != nil {
		return nil, err
	}
	var maxColIndex int
	var colIndexs []int
	for _, colIndex := range colNames {
		if colIndex > maxColIndex {
			maxColIndex = colIndex
		}
		colIndexs = append(colIndexs, colIndex)
	}
	sort.Ints(colIndexs)
	for index := 0; rows.Next(); index++ {
		columns, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		if sub := maxColIndex + 1 - len(columns); sub > 0 {
			columns = append(columns, make([]string, sub)...)
		}
		var r = make([]string, 0, maxColIndex)
		for _, colIndex := range colIndexs {
			r = append(r, columns[colIndex])
		}
		result = append(result, r)
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

// colNames k=表格的列名，v=表示第几列的下标
func getColNames(rows *excelize.Rows, f ColNameIndex) (colNames map[string]int, err error) {
	endAt, callback := f()
	var rowsSlice [][]string
	for index := 0; index < endAt && rows.Next(); index++ {
		columns, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		rowsSlice = append(rowsSlice, columns)
	}
	colNames, err = callback(rowsSlice)
	return colNames, err
}

// 注册自定义的解析器
func RegisterDecoder(funs ...Decoder) {
	decoders = append(decoders, funs...)
}
