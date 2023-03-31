package excel

import (
	"encoding/json"
	"errors"
	"reflect"
)

var UnsupportypeErr = errors.New("unsupported type")

func sliceDecode(rv reflect.Value, str string) error {
	var err error
	switch v := rv.Interface().(type) {
	case []uint8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.AppendSlice(rv, reflect.ValueOf(v)))
	case []uint16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.AppendSlice(rv, reflect.ValueOf(v)))
	case []uint32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.AppendSlice(rv, reflect.ValueOf(v)))
	case []uint64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.AppendSlice(rv, reflect.ValueOf(v)))
	case []uint:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.AppendSlice(rv, reflect.ValueOf(v)))

	case []int8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.AppendSlice(rv, reflect.ValueOf(v)))
	case []int16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.AppendSlice(rv, reflect.ValueOf(v)))
	case []int32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.AppendSlice(rv, reflect.ValueOf(v)))
	case []int64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.AppendSlice(rv, reflect.ValueOf(v)))
	case []int:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.AppendSlice(rv, reflect.ValueOf(v)))

	case []float32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.AppendSlice(rv, reflect.ValueOf(v)))
	case []float64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.AppendSlice(rv, reflect.ValueOf(v)))

	case []string:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.AppendSlice(rv, reflect.ValueOf(v)))

	default:
		return UnsupportypeErr
	}
	return err
}

func arrayDecode(rv reflect.Value, str string) error {
	var err error
	switch v := rv.Interface().(type) {
	case [1]uint8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [2]uint8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [3]uint8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [4]uint8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [5]uint8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [6]uint8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [7]uint8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [8]uint8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [9]uint8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [10]uint8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))

	case [1]uint16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [2]uint16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [3]uint16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [4]uint16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [5]uint16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [6]uint16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [7]uint16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [8]uint16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [9]uint16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [10]uint16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))

	case [1]uint32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [2]uint32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [3]uint32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [4]uint32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [5]uint32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [6]uint32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [7]uint32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [8]uint32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [9]uint32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [10]uint32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))

	case [1]uint64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [2]uint64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [3]uint64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [4]uint64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [5]uint64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [6]uint64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [7]uint64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [8]uint64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [9]uint64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [10]uint64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))

	case [1]uint:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [2]uint:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [3]uint:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [4]uint:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [5]uint:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [6]uint:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [7]uint:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [8]uint:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [9]uint:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [10]uint:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))

	case [1]int8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [2]int8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [3]int8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [4]int8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [5]int8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [6]int8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [7]int8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [8]int8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [9]int8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [10]int8:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))

	case [1]int16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [2]int16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [3]int16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [4]int16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [5]int16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [6]int16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [7]int16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [8]int16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [9]int16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [10]int16:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))

	case [1]int32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [2]int32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [3]int32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [4]int32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [5]int32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [6]int32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [7]int32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [8]int32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [9]int32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [10]int32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))

	case [1]int64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [2]int64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [3]int64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [4]int64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [5]int64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [6]int64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [7]int64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [8]int64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [9]int64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [10]int64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))

	case [1]int:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [2]int:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [3]int:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [4]int:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [5]int:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [6]int:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [7]int:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [8]int:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [9]int:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [10]int:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))

	case [1]float32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [2]float32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [3]float32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [4]float32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [5]float32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [6]float32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [7]float32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [8]float32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [9]float32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [10]float32:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))

	case [1]float64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [2]float64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [3]float64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [4]float64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [5]float64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [6]float64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [7]float64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [8]float64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [9]float64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [10]float64:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))

	case [1]string:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [2]string:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [3]string:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [4]string:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [5]string:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [6]string:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [7]string:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [8]string:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [9]string:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))
	case [10]string:
		err = json.Unmarshal([]byte(str), &v)
		rv.Set(reflect.ValueOf(v))

	default:
		return UnsupportypeErr
	}
	return err
}
