package excel

import (
	"reflect"
	"strconv"
	"testing"
)

type MyInt int32

func (m *MyInt) Scan(value any) error {
	v, err := strconv.ParseInt(value.(string), 10, 64)
	*m = MyInt(v)
	return err
}

type Item struct {
	Id     uint32
	Name   string
	Price  float64
	Price2 MyInt
	Price3 []uint32
	Price4 [2]float64
}

type TestDecode struct {
	Id     uint32
	Slice1 [][]uint
	Slice2 [][]int32
	Arr1   [2]uint8
	Arr2   [2]int8
}

func TestRead(t *testing.T) {
	f, err := OpenFile("test.xlsx")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		Sheet    string
		Template any
		want     []any
	}{
		{
			Sheet: "Item", Template: Item{}, want: []any{
				&Item{Id: 1, Name: "test1", Price: 0.01, Price2: 3, Price3: []uint32{2, 3, 4}, Price4: [2]float64{7, 0}},
				&Item{Id: 2, Name: "test2", Price: 10, Price2: 4, Price3: []uint32{5, 6, 7}, Price4: [2]float64{0.01, 0.9999}},
				&Item{Id: 3, Name: "test3", Price: 9.9, Price2: 5, Price3: nil, Price4: [2]float64{3.1415, 0}},
			},
		},
		{
			Sheet: "TestDecode", Template: TestDecode{}, want: []any{
				&TestDecode{Id: 1, Slice1: [][]uint{{1, 2, 3}, {4, 5, 6}}, Slice2: [][]int32{{-1, -2, -3}, {-4, -5, -6}}, Arr1: [2]uint8{1, 2}, Arr2: [2]int8{-1, -2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Sheet, func(t *testing.T) {
			itemRows, err := f.Rows(tt.Sheet)
			if err != nil {
				t.Fatal(err)
			}
			result, err := Read(itemRows, tt.Template, SimpleColNameIndex(2))
			if err != nil {
				t.Fatal(err)
			}
			if len(tt.want) != len(result) {
				t.Fatalf("want len:%d, got len:%d", len(tt.want), len(result))
			}
			for index := range tt.want {
				if !reflect.DeepEqual(tt.want[index], result[index]) {
					t.Fatalf("want:%+v got:%+v", tt.want[index], result[index])
				}
			}
		})
	}

}
