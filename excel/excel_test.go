package excel

import (
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
	itemRows, err := f.Rows("Item")
	if err != nil {
		t.Fatal(err)
	}
	Read(itemRows, Item{}, 1)
}
