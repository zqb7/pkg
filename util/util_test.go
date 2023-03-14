package util

import (
	"reflect"
	"testing"
)

func TestInterfaceTypePtr(t *testing.T) {
	tests := []struct {
		name string
		args interface{}
		want interface{}
	}{
		{args: struct{}{}, want: &struct{}{}},
		{args: struct{ A int }{A: 1}, want: &struct{ A int }{A: 1}},
		{args: int(0), want: new(int)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceTypePtr(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InterfaceTypePtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
