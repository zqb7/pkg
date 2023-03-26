package excel

import (
	"testing"

	"github.com/xuri/excelize/v2"
)

type Item struct {
	Id    uint32
	Name  string
	Price float64
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
