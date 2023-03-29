package excel

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/xuri/excelize/v2"
)

type MyInt int32

func (m *MyInt) Scan(s string) error {
	v, err := strconv.ParseInt(s, 10, 64)
	*m = MyInt(v)
	return err
}

type Item struct {
	Id     uint32
	Name   string
	Price  float64
	Price2 MyInt
}

func TestRead(t *testing.T) {
	f, err := excelize.OpenFile("test.xlsx")
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
				&Item{Id: 1, Name: "test1", Price: 0.01, Price2: 3},
				&Item{Id: 2, Name: "test2", Price: 10, Price2: 4},
				&Item{Id: 3, Name: "test3", Price: 9.9, Price2: 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Sheet, func(t *testing.T) {
			itemRows, err := f.Rows(tt.Sheet)
			if err != nil {
				t.Fatal(err)
			}
			result, err := Read(itemRows, Item{}, SimpleGetFieldInfo(2))
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
